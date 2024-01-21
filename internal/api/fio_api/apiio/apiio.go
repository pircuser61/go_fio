package apiio

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/pircuser61/go_fio/internal/api/fio_api"
)

type ApiIo struct {
	log *slog.Logger
}

func GetApi(logger *slog.Logger) fio_api.Api {
	return ApiIo{log: logger}
}

func (i ApiIo) GetGender(name string) (string, error) {
	base, err := url.Parse("https://api.genderize.io/")
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("name", name)
	base.RawQuery = params.Encode()
	urlStr := base.String()
	i.log.Debug("api: GetGender", slog.String("name", name), slog.String("url", urlStr))

	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GetGender http.Status %d", resp.StatusCode)
	}

	type genderResponse struct {
		Count       int
		Name        string
		Gender      string
		Probability float32
	}
	var data genderResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	i.log.Debug("api: GetGender", slog.Any("response", data))
	return data.Gender, nil
}

func (i ApiIo) GetAge(name string) (int, error) {
	base, err := url.Parse("https://api.agify.io/")
	if err != nil {
		return 0, err
	}
	params := url.Values{}
	params.Add("name", name)
	base.RawQuery = params.Encode()
	urlStr := base.String()
	i.log.Debug("api: GetAge", slog.String("name", name), slog.String("url", urlStr))

	resp, err := http.Get(urlStr)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("GetAge http.Status %d", resp.StatusCode)
	}

	type ageResponse struct {
		Count int
		Name  string
		Age   int
	}
	var data ageResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, err
	}
	i.log.Debug("api: GetAge", slog.Any("response", data))
	return data.Age, nil
}
func (i ApiIo) GetNationality(name string) (string, error) {
	base, err := url.Parse("https://api.nationalize.io/")
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("name", name)
	base.RawQuery = params.Encode()
	urlStr := base.String()
	i.log.Debug("api: GetNationality", slog.String("name", name), slog.String("url", urlStr))

	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GetNationality http.Status %d", resp.StatusCode)
	}
	type Country struct {
		Country_id  string
		Probability float32
	}
	type nationResponse struct {
		Count   int
		Name    string
		Country []Country
	}
	var data nationResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	i.log.Debug("api: GetNationality", slog.Any("response", data))
	if len(data.Country) > 0 {
		return data.Country[0].Country_id, nil
	}
	return "", nil
}
