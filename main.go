package main

import (
	"go-grow/config"
	handler "go-grow/delivery/http"
	"go-grow/repository"
	"go-grow/router"
	"go-grow/usecase"
	"log"
)

func main() {
    config.GetEnvConfig()

    /*
    Config for JSON

    config.GetJSONConfig()
    g := viper.Get("server.port")
    fmt.Println(g)
    
    */

    databaseConnection, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("could not initialize database connection: %s", err)
    }

    baseRepository := repository.NewBaseRepository(databaseConnection)

    userUsecase := usecase.NewUserUsecase(baseRepository)
    userHandler := handler.NewUserHandler(userUsecase)

    router.InitRouter(userHandler)
    router.Start()
}
