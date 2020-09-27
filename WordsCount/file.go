package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	outputfile       = "testoutput.txt"
	pathofFile       = "testfile.txt"
	pathoftestFile   = "testoutputfile.txt"
	pathofOutputfile = "output.txt"
	inputFile        = "mobydick.txt"
)

type keyValue struct {
	Key   string
	Value int
}

func main() {

	// Removing Output File if already Exists
	_ = fileExists(pathofOutputfile)

	// calling scanWords() to read a file word by word
	words, _ := scanWords(inputFile)

	//Calling wordCount Method to get the Count of every word in Converted string from a file
	countword := wordCount(words)

	// Calling sortKeysOnValue sorts the keys on the descending order of its values
	sortKeysOnValue(countword)
}

//exists function removes Output File if already Exists
func fileExists(pathofOutputfile string) error {
	_, err := os.Stat(pathofOutputfile)
	if !os.IsNotExist(err) {
		err := os.Remove(pathofOutputfile)
		if err != nil {
			return err
		}
	}
	return nil
}

// sortKeysOnValue sorts the keys on the descending order of its values
func sortKeysOnValue(countword map[string]int) {
	var ss []keyValue
	var countofRecords int

	for key, value := range countword {
		ss = append(ss, keyValue{key, value})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		countofRecords += 1
		if countofRecords <= 20 {
			_ = writeContentsInFile(kv.Key, kv.Value)
		}
	}
}

//writeContentsInFile writes the 20 top Most frequently used words & their frequency in a output File
func writeContentsInFile(key string, value int) error {

	f, err := os.OpenFile(pathofOutputfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

	fmt.Fprintln(f, key, value)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

// scanWords() Reads file word by word and return whole file as an array of string.
func scanWords(pathofFile string) (string, error) {

	file, err := os.Open(pathofFile)
	if err != nil {
		return " ", err
	}
	defer file.Close()

	//NewScanner returns a new Scanner to read from file.
	scanner := bufio.NewScanner(file)
	//ScanWords function  returns each space-separated word of text, with surrounding spaces deleted
	scanner.Split(bufio.ScanWords)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	completeFileAsaString := strings.Join(words, " ")
	return completeFileAsaString, nil
}

//wordCount() returns the count of all the words along with their frequency of occurence
func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}