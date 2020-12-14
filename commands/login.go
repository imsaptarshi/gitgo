package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/saptarshibasu15/gitgo/utils"
)

func Login() {
	resp, err := http.PostForm("https://github.com/login/device/code", url.Values{
		"client_id": {"625394a947368b4103bf"},
		"scope":     {"repo"}})

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	auth_code := strings.Split(string(body), "&")[3]
	deviceCode := strings.Split(string(body), "=")[0]
	auth_code = strings.Split(auth_code, "=")[1]
	fmt.Printf("Open https://github.com/login/device and type this code:\n%s", utils.Green(auth_code))

	// Doesn't Work
	// TODO
	time.Sleep(time.Second * 20)

	resp, err = http.PostForm("https://github.com/login/oauth/access_token", url.Values{
		"client_id":   {"625394a947368b4103bf"},
		"device_code": {auth_code},
		"grant_type":  {fmt.Sprintf("urn:ietf:params:oauth:repo:%s", deviceCode)},
	})

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(body))
	resp.Body.Close()
}
