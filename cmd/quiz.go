/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A simple quiz",
	Long: `Golang quiz which requires to fill the following requirements:
	- Create a multiple choice quiz (get data from an quiz API).
	- User can select one answer only.
	- Display how many correct answers user got.
	- Show some statistics.`,
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func init() {
	rootCmd.AddCommand(quizCmd)

}

type Payload struct {
	Results string `json:"results"`
	// Results      Results `json:"results"`
}

// type Results struct {
// 	Question string `json:"question"`
// }

func test() {

	// Getting questions from Open Trivia API w/ error handling
	response, err := http.Get("https://opentdb.com/api.php?amount=10&difficulty=medium&type=multiple")
	if err != nil {
		log.Printf("HTTP request failed - %v", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		log.Println(string(data))

		// Unstructured data approach
		// var reading map[string]interface{}
		// err = json.Unmarshal([]byte(data), &reading)
		// fmt.Printf("%+v\n", reading)

		var questions []Payload
		// var questions []Payload
		err := json.Unmarshal(data, &questions)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v\n", questions)

	}
	// Need to unmarshall data

}
