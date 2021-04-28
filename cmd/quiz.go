package cmd

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

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

		var q_data = getQuizData()
		organiseData(q_data)

	},
}

func init() {
	rootCmd.AddCommand(quizCmd)

}

/*
	Created structs to unmarshal json data in from the Open Trivia API.
	Question data struct contains the response code and array of questions data stored in results json tag.
*/
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

/*
	Function which gets a random set of 10 questions from the Open Trivia API and
	returns a structure of the data.
*/
func getQuizData() *QuestionsData {

	// Getting questions from Open Trivia API w/ error handling
	response, err_0 := http.Get("https://opentdb.com/api.php?amount=10&category=15&difficulty=medium&type=multiple")
	if err_0 != nil {
		log.Printf("HTTP request failed - %v", err_0)
	}
	data, _ := ioutil.ReadAll(response.Body)

	// Print the result of API call
	log.Println(string(data))

	var q_data QuestionsData

	err_1 := json.Unmarshal(data, &q_data)
	if err_1 != nil {
		log.Println(err_1)
	}
	// :catjam:
	//fmt.Printf("%v\n", html.UnescapeString(q_data.Results[0].Question))

	return &q_data

}

/*
	Function which organises quiz data into arrays in order print them easily in
	the quiz.
*/
func organiseData(q *QuestionsData) ([10]string, [10][4]string) {

	// Testing prints
	fmt.Printf("Length of Array:\t%v\n", len(q.Results))
	fmt.Println(html.UnescapeString(q.Results[9].Question))

	var questions [10]string
	var answers [10][4]string

	// Outer loop is used to organise quiz questions
	for i := 0; i < len(q.Results); i++ {
		questions[i] = html.UnescapeString(q.Results[i].Question)
		// Inner loop is used to organise quiz possible answers
		for k := 0; k < len(q.Results[i].IncorrectAnswers); k++ {
			answers[i][0] = html.UnescapeString(q.Results[i].CorrectAnswer)
			answers[i][k+1] = html.UnescapeString(q.Results[i].IncorrectAnswers[k])
		}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(4, func(h, j int) { answers[i][h], answers[i][j] = answers[i][j], answers[i][h] })
	}

	fmt.Printf("Question: %v\nPossible Answers: %v, %v, %v, %v", questions[9], answers[9][0], answers[9][1], answers[9][2], answers[9][3])

	return questions, answers
}
