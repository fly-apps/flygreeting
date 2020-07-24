package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguage(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/language/ch", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"country\":\"ch\",\"languages\":[\"de\",\"fr\",\"it\",\"rm\"]}", w.Body.String())
}

func TestCountry(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/country/ch", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"country\":{\"CountryCode\":\"CH\",\"CountrynameEnglish\":\"Switzerland\",\"CountrynameLocal\":\"Schweiz,Suisse,Svizzera,Svizra\",\"Languages\":\"de,fr,it,rm\"}}", w.Body.String())
}

func TestGreeting(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/greeting/fr", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"language\":\"fr\",\"formal\":\"Bonjour\",\"informal\":\"Salut\"}", w.Body.String())
}
