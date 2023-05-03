package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/stulepbergen/API-Client-for-Codeforces/models"
)

var methods = []string{"user.info", "user.rating", "problemset.recentStatus"}

func main() {
	fmt.Print("Choose method name: 0(user.info), 1(user.rating), 2(problemset.recentStatus) -> ")

	var methodInd int
	fmt.Scan(&methodInd)

	var parameter string

	var count int
	var handle string
	if methodInd == 2 {
		fmt.Print("Enter count: ")
		fmt.Scan(&count)
		parameter = "count=" + strconv.Itoa(count)
	} else {
		fmt.Print("Enter handle: ")
		fmt.Scan(&handle)
		if methodInd == 1 {
			parameter = "handle="
		} else {
			parameter = "handles="
		}
		parameter += handle
	}

	client := &http.Client{}

	resp, err := client.Get("https://codeforces.com/api/" + methods[methodInd] + "?" + parameter)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if methodInd == 0 {
		UserInfo(body)
	} else if methodInd == 1 {
		UserRating(body)
	} else {
		ProblemsetRecentStatus(body)
	}

}

func UserInfo(body []byte) {
	var UI models.UsersInfo
	err := json.Unmarshal(body, &UI)
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range UI.Result {
		fmt.Println(info)
	}
}

func UserRating(body []byte) {
	var UR models.UsersRating
	err := json.Unmarshal(body, &UR)
	if err != nil {
		log.Fatal(err)
	}

	for _, ratings := range UR.Result {
		fmt.Println(ratings)
	}
}

func ProblemsetRecentStatus(body []byte) {
	var PS models.ProblemSet
	err := json.Unmarshal(body, &PS)
	if err != nil {
		log.Fatal(err)
	}

	for _, problems := range PS.Result {
		fmt.Println(problems)
	}
}
