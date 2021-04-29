package cmd

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
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

		startQuiz()

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

// Global Variables
var q_data QuestionsData
var questions [10]string
var answers [10][4]string

/*
	Menu function
*/
func menu() {

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
	//log.Println(string(data))

	err_1 := json.Unmarshal(data, &q_data)
	if err_1 != nil {
		log.Println(err_1)
	}

	return &q_data

}

/*
	Function which organises quiz data into arrays in order print them easily in
	the quiz.
*/
func organiseData(q *QuestionsData) ([10]string, [10][4]string) {

	// Outer loop is used to organise quiz questions
	for i := 0; i < len(q.Results); i++ {
		questions[i] = html.UnescapeString(q.Results[i].Question)
		// Inner loop is used to organise quiz possible answers
		for j := 0; j < len(q.Results[i].IncorrectAnswers); j++ {
			answers[i][0] = html.UnescapeString(q.Results[i].CorrectAnswer)
			answers[i][j+1] = html.UnescapeString(q.Results[i].IncorrectAnswers[j])
		}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(4, func(k, l int) { answers[i][k], answers[i][l] = answers[i][l], answers[i][k] })
	}

	//fmt.Printf("Question: %v\nPossible Answers: %v, %v, %v, %v\n", questions[9], answers[9][0], answers[9][1], answers[9][2], answers[9][3])

	return questions, answers
}

/*
	Starts quiz by looping through questions.
*/
func startQuiz() {

	fmt.Println("------------------------Starting Quiz------------------------")

	getQuizData()
	questions, answers := organiseData(&q_data)
	var c_answers, inc_answers int = 0, 0

	// Loop through questions and prints them
	for i := 0; i < len(questions); i++ {
		fmt.Printf("\nQuestion %v\n%v\n", i+1, questions[i])
		// Loops through possible answers and prints them
		for k := 0; k < len(answers[i]); k++ {
			fmt.Printf("%v) %v\n", k+1, answers[i][k])
		}

		var choice int
		fmt.Println("Input a number between 1-4 to answer1.")
		fmt.Scanln(&choice)

		choice = checkRange(choice)

		var actual_answer string = answers[i][choice-1]
		// Call function to check answers and returns the updated score
		c_answers, inc_answers = checkAnswer(actual_answer, i, c_answers, inc_answers)

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

	}
	fmt.Println("------------------------Quiz Finished------------------------")
	fmt.Printf("You got %v/10 correct answers\n", c_answers)
}

/*
	Error checking function for the user input on selecting answer.
*/
func checkRange(choice int) int {
	if (choice >= 1) && (choice <= 4) {
		fmt.Print(choice)
		return choice
	} else {
		fmt.Println("Input a number between 1-4 to answer2.")
		fmt.Scanln(&choice)
		return checkRange(choice)
	}
}

/*
	Function which checks answer and returns the number of correct and incorrect answers.
*/
func checkAnswer(choice string, question_number int, c_answers int, inc_answers int) (int, int) {
	if choice == html.UnescapeString(q_data.Results[question_number].CorrectAnswer) {
		return c_answers + 1, inc_answers
	} else {
		return c_answers, inc_answers + 1
	}
}
