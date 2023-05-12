package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("gh-pull-request.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	d := json.NewDecoder(f)
	type entry struct {
		Year    int
		Quarter int
		Name    string
		Count   int
	}
	var entries []entry
	err = d.Decode(&entries)
	if err != nil {
		log.Fatal(err)
	}

	of, err := os.Create("gh-pull-request.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()
	e := csv.NewWriter(of)
	e.Write([]string{"Year", "Quarter", "Name", "Count"})
	defer e.Flush()
	for _, entry := range entries {
		e.Write([]string{
			strconv.Itoa(entry.Year),
			strconv.Itoa(entry.Quarter),
			entry.Name,
			strconv.Itoa(entry.Count),
		})
	}
}
