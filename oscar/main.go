package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./oscar_age_male.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	nameCount := map[string]int{}

	for _, record := range records {
		nameCount[record[3]]++
	}

	maxVal := 1
	for _, v := range nameCount {
		if ( v > maxVal ) {
			maxVal = v
		}
	}
	for k, v := range nameCount {
		if( v == maxVal ) {
			fmt.Println(k)
		}
	}
	
}
