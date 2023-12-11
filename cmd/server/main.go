package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rafaelsouzaribeiro/9-API/configs"
	_ "github.com/rafaelsouzaribeiro/9-API/docs"
	"github.com/rafaelsouzaribeiro/9-API/internal/entity"
	"github.com/rafaelsouzaribeiro/9-API/internal/infra/database"
	"github.com/rafaelsouzaribeiro/9-API/internal/infra/webservice/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           api go standard
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Rafael Fernando
// @contact.url    https://github.com/rafaelsouzaribeiro
// @contact.email  rafaelribeirosouza86@gmail.com

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDb := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDb)

	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Se a aplicação cai ele não deixa cair
	router.Use(middleware.Recoverer)
	//router.Use(LogRequest)
	router.Use(middleware.WithValue("jwt", config.TokenAuth))
	router.Use(middleware.WithValue("jwtExpiresin", config.JwtExpiresIn))

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(&config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	//http.HandleFunc("/products", productHandler.CreateProduct)

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/generate_token", userHandler.GetJwt)

	http.ListenAndServe(":8080", router)
}

// Request -> middleware -> Handler->respose

// func LogRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("Request: %s %s", r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }
