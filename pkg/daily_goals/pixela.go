package daily_goals

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type PostPixelData struct {
	Date     time.Time
	Quantity int
}

type postPixelRequest struct {
	Date     string `json:"date"`
	Quantity string `json:"quantity"`
}

type postPixelResponse struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}

func convert(data PostPixelData) postPixelRequest {
	return postPixelRequest{
		Date:     data.Date.Format("20060102"),
		Quantity: strconv.Itoa(data.Quantity),
	}
}

func PostPixel(data PostPixelData, graphID string, credential PixelaCredential) error {
	reqBody, err := json.Marshal(convert(data))
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://pixe.la/v1/users/"+credential.Username+"/graphs/"+graphID, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("X-USER-TOKEN", credential.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result postPixelResponse
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return err
	}

	if !result.IsSuccess {
		return errors.New(result.Message)
	}

	return nil
}
