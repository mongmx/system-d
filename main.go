package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mongmx/system-d/application/domain/auth"
	"github.com/mongmx/system-d/application/domain/member"
	"github.com/mongmx/system-d/application/infrastructure/postgres"
	"github.com/mongmx/system-d/application/middleware"
	"golang.org/x/sync/errgroup"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	//env     = "development"
	sslMode = "disable"
	dbPort  = "5432"
	dbHost  = "localhost"
	dbUser  = "root"
	dbPass  = "root"
	dbName  = "system-d"

	g errgroup.Group
)

func main() {
	db, err := connectPostgres()
	must(err)
	defer func() {
		err := db.Close()
		must(err)
	}()

	serverAPI := &http.Server{
		Addr:         ":8080",
		Handler:      routerAPI(db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	serverMetrics := &http.Server{
		Addr:         ":8081",
		Handler:      routerMetrics(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		log.Println("API server listen on :8080")
		return serverAPI.ListenAndServe()
	})
	g.Go(func() error {
		log.Println("Metrics server listen on :8081")
		return serverMetrics.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func routerAPI(db *sql.DB) http.Handler {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(ginprom.PromMiddleware(nil))
	e.Use(middleware.TraceMiddleware())

	memberRepo, err := postgres.NewMemberRepository(db)
	must(err)
	memberService, err := member.NewService(memberRepo)
	must(err)

	authRepo, err := postgres.NewMemberRepository(db)
	must(err)
	authService, err := auth.NewService(authRepo, nil)
	must(err)

	member.Routes(e, memberService, authService)
	auth.Routes(e, authService)

	return e
}

func routerMetrics() http.Handler {
	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	return e
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connectPostgres() (*sql.DB, error) {
	conn := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		dbName, dbUser, dbPass, dbHost, dbPort, sslMode,
	)
	return sql.Open("postgres", conn)
}

func connectRedis()  {
	
}