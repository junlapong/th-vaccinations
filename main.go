package main

import (
	"encoding/csv"
	"fmt"
	"github.com/dustin/go-humanize"
	"net/http"
	"os"
	"strconv"
)

type Vaccinate struct {
	Location              string
	Date                  string
	Vaccine               string
	SourceURL             string
	TotalVaccinations     string
	PeopleVaccinated      string
	PeopleFullyVaccinated string
	Perday                string
}

const csvURL = "https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/vaccinations/country_data/Thailand.csv"

func main() {

	home := os.Getenv("HOME")
	lines, err := readCsv(home + "/tmp/Thailand.csv")

	if err != nil {
		panic(err)
	}

	var prev int
	prev = 0

	for _, line := range lines {
		total, _ := strconv.Atoi(line[4])
		today := total - prev
		t := humanize.Comma(int64(total))
		p := humanize.Comma(int64(today))

		data := Vaccinate{
			Date:                  line[1],
			TotalVaccinations:     line[4],
			PeopleVaccinated:      line[5],
			PeopleFullyVaccinated: line[6],
			Perday:                p,
		}

		prev = total
		//fmt.Printf("%v\n", data)
		fmt.Printf("Date: %s Total: %10s Perday: %7s\n", data.Date, t, p)
	}

	//etag := getEtag()
	//fmt.Printf("Etag: %s\n", etag)
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
