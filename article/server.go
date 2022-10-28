package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/sumit-dhull-97/assignment/article/service"
	"github.com/sumit-dhull-97/assignment/article/store/postgres"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sumit-dhull-97/assignment/article/graph"
	"github.com/sumit-dhull-97/assignment/article/graph/generated"
)

const defaultPort = ":3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	go startListening()

	router := gin.Default()
	router.GET("/", playgroundHandler())
	router.POST("/query", graphqlHandler())
	router.Run(port)

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

func startListening() {
	host := os.Getenv("DD_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	var conninfo string = "postgres://goland:goland@%s:5432/postgres?sslmode=disable"
	conninfo = fmt.Sprintf(conninfo, host)

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	listener := pq.NewListener(conninfo, 10*time.Second, time.Minute, reportProblem)
	err := listener.Listen("events")
	if err != nil {
		panic(err)
	}

	fmt.Println("Start monitoring PostgreSQL...")
	for {
		waitForNotification(listener)
	}
}

func waitForNotification(l *pq.Listener) {
	select {
	case n := <-l.Notify:
		fmt.Println("Received data from channel [", n.Channel, "] :")

		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, []byte(n.Extra), "", "\t")
		if err != nil {
			fmt.Println("Error processing JSON: ", err)
			return
		}

		jsonString := prettyJSON.String()

		title := gjson.Get(jsonString, "data.title").String()
		userId := gjson.Get(jsonString, "data.user_id").String()
		hashtags := gjson.Get(jsonString, "data.hashtags").Array()

		fmt.Printf("%s \n%s \n%s \n", title, userId, hashtags)

		return
	case <-time.After(90 * time.Second):
		fmt.Println("Received no events for 90 seconds, checking connection")
		go func() {
			l.Ping()
		}()
		return
	}
}
