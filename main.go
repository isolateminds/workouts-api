package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/dannyvidal/workout-service/internal/fitness"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const DBNAME = "workouts.db"

func scanQuery(exercises *sql.Rows) []fitness.FitnessExercise {
	fe := []fitness.FitnessExercise{}
	for exercises.Next() {
		var e fitness.FitnessExercise
		err := exercises.Scan(&e.Id, &e.BodyPart, &e.Equipment, &e.GifURL, &e.Name, &e.Target)
		if err != nil {
			panic(err)
		}
		fe = append(fe, e)
	}
	return fe
}
func queryWorkouts(db *sql.DB, search string) []fitness.FitnessExercise {
	SQL := `SELECT * from workouts WHERE name LIKE ?`
	statement, err := db.Prepare(SQL)
	if err != nil {
		panic(err)
	}
	exercises, err := statement.Query(fmt.Sprintf("%%%v%%", search))
	if err != nil {
		panic(err)
	}
	return scanQuery(exercises)
}

func main() {
	gifdir, err := filepath.Abs("./gifs")
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := sql.Open("sqlite3", DBNAME)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	router.Static("/api/gifs", gifdir)

	router.GET("/api/workouts", func(ctx *gin.Context) {
		search := ctx.Query("search")
		fmt.Println(search)
		exercises := queryWorkouts(db, search)
		ctx.JSON(200, exercises)
	})

	router.Run(":8080")

}
