package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*
	Define Movie Data structure
*/
type Movie struct {
	Ratings []Ratings `json:"Ratings"`
}

/*
	Define Movie Ratings Data structure
*/
type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

// json data type
var obj Movie

/*
	inputArg() function Validate the input parameters and
	return movie name.
*/
func inputArg() string {
	var inputName string = "Avatar"
	if len(os.Args) == 2 {
		inputName = os.Args[1]
	}
	return inputName
}

/*
	GetAPIKey() function Read the API key from movieAPIKey file and
	return key value
*/
func GetAPIKey() string {
	//file, err := os.Open("/Users/sichekur/.movieAPIKey")
	//file, err := os.Open("C:\\Users\\sichekur\\go\\src\\github.com\\test\\assignment\\.movieAPIKey")
	file, err := os.Open("/SRC/.movieAPIKey")
	//file, err := os.Open(".movieAPIKey")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return string(b)
}

// Return the movie url
func fetchURLInfo(name string) string {
	apiKey := GetAPIKey()
	req, err := http.NewRequest("GET", "http://www.omdbapi.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	q := req.URL.Query()
	q.Add("t", name)
	q.Add("apikey", strings.TrimSpace(apiKey))
	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}

// return Movie info, it's json format
func GetMovieInfo(name string) string {
	movieUrl := fetchURLInfo(name)
	resp, err := http.Get(movieUrl)
	//fmt.Println("Movie Rating Source: " + movieUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

/*
	Validate Movie ratings, consider Rotten Tomatoes rating and,
	ignore other movie rating sources
*/
func IsValidSource(source string) bool {
	switch source {
	case
		"Rotten Tomatoes":
		return true
	}
	return false
}

func main() {
	movieName := inputArg()
	rating := GetMovieInfo(movieName)
	json.Unmarshal([]byte(rating), &obj)

	for i := 0; i < len(obj.Ratings); i++ {

		if IsValidSource(obj.Ratings[i].Source) {
			fmt.Println("Movie Rating Source: " + obj.Ratings[i].Source)
			fmt.Println("\nMovie Rating: " + obj.Ratings[i].Value)
		}
	}

}
