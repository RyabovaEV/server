package main

import (
	"fmt"
	"net/http"
	"server/configs"
	"server/internal/auth"
	"server/internal/link"
	"server/pkg/db"
	"server/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// repositories
	linkRepositories := link.NewLinkRepository(db)

	// Handler
	auth.NewAuthHendler(router, auth.AuthHendlerDeps{
		Config: conf,
	})
	link.NewLinkHendler(router, link.LinkHendlerDeps{
		LinkRepository: linkRepositories,
	})

	//Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()

}
