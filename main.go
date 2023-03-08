package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	log.Debugf("application is starting and listening on 8080")

	log.Debug("starting server")

	// testing log outputs
	log.Trace("testing trace log output")
	log.Warn("testing warning log output")
	log.Error("testing error log output")

	log.Error(http.ListenAndServe(":8080", http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {
			log.Infof("%s, %+v, %s", r.URL.Path, r.Header, r.RemoteAddr)

			hostname, _ := os.Hostname()
			json.NewEncoder(w).Encode(map[string]any{
				"current time in UTC": time.Now().UTC(),
				"hostname is":         hostname,
			})
		},
	)))
}
