package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Entity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	var (
		host     = "127.0.0.1"
		port     = 2881
		dbName   = "test"
		username = "root@test"
		password = ""
	)

	// For details about 'dataSourceName', see
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err = db.Exec("DROP TABLE IF EXISTS `t_test`"); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec("CREATE TABLE `t_test` (" +
		"    `id`   int(10) NOT NULL AUTO_INCREMENT," +
		"    `name` varchar(20) DEFAULT NULL," +
		"    PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE = utf8_bin"); err != nil {
		log.Fatal(err)
	}
	if _, err = db.Exec("INSERT INTO `t_test` VALUES (default, 'Hello OceanBase')"); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM `t_test`")
	if err != nil {
		log.Fatal(err)
	}

	var result []Entity
	for rows.Next() {
		var entity Entity
		if err = rows.Scan(&entity.ID, &entity.Name); err != nil {
			log.Fatal(err)
		}
		result = append(result, entity)
	}

	log.Printf("Query got result: +%v", result)
}
