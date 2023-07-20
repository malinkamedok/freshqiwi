package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"golang.org/x/exp/slices"
	"golang.org/x/net/html/charset"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ValCurs struct {
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func parseFlags() (string, string) {
	var code, date string
	flag.StringVar(&code, "code", "", "enter currency code")

	flag.StringVar(&date, "date", "", "enter date")

	flag.Parse()

	return code, date
}

func checkValuteCorrect(code string) {
	valutes := []string{"USD", "AUD", "AZN", "GBP", "AMD", "BYN", "BGN", "BRL", "HUF", "VND", "HKD", "GEL", "DKK", "AED", "EUR", "EGP", "INR", "IDR", "KZT", "CAD", "QAR", "KGS", "CNY", "MDL", "NZD", "NOK", "PLN", "RON", "XDR", "SGD", "TJS", "THB", "TRY", "TMT", "UZS", "UAH", "CZK", "SEK", "CHF", "RSD", "ZAR", "KRW", "JPY"}
	if slices.Contains(valutes, code) {
		//В ближайшем времени подобный метод появится в стандартной библиотеке
		//https://pkg.go.dev/slices
	} else {
		fmt.Println("currency code is incorrect. exiting")
		os.Exit(1)
	}
}

func parseAndFormatTime(dateStr string) string {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Print(err.Error())
	}

	dateFormatted := date.Format("02/01/2006")

	return dateFormatted
}

func initClientAndRequest(date string) (http.Client, *http.Request) {
	client := http.Client{}

	url := "https://www.cbr.ru/scripts/XML_daily.asp?date_req=" + date

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	//Без указания юзер агента получаем ошибку 403 - Forbidden
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	return client, req
}

func decodeResponse(response *http.Response) *ValCurs {
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	quotes := new(ValCurs)
	reader := bytes.NewReader(responseData)
	decoder := xml.NewDecoder(reader)
	//Кодировка XML в исходном документе - Windows-1251
	//Необходима конвертация в UTF-8
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(quotes)
	if err != nil {
		fmt.Print(err.Error())
	}

	return quotes
}

func findAndPrintValute(quotes *ValCurs, code string) {
	var exists bool
	for _, v := range quotes.Valutes {
		if v.CharCode == code {
			fmt.Printf("%s (%s): %s\n", v.CharCode, v.Name, v.Value)
			exists = true
		}
	}
	if !exists {
		fmt.Println("Данные по указанным параметрам отсутствуют")
	}
}

func main() {
	code, dateStr := parseFlags()

	checkValuteCorrect(code)

	date := parseAndFormatTime(dateStr)

	client, req := initClientAndRequest(date)

	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	quotes := decodeResponse(response)

	findAndPrintValute(quotes, code)
}
