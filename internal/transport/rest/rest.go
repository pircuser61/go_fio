package rest

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	config "github.com/pircuser61/go_fio/config"
	"github.com/pircuser61/go_fio/internal/models"
	srv "github.com/pircuser61/go_fio/internal/services"
)

type response struct {
	Time   int64
	Error  bool
	ErrMsg string
	Body   any
}

var log *slog.Logger

func RunHttpServer(ctx context.Context, logger *slog.Logger) error {
	log = logger
	port := config.GetHTTPPort()

	router := mux.NewRouter()
	router.HandleFunc("/", PersonList).Methods("GET")
	router.HandleFunc("/", PersonCreate).Methods("POST")
	router.HandleFunc("/{id}", PersonGet).Methods("GET")
	router.HandleFunc("/{id}", PersonUpdate).Methods("PUT")
	router.HandleFunc("/{id}", PersonDelete).Methods("DELETE")

	log.Info("http start", slog.String("port", port))
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Error("http exit with error", err)
	} else {
		log.Error("http exit")
	}
	return nil
}

func PersonList(rw http.ResponseWriter, req *http.Request) {
	log.Debug("rest: person list")
	res, err := srv.PersonList(context.Background())
	makeResp(rw, 0, res, err)
}

func PersonCreate(rw http.ResponseWriter, req *http.Request) {
	log.Debug("rest: person create")
	defer req.Body.Close()
	var person models.Person
	err := json.NewDecoder(req.Body).Decode(&person)
	if err == nil {
		log.Debug("rest: params", slog.String("name", person.Name))
		person.Id, err = srv.PersonCreate(context.Background(), &person)
	}
	makeResp(rw, 0, person.Id, err)
}

func PersonUpdate(rw http.ResponseWriter, req *http.Request) {
	log.Debug("rest: person update")
	defer req.Body.Close()
	var person models.Person
	vars := mux.Vars(req)
	strId := vars["id"]
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		makeResp(rw, 0, nil, err)
		return
	}

	err = json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		makeResp(rw, 0, nil, err)
		return
	}

	person.Id = uint32(id)
	log.Debug("rest: params", slog.Uint64("Id", id), slog.String("name", person.Name))
	err = srv.PersonUpdate(context.Background(), &person)

	makeResp(rw, 0, nil, err)
}

func PersonGet(rw http.ResponseWriter, req *http.Request) {
	log.Debug("rest: person get")
	var res *models.Person
	vars := mux.Vars(req)
	strId := vars["id"]
	id, err := strconv.ParseUint(strId, 10, 32)
	if err == nil {
		res, err = srv.PersonGet(context.Background(), uint32(id))
	}
	makeResp(rw, 0, res, err)
}

func PersonDelete(rw http.ResponseWriter, req *http.Request) {
	log.Debug("rest: person delete")
	vars := mux.Vars(req)
	strId := vars["id"]
	id, err := strconv.ParseUint(strId, 10, 32)
	if err == nil {
		err = srv.PersonDelete(context.Background(), uint32(id))
	}
	makeResp(rw, 0, nil, err)
}

func makeResp(rw http.ResponseWriter, tm int64, body any, err error) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if err != nil {
		_ = json.NewEncoder(rw).Encode(response{Error: true, ErrMsg: err.Error()})
	} else {
		_ = json.NewEncoder(rw).Encode(response{Time: tm, Body: body})
	}

}
