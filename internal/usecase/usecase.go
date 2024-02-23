package usecase

import (
	"github.com/okashoi/daily-goals/pkg/daily_goals"
	"time"
)

type Input struct {
	Date       time.Time
	SleepScore int
}

func DoUsecase(input Input, config daily_goals.Config) error {
	data := daily_goals.PostPixelData{
		Date:     input.Date,
		Quantity: input.SleepScore,
	}

	err := daily_goals.PostPixel(data, "sleep-score", config.PixelaCredential)
	if err != nil {
		return err
	}

	return nil
}
