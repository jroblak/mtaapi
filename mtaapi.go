package main

import (
	"bytes"
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/cors"
	"io/ioutil"
	"net/http"
	"strings"
)

type MTAResponse struct {
	BT struct {
		Line []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"line"`
	} `json:"BT"`
	LIRR struct {
		Line []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"line"`
	} `json:"LIRR"`
	MetroNorth struct {
		Line []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"line"`
	} `json:"MetroNorth"`
	Bus struct {
		Line []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"line"`
	} `json:"bus"`
	Responsecode string `json:"responsecode"`
	Subway       struct {
		Line []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"line"`
	} `json:"subway"`
}

func main() {
	server := martini.Classic()

	server.Use(cors.Allow(&cors.Options{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	server.Get("/mta-api/subway", func() string {
		var buffer bytes.Buffer
		var r MTAResponse
		var bytes []byte

		resp, err := http.Get("http://www.mta.info/service_status_json/23458509")

		if err != nil {
			buffer.WriteString(err.Error())
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			buffer.WriteString(err.Error())
		}

		stringed := string(body[:])
		stringed = strings.Replace(stringed, "\\u0022", "\"", -1)
		stringed = stringed[3 : len(stringed)-1]

		bytes = []byte(stringed)

		e := json.Unmarshal(bytes, &r)
		if e != nil {
			buffer.WriteString(e.Error())
		}

		b, err := json.Marshal(r.Subway)

		return string(b[:])
	})
	server.Run()
}
