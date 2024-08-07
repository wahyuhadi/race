package main

import (
	"errors"
	"flag"
	"log"
	"net/url"
	"os"

	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/race/models"
	"github.com/wahyuhadi/race/parser"
	"github.com/wahyuhadi/race/race"
)

var file = flag.String("f", "", "http request file ")
var total = flag.Int("total", 1000, "total req /s")
var duration = flag.Int("duration", 10, "duration of attack in seccond")
var urls = flag.String("url", "", "url")

const (
	fileNotFond = "http raw file not found"
)

func ParseOptions() (opts *models.Opt) {
	flag.Parse()
	return &models.Opt{
		File:     *file,
		URL:      *urls,
		Duration: *duration,
		TotalReq: *total,
	}
}

func CheckOptions(opts *models.Opt) (err error) {
	if opts.File == "" {
		return errors.New(fileNotFond)
	}

	if opts.URL == "" {
		return errors.New("url is empty")
	}
	return nil
}

func main() {
	opts := ParseOptions()
	err := CheckOptions(opts)
	if err != nil {
		gologger.Info().Str("state", "errored").Str("status", "error").Msg(err.Error())
		return
	}

	f, err := os.ReadFile(opts.File)
	if err != nil {
		log.Fatal(err)
	}

	request, err := parser.ReadHTTPFromFile(string(f), opts.URL)
	if err != nil {
		gologger.Info().Str("status", "error1").Msg(err.Error())
		return
	}

	u, e := url.Parse(opts.URL + request.Path)
	if e != nil {
		gologger.Info().Str("status", "error").Msg(e.Error())
		return
	}
	request.Url = opts.URL + request.Path
	request.Scheme = u.Scheme
	request.Port = "80"
	if u.Scheme == "https" {
		request.Port = "443"
	}
	request.Query = u.RawQuery
	race.Run(*request, *opts)
}
