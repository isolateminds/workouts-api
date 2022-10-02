package main

import (
	"log"
	"path/filepath"

	"github.com/isolateminds/workout-api/internal/fitness"
	"github.com/gin-gonic/gin"
)

const DBNAME = "workouts.db"

func main() {
	// Get absolute file path of workout gifs
	gifs, err := filepath.Abs("./gifs")
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	router.Static("/api/gifs", gifs)
	router.GET("/api/workouts", fitness.SearchWorkouts)
	router.Run(":8080")

}
