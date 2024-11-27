package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"github.com/swallowstalker/online-book-store/modules/bookstore/handler"
	"github.com/swallowstalker/online-book-store/modules/bookstore/middleware"
	"github.com/swallowstalker/online-book-store/modules/bookstore/repository"
	"github.com/swallowstalker/online-book-store/modules/bookstore/repository/db"
	"github.com/swallowstalker/online-book-store/modules/bookstore/service"
)

type Config struct {
	AppPort    int    `env:"APP_PORT,default=8080"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
}

func main() {
	_ = godotenv.Load(".env")

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		panic(err)
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		panic(err)
	}

	querier := db.New(pool)
	repoWrapper := repository.NewDbWrapperRepo(querier)
	userService := service.NewUserService(repoWrapper)
	bookService := service.NewBookService(repoWrapper)
	orderService := service.NewOrderService(repoWrapper)
	h := handler.NewHandler(userService, bookService, orderService)
	m := middleware.NewAuthMiddleware(repoWrapper)

	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/v1/users", h.CreateUser)
	router.HandlerFunc(http.MethodGet, "/v1/books", h.GetBooks)
	router.HandlerFunc(http.MethodPost, "/v1/orders", m.CheckTokenMiddleware(h.CreateOrder))
	router.HandlerFunc(http.MethodGet, "/v1/orders", m.CheckTokenMiddleware(h.GetMyOrders))

	fmt.Println("server started")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.AppPort), router); err != nil {
		fmt.Println("server stopped")
		panic(err)
	}
}
