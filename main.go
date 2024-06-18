package main

import (
	"database/sql"
	_ "github.com/ISNewton/database/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type apiConfig struct {
	DB *internal.Queries
}

func main() {

	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load Port
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Set PORT environment variable")
	}

	//DB
	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == "" {
		log.Fatal("Set DATABASE_URL environment variable")
	}

	connection, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	queries := database.New(connection)

	apiCfg := apiConfig{
		DB: queries,
	}

	//router
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	serverRouter := chi.NewRouter()
	serverRouter.Get("/health", handleReadiness)
	serverRouter.Get("/error", handleError)

	router.Mount("/server", serverRouter)

	userRouter := chi.NewRouter()
	userRouter.Post("/", apiCfg.handleCreateUser)

	router.Mount("/users", userRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Server up and running")

	serverError := server.ListenAndServe()

	if serverError != nil {
		log.Fatal(err)
	}

}
