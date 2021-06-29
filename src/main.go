package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
	suply "../src/calculator/supply"
	"../src/calculator/applying"
	"../src/utills"
)

func main() {
	r := mux.NewRouter()
	supply := suply.CalculatorUseNew(utills.Timeouts.ContextTimeout)
	applying.NewCalculatorHandler(r, supply)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         utills.ServerUrl,
		WriteTimeout: utills.Timeouts.WriteTimeout,
		ReadTimeout:  utills.Timeouts.ReadTimeout,
	}
	log.Info().Msgf("Server started at " + utills.ServerUrl)
	log.Error().Msgf(srv.ListenAndServe().Error())

}
