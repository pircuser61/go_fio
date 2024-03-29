package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	config "github.com/pircuser61/go_fio/config"
	"github.com/pircuser61/go_fio/internal/models"
)

type ReqPerson struct {
	Name       string
	Surname    string
	Patronymic string
}
type Response struct {
	Time   int64
	Error  bool
	ErrMsg string
}

var sUrl string

func main() {
	var line string

	port := config.GetHTTPPort()
	sUrl = "http://127.0.0.1" + port
	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nPerson>")

		if !in.Scan() {
			fmt.Println("Scan error")
			continue
		}
		line = in.Text()
		cmd := strings.Split(line, " ")[0]
		switch cmd {
		case "й":
			fallthrough
		case "quit":
			fallthrough
		case "q":
			return

		case "list":
			listPerson(line)
		case "add":
			addPerson(line)
		case "put":
			fallthrough
		case "update":
			updatePerson(line)
		case "get":
			getPerson(line)
		case "del":
			delPerson(line)
		default:
			fmt.Printf("Unknown command <%s>\n", line)
		}
	}
}

func listPerson(line string) {

	base, err := url.Parse(sUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	params := strings.Split(line, " ")
	reqParams := url.Values{}
	for i := 3; i <= len(params); i += 2 {
		reqParams.Add(params[i-2], params[i-1])
	}
	base.RawQuery = reqParams.Encode()
	fmt.Println(base.String())
	resp, err := http.DefaultClient.Get(base.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	type ListResponse struct {
		Response
		Body []*models.Person
	}
	var data ListResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Json parse error:", err)
		return
	}
	if data.Error {
		fmt.Println("ERROR", data.ErrMsg)
	} else {
		for _, val := range data.Body {
			fmt.Println(*val)
		}
	}
}

func addPerson(line string) {
	params := strings.Split(line, " ")

	var req ReqPerson
	switch len(params) {
	case 3:
		req.Name = params[1]
		req.Surname = params[2]
	case 4:
		req.Name = params[1]
		req.Surname = params[2]
		req.Patronymic = params[3]
	default:
		fmt.Printf("invalid args %d items <%v>", len(params), params)
		return
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.DefaultClient.Post(sUrl, "text/json", bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	type AddRespnose struct {
		Response
		Body any
	}
	var data AddRespnose
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Json parse error:", err)
		return
	}
	if data.Error {
		fmt.Println("ERROR", data.ErrMsg)
	} else {
		fmt.Println("new id:", data.Body)
	}
}

func updatePerson(line string) {
	params := strings.Split(line, " ")

	var reqParam ReqPerson
	switch len(params) {
	case 4:
		reqParam.Name = params[2]
		reqParam.Surname = params[3]
	case 5:
		reqParam.Name = params[2]
		reqParam.Surname = params[3]
		reqParam.Patronymic = params[4]
	default:
		fmt.Printf("invalid args %d items <%v>", len(params), params)
		return
	}
	jsonBody, err := json.Marshal(reqParam)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, sUrl+"/"+params[1], bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request error", err.Error())
		return
	}

	defer resp.Body.Close()
	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Json parse error:", err)
		return
	}
	if data.Error {
		fmt.Println("ERROR", data.ErrMsg)
	} else {
		fmt.Println("done")
	}
}

func getPerson(line string) {
	params := strings.Split(line, " ")
	if len(params) != 2 {
		fmt.Printf("invalid args %d items <%v>", len(params), params)
		return
	}

	resp, err := http.DefaultClient.Get(sUrl + "/" + params[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	type GetRespnose struct {
		Response
		Body models.Person
	}
	var data GetRespnose
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Json parse error:", err)
		return
	}
	if data.Error {
		fmt.Println("ERROR", data.ErrMsg)
	} else {
		fmt.Println(data.Body)
	}
}

func delPerson(line string) {
	params := strings.Split(line, " ")
	if len(params) != 2 {
		fmt.Printf("invalid args %d items <%v>", len(params), params)
		return
	}

	req, err := http.NewRequest(http.MethodDelete, sUrl+"/"+params[1], nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Json parse error:", err)
		return
	}
	if data.Error {
		fmt.Println("ERROR", data.ErrMsg)
	} else {
		fmt.Println("done")
	}

}
