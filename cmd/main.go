package main

import (
	"net/http"

	"example.com/m/v2/internal/api/rest"
	stats2 "example.com/m/v2/internal/business/stats"
	"example.com/m/v2/internal/clients/nsq"
	"example.com/m/v2/internal/repository"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	nsqTopic  = "nsq_test"
	nsqTarget = "127.0.0.1:4151"
	appAddr   = ":80"
	logLevel  = log.DebugLevel
)

func main() {
	log.SetLevel(logLevel)

	nsq, err := nsq.NewClient(nsqTopic, nsqTarget)
	if err != nil {
		panic(err)
	}
	defer nsq.Stop()

	repoWallet := repository.NewRepo()
	manStats := stats2.NewManager(*repoWallet)
	handler, err := rest.NewHandler(manStats)
	if err != nil {
		panic(err)
	}

	nsq.StartConsume(manStats.EventHandler)

	r := mux.NewRouter()
	r.HandleFunc("/stats/wallets", handler.StatsHandler).Methods(http.MethodGet)

	log.Infof("app started on: %s", appAddr)
	defer func() {
		log.Infof("app finished on: %s", appAddr)
	}()
	err = http.ListenAndServe(appAddr, r)
	panic(err)
}
