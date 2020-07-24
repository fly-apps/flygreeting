package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
 * Upstream source: https://wiki.openstreetmap.org/wiki/Nominatim/Country_Codes
 */

type countryLanguage struct {
	CountryCode string `json:"CountryCode"`
	EnglishName string `json:"CountrynameEnglish"`
	LocalName   string `json:"CountrynameLocal"`
	Languages   string `json:"Languages"`
}

var countryToLanguageMapping map[string]countryLanguage

func initCountryLanguageMapping() {
	langfile, err := os.Open("countrylangmap.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var countryToLanguageArray []countryLanguage

	bytes, err := ioutil.ReadAll(langfile)

	json.Unmarshal(bytes, &countryToLanguageArray)

	countryToLanguageMapping = make(map[string]countryLanguage, len(countryToLanguageArray))

	for _, r := range countryToLanguageArray {
		countryToLanguageMapping[r.CountryCode] = r
	}

}
