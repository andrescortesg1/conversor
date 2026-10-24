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

	convertedValue, err := ProcessValue(in, out, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value, in, "->", convertedValue, out)
	fmt.Println("Terminando...")
}
