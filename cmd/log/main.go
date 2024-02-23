package main

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/okashoi/daily-goals/internal/usecase"
	"github.com/okashoi/daily-goals/pkg/daily_goals"
	"log"
	"os"
	"strconv"
	"time"
)

var expectedLayouts = []string{
	"20060102",
	"2006-01-02",
	"2006/01/02",
}

func parseDate(s string) (time.Time, error) {
	var t time.Time

	for _, layout := range expectedLayouts {
		t, err := time.Parse(layout, s)
		if err != nil {
			continue
		}
		return t, nil
	}

	return t, errors.New("invalid date format: supported layouts are 20240101, 2024/01/01 and 2024-01-01")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) != 3 {
		log.Fatal("usage: log <date> <sleep score>")
	}

	date, err := parseDate(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	score, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	input := usecase.Input{
		Date:       date,
		SleepScore: score,
	}

	config := daily_goals.Config{
		PixelaCredential: daily_goals.PixelaCredential{
			Token:    os.Getenv("PIXELA_TOKEN"),
			Username: os.Getenv("PIXELA_USERNAME"),
		},
	}

	err = usecase.DoUsecase(input, config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("succeeded")
}
