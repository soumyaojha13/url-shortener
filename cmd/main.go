package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"url-shortener/internal/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/metrics"
	"url-shortener/internal/repository"
	"url-shortener/internal/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Postgres
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	// Retry DB connection
	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Println("Waiting for DB...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("DB not connected:", err)
	}

	// Redis
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis not connected:", err)
	}

	// Repositories
	pgRepo := &repository.PostgresRepo{DB: db}
	redisRepo := &repository.RedisRepo{Client: rdb}

	// Usecase
	urlUsecase := usecase.NewURLUsecase(pgRepo, redisRepo)

	// Router
	r := gin.Default()

	// Metrics
	metrics.Init(r)

	// Routes
	h := handler.NewHandler(urlUsecase)
	r.POST("/shorten", h.CreateURL)
	r.GET("/:code", h.Redirect)

	// Run server
	r.Run(":" + cfg.AppPort)
}
