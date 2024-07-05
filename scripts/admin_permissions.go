package scripts

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"trec/config"

	"go.uber.org/zap"
)

func ImportAdminPermission(db *sql.DB, logger *zap.Logger, configs *config.Configurations) {
	userOid := configs.App.AzureAdminOid
	permissionFile := "./master_db/permissions.csv"
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connStr == "" {
		permissionFile = "../master_db/permissions.csv"
	}

	pFile, err := os.Open(permissionFile)
	if err != nil {
		log.Fatal(err)
	}
	defer pFile.Close()
	pReader := csv.NewReader(pFile)
	pRecords, err := pReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var userId string
	var deleted_at any
	err = db.QueryRow("SELECT id, deleted_at FROM users WHERE oid = $1", userOid).Scan(&userId, &deleted_at)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("", zap.Error(err))
		return
	}
	if deleted_at != nil {
		logger.Error(fmt.Sprintf("User %s is deleted", userOid))
		return
	}
	var id string
	err = db.QueryRow("SELECT id FROM entity_permissions WHERE entity_id = $1 AND entity_type = 'user'", userId).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error reading user permissions: %v", err)
		return
	}
	if id != "" {
		log.Println("[SCRIPTS] Admin permissions already imported")
		return
	}
	for _, record := range pRecords[1:] {
		entityId := userId
		entityType := "user"
		permission := record[0]
		sqlStatement := `
		INSERT INTO public.entity_permissions (entity_id, entity_type, permission_id, for_owner, for_team, for_all)
		VALUES ($1, $2, $3, false, false, true)`
		_, err = db.Exec(sqlStatement, entityId, entityType, permission)
		if err != nil {
			log.Fatalf("Unable to insert entity permissions. %v", err)
		}
	}
	log.Println("[SCRIPTS] Admin permissions imported successfully")
}
