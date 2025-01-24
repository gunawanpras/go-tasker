package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gunawanpras/go-tasker/config"
	"github.com/gunawanpras/go-tasker/tasker"
)

func main() {
	conf := config.LoadConfig("./config/config.yaml")
	log.Println("Load configuration...")

	scheduler := tasker.NewScheduler(conf)
	scheduler.AddTask("Backup production db", time.Duration(conf.TaskScheduler.BackProductionDb.Interval)*time.Second, func() {
		log.Println("Starting database backup...")
		log.Println("Database backup completed successfully.")
	})
	scheduler.AddTask("Clean up expired session", time.Duration(conf.TaskScheduler.CleanUp.Interval)*time.Second, func() {
		log.Println("Starting session cleanup...")
		log.Println("Session cleanup completed.")
	})
	scheduler.AddTask("Cache Eviction", time.Duration(conf.TaskScheduler.CacheEviction.Interval)*time.Second, func() {
		log.Println("Clearing old cache...")
		log.Println("Cache cleared.")
	})
	scheduler.AddTask("Generate Daily Reports", time.Duration(conf.TaskScheduler.GenerateReport.Interval)*time.Second, func() {
		log.Println("Generating daily reports...")
		log.Println("Daily reports generated.")
	})

	go scheduler.Run()

	// Start HTTP server
	log.Println("Starting server on port", conf.Server.Port)
	if err := http.ListenAndServe(":"+conf.Server.Port, nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
