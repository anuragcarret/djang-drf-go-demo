package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"demo/apps/user_app"

	"github.com/anuragcarret/djang-drf-go/framework/admin"
	"github.com/anuragcarret/djang-drf-go/framework/core"
	"github.com/anuragcarret/djang-drf-go/framework/drf"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run manage.go [runserver|shell]")
		return
	}

	command := os.Args[1]

	switch command {
	case "runserver":
		runServer()
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func runServer() {
	// Initialize all registered apps
	if err := core.InitializeAllApps(); err != nil {
		log.Fatalf("Critical error during app initialization: %v", err)
	}

	fmt.Println("Starting django-drf-go development server at http://127.0.0.1:8000/")
	fmt.Println("Quit the server with CONTROL-C.")

	// Registry info
	fmt.Printf("Registered Apps: %v\n", core.ListApps())
	fmt.Printf("Admin Status: %s\n", admin.RenderDashboard())

	// Basic Router
	router := drf.NewRouter()

	// Register UserViewSet from user_app
	userViewSet := &user_app.UserViewSet{}
	router.Register("/api/users", userViewSet)

	log.Fatal(http.ListenAndServe(":8000", router))
}
