package model

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const (
	maxLifetime  int = 10
	maxOpenConns int = 10
	maxIdleConns int = 10
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAllContracts() []string {
	contracts := []string{}

	viper.SetConfigFile("configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
	}

	host := viper.GetString("database.host")
	port, _ := strconv.Atoi(viper.GetString("database.port"))
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	database := viper.GetString("database.database")

	// Initialize connection string.
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	err = db.Ping()
	checkError(err)

	// Variables for printing column data when scanned.
	var (
		address string
	)

	// Read some data from the table.
	rows, err := db.Query("SELECT address from contracts;")
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&address)
		checkError(err)
		contracts = append(contracts, address)
	}
	err = rows.Err()
	checkError(err)

	return contracts
}
