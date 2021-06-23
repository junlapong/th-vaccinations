package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
)

type Vaccinate struct {
	Location              string
	Date                  string
	Vaccine               string
	SourceURL             string
	TotalVaccinations     string
	PeopleVaccinated      string
	PeopleFullyVaccinated string
}

const csvURL = "https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/vaccinations/country_data/Thailand.csv"

func main() {

	lines, err := readCsv("./data/Thailand.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines turn into object
	for _, line := range lines {
		data := Vaccinate{
			Date:                  line[1],
			TotalVaccinations:     line[4],
			PeopleVaccinated:      line[5],
			PeopleFullyVaccinated: line[6],
		}
		fmt.Printf("%v\n", data)
	}

	etag := getEtag()
	fmt.Printf("Etag: %s\n", etag)
}

func readCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func getEtag() (etag string) {

	res, err := http.Head(csvURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Content-Length: %d\n", res.ContentLength)

	etag = res.Header.Get("Etag")
	return
}
