package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ApiResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func fetchApi(base, target string) (float64, error) {
	api := "https://api.exchangeratesapi.io/latest"
	url := api + "?symbols=" + target + "&base=" + base
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 400 {
		return 0, fmt.Errorf("API error: Probably misspelled one of currencies: %s, %s", base, target)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result ApiResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return result.Rates[target], nil
}

func calculateEquivalent(base, target string, amount float64) (float64, float64, error) {
	rate, err := fetchApi(base, target)
	if err != nil {
		return 0, 0, err
	}

	return amount * rate, rate, nil
}

func main() {
	args := os.Args[1:]

	helpFlag := flag.Bool("h", false, "Show help")
	flag.Parse()

	if *helpFlag || len(args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("currconv 2.5 eur usd")
		fmt.Println("currconv 100 USD to pln")
		fmt.Println("currconv 25 GBP USD")
		return
	}

	amountOfBaseCurr, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		panic(err)
	}

	baseCurrency := strings.ToUpper(args[1])
	targetCurrency := strings.ToUpper(args[len(args)-1])

	amountOfTargetCurr, rate, err := calculateEquivalent(baseCurrency, targetCurrency, amountOfBaseCurr)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%f %s = %f %s\n", amountOfBaseCurr, baseCurrency, amountOfTargetCurr, targetCurrency)
	fmt.Printf("1 %s = %f %s\n", baseCurrency, rate, targetCurrency)
}
