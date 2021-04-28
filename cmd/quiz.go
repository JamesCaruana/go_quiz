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
	"html"
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

// Created structs to unmarshal json data in from the Open Trivia API.
// Question data struct contains the response code and array of questions data stored in results json tag.
type QuestionsData struct {
	ResponseCode int        `json:"response_code"`
	Results      []Question `json:"results"`
}

// Question struct contains the question, correct answer and incorrect answers
type Question struct {
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

func test() {

	// Getting questions from Open Trivia API w/ error handling
	response, err_0 := http.Get("https://opentdb.com/api.php?amount=10&difficulty=medium&type=multiple")
	if err_0 != nil {
		log.Printf("HTTP request failed - %v", err_0)
	}
	data, _ := ioutil.ReadAll(response.Body)
	log.Println(string(data))
	var q_data QuestionsData

	err_1 := json.Unmarshal(data, &q_data)
	if err_1 != nil {
		log.Println(err_1)
	}
	// :catjam:
	fmt.Printf("%v\n", html.UnescapeString(q_data.Results[1].Question))

}
