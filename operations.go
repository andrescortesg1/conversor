package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"slices"
	"strconv"
)

const (
	currencyAPI = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.min.json"
	eurAPI = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.min.json"
	usdAPI = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/usd.min.json"
	jpyAPI = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json"
	copAPI = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/cop.min.json"
)

var supportedCurrencies = []string{"usd", "eur", "jpy", "cop"}	

func APICheck(apiLink string) (bool, error) {
	var outputData map[string]any

	r, err := http.Get(apiLink)
	if err != nil {
		return false, err
	}

	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		return false, err
	}
	defer func(){
		err = r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err = json.Unmarshal(responseData, &outputData); err != nil {
		return false, err
	}
	
	return true, nil
}

// apiLink: link to the API call
// code: currency code
// returns map with currency values
func APIGet(apiLink string, code string) (map[string]any, error) {
	var outputData map[string]any

	r, err := http.Get(apiLink)
	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer func(){
		err = r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err = json.Unmarshal(responseData, &outputData); err != nil {
		return nil, err
	}

	result, ok := outputData[code].(map[string]any)
	if !ok {
		return nil, errors.New("Invalid currency code")
	}

	return result, nil
}

// list: curated currency list 
// if nil: returns float64
// else: returns string
func ReadData(list []string, reader io.Reader) (any, error) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if list == nil { // returns float64
		input, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return 0, err
		}
		if input < 0 {
			return 0, errors.New("Invalid value, must not be negative")
		}	
		return input, nil
	} else { // returns string
		if r := slices.Compare(list, supportedCurrencies); r != 0 {
			return nil, errors.New("Unsupported currency list")
		}
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return "", err
		}
		if input > 4 || input < 1 {
			return "", errors.New("Invalid value, must be between 1 and 4")
		}
		input = input - 1
		return list[input], nil
	}
}
