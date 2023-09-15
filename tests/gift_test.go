package tests

import (
	"apiGolang/apiSchema/userSchema"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"net/http"
	"testing"
)

func TestGift(t *testing.T) {
	reqGift := userSchema.GiftRequest{
		Mobile:      "+989155099394",
		VoucherCode: "aaaa",
	}

	b, err := json.Marshal(reqGift)
	log.Info("reqGift: ",bytes.NewBuffer(b))
	req, err := http.NewRequest("POST", "http://127.0.0.1:7575/user/gift", bytes.NewBuffer(b))

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

	var respM struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data"`
	}
	err = json.Unmarshal(body, &respM)
	if err != nil {
		panic(err)
	}

	fmt.Println(respM.Status)
	fmt.Println(respM.Data)
}
