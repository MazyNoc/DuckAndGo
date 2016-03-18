package main

import (
	"log"
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"annat.nu/data/metric"
	"annat.nu/data/sample"
	"annat.nu/persistence"
	"annat.nu/data/upload"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, time to hunt ducks (ps. don't use the path %s! )", r.URL.Path[1:])
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := metric.Get("http://" + r.Host)
	addHeaders(w)
	fmt.Fprintf(w, metrics.Json())
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	sample := sample.Init()
	fmt.Fprintf(w, sample.Json())
}

func nodepingHandler(w http.ResponseWriter, r *http.Request) {
	sample := persistence.GetNodePing()
	addHeaders(w)
	fmt.Fprintf(w, sample.Json())
}

func uploadNodepingHandler(w http.ResponseWriter, r *http.Request) {
	var temp upload.Nodeping
	json.NewDecoder(r.Body).Decode(&temp)
	nodePing := sample.Sample{"nodeping", "status", sample.Datum{sample.ConvertStatus(temp.Value), time.Now()}}
	persistence.SetNodePing(nodePing)
}

func addHeaders(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
}

func main() {
	log.Print("starting")

	http.HandleFunc("/", handler)
	http.HandleFunc("/metrics", metricsHandler)
	http.HandleFunc("/sample", sampleHandler)
	http.HandleFunc("/nodeping", nodepingHandler)

	http.HandleFunc("/upload/nodeping", uploadNodepingHandler)
	http.ListenAndServe(":8080", nil)
}