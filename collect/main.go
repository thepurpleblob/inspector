package main

import (
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"resty.dev/v3"
	"strconv"
	"time"
)

var token = ""
var endpoint = ""
var dsn = ""

func main() {
	getconfig()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database", err.Error())
	}
	fmt.Println("Connected to database")

	getloads(db)
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

	token = viper.Get("token").(string)
	endpoint = viper.Get("endpoint").(string)
	dsn = viper.Get("dsn").(string)
}

// Get the server loads with Gorm
func getloads(db *gorm.DB) {
	client := resty.New()
	defer client.Close()

	type LoadItem struct {
		Ip     string
		Load1  string
		Load5  string
		Load15 string
	}

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

	db.AutoMigrate(&Ip{})
	db.AutoMigrate(&Load{})

	var loads []LoadItem

	// Get json list of load averages from web service.
	url := endpoint + "?moodlewsrestformat=json&wstoken=" + token + "&wsfunction=local_loadtest_get_loads"
	res, err := client.R().
		EnableTrace().
		Get(url)
	if err != nil {
		fmt.Printf("Error getting load averages from Moodle, %s", err)
		os.Exit(1)
	}

	// decode json
	bytejson := []byte(res.String())
	jsonerr := json.Unmarshal(bytejson, &loads)
	if jsonerr != nil {
		fmt.Printf("Error parsing load averages from Moodle, %s", err)
		os.Exit(1)
	}

	// Process load data for each front end
	for _, load := range loads {
		fmt.Println(load.Ip, load.Load1, load.Load5, load.Load15)
		var ip Ip
		result := db.Limit(1).Find(&ip, "ip = ?", load.Ip)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Error getting ip, %s", result.Error)
			os.Exit(1)
		}
		if result.RowsAffected == 0 {
			ip.Ip = load.Ip
			db.Create(&ip)
		}
		newload := Load{
			IpId: ip.Id,
		}
		newload.Load1, _ = strconv.ParseFloat(load.Load1, 32)
		newload.Load5, _ = strconv.ParseFloat(load.Load5, 32)
		newload.Load15, _ = strconv.ParseFloat(load.Load15, 32)

		db.Create(&newload)
	}
}
