package helper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

var jsonData map[string]string

func GetUrlMapFromJson() map[string]string {
	if len(jsonData) > 0 {
		return jsonData
	}

	file, error := ioutil.ReadFile("./urls.json")

	if error != nil {
		log.Fatal("Error while opening ./urls.json")
	}

	error = json.Unmarshal(file, &jsonData)

	if error != nil {
		log.Fatal("Error while unmarshalling")
	}

	return jsonData
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Table    string
}

func GetUrlMapFromDB(config DBConfig) map[string]string {
	if len(jsonData) > 0 {
		return jsonData
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	rows, error := db.Query(fmt.Sprintf("SELECT * FROM %s", config.Table))

	if error != nil {
		panic(error)
	}

	jsonData = make(map[string]string)

	defer rows.Close()

	for rows.Next() {
		var name string
		var url string

		err = rows.Scan(&name, &url)
		if err != nil {
			panic(err)
		}

		jsonData[name] = url
	}

	if err != nil {
		panic(err)
	}

	defer db.Close()

	return jsonData
}
