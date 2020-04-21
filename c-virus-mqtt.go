package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	//	For working correctly please use first:
	//	go get github.com/eclipse/paho.mqtt.golang
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Country struct {
	Country     string `json:"country"`
	Cases       int64  `json:"cases"`
	TodayCases  int    `json:"todayCases"`
	Deaths      int    `json:"deaths"`
	TodayDeaths int    `json:"todayDeaths"`
	Recovered   int    `json:"recovered"`
	Active      int    `json:"active"`
	Critical    int    `json:"critical"`
	Updated     int64  `json:"updated"`
}

func readJSONFromUrl(url string) ([]Country, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var countryList []Country
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &countryList); err != nil {
		return nil, err
	}

	return countryList, nil
}

func getFormattedTimeFromEpochMillis(z int64, zone string) string {

	unixTime := time.Unix(z/1000, 0)
	timeZoneLocation, err := time.LoadLocation(zone)
	if err != nil {
		fmt.Println("Error loading timezone:", err)
	}

	timeInZone := unixTime.In(timeZoneLocation)
	timeInZoneStyleOne := timeInZone.Format("2006-02-01 15:04:05")

	return timeInZoneStyleOne
}

func main() {
	url := "https://corona.lmao.ninja/v2/countries"

	broker := flag.String("broker", "tcp://localhost:1883", "The broker URI. ex: tcp://localhost:1883")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	id := flag.String("id", "CV-Stats", "The ClientID (optional)")
	topic1 := flag.String("topic", "/coronavirus", "Topics start at (optional)")
	forCountry1 := flag.String("country", "Ukraine", "For witch country (optional)")
	timez1 := flag.String("timezone", "Europe/Kiev", "Timezone for updated date (optional)")
	flag.Parse()

	topic := *topic1
	forCountry := *forCountry1
	timez := *timez1

	opts := MQTT.NewClientOptions()
	opts.AddBroker(*broker)
	opts.SetClientID(*id)
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetCleanSession(false)

	countryList, err := readJSONFromUrl(url)
	if err != nil {
		panic(err)
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, row := range countryList {
		if row.Country == forCountry {
			cases := client.Publish(topic+"/"+forCountry+"/cases", 0, false, strconv.FormatInt(int64(row.Cases), 10))
			cases.Wait()
			todayCases := client.Publish(topic+"/"+forCountry+"/todayCases", 0, false, strconv.FormatInt(int64(row.TodayCases), 10))
			todayCases.Wait()
			deaths := client.Publish(topic+"/"+forCountry+"/deaths", 0, false, strconv.FormatInt(int64(row.Deaths), 10))
			deaths.Wait()
			todayDeaths := client.Publish(topic+"/"+forCountry+"/todayDeaths", 0, false, strconv.FormatInt(int64(row.TodayDeaths), 10))
			todayDeaths.Wait()
			recovered := client.Publish(topic+"/"+forCountry+"/recovered", 0, false, strconv.FormatInt(int64(row.Recovered), 10))
			recovered.Wait()
			updated := client.Publish(topic+"/"+forCountry+"/updated", 0, false, getFormattedTimeFromEpochMillis(row.Updated, timez))
			updated.Wait()
		}
	}
	client.Disconnect(250)
}
