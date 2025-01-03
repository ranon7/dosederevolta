package main

import (
	"context"
	"fmt"
	"net/http"
)

type API struct{}

func (api *API) Routes(ctx context.Context) http.Handler {
	routes := http.NewServeMux()

	routes.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	return routes
}
