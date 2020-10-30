package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ApiResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func fetchApi(base, target string) float64 {
	api := "https://api.exchangeratesapi.io/latest"
	url := api + "?symbols=" + target + "&base=" + base
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result ApiResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return result.Rates[target]
}

func calculateEquivalent(base, target string, amount float64) (float64, float64) {
	rate := fetchApi(base, target)
	return amount * rate, rate
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	amountOfBaseCurr, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		panic(err)
	}

	baseCurrency := args[1]
	targetCurrency := args[len(args)-1]

	amountOfTargetCurr, rate := calculateEquivalent(baseCurrency, targetCurrency, amountOfBaseCurr)

	fmt.Printf("%f %s = %f %s\n", amountOfBaseCurr, baseCurrency, amountOfTargetCurr, targetCurrency)
	fmt.Printf("1 %s = %f %s\n", baseCurrency, rate, targetCurrency)
}
