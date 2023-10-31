package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/gksingh1910/asset-token-trade/config"
	route "github.com/gksingh1910/asset-token-trade/route"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Couldn't intiate the config file %v", err.Error())
	}
	// Check if config is accessible
	fmt.Printf("\n %v", config.Get().HTTPServerConfig.Host)
	x := fmt.Sprintf("%v:%v", config.Get().HTTPServerConfig.Host, config.Get().HTTPServerConfig.Port)
	print(x)
	r := chi.NewRouter()
	srv := http.Server{
		Addr:    fmt.Sprintf("%v:%v", config.Get().HTTPServerConfig.Host, config.Get().HTTPServerConfig.Port),
		Handler: r,
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hELLO")) })

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/user", route.)
	})

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server")
	}

}
