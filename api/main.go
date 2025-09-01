package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Ip struct {
	Id        int64
	Ip        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Load struct {
	Id        int64
	IpId      int64
	Load1     float64
	Load5     float64
	Load15    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

var dsn = ""

var db *gorm.DB

func main() {
	getconfig()
	dbconnect()

	// Mysterious CORS shit.
	corsObj := handlers.AllowedOrigins([]string{"*"})

	r := mux.NewRouter()
	r.HandleFunc("/getips", IpsHandler)
	r.HandleFunc("/getlatestloads", LatestHandler)
	r.HandleFunc("/getloads/{starttimestamp}/{finishtimestamp}", LoadsHandler)
	log.Fatal(http.ListenAndServe(":8900", handlers.CORS(corsObj)(r)))
}

// LoadsHandler Handler for Loads
func LoadsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	starttimestamp, _ := strconv.ParseInt(vars["starttimestamp"], 10, 64)
	endtimestamp, _ := strconv.ParseInt(vars["finishtimestamp"], 10, 64)

	loadsjson := getloads(time.Unix(starttimestamp, 0), time.Unix(endtimestamp, 0))

	w.WriteHeader(http.StatusOK)
	w.Write(loadsjson)
}

// LatestHandler Handler for latest Loads
func LatestHandler(w http.ResponseWriter, r *http.Request) {

	loadsjson := getlatestloads()

	w.WriteHeader(http.StatusOK)
	w.Write(loadsjson)
}

// Get list of IP addresses from database
func IpsHandler(w http.ResponseWriter, r *http.Request) {

	ips := getIps()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ips)
}

// Get the configuration from config.json
func getconfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
		os.Exit(1)
	}

	dsn = viper.Get("dsn").(string)
}

// Connect to the database
func dbconnect() {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database", err.Error())
	}

	db.AutoMigrate(&Ip{})
	db.AutoMigrate(&Load{})
}

// Get the load data
// Read all load data between specified time stamps
// and return as json
func getloads(start time.Time, end time.Time) []byte {

	var loads []Load
	result := db.Where("created_at > ? AND created_at < ?", start, end).Find(&loads)
	if result.Error != nil {
		fmt.Println("Error getting loads from database", result.Error)
		os.Exit(1)
	}
	json, err := json.Marshal(loads)
	if err != nil {
		fmt.Println("Error converting loads to JSON", err.Error())
		os.Exit(1)
	}

	return json
}

// Get the latest load data for each FE
// Get the latest load record for each FE
// Return array as json
func getlatestloads() []byte {

	// Get list of ips
	var ips []Ip
	result := db.Find(&ips)
	if result.Error != nil {
		fmt.Println("Error getting ips from database", result.Error)
		os.Exit(1)
	}

	// Iterate over ips getting latest load record for each
	var loads []Load
	for _, ip := range ips {
		ipid := ip.Id
		var load Load
		result := db.Where("id = (?)", db.Table("loads").Select("max(id)").Where("ip_id = ?", ipid)).First(&load)
		if result.Error != nil {
			fmt.Println("Error getting load from database", result.Error)
			os.Exit(1)
		}
		loads = append(loads, load)
	}

	json, err := json.Marshal(loads)
	if err != nil {
		fmt.Println("Error converting loads to JSON", err.Error())
		os.Exit(1)
	}

	return json
}

// Get the list of IPs
func getIps() []byte {
	var ips []Ip

	result := db.Find(&ips)
	if result.Error != nil {
		fmt.Println("Error getting ips from database", result.Error)
		os.Exit(1)
	}
	json, err := json.Marshal(ips)
	if err != nil {
		fmt.Println("Error converting ips to JSON", err.Error())
		os.Exit(1)
	}

	return json
}
