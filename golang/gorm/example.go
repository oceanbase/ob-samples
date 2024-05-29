package main

import (
	"flag"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	stdmysql "github.com/go-sql-driver/mysql"
)

var (
	host     string
	port     int
	username string
	password string
	database string
)

func init() {
	flag.StringVar(&host, "host", "", "host")
	flag.IntVar(&port, "port", 0, "port")
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&database, "database", "", "database")

	flag.Parse()
}

type Product struct {
	ID    uint `gorm:"primaryKey;default:auto_random()"`
	Code  string
	Price uint
}

func main() {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	conf := stdmysql.Config{
		Addr:   fmt.Sprintf("%s:%d", host, port),
		User:   username,
		Passwd: password,
		DBName: database,
		//TODO: set parameters
	}
	dial := mysql.New(mysql.Config{DSNConfig: &conf})
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&Product{}); err != nil {
		panic(err)
	}

	insertProduct := &Product{Code: "D42", Price: 100}

	db.Create(insertProduct)
	fmt.Printf("insert ID: %d, Code: %s, Price: %d\n",
		insertProduct.ID, insertProduct.Code, insertProduct.Price)

	readProduct := &Product{}
	db.First(&readProduct, "code = ?", "D42") // find product with code D42

	fmt.Printf("read ID: %d, Code: %s, Price: %d\n",
		readProduct.ID, readProduct.Code, readProduct.Price)
}
