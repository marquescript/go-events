package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marquescript/go-events/config"
	_ "github.com/marquescript/go-events/docs"
	"github.com/marquescript/go-events/internal/infra/database"
	"github.com/marquescript/go-events/internal/infra/factory"
	"github.com/marquescript/go-events/internal/infra/http/middlewares"
	"github.com/marquescript/go-events/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title						Go Events API
// @version					1.0
// @description				API para gerenciamento de eventos
// @host						localhost:8080
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("erro ao carregar configurações: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("erro ao abrir conexão com banco de dados: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("erro ao conectar com banco de dados: %v", err)
	}

	log.Println("Conexão com banco de dados estabelecida com sucesso!")

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("erro ao executar migrations: %v", err)
	}

	log.Println("Migrations executadas com sucesso!")

	userService := service.NewUserService(database.NewUser(db))

	eventFactory := factory.NewEventFactory(db)
	userFactory := factory.NewUserFactory(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", cfg.JWTExpiresIn))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.Route("/events", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Use(middlewares.VerifyUserMiddleware(userService))

		r.Post("/", eventFactory.Handler.CreateEvent)
		r.Get("/{id}", eventFactory.Handler.FindEvent)
		r.Get("/", eventFactory.Handler.FindAllEvents)
		r.Put("/{eventId}", eventFactory.Handler.UpdateEvent)
		r.Delete("/{eventId}", eventFactory.Handler.DeleteEvent)

	})

	r.Post("/users", userFactory.Handler.CreateUser)
	r.Post("/sign-in", userFactory.Handler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	fmt.Println("docs in http://localhost:8080/docs/index.html")
	http.ListenAndServe(":8080", r)
}
