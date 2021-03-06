package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
	"github.com/jasonlvhit/gocron"

	"github.com/paul67567/marketplace"
	"github.com/paul67567/settings"
)

var (
	APP_SETTINGS = settings.GetSettings()
)

func runCrons() {
	if !APP_SETTINGS.Debug {
		marketplace.StartTransactionsCron()
		marketplace.StartWalletsCron()
		marketplace.StartStatsCron()
		marketplace.StartSERPCron()
		marketplace.StartMessageboardCron()
	}

	marketplace.StartCurrencyCron()

	<-gocron.Start()

}

func runWebserver() {
	// Root router
	rootRouter := web.New(marketplace.Context{})

	rootRouter.OptionsHandler((*marketplace.Context).OptionsHandler)

	rootRouter.Middleware(web.StaticMiddleware("public"))
	// Marketplace router
	marketplace.ConfigureRouter(rootRouter.Subrouter(marketplace.Context{}, "/"))
	// Start HTTP server

	startHttpServer := func() {
		address := fmt.Sprintf("%s:%s", APP_SETTINGS.Host, APP_SETTINGS.Port)
		println(fmt.Sprintf("Running a HTTP server or, %s", address))
		http.ListenAndServe(address, rootRouter)
	}
	// Start HTTPs server
	startHttpServer()

}

func runServer() {
	go runCrons()
	runWebserver()
}
