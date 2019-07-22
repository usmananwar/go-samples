package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestQuiz1(t *testing.T) {

	filename := filepath.Clean("C:/Users/Usman/go/src/usman.com/go-samples/gophercises/ex1/problems.csv")
	fmt.Println(filename)

	quiz1(filename)


}