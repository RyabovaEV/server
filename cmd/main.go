package main

import (
	"fmt"
	"net/http"
	"server/configs"
	"server/internal/auth"
	"server/internal/link"
	"server/internal/user"
	"server/pkg/db"
	"server/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	// repositories
	linkRepositories := link.NewLinkRepository(db)
	userRepository := user.NewRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)

	// Handler
	auth.NewAuthHendler(router, auth.AuthHendlerDeps{
		Config:  conf,
		Service: authService,
	})
	link.NewLinkHendler(router, link.LinkHendlerDeps{
		LinkRepository: linkRepositories,
		Config:         conf,
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
