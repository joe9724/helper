package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Sign struct {
	Id         int     `json:"id"`
	WorkerId   string  `json:"workerId"`
	WorkerName string  `json:"workerName"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Location   string  `json:"location"`
	WorkTime   string  `json:"workTime"`
}

type Collection struct {
	Id           int     `json:"id"`
	DispeopleId  string  `json:"dispeopleId"`
	AccessRecord string  `json:"accessRecord"`
	Photos       string  `json:"photos"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Location     string  `json:"location"`
	AccessTime   string  `json:"accessTime"`
}

type Privates struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Scope    int    `json:"scope"`
	UserId   int    `json:"userId"`
}

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:root811123@tcp(106.14.2.153:3306)/help")
	db.SetMaxOpenConns(20000)
	db.SetMaxIdleConns(10000)
	db.Ping()
	fmt.Println("server start success...")
}

func main() {
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/sign", _Sign)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func _Sign(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("select id,workerId,workerName,latitude,longitude,location,workTime from sign")
	checkErr(err)
	var Arr []Sign
	for rows.Next() {
		var id int
		var workerId string
		var workerName string
		var latitude float64
		var longitude float64
		var location string
		var workTime string

		err = rows.Scan(&id, &workerId, &workerName, &latitude, &longitude, &location, &workTime)
		checkErr(err)

		var __sign Sign
		__sign.Id = id
		__sign.WorkerId = workerId
		__sign.WorkerName = workerName
		__sign.Latitude = latitude
		__sign.Longitude = longitude
		__sign.Location = location
		__sign.WorkerName = workTime

		Arr = append(Arr, __sign)

	}

	//
	data, err := json.Marshal(Arr)
	if err != nil {
		log.Fatal("err get data: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "NanjingYouzi")
	defer func() {
		io.WriteString(w, string(data))
		//fmt.Print(string(data))
		//w.Write(data)
	}()
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
