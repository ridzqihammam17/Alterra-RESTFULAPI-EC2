package midtrans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type JSON map[string]interface{}

func RequestPayment(OrderId string, amount int) (redirectURL string, err error) {

	url := "https://app.sandbox.midtrans.com/snap/v1/transactions"
	method := "POST"

	payload, err := json.Marshal(JSON{
		"transaction_details": JSON{"order_id": OrderId, "gross_amount": amount},
		"credit_card":         JSON{"secure": true},
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic U0ItTWlkLXNlcnZlci1LTjFYOFVBckRIdERvcEc1aVF3d1c2Zi0=")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))
	temp2 := string(body)
	temp1 := strings.Index(temp2, "https")
	return string(body[temp1 : len(body)-2]), nil
}

func StatusPayment(OrderId string) (redirectURL string, err error) {
	url := "https://api.sandbox.midtrans.com/v2/" + OrderId + "/status"
	method := "GET"

	payload := strings.NewReader("\n\n")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic U0ItTWlkLXNlcnZlci1LTjFYOFVBckRIdERvcEc1aVF3d1c2Zi0=")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return string(body), nil
}
