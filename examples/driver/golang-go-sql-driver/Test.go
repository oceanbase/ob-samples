package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type Entity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	workspace = "/workspace/ob-example"
	sqlFile   = "tests/sql/test.sql"
	tableName = "t_test"
)

func main() {
	// For details about 'dataSourceName', see
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	var (
		host     = "127.0.0.1"
		port     = 2881
		dbName   = "test"
		username = "root@test"
		password = ""
	)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filePath := filepath.Join(workspace, sqlFile)
	log.Println("Load sql file from: " + filePath)

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sqlFileContent := string(bytes)
	log.Println("Exec sql:\n" + sqlFileContent)

	if _, err = db.Exec(sqlFileContent); err != nil {
		log.Fatal(err)
	}

	selectSql := "SELECT * FROM " + tableName
	log.Println("Query sql: " + selectSql)

	rows, err := db.Query(selectSql)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get rows:")
	count := 0
	for rows.Next() {
		var entity Entity
		if err = rows.Scan(&entity.ID, &entity.Name); err != nil {
			log.Fatal(err)
		}
		b, _ := json.Marshal(entity)
		log.Printf("## row %d: %s", count, string(b))
		count++
	}
}
