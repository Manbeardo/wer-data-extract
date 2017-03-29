package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

func main() {
	dirname := ""
	filename := ""
	flag.StringVar(&filename, "file", "", "the file to read")
	flag.StringVar(&dirname, "dir", "", "the directory to read")
	flag.Parse()
	filenames := []string{}
	if dirname != "" {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			log.Fatalln("error reading export directory: ", err)
		}
		for _, file := range files {
			filenames = append(filenames, path.Join(dirname, file.Name()))
		}
	}
	if filename != "" {
		filenames = append(filenames, filename)
	}
	events := []Event{}
	for _, filename := range filenames {
		event, err := parseEvent(filename)
		if err != nil {
			log.Fatalln("error parsing event: ", err)
		}
		events = append(events, event)
	}
	for _, event := range events {
		printEvent(event)
	}
}

func printEvent(event Event) {
	personID2team := map[string]Team{}
	for _, team := range event.Participation.Teams {
		for _, member := range team.Members {
			personID2team[member.ID] = team
		}
	}
	csvWriter := csv.NewWriter(os.Stdout)
	for _, person := range event.Participation.People {
		fullName := fmt.Sprintf("%v %v", person.FirstName, person.LastName)
		eliminationRound := personID2team[person.ID].EliminationRound
		csvWriter.Write([]string{
			event.EventGUID,
			person.ID,
			fullName,
			strconv.Itoa(eliminationRound),
			event.StartDate,
		})
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Fatalln("error writing csv: ", err)
	}
}
