package main

import (
	"net/http"

	"morphbreed.com/url/helper"
	"morphbreed.com/url/pkg/handlers"
)

func main() {
	// urlMap := utils.GetUrlMapFromJson()
	config := helper.DBConfig{Host: "localhost", Port: 5432, User: "postgres", Password: "muh", DBName: "urlshortener", Table: "urls"}
	urlMap := helper.GetUrlMapFromDB(config)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		handlers.HandleURLShortenerRequest(rw, r, urlMap)
	})

	http.ListenAndServe(":3000", nil)
}
