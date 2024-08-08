package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"trec/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TalenaController interface {
	TalenaLogin(c *gin.Context)
}

// talenaCtrlImpl is the implementetion of TalenaController.
type talenaCtrlImpl struct {
	talenaConfig config.TalenaConfig
	logger       *zap.Logger
}

// NewTalenaController creates a new TalenaController.
func NewTalenaController(talenaConfig config.TalenaConfig, logger *zap.Logger) TalenaController {
	return &talenaCtrlImpl{
		talenaConfig: talenaConfig,
		logger:       logger,
	}
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Body struct {
	AccessToken string `json:"access_token"`
	User        struct {
		ID         string `json:"id"`
		Email      string `json:"email"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Role       string `json:"role"`
		Avatar     string `json:"avatar"`
		Department string `json:"department"`
		Workspace  struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"workspace"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
}

type LoginOutput struct {
	AccessToken string `json:"access_token"`
}

// talenaLogin is the handler for the /talena/login endpoint.
func (ctl *talenaCtrlImpl) TalenaLogin(c *gin.Context) {
	client := &http.Client{}
	host := ctl.talenaConfig.Host
	Email := ctl.talenaConfig.Email
	password := ctl.talenaConfig.Password
	login := LoginInput{
		Email:    Email,
		Password: password,
	}
	reqBody, err := json.Marshal(login)
	if err != nil {
		ctl.logger.Error("failed to marshal login input", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	req, err := http.NewRequest("POST", host+"/auth/login", bytes.NewBuffer(reqBody))
	if err != nil {
		ctl.logger.Error("failed to create request", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		ctl.logger.Error("failed to send request", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if res.StatusCode != http.StatusOK {
		ctl.logger.Error("failed to login", zap.Int("status", res.StatusCode))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		ctl.logger.Error("failed to send request", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	resBody := Body{}
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		ctl.logger.Error("failed to unmarshal response", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, LoginOutput{
		AccessToken: resBody.AccessToken,
	})
}
