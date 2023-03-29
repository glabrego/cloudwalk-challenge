package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/glabrego/cloudwalk-challenge/lib/parser"
	"github.com/glabrego/cloudwalk-challenge/lib/report"
	"log"
	"os"
	"regexp"
)

func main() {
	file, error := os.Open(os.Args[1])
	if error != nil {
		log.Fatal(error)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matches := []parser.Match{}
	regex, error := regexp.Compile(`\d+:\d+ (\w+): (.*)`)
	if error != nil {
		log.Fatal(error)
		os.Exit(1)
	}
	for scanner.Scan() {
		matched := regex.MatchString(scanner.Text())
		if !matched {
			continue
		}
		gameID := len(matches) - 1
		parseError := parser.ParseLine(gameID, &matches, scanner.Text())
		if parseError != nil {
			log.Fatal(parseError)
			os.Exit(1)
		}
	}

	report := report.ReportMatches(matches)

	reportJson, error := json.Marshal(report)
	if error != nil {
		log.Fatal(error)
		os.Exit(1)
	}

	fmt.Println(string(reportJson))
	os.Exit(0)
}
