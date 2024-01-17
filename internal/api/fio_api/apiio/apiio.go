package apiio

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/pircuser61/go_fio/internal/api/fio_api"
)

type ApiIo struct {
	log *slog.Logger
}

func GetApi(logger *slog.Logger) fio_api.Api {
	return ApiIo{log: logger}
}

func (i ApiIo) GetGender(name string) (string, error) {
	i.log.Debug("api: GetGender")

	req, err := http.NewRequest("GET", "https://api.agify.io/?name=Dmitriy", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("name", name)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	i.log.Debug("api: GetGender response", slog.Int("code", res.StatusCode))

	var respJson interface{}
	json.NewDecoder(res.Body).Decode(&respJson)
	i.log.Debug("api: GetGender response", slog.Any("body", respJson))

	return "", nil
}

func (i ApiIo) GetAge(name string) (int, error) {
	i.log.Debug("api: GetAge")
	return 0, nil
}
func (i ApiIo) GetNationality(name string) (string, error) {
	i.log.Debug("api: GetNationality")
	return "", nil
}
