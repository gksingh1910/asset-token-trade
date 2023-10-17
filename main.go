package main

import (
	"fmt"
	"net/http"

	config "github.com/gksingh1910/asset-token-trade/config"
	"github.com/go-chi/chi"
)

type AssetToken struct {
	assetId        string
	assetName      string
	assetTokenCode string
}

func main() {
	config.Init()
	r := chi.NewRouter()
	fmt.Printf("Server host is %v ", config.Get().SrvConfig.Host)
	srv := http.Server{
		Addr:    config.Get().SrvConfig.Host,
		Handler: r,
	}

	r1 := chi.NewRouter()
	r1.Get("/hello", func(w http.ResponseWriter, r *http.Request) { fmt.Printf("Welcome Darling !!") })

	// Mount example
	r.Mount("/v1", r1)

	// Route example
	r.Route("/app", func(r chi.Router) {
		r.Get("/api1", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("api1 response")) })
		r.Get("/api2", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("api2 response")) })
		r.Get("/api3", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("api3 response")) })
	})

	r.Route("/dapp", func(r chi.Router) {
		r.Get("/api1", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("api1 response")) })
		r.Get("/api2", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("api2 response")) })
		r.Get("/api3", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("api3 response")) })
	})

	//Versioning with route and mount
	r.Route("/v2", func(r chi.Router) {
		r.Mount("/user", userHandler())
	})

	srv.ListenAndServe()
}

func userHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Lists all users")) })
	r.Post("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Add users")) })
	r.Route("/{user_id}", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Gets user by id")) })
		r.Put("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Put user by id")) })
	})
	return r
}
