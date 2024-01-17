package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	config "github.com/pircuser61/go_fio/config"
)

var url string

func main() {
	var line string

	port, err := config.GetHTTPPort()
	if err != nil {
		fmt.Println(err)
		return
	}
	url = "http://127.0.0.1" + port
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
			listPerson()
		case "add":
			addPerson(line)
		case "asyncAdd":
		case "update":
		case "get":
			getPerson(line)
		case "delete":
		default:
			fmt.Printf("Unknown command <%s>\n", line)
		}
	}
}

func listPerson() {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

func addPerson(line string) {
	params := strings.Split(line, " ")

	type Req struct {
		name    string
		surname string
	}
	var req Req
	switch len(params) {
	case 3:
		req.name = params[1]
		req.surname = params[2]
	case 4:
		req.name = params[1]
		req.surname = params[2]
	default:
		fmt.Printf("invalid args %d items <%v>", len(params), params)
		return
	}
	jsonBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.DefaultClient.Post(url, "text/json", bytes.NewReader(jsonBody))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}

func getPerson(line string) {
	params := strings.Split(line, " ")
	if len(params) != 2 {
		fmt.Printf("invalid args %d items <%v>", len(params), params)
		return
	}

	resp, err := http.DefaultClient.Get(url + "/" + params[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}
