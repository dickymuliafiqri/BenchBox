package bench

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/LalatinaHub/LatinaSub-go/provider"
	"github.com/dickymuliafiqri/BenchBox/modules/benchmark"
	singbox "github.com/dickymuliafiqri/BenchBox/modules/sing-box"
)

type ResultType struct {
	Node   string         `json:"node"`
	Result map[string]int `json:"result"`
}

func PostBench(w http.ResponseWriter, r *http.Request) {
	result := []ResultType{}
	nodes := strings.Split(r.PostFormValue("url"), ",")
	outbounds, err := provider.Parse(strings.Join(nodes, "\n"))
	if err != nil {
		fmt.Printf("Error while parsing node: %s", err)
	}

	for i, outbound := range outbounds {
		opt, listenPort := singbox.GenerateConfig(&outbound)
		box, err := singbox.Create(opt)
		if err != nil {
			fmt.Printf("Error while starting singbox instances: %s", err)
			continue
		}
		defer box.Close()

		// Wait singbox to fully started
		time.Sleep(1 * time.Second)

		result = append(result, ResultType{
			Node:   nodes[i],
			Result: benchmark.StartBenchmark(listenPort),
		})
	}

	text, err := json.Marshal(result)
	if err != nil {
		text = []byte(err.Error())
	}

	io.WriteString(w, string(text))
}
