package main

import (
	"code.google.com/p/goprotobuf/proto"
	"github.com/codegangsta/martini"
	"io/ioutil"
	"net/http"
)

func main() {
	server := martini.Classic()
	server.Get("/mtaapi", func() string {
		resp, err := http.Get("http://datamine.mta.info/mta_esi.php?key=ef80ff678028e7cc96ba474e6194b7d6")
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		transit := &FeedMessage{}
		_ = proto.Unmarshal(body, transit)
		s := transit.String()
		return "<h1>" + s + "</h1>"
	})
	server.Run()
}
