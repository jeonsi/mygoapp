// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for index := range lines {
		fmt.Println(index)
	}
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}

	countsMap := make(map[string]int)
	for _, line := range lines {
		countsMap[line]++
	}
	for name, count := range countsMap {
		fmt.Printf("Votes for %v: %v\n", name, count)
	}
}
