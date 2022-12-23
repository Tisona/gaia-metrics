package metrics

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func collect(metrics string) string {
	resp, err := http.Get(fmt.Sprintf("http://localhost:26657/%s", metrics))

	if err != nil {
		log.Print(err)

		return string("")
	}

	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}

func CollectMetrics() (float64, float64, float64) {
	var latest_block_height float64 = -1
	var desync float64 = -1
	var n_peers float64 = -1

	status := collect("status")
	net_info := collect("net_info")

	if status != "" {

	}
	if net_info != "" {

	}

	return latest_block_height, desync, n_peers
}
