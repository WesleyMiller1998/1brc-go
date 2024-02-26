package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Info struct {
	City  string
	Total float64
	Count float64
	Min   float64
	Max   float64
}

func main() {
	file, err := os.Open("measurements_10000.txt")
	if err != nil {
		slog.Error("failed to open file", "err", err)
		os.Exit(1)
	}
	defer file.Close()

	cityInfo := make(map[string]*Info)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, ";")

		temp, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			slog.Error("failed to parse integer", "err", err)
			continue
		}

		info, found := cityInfo[parts[0]]
		if !found {
			cityInfo[parts[0]] = &Info{
				City:  parts[0],
				Total: temp,
				Count: 1,
				Min:   temp,
				Max:   temp,
			}

			continue
		}

		info.Total += temp
		info.Count += 1
		if temp > info.Max {
			info.Max = temp
		}
		if temp < info.Min {
			info.Min = temp
		}
	}

	cities := make([]Info, len(cityInfo))

	for _, info := range cityInfo {
		cities = append(cities, *info)
	}

	sort.Slice(cities, func(i, j int) bool {
		return cities[i].City < cities[j].City
	})

	fmt.Printf("{ %s=%.2f/%.2f/%.2f", cities[0].City, cities[0].Min, cities[0].Total/cities[0].Count, cities[0].Max)
	for _, info := range cities {
		fmt.Printf(", %s=%.2f/%.2f/%.2f", info.City, info.Min, info.Total/info.Count, info.Max)
	}
	fmt.Println(" }")
}
