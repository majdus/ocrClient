package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SendPostRequest(jsonData []byte) ([]byte) {
	url:= "http://18.216.87.211:9292/ocr"

	fmt.Println("URL:>", url)
	fmt.Println("Request:>", string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status != "200 OK" {
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response body:", string(body))
	}

	return body
}



