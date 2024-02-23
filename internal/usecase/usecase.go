package usecase

import (
	"errors"
	"github.com/okashoi/daily-goals/pkg/daily_goals"
	"time"
)

type Input struct {
	Date       time.Time
	SleepScore int
}

func DoUsecase(input Input, config daily_goals.Config) error {
	if input.SleepScore < 0 || input.SleepScore > 100 {
		return errors.New("sleep score must be between 0 and 100")
	}

	if input.Date.After(time.Now()) {
		return errors.New("date must be before tomorrow")
	}

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
