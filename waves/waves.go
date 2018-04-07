package waves

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)

// post Пост запрос с параметрами в тело
func post(url1 string, data url.Values) string {
	form := data
	body1 := bytes.NewBufferString(form.Encode())
	req, err := http.NewRequest("POST", url1, body1)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// CreateSeed Получаем зашифрованный seed
func CreateSeed(userID string, name string) string {
	postData := url.Values{
		"userID": {userID},
		"name":   {name},
	}
	prvtKey := post("http://localhost:3001/createSeed", postData)
	return prvtKey
}

// GetAddress Получение адреса по зашифрованному сиду
func GetAddress(userID string, seed string) string {
	postData := url.Values{
		"userID": {userID},
		"seed":   {seed},
	}
	address := post("http://localhost:3000/getAddress", postData)
	return address
}

// GetBalance Получение баланса токена на блокчейне Waves
func GetBalance(address string, currency string) string {
	postData := url.Values{
		"address":  {address},
		"currency": {currency},
	}
	balance := post("http://localhost:3001/getBalance", postData)
	return balance
}

// GetWavesBalance Получение баланса WAVES
func GetWavesBalance(address string) gjson.Result {
	resp, err := http.Get("https://nodes.wavesnodes.com/addresses/balance/" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	balance := gjson.Get(string(body), "balance")
	return balance
}

// DecryptSeed Расшифровываем Seed
func DecryptSeed(userID string, encryptedSeed string) string {
	postData := url.Values{
		"userID":        {userID},
		"encryptedSeed": {encryptedSeed},
	}
	seed := post("http://localhost:3001/decryptSeed", postData)
	return seed
}
