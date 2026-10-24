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
		wantErr bool
	}{
		{ 	//  case: usd
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/usd.min.json",
			"usd",
			false,
		},
		{ 	//  case: eur
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.min.json",
			"eur",
			false,
		},
		{ 	//  case: jpy
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"jpy",
			false,
		},
		{ 	//  case: cop
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/cop.min.json",
			"cop",
			false,
		},
		{ 	//  case: different currencies 
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"usd",
			true,
		},
		{ 	//  case: no api link 
			"",
			"usd",
			true,
		},
		{ 	//  case: no code 
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"",
			true,
		},
	}

	for _, tt := range tests {
		_, err := APIGet(tt.apiLink, tt.code)

		if (err != nil) != tt.wantErr {
			t.Errorf("APIGet(): error: %v - wantErr: %v", err, tt.wantErr)
		}
	}
}

func TestReadData(t *testing.T) {
	tests := []struct{
		list []string
		input io.Reader
		wantErr bool
	}{
		// --- CURRENCY CODE TESTS ---
		{ 	//  case: get supported currencies
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			false,
		},
		{ 	//  case: unsupported valid currencies
			[]string{"btc", "usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			true,
		},
		{ 	//  case: invalid currencies
			[]string{"udea", "usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			true,
		},
		{ 	//  case: value < 1 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("0"),
			true,
		},
		{ 	//  case: value > 1 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("6"),
			true,
		},
		{ 	//  case: invalid input 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("udea"),
			true,
		},
		{ 	//  case: no input 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader(""),
			true,
		},
		// --- VALUE TESTS ---
		{ 	//  case: valid input 
			nil,
			strings.NewReader("10"),
			false,
		},
		{ 	//  case: value < 0 
			nil,
			strings.NewReader("-1"),
			true,
		},
		{ 	//  case: invalid input 
			nil,
			strings.NewReader("udea"),
			true,
		},
		{ 	//  case: no input 
			nil,
			strings.NewReader(""),
			true,
		},
	}

	for _, tt := range tests {
		_, err := ReadData(tt.list, tt.input)

		if (err != nil) != tt.wantErr {
			t.Errorf("ReadData(): error: %v - wantErr: %v", err, tt.wantErr)
		}
	}
}
