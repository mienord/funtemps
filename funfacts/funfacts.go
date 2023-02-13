package funfacts

/**
  Implementer funfacts-funksjon:
    GetFunFacts(about string) []string
      hvor about kan ha en av tre testverdier, -
        sun, luna eller terra

  Sett inn alle Funfucts i en struktur
  type FunFacts struct {
      Sun []string
      Luna []string
      Terra []string
  }
*/


import (
	"fmt"
)

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

//tror ikke vi skal ha denne - en for loop som printer ut alle funfactsa etter hverandre
func main() {
	funFacts := GetFunFacts("sun")
	for i, fact := range funFacts {
		fmt.Println(i+1, fact)
	}
}

// GetFunFacts returns a slice of fun facts about the Sun, Moon, or Earth.
func GetFunFacts(about string) []string {
	var funFacts FunFacts

	funFacts.Sun = []string{
		"Temperaturen i Solens kjerne er: ",
		"Temperaturen på det ytre laget laget av Solen er: ",
	}

	funFacts.Luna = []string{
		"Temperaturen på Månens overflate om natten er: ",
		"Temperaturen på Månens overflate om dagen er: ",
	}

	funFacts.Terra = []string{
		"Høyeste temperatur målt på Jordens overflate er: ",
		"Laveste temperatur målt på Jordens overflate er: ",
		"Temperaturen i Jordens indre kjerne er: ",
	}

	//denne koden bruker en switch statement for å bestemme hvilken funfact som skal returneres basert på verdien til about 
	switch about {
	case "sun":
		return funFacts.Sun
	case "luna":
		return funFacts.Luna
	case "terra":
		return funFacts.Terra
	//hvis ikke about ikke er lik sun, terra eller luna vil det returnere denne feilmld
	default:
		return []string{"Invalid argument. Choose one of: sun, luna, terra"}
	}
}

