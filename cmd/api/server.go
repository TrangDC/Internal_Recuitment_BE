package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"trec/config"
	"trec/ent"
	"trec/internal/azuread"
	"trec/internal/azurestorage"
	"trec/internal/pg"
	"trec/internal/servicebus"
	"trec/middleware"
	"trec/models"
	"trec/resolver"
	"trec/rest"
	"trec/scripts"
	"trec/service"

	"entgo.io/contrib/entgql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/golang-module/carbon/v2"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewServerCmd(configs *config.Configurations, logger *zap.Logger, i18n models.I18n) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with graphql",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()

			carbon.SetDefault(carbon.Default{
				WeekStartsAt: carbon.Monday,
			})

			defer func() {
				err := recover()
				if err != nil {
					logger.Fatal("recover error", zap.Any("error", err))
				}
			}()

			// Create postgresql connection
			db, err := pg.NewDBConnection(configs.Postgres, logger)
			if err != nil {
				logger.Error("Getting error connect to postgresql database", zap.Error(err))
				os.Exit(1)
			}
			defer db.Close()
			// Create ent client
			entdrv := entsql.OpenDB("postgres", db)

			entOptions := []ent.Option{
				ent.Driver(entdrv),
			}
			if configs.App.Debug {
				entOptions = append(entOptions, ent.Debug())
			}

			entClient := ent.NewClient(entOptions...)
			defer entClient.Close()

			// Object Storage with Azure Storage Container
			objectStorageClient, err := azurestorage.NewAzureStorage(configs.AzureStorage)
			if err != nil {
				logger.Error("Getting error connect to azure storage", zap.Error(err))
				os.Exit(1)
			}

			serviceBusClient, err := servicebus.NewServiceBus(configs.ServiceBus, entClient)
			if err != nil {
				logger.Error("Getting error connect to service bus", zap.Error(err))
				os.Exit(1)
			}

			// Create validator
			validator := validator.New()
			// Add translator for validator
			en := en.New()
			uni := ut.New(en, en)
			validationTranslator, _ := uni.GetTranslator("en")
			// Register default translation for validator
			err = en_translations.RegisterDefaultTranslations(validator, validationTranslator)
			if err != nil {
				logger.Error("Getting error from register default translation", zap.Error(err))
				os.Exit(1)
			}

			// Create sessions store with cookie
			sessionStore := sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))
			sessionStore.MaxAge(60)

			// Create AzureAD OAuth client
			var azureADOAuthClient azuread.AzureADOAuth
			// Authentication with AzureAD
			if configs.AzureADOAuth.Enabled {
				azureADOAuthClient, err = azuread.NewAzureADOAuth(configs.AzureADOAuth, sessionStore)
				if err != nil {
					logger.Error("Getting error create to azuread oauth", zap.Error(err))
					os.Exit(1)
				}
			}
			serviceRegistry := service.NewService(azureADOAuthClient, objectStorageClient, serviceBusClient, i18n, entClient, logger, configs)
			restController := rest.NewRestController(serviceRegistry, configs, logger)

			// GraphQL schema resolver handler.
			resolverHandler := handler.NewDefaultServer(resolver.NewSchema(serviceRegistry, entClient, validator, validationTranslator, logger))
			// Handle all transaction for resolver.
			resolverHandler.Use(entgql.Transactioner{TxOpener: entClient})

			// Handler for GraphQL Playground
			playgroundHandler := playground.Handler("GraphQL Playground", "/graphql")

			if !configs.App.Debug {
				gin.SetMode(gin.ReleaseMode)
			}
			r := gin.New()
			// Handle a method not allowed.
			r.HandleMethodNotAllowed = true

			// Use middlewares
			r.Use(
				ginzap.Ginzap(logger, time.RFC3339, true),
				ginzap.RecoveryWithZap(logger, true),
				middleware.CorsMiddleware(),
				middleware.RequestCtxMiddleware(),
			)

			readyRouter := r.Group("ready")
			{
				readyRouter.GET("/readiness", middleware.ReadinessCheckMiddleware(db, logger))
				readyRouter.GET("/liveliness", middleware.LivelinessCheckMiddleware(db, logger))
			}

			if configs.AzureADOAuth.Enabled {
				authRouter := r.Group("/auth")
				{
					authRouter.GET("/login", restController.Auth().OAuthLogin)
					authRouter.GET("/callback", restController.Auth().OAuthCallback)
					authRouter.POST("/refresh", restController.Auth().RefreshToken)
					// This one is just debug purpose
					// authRouter.GET("/redirect", func(c *gin.Context) {
					// 	c.JSON(http.StatusOK, gin.H{"raw_query": c.Request.URL.RawQuery})
					// })
				}
			}

			talenaRouter := r.Group("/talena")
			{
				talenaRouter.GET("/login", restController.Talena().TalenaLogin)
			}

			graphqlRouter := r.Group("/graphql")
			{
				if configs.AzureADOAuth.Enabled {
					graphqlRouter.Use(middleware.AuthenticateMiddleware(azureADOAuthClient, db, logger))
				}

				graphqlRouter.Use(middleware.ValidPermission(db))

				graphqlRouter.POST("", func(c *gin.Context) {
					resolverHandler.ServeHTTP(c.Writer, c.Request)
				})

				graphqlRouter.OPTIONS("", func(c *gin.Context) {
					c.Status(200)
				})
			}

			if configs.App.Debug {
				// Enable playground for query/testing with debug
				r.GET("/", func(c *gin.Context) {
					playgroundHandler.ServeHTTP(c.Writer, c.Request)
				})
			}

			scripts.ImportMasterDB(db, logger, configs)
			scripts.ImportAdminPermission(db, logger, configs)

			server := &http.Server{
				ReadTimeout:  15 * time.Second,
				WriteTimeout: 30 * time.Second,
				Addr:         ":8000",
				Handler:      r,
			}

			// Graceful shutdown
			idleConnectionsClosed := make(chan struct{})

			messages := make(chan models.Messages, 10000) // Use buffered channel to avoid blocking
			go serviceBusClient.ListenToEmailSubscription(messages)
			go serviceBusClient.ListenToInterviewScheduleSubscription(messages)
			go serviceBusClient.ProcessMessages(ctx, messages)

			go func() {
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

				<-c

				ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
				defer cancel()
				// A interrupt signal has sent to us, let's shutdown server with gracefully
				logger.Debug("Stopping server...")

				if err := server.Shutdown(ctx); err != nil {
					logger.Error("Graceful shutdown has failed with error: %s", zap.Error(err))
				}

				if err := db.Close(); err != nil {
					logger.Error("Closing db connection has error", zap.Error(err))
				}

				close(idleConnectionsClosed)
			}()

			go func() {
				logger.Debug("Listing on the port: 8000")
				if err := server.ListenAndServe(); err != http.ErrServerClosed {
					logger.Error("Run server has error", zap.Error(err))
					// Exit the application if run fail
					os.Exit(1)
				} else {
					logger.Info("Server was closed by shutdown gracefully")
				}
			}()

			<-idleConnectionsClosed
		},
	}
}
