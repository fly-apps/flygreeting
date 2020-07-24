package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	initCountryLanguageMapping()
	initGreetingMapping()
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%d %s %s\"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	var ok bool

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	r.GET("/v1/countries/", countryList)
	r.GET("/v1/languages/", languageList)
	r.GET("/v1/language/:countrycode", languageLookup)
	r.GET("/v1/country/:countrycode", countryLookup)
	r.GET("/v1/greeting/:languagecode", greetingLookup)

	r.Run(":" + port)
}

func countryList(c *gin.Context) {
	keys := make([]string, 0, len(countryToLanguageMapping))
	for k := range countryToLanguageMapping {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	c.JSON(http.StatusOK, gin.H{"countries": keys})
	return
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func languageList(c *gin.Context) {
	languages := make([]string, 0, len(countryToLanguageMapping))
	for _, v := range countryToLanguageMapping {
		for _, l := range strings.Split(v.Languages, ",") {
			languages = appendIfMissing(languages, l)
		}
	}
	sort.Slice(languages, func(i, j int) bool {
		return languages[i] < languages[j]
	})

	c.JSON(http.StatusOK, gin.H{"languages": languages})
	return
}

func languageLookup(c *gin.Context) {
	country := c.Param("countrycode")
	countryEntry, ok := countryToLanguageMapping[strings.ToUpper(country)]
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	langsarray := strings.Split(countryEntry.Languages, ",")
	c.JSON(http.StatusOK, gin.H{"country": country, "languages": langsarray})
	return
}

func countryLookup(c *gin.Context) {
	country := c.Param("countrycode")
	countryEntry, ok := countryToLanguageMapping[strings.ToUpper(country)]
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"country": countryEntry})
	return
}

func greetingLookup(c *gin.Context) {
	language := c.Param("languagecode")
	mygreetings, ok := greetingMapping[language]
	if !ok {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, mygreetings)
	return
}
