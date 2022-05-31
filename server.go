package main

import (
	"dataloader/dataloader"
	"dataloader/entity"
	"dataloader/graph"
	"dataloader/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger = db.Logger.LogMode(logger.Info)

	err = db.AutoMigrate(&entity.User{}, &entity.Todo{})
	if err != nil {
		panic(err)
	}

	loader := dataloader.New(db)

	dataloaderSrv := dataloader.Middleware(loader, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					DB: db,
				},
			},
		),
	))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
