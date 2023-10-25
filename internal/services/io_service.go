package services

import (
	"context"
	"effective-mobile/internal/domain/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type IoSvc struct {
}

func NewIoSvc() *IoSvc {
	return &IoSvc{}
}

const agifyURL = "https://api.agify.io"
const genderizeURL = "https://api.genderize.io"
const nationalizeURL = "https://api.nationalize.io"

// age
type agifyResponseBody struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

// gender
type genderizeRespBody struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

// nationality
type nationalizeRespBody struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

type Country struct {
	CountryId   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type RespBodies interface {
	agifyResponseBody | genderizeRespBody | nationalizeRespBody
}

const paramKey = "name"

func sendRequest[B RespBodies](name string, apiURL string, wg *sync.WaitGroup, errChan chan error, respBody *B) {

	defer wg.Done()

	params := url.Values{}
	params.Add(paramKey, name)
	url, err := url.Parse(apiURL)
	if err != nil {
		errChan <- err
		return
	}
	url.RawQuery = params.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		fmt.Println("ERR")
		errChan <- err
		return
	}

	fmt.Println(respBody)

}

// количество горутин
const calls = 3

func (ios *IoSvc) FillAddPersonRequest(ctx context.Context, req dto.AddPersonRequest) (dto.AddPersonRequest, error) {
	wg := sync.WaitGroup{}
	wg.Add(calls)
	var ageResp agifyResponseBody
	var genderResp genderizeRespBody
	var nationalityResp nationalizeRespBody

	errChan := make(chan error, 5)
	go sendRequest[agifyResponseBody](req.Name, agifyURL, &wg, errChan, &ageResp)
	go sendRequest[genderizeRespBody](req.Name, genderizeURL, &wg, errChan, &genderResp)
	go sendRequest[nationalizeRespBody](req.Name, nationalizeURL, &wg, errChan, &nationalityResp)

	wg.Wait()

	close(errChan)
	if err, ok := <-errChan; ok {
		return dto.AddPersonRequest{}, err
	}

	req.Age = ageResp.Age
	req.Gender = genderResp.Gender
	req.Nationality = nationalityResp.Country[0].CountryId

	return req, nil
}
