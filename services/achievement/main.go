package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"achievement-leveling/achievement/config"
	"achievement-leveling/achievement/handler"
	"achievement-leveling/achievement/migration"
	"achievement-leveling/achievement/repository"
	"achievement-leveling/achievement/service"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	var db *gorm.DB
	var err error
	for i := 0; i < 30; i++ {
		db, err = gorm.Open(mysql.Open(cfg.DBDsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("waiting for database... attempt %d/30", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("connected to database")

	migration.AutoMigrate(db)

	userRepo := repository.NewUserRepo(db)
	achRepo := repository.NewAchievementRepo(db)
	progressRepo := repository.NewProgressRepo(db)

	userSvc := service.NewUserService(userRepo)
	achSvc := service.NewAchievementService(achRepo)
	levelSvc := service.NewLevelingService(progressRepo, userRepo, achRepo)

	userHandler := handler.NewUserHandler(userSvc)
	achHandler := handler.NewAchievementHandler(achSvc)
	progressHandler := handler.NewProgressHandler(levelSvc)

	r := chi.NewRouter()
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"achievement"}`))
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/users", userHandler.UpsertUser)
		r.Get("/users/{id}", userHandler.GetUser)

		r.Get("/achievements", achHandler.GetAll)
		r.Get("/achievements/{id}", achHandler.GetByID)
		r.Get("/categories", achHandler.GetCategories)

		r.Post("/progress/complete", progressHandler.CompleteAchievement)
		r.Get("/progress", progressHandler.GetUserProgress)
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("achievement service listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
