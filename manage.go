package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"demo/accounts"
	_ "demo/admin_config" // Register admin models
	"demo/appconfig"
	"demo/blog"
	_ "demo/migrations" // Register migrations
	"demo/posts"

	framework_admin "github.com/anuragcarret/djang-drf-go/admin"
	"github.com/anuragcarret/djang-drf-go/contrib/auth"
	auth_urls "github.com/anuragcarret/djang-drf-go/contrib/auth/urls"
	"github.com/anuragcarret/djang-drf-go/core/apps"
	"github.com/anuragcarret/djang-drf-go/core/management"
	"github.com/anuragcarret/djang-drf-go/core/settings"
	"github.com/anuragcarret/djang-drf-go/core/urls"
	"github.com/anuragcarret/djang-drf-go/orm/db"
	cmd "github.com/anuragcarret/djang-drf-go/orm/management/commands"
)

func main() {
	// 1. Explicitly register all apps
	apps.Apps.Register(&auth.AuthApp{})
	apps.Apps.Register(&appconfig.DemoApp{})
	apps.Apps.Register(&accounts.AccountsApp{})
	apps.Apps.Register(&posts.PostsApp{})
	apps.Apps.Register(&blog.BlogApp{})

	// 2. Populate App Registry (triggers Ready and model associations)
	if err := apps.Apps.Populate(); err != nil {
		log.Fatalf("Failed to populate apps: %v", err)
	}

	// 3. Load Settings
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	conf, err := settings.LoadChain("config/settings", env)
	if err != nil {
		log.Fatalf("Failed to load settings: %v", err)
	}
	settings.Initialize(conf)
	fmt.Printf("Settings loaded for environment: %s (Debug: %v)\n", env, conf.Debug)

	// 4. Database Connection (from settings)
	dbConf := conf.Databases["default"]
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name)
	database, err := db.NewDB(dbConf.Engine, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer database.Close()

	// 4. Register Commands
	management.Commands.Register(cmd.NewMigrateCommand(database))
	management.Commands.Register(cmd.NewMakemigrationsCommand(database))
	management.Commands.Register(cmd.NewCreateSuperuserCommand[*accounts.Account](database))
	management.Commands.Register(&RunserverCommand{DB: database})

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run manage.go [migrate|makemigrations|createsuperuser|runserver]")
		return
	}

	if err := management.Execute(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

type RunserverCommand struct {
	DB *db.DB
}

func (c *RunserverCommand) Name() string { return "runserver" }
func (c *RunserverCommand) Help() string { return "Start the HTTP server" }
func (c *RunserverCommand) Run(ctx context.Context, args []string) error {
	router := urls.NewRouter()

	// Register App Routers
	router.Include("/admin", framework_admin.DefaultSite.URLs(c.DB), "admin")
	router.Include("/auth", auth_urls.RegisterRoutes[*accounts.Account](c.DB), "auth")
	router.Include("/accounts", accounts.RegisterRoutes(c.DB), "accounts")
	router.Include("/posts", posts.RegisterRoutes(c.DB), "posts")

	fmt.Println("Starting modular demo server on :8080...")
	handler := auth.AuthenticationMiddleware(router)
	return http.ListenAndServe(":8080", handler)
}
