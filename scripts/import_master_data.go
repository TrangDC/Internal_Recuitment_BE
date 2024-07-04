package scripts

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"trec/config"

	"go.uber.org/zap"
)

func ImportMasterDB(db *sql.DB, logger *zap.Logger, configs *config.Configurations) {
	permissionGroupFile := "./master_db/permission_groups.csv"
	permissionFile := "./master_db/permissions.csv"
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connStr == "" {
		permissionGroupFile = "../master_db/permission_groups.csv"
		permissionFile = "../master_db/permissions.csv"
	}

	pGFile, err := os.Open(permissionGroupFile)
	if err != nil {
		log.Fatal(err)
	}
	defer pGFile.Close()
	pGReader := csv.NewReader(pGFile)
	pGRecords, err := pGReader.ReadAll()
	if err != nil {
		log.Fatal(err)
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

	truncateStmt := `TRUNCATE TABLE public.permission_groups, public.permissions`
	_, err = db.Exec(truncateStmt)
	if err != nil {
		log.Fatalf("Unable to truncate the tables. %v", err)
	}

	for _, record := range pGRecords[1:] { // Bỏ qua hàng tiêu đề
		id := record[0]
		title := record[1]
		parentId := record[2]
		groupType := record[3]
		orderId := record[4]
		if parentId == "" {
			sqlStatement := `
			INSERT INTO public.permission_groups (id, title, group_type, order_id)
			VALUES ($1, $2, $3, $4)`
			_, err = db.Exec(sqlStatement, id, title, groupType, orderId)
			if err != nil {
				log.Fatalf("Unable to execute the query. %v", err)
			}
		} else {
			sqlStatement := `
			INSERT INTO public.permission_groups (id, title, parent_id, group_type, order_id)
			VALUES ($1, $2, $3, $4, $5)`
			_, err = db.Exec(sqlStatement, id, title, parentId, groupType, orderId)
			if err != nil {
				log.Fatalf("Unable to execute the query. %v", err)
			}
		}
	}

	for _, record := range pRecords[1:] { // Bỏ qua hàng tiêu đề
		id := record[0]
		title := record[1]
		groupId := record[2]
		forOwner := record[3]
		forTeam := record[4]
		forAll := record[5]
		operationName := record[6]
		parentId := record[7]
		orderId := record[8]
		if parentId == "" {
			sqlStatement := `
		INSERT INTO public.permissions (id, title, group_id, for_owner, for_team, for_all, operation_name, order_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
			_, err = db.Exec(sqlStatement, id, title, groupId, forOwner, forTeam, forAll, operationName, orderId)
			if err != nil {
				log.Fatalf("Unable to execute the query. %v", err)
			}
		} else {
			sqlStatement := `
		INSERT INTO public.permissions (id, title, group_id, for_owner, for_team, for_all, operation_name, parent_id, order_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
			_, err = db.Exec(sqlStatement, id, title, groupId, forOwner, forTeam, forAll, operationName, parentId, orderId)
			if err != nil {
				log.Fatalf("Unable to execute the query. %v", err)
			}
		}
	}
}
