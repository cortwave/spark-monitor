package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var appsCount int
var host string
var publisherURL string
var channel string
var period int
var longExecution int64

type stage struct {
	Status          string `json:"status"`
	ExecutorRunTime int64  `json:"executorRunTime"`
}

type app struct {
	ID string `json:"id"`
}

func main() {
	initVars()
	for _ = range time.Tick(time.Duration(period) * time.Second) {
		checkApps()
	}
}

func initVars() {
	appsCount = initInt("APP_COUNT")
	host = os.Getenv("DRIVER_ADDRESS")
	publisherAddress := initString("PUBLISHER_ADDRESS")
	publisherURL = "http://" + publisherAddress + "/publish"
	channel = os.Getenv("CHANNEL")
	period = initInt("PERIOD")
	longExecution = int64(initInt("LONG_EXECUTION")) * 1000
}

func initString(name string) string {
	val := os.Getenv(name)
	if val == "" {
		panic("No publisher provided")
	}
	return val
}

func initInt(name string) int {
	val, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		panic(err)
	}
	return val
}

func checkApps() {
	apps := []app{}
	err := get("http://"+host+"/api/v1/applications?status=running", &apps)
	if err != nil {
		send("Can't get applications list")
	} else {
		if len(apps) < appsCount {
			send("Expected " + strconv.Itoa(appsCount) + " apps, actual " + strconv.Itoa(len(apps)) + " apps")
		}
		for _, app := range apps {
			checkApp(app.ID)
		}
	}
}

func checkApp(appID string) {
	stages := []stage{}
	err := get("http://"+host+"/api/v1/applications/"+appID+"/stages", &stages)
	if err != nil {
		send("Can't get info about app: " + appID)
	} else {
		for _, stage := range stages {
			if stage.ExecutorRunTime > longExecution {
				send("Long executing stage in app " + appID)
			}
		}
	}
}

func send(text string) {
	textToSend := time.Now().Format(time.Kitchen) + ": " + text
	sendNotification(textToSend, publisherURL)
}

func get(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}
