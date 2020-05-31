package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/oryono/banking-go/graph"
	"github.com/oryono/banking-go/graph/generated"
	_ "log"
	_ "net/http"
	_ "os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/oryono/banking-go/graph"
	_ "github.com/oryono/banking-go/graph/generated"
)

const defaultPort = "8081"

var DB *gorm.DB

func initDB()  {
	var err error
	postgresUser := "postgres"
	postgresPassword := "postgres"
	postgresDb := "banking_graph_dev"
	postgresHost := "localhost"
	postgresPort := "5432"

	connection := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", postgresHost, postgresPort, postgresUser, postgresDb, postgresPassword)
	db, err := gorm.Open("postgres", connection)

	if err != nil {
		println(err)
		panic("Failed to connect to the database")
	}

	DB = db
	db.LogMode(true)
	db.AutoMigrate()

}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: DB}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	initDB()
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}