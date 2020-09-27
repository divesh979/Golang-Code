package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func TestFileExists(t *testing.T) {
	err := fileExists(outputfile)
	if err != nil {
		t.Error()
	}
}

func TestScanWords(t *testing.T) {
	expectedwords := "My name is Divesh I work as a Software Engineer in Noida I love Playing Cricket I love watching Cricket"
	actualwords, err := scanWords(pathofFile)
	if err != nil {
		t.Error()
	}
	if actualwords != expectedwords {
		t.Errorf("Expected Words and actual Words didt match. Expected: %s. Acutal: %s", expectedwords, actualwords)
	}
}

func TestWordCount(t *testing.T) {
	words := "My name is Divesh I work as a Software Engineer in Noida I love Playing Cricket I love watching Cricket"
	var expectedCount = map[string]int{"Cricket": 2, "Divesh": 1, "Engineer": 1, "I": 3, "My": 1, "Noida": 1, "Playing": 1, "Software": 1, "a": 1, "as": 1,
		"in": 1, "is": 1, "love": 2, "name": 1, "watching": 1, "work": 1}
	countWords := wordCount(words)
	if !reflect.DeepEqual(countWords, expectedCount) {
		t.Errorf("Invalid listed files. Expected: %v. Actual: %v", expectedCount, countWords)
	}
}

func TestSortKeysonValue(t *testing.T) {

	var countWord = map[string]int{"Cricket": 2, "Divesh": 1, "Engineer": 1, "I": 3, "My": 1, "Noida": 1, "Playing": 1, "Software": 1, "a": 1, "as": 1,
		"in": 1, "is": 1, "love": 2, "name": 1, "watching": 1, "work": 1}
	sortKeysOnValue(countWord)

	expectedOutputFileContents, err := ioutil.ReadFile(pathoftestFile)

	if err != nil {
		log.Fatal(err)
	}

	actualOutputFileContents, err1 := ioutil.ReadFile(pathofOutputfile)

	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(bytes.Equal(expectedOutputFileContents, actualOutputFileContents))

}
