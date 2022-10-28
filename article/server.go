package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/sumit-dhull-97/assignment/article/service"
	"github.com/sumit-dhull-97/assignment/article/store/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sumit-dhull-97/assignment/article/graph"
	"github.com/sumit-dhull-97/assignment/article/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.Default()
	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler())
	router.Run()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func graphqlHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		pg := postgres.GetDBConnection(&ctx)
		auth := &postgres.Article{DB: pg}

		serv := &service.ArticleService{Store: auth}

		h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Service: serv}}))
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	}
}
