package race

import (
	"fmt"
	"log"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	"github.com/wahyuhadi/race/models"
	"github.com/wahyuhadi/race/parser"
)

func Run(c parser.Request, opt models.Opt) {
	rate := vegeta.Rate{Freq: opt.TotalReq, Per: time.Second}
	duration := time.Duration(opt.Duration) * time.Second
	fmt.Println(c.Url)

	runner := vegeta.Target{
		Method: c.Method,
		URL:    c.Url,
		Header: c.Headers,
	}
	if c.Body != nil {
		runner.Body = c.Body
	}

	targeter := vegeta.NewStaticTargeter(runner)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	var total = 0
	var StatusOK = 0
	for res := range attacker.Attack(targeter, rate, duration, "") {
		log.Println("[+] get status code ", res.Code, string(res.Body))
		if res.Code == 200 {
			StatusOK += 1
		}
		total += 1
		metrics.Add(res)
	}
	metrics.Close()
	fmt.Println(fmt.Sprintf("total req %v status ok %v", total, StatusOK))
	fmt.Printf("99th percentile: %s %s\n", metrics.Latencies.P99, metrics.Latencies.Total)
}
