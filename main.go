/* app */
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
//	var currencyValue map[string]any
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
	if err != nil {
		log.Fatal(err)
	}

	convertedValue, err := processValue(in, out, value)

	/*
	// to be function: 
	// input: 
	//  in - string
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
	if err != nil {
		log.Fatal(err)
	}

	convertedValue := value.(float64) * currencyValue[out.(string)].(float64)
	// WIP function out: 
	// convertedValue - float64	
	*/

	fmt.Println(value, in, "->", convertedValue, out)
	fmt.Println("Terminando...")
}

/*
func processValue(in, out ,value any) (float64, error){
	var currencyValue map[string]any
	var err error
	// to be function: 
	// input: 
	//  in - string
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
		return 0, errors.New("Entrada invalida, terminando...")
	}
	if err != nil {
		return 0, err
	}

	convertedValue := value.(float64) * currencyValue[out.(string)].(float64)

	return convertedValue, nil
}
*/
