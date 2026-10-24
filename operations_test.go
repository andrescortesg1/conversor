/*
GIVEN
WHEN
THEN
*/
package main

import (
	"io"
	"strings"
	"testing"
)

func TestAPICheck(t *testing.T) {
	tests := []struct{
		api string
		caseName string
		wantErr bool
	}{
		{ 	//  case: valid api link
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.json",
			"valid api link",
			false,
		},
		{ 	//  case: wrong address 
			"https://www.google.com",
			"wrong address",
			true,
		},
		{ 	//  case: no address 
			"",
			"no address",
			true,
		},
	}

	for _, tt := range tests {
		_, err := APICheck(tt.api)

		if (err != nil) != tt.wantErr {
			t.Errorf("APICheck(): case: %v - error: %v - wantErr: %v", tt.caseName, err, tt.wantErr)
		} else {
			t.Logf("APICheck(): case: %v PASS", tt.caseName)
		}
	}

}

func TestAPIGet(t *testing.T) {
	tests := []struct{
		apiLink string
		code string
		caseName string
		wantErr bool
	}{
		{ 	//  case: usd
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/usd.min.json",
			"usd",
			"usd",
			false,
		},
		{ 	//  case: eur
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.min.json",
			"eur",
			"eur",
			false,
		},
		{ 	//  case: jpy
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"jpy",
			"jpy",
			false,
		},
		{ 	//  case: cop
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/cop.min.json",
			"cop",
			"cop",
			false,
		},
		{ 	//  case: different currencies 
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"usd",
			"different currencies",
			true,
		},
		{ 	//  case: no api link 
			"",
			"usd",
			"no api link",
			true,
		},
		{ 	//  case: no code 
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"",
			"no code",
			true,
		},
	}

	for _, tt := range tests {
		_, err := APIGet(tt.apiLink, tt.code)

		if (err != nil) != tt.wantErr {
			t.Errorf("APIGet(): case: %v - error: %v - wantErr: %v", tt.caseName, err, tt.wantErr)
		} else {
			t.Logf("APIGet(): case: %v PASS", tt.caseName)
		}
	}
}

func TestReadData(t *testing.T) {
	tests := []struct{
		list []string
		input io.Reader
		caseName string
		wantErr bool
	}{
		// --- CURRENCY CODE TESTS ---
		{ 	//  case: get supported currencies
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			"get supported currencies",
			false,
		},
		{ 	//  case: unsupported valid currencies
			[]string{"btc", "usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			"unsupported valid currencies",
			true,
		},
		{ 	//  case: invalid currencies
			[]string{"udea", "usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			"invalid currencies",
			true,
		},
		{ 	//  case: value < 1 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("0"),
			"value < 1",
			true,
		},
		{ 	//  case: value > 1 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("6"),
			"value > 1",
			true,
		},
		{ 	//  case: invalid input 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("udea"),
			"invalid input",
			true,
		},
		{ 	//  case: no input 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader(""),
			"no input",
			true,
		},
		// --- VALUE TESTS ---
		{ 	//  case: valid input 
			nil,
			strings.NewReader("10"),
			"valid input",
			false,
		},
		{ 	//  case: value < 0 
			nil,
			strings.NewReader("-1"),
			"value < 0",
			true,
		},
		{ 	//  case: invalid input 
			nil,
			strings.NewReader("udea"),
			"invalid input",
			true,
		},
		{ 	//  case: no input 
			nil,
			strings.NewReader(""),
			"no input",
			true,
		},
	}

	for _, tt := range tests {
		_, err := ReadData(tt.list, tt.input)

		if (err != nil) != tt.wantErr {
			t.Errorf("ReadData(): case: %v - error: %v - wantErr: %v", tt.caseName, err, tt.wantErr)
		} else {
			t.Logf("ReadData(): case: %v PASS", tt.caseName)
		}
	}
}

func TestProcessValue(t *testing.T) {
	tests := []struct{
		in any
		out any
		value any 
		caseName string 
		wantErr bool
	}{
		{ // case: valid value to process
			"usd",
			"cop",
			1.0,
			"valid value to process USD",
			false,
		},
		{ // case: valid value to process
			"eur",
			"cop",
			1.0,
			"valid value to process EUR",
			false,
		},
		{ // case: valid value to process
			"jpy",
			"cop",
			1.0,
			"valid value to process JPY",
			false,
		},
		{ // case: valid value to process
			"cop",
			"usd",
			1.0,
			"valid value to process COP",
			false,
		},
		{ // case: missing currency code
			"",
			"cop",
			1.0,
			"missing currency code",
			true,
		},
		{ // case: missing all currency codes
			"",
			"",
			1.0,
			"missing all currency codes",
			true,
		},
		{ // case: negative value
			"usd",
			"cop",
			-1.0,
			"negative value",
			true,
		},
		{
			// case: wrong IN currency code 
			"udea",
			"cop",
			1.0,
			"wrong IN currency code",
			true,
		},

		{
			// case: wrong OUT currency code 
			"usd",
			"udea",
			1.0,
			"wrong OUT currency code",
			true,
		},
	}

	for _, tt := range tests {
		_, err := ProcessValue(tt.in, tt.out, tt.value)

		if (err != nil) != tt.wantErr {
			t.Errorf("ProcessValue(): case: %v - error: %v - wantErr: %v", tt.caseName, err, tt.wantErr)
		} else {
			t.Logf("ProcessValue(): case: %v PASS", tt.caseName)
		}
	}
}
