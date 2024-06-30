package main

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	permissionGroupFile := "./master_db/permission_groups.csv"
	permissionFile := "./master_db/permissions.csv"
	connStr := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connStr == "" {
		connStr = "postgres://postgres:Hoa**3264@localhost:5432/trec?sslmode=disable"
		permissionGroupFile = "../master_db/permission_groups.csv"
		permissionFile = "../master_db/permissions.csv"
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
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

	truncateStmt := `TRUNCATE TABLE public.permission_groups, public.permissions CASCADE`
	_, err = db.Exec(truncateStmt)
	if err != nil {
		log.Fatalf("Unable to truncate the tables. %v", err)
	}

	for _, record := range pGRecords[1:] { // Bỏ qua hàng tiêu đề
		if record[2] == "" {
			id := record[0]
			title := record[1]
			sqlStatement := `
			INSERT INTO public.permission_groups (id, title)
			VALUES ($1, $2)`
			_, err = db.Exec(sqlStatement, id, title)
			if err != nil {
				log.Fatalf("Unable to execute the query. %v", err)
			}
		} else {
			id := record[0]
			title := record[1]
			parentId := record[2]
			sqlStatement := `
			INSERT INTO public.permission_groups (id, title, parent_id)
			VALUES ($1, $2, $3)`
			_, err = db.Exec(sqlStatement, id, title, parentId)
			if err != nil {
				log.Fatalf("Unable to execute the query. %v", err)
			}
		}
	}

	for _, record := range pRecords[1:] { // Bỏ qua hàng tiêu đề
		id := record[0]
		title := record[1]
		groupId := record[2]

		sqlStatement := `
				INSERT INTO public.permissions (id, title, group_id)
				VALUES ($1, $2, $3)`
		_, err = db.Exec(sqlStatement, id, title, groupId)
		if err != nil {
			log.Fatalf("Unable to execute the query. %v", err)
		}
	}
}
