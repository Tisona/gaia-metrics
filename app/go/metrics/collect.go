package metrics

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

const GAIA_HOST = "http://localhost:26657"

func readAPI(resource string) string {
	resp, err := http.Get(fmt.Sprintf("%s/%s", GAIA_HOST, resource))

	if err != nil {
		log.Print(err)

		return string("")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func CollectMetrics() (float64, float64, float64) {
	var latest_block_height float64 = -1
	var desync float64 = -1
	var n_peers float64 = -1

	net_info := readAPI("net_info")
	status := readAPI("status")

	if net_info != "" {
		n_peers = gjson.Get(net_info, "result.n_peers").Float()

		log.Printf(fmt.Sprintf("peers: %f", n_peers))
	}
	if status != "" {
		latest_block_height = gjson.Get(status, "result.sync_info.latest_block_height").Float()
		latest_block_time := gjson.Get(status, "result.sync_info.latest_block_time")

		timeGaia, _ := time.Parse(time.RFC3339Nano, latest_block_time.Str)
		desync = float64(time.Now().Unix() - timeGaia.Unix())

		log.Printf(fmt.Sprintf("latest_block_height: %f", latest_block_height))
		log.Printf(fmt.Sprintf("desync: %f", desync))
	}

	return latest_block_height, desync, n_peers
}
