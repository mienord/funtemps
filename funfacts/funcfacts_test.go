package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string // her må du skrive riktig type for input
		want  []string // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
	  {input: "sun" , want: []string {"Temperaturen i Solens kjerne er: ", "Temperaturen på det ytre laget laget av Solen er: "}},
	  {input: "terra", want: []string {"Høyeste temperatur målt på Jordens overflate er: ", "Laveste temperatur målt på Jordens overflate er: ", "Temperaturen i Jordens indre kjerne er: "}},
	  {input: "luna", want: []string {"Temperaturen på Månens overflate om natten er: ", "Temperaturen på Månens overflate om dagen er: "}},
	}

	
	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

//bruker en switch statement til å matche input verdien med enten sun, terra, luna - også returnerer
func GoGetFunFacts(input string) []string {
	switch input {
	case "sun":
		return []string {"Temperaturen i Solens kjerne er: ", "Temperaturen på det ytre laget laget av Solen er: ",}
	case "terra":
		return []string {"Høyeste temperatur målt på Jordens overflate er: ", "Laveste temperatur målt på Jordens overflate er: ", "Temperaturen i Jordens indre kjerne er: ",}
	case "luna":
		return []string {"Temperaturen på Månens overflate om natten er: ", "Temperaturen på Månens overflate om dagen er: "}
	default:
		return []string{}
	}
}
