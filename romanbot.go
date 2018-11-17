// Copyright (c) 2018 Robert Clarke.

// RomanBot 0.2
// Tweets today's Roman Date in DMY and MDY format
// Based on BuoyBot https://github.com/johnbeil/BuoyBot

package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/StefanSchroeder/Golang-Roman"
	"github.com/dustin/go-humanize"
	"log"
	"os"
	"time"
)

// RomanDate struct stores a date and converted day, month, and year
type RomanDate struct {
	Date  time.Time
	Day   string
	Month string
	Year  string
}

// Config struct stores Twitter credentials
type Config struct {
	UserName       string `json:"UserName"`
	ConsumerKey    string `json:"ConsumerKey"`
	ConsumerSecret string `json:"ConsumerSecret"`
	Token          string `json:"Token"`
	TokenSecret    string `json:"TokenSecret"`
}

func main() {
	fmt.Println("Starting RomanBot...")

	// Load configuration
	config := Config{}
	loadConfig(&config)

	// Get today's date
	romanDate := getRomanDate()

	// Format date
	formattedDate := formatRomanDate(romanDate)

	tweetFormattedDate(config, formattedDate)

	// Shutdown RomanBot
	fmt.Println("Exiting RomanBot...")
}

// Given config and tweet text, tweets latest update
func tweetFormattedDate(config Config, date string) {
	fmt.Println("Preparing to tweet formatted date...")
	api := anaconda.NewTwitterApiWithCredentials(config.Token, config.TokenSecret, config.ConsumerKey, config.ConsumerSecret)
	tweet, err := api.PostTweet(date, nil)
	if err != nil {
		fmt.Println("update error:", err)
	} else {
		fmt.Println("Tweet posted:")
		fmt.Println(tweet.Text)
	}

	//fmt.Println("Should tweet:", date)
}

// Given path to config.js file, loads credentials
func loadConfig(config *Config) {
	// Load path to config from CONFIGPATH environment variable
	configpath := os.Getenv("CONFIGPATH")
	fmt.Println("Loading config.json from:", configpath)
	file, _ := os.Open(configpath)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error loading config.json:", err)
	}
}

// Given a date and the date components in roman numerals, returns formatted text for tweet
func formatRomanDate(rd RomanDate) string {
	dmyDate := humanize.Ordinal(rd.Date.Day()) + " " + rd.Date.Format("January 2006")
	mdyDate := rd.Date.Format("January 2 2006")

	output := fmt.Sprintf("%s (%s.%s.%s) / %s (%s.%s.%s)",
		dmyDate, rd.Day, rd.Month, rd.Year,
		mdyDate, rd.Month, rd.Day, rd.Year)

	return output
}

// getRomanDate converts the components of today's date into roman numerals and returns a RomanDate struct
func getRomanDate() RomanDate {
	var romanDate RomanDate

	romanDate.Date = time.Now()

	year, month, day := romanDate.Date.Date()

	romanDate.Day = roman.Roman(day)
	romanDate.Month = roman.Roman(int(month))
	romanDate.Year = roman.Roman(year)

	return romanDate
}
