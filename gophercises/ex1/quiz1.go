package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)


type Row struct {
	question string
	answer string
}


func readCsvFile(filename string) ([]Row,error) {
	var rows []Row
	file,err := os.Open(filename);
	defer file.Close()
	if (err!= nil) {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	for {
		line, err:= reader.Read()
		if (err==io.EOF) {
			break;
		} else if (err!=nil) {
			return nil, err
		}

		row:= Row{
			question:line[0],
			answer:line[1],
		}
		rows=append(rows, row)
	}

	return rows, nil
}


func main() {

	filename := filepath.Clean("C:/Users/Usman/go/src/usman.com/go-samples/gophercises/ex1/problems.csv")
	fmt.Println(filename)
	quiz1(filename)
}



func quiz1(filename string) []bool {
	rows, err := readCsvFile(filename)
	if (err != nil) {
		log.Fatalln(err)
	}

	var result []bool;
	for i:=0;i<5;i++ {

		startTime := time.Now()

		randomQuestionIndex := rand.Intn(len(rows))
		row := rows[randomQuestionIndex];
		fmt.Printf("Type the anserwer for %s ", row.question)
		var input string
		fmt.Scan(&input)

		takenTime := time.Since(startTime)
		if input == row.answer {
			result = append(result,true)
			fmt.Printf("Correct answer :) It took:  %s \n", takenTime)
		} else {
			result = append(result,false)
			fmt.Printf("Wrong answer!!! It took:  %s \n", takenTime)
		}
	}


	return result;
}





