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
		wantErr bool
	}{
		{ 	// good case: valid api link
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.json",
			false,
		},
		{ 	// bad case: wrong address 
			"https://www.google.com",
			true,
		},
		{ 	// bad case: no address 
			"",
			true,
		},
	}

	for _, tt := range tests {
		_, err := APICheck(tt.api)

		if (err != nil) != tt.wantErr {
			t.Errorf("APICheck(): error: %v - wantErr: %v", err, tt.wantErr)
		}
	}

}

func TestAPIGet(t *testing.T) {
	tests := []struct{
		apiLink string
		code string
		wantErr bool
	}{
		{ 	// good case: usd
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/usd.min.json",
			"usd",
			false,
		},
		{ 	// good case: eur
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/eur.min.json",
			"eur",
			false,
		},
		{ 	// good case: jpy
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"jpy",
			false,
		},
		{ 	// good case: cop
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/cop.min.json",
			"cop",
			false,
		},
		{ 	// bad case: different currencies 
			"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/jpy.min.json",
			"usd",
			true,
		},
		{ 	// bad case: no api link 
			"",
			"usd",
			true,
		},
		{ 	// bad case: no code 
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
		{ 	// good case: get supported currencies
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			false,
		},
		{ 	// good case: unsupported valid currencies
			[]string{"btc", "usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			true,
		},
		{ 	// bad case: invalid currencies
			[]string{"udea", "usd", "eur", "jpy", "cop"},
			strings.NewReader("1"),
			true,
		},
		{ 	// bad case: value < 1 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("0"),
			true,
		},
		{ 	// bad case: value > 1 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("6"),
			true,
		},
		{ 	// bad case: invalid input 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader("udea"),
			true,
		},
		{ 	// bad case: no input 
			[]string{"usd", "eur", "jpy", "cop"},
			strings.NewReader(""),
			true,
		},
		// --- VALUE TESTS ---
		{ 	// good case: valid input 
			nil,
			strings.NewReader("10"),
			false,
		},
		{ 	// bad case: value < 0 
			nil,
			strings.NewReader("-1"),
			true,
		},
		{ 	// bad case: invalid input 
			nil,
			strings.NewReader("udea"),
			true,
		},
		{ 	// bad case: no input 
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
