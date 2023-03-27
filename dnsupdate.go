package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/joho/godotenv"
	cron "github.com/robfig/cron/v3"
)

var interval string
var client *http.Client

func main() {
	c := cron.New()

	if interval == "" {
		interval = "*/59 * * * *"
	}
	//Run once on startup
	updateIP()
	//Run cronjob moving forward
	c.AddFunc(interval, updateIP)
	c.Start()
	log.Println("=====cron system started======")

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
}

func updateIP() {

	err := godotenv.Load()
	if err != nil {
		log.Print("no env file found ")
	}

	client = &http.Client{
		Timeout: 5 * time.Second,
	}

	domain := os.Getenv("DOMAIN")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	log.Printf("Username: %s Password: %s", username, password)
	ip := getip()
	if ip == "" {
		log.Println("IP address could not be found skipping update")
		return
	}
	if username == "" || password == "" {
		log.Println("username or password is missing skipping update")
		return
	}

	values := url.Values{}
	values.Set("hostname", domain)
	values.Set("myip", ip)

	URL := url.URL{
		Scheme: "https",
		Host:   "domains.google.com",
		Path:   "nic/update",
	}

	URL.RawQuery = values.Encode()
	log.Printf("URL being used %s", URL.String())

	req, err := http.NewRequest(http.MethodPost, URL.String(), nil)
	if err != nil {
		log.Println(err)
	}

	req.SetBasicAuth(username, password)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Error updating ip from Google %v \n", res.StatusCode)
		} else {
			log.Printf("Success updating ip from Google %v \n", string(b))
		}

	} else {
		log.Printf("Error updating ip from Google %v \n", res.StatusCode)
	}

}

func getip() string {
	req, err := http.NewRequest(http.MethodGet, "http://checkip.amazonaws.com/", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return strings.Replace(string(ip), "\n", "", 1)
}
