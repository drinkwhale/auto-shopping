package server

import (
	"log"
	"net/http"
	"os"

	"github.com/drinkwhale/auto-shopping/internal/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

// Run 함수는 의존성을 주입받아 서버를 시작합니다.
func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{}
	c := graph.Config{Resolvers: resolver}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 GraphQL Playground available at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
