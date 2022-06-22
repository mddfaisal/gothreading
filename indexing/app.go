package main

import (
	"bytes"
	"fmt"
	"indexing/db"
	"indexing/models/indexdata"
	"indexing/utils"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var (
	nextUsersId int64
	records     = 1000
	url         = "http://localhost:9200/fakerdata/_doc/"
)

func getUsersId() int64 {
	var uid int64
	query := `select id from users where id > ` + fmt.Sprintf("%d", nextUsersId) + ` limit 1`
	result := db.Db().QueryRow(query)
	result.Scan(&uid)
	return uid
}

func getIndexingData(i int64) string {
	index := indexdata.NewIndexingData(i)
	index.FindById()
	return index.AsJson()
}

func doIndexing(data string) {
	req, err := http.NewRequest(http.MethodPost, url+fmt.Sprintf("%v", nextUsersId), bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	utils.Panic(err)
	client := http.Client{}
	resp, err := client.Do(req)
	utils.Panic(err)
	fmt.Println(resp.Status, resp.StatusCode)
	err = resp.Body.Close()
	utils.Panic(err)
}

func main() {
	utils.Panic(godotenv.Load())
	now := time.Now()
	for i := 0; i <= records; i++ {
		// get user id
		nextUsersId = getUsersId()
		// fetch data
		indexingData := getIndexingData(nextUsersId)
		// index data
		doIndexing(indexingData)
	}
	fmt.Println("Elapced:", time.Since(now))
}
