package scores

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
)

func Subscribe(channel chan<- ScoreEvent) {
	resp, _ := http.Get("http://live-test-scores.herokuapp.com/scores")
	reader := bufio.NewReader(resp.Body)

	for {
		line, _ := reader.ReadString('\n')
		trimmed := strings.TrimSpace(line)
		length := len(trimmed)

		if length > 0 {
			parts := strings.Split(trimmed, ": ")

			if parts[0] != "event" {
				var score ScoreEvent
				json.Unmarshal([]byte(parts[1]), &score)
				channel <- score
			}
		}
	}
}
