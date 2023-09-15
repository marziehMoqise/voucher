package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"net/http"
)

const baseUrl = "http://127.0.0.1:7575"

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func SendTestReq(apiReq interface{}, apiUrl string) (res Response) {
	log.Info("apiReq: ", apiReq)
	b, err := json.Marshal(apiReq)

	req, err := http.NewRequest("POST",  baseUrl + apiUrl, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp == nil {
		panic("invalid response")
	}

	fmt.Println("statusCode", resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	if len(body) == 0 {
		fmt.Println("empty body")
		return
	}

	if err = resp.Body.Close(); err != nil {
		panic(err)
	}


	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Status)
	fmt.Println(res.Data)

	return res
}