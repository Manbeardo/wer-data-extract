package main

import (
	"encoding/xml"
	"io/ioutil"
)

type Person struct {
	ID        string `xml:"id,attr"`
	FirstName string `xml:"first,attr"`
	LastName  string `xml:"last,attr"`
}

type TeamMember struct {
	ID string `xml:"person,attr"`
}

type Team struct {
	EliminationRound int          `xml:"eliminationround,attr"`
	Members          []TeamMember `xml:"member"`
}

type Participation struct {
	People []Person `xml:"person"`
	Teams  []Team   `xml:"team"`
}

type Event struct {
	StartDate     string        `xml:"startdate,attr"`
	EventGUID     string        `xml:"eventguid,attr"`
	Participation Participation `xml:"participation"`
}

func parseEvent(filename string) (Event, error) {
	event := Event{}
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return event, err
	}
	if err := xml.Unmarshal(contents, &event); err != nil {
		return event, err
	}
	return event, nil
}
