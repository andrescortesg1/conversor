// main app
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var currencyValue map[string]any
//	supportedCurrencies := []string{"usd", "eur", "jpy", "cop"}	

	ok, err := APICheck(currencyAPI)
	if !ok {
		log.Fatal(err)
	}

	fmt.Println("--- Conversor de monedas ---")
	fmt.Printf("Monedas disponibles: \n")
	fmt.Printf("\t1- usd: Dolar de EEUU\n")
	fmt.Printf("\t2- eur: Euro\n")
	fmt.Printf("\t3- jpy: Yen Japones\n")
	fmt.Printf("\t4- cop: Peso Colombiano\n")

	fmt.Printf("Elija la moneda de origen: ")
	in, err := ReadData(supportedCurrencies, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Elija la moneda de destino: ")
	out, err := ReadData(supportedCurrencies, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ingrese el valor a calcular: ")
	value, err := ReadData(nil, os.Stdin)

	// get converted value via API call 
	switch in {
	case "usd":
		currencyValue, err = APIGet(usdAPI, "usd")
	case "eur":
		currencyValue, err = APIGet(eurAPI, "eur")
	case "jpy":
		currencyValue, err = APIGet(jpyAPI, "jpy")
	case "cop":
		currencyValue, err = APIGet(copAPI, "cop")
	default:
		fmt.Println("Entrada invalida, terminando...")
	}

	convertedValue := value.(float64) * currencyValue[out.(string)].(float64)
	fmt.Println(value, in, "->", convertedValue, out)
	fmt.Println("Terminando...")
}

// DEPRECATED: old repeated code implementation, do not touch
// for reference
/*
func ReadValue() (float64, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		return 0, err
	}
	if input < 0 {
		return 0, errors.New("Invalid value, must not be negative")
	}	
	return input, nil
}

func ReadSelection(list []string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
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
*/


