package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Info struct {
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

	for city, info := range cityInfo {
		fmt.Printf("%s: %v\n", city, info)
	}
}
