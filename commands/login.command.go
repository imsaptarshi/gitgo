package commands

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Login() {
	client := http.Client{}

	params := url.Values{}

	params.Add("client_id", "625394a947368b4103bf")

	req, err := http.NewRequest("POST", "https://github.com/login/device/code", strings.NewReader(params.Encode()))

	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	fmt.Printf("%v", resp)
}
