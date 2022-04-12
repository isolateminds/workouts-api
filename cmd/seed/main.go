package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dannyvidal/workout-service/internal/fitness"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBNAME = "workouts.db"
	FE_CSV = "./fitness_exercises.csv"
)

func createWorkoutTable(db *sql.DB) {
	SQL := `CREATE TABLE workouts (
		"id" integer NOT NULL PRIMARY KEY,
		"bodypart"  TEXT,
		"equipment" TEXT,
		"gif"       TEXT,
		"name"      TEXT,
		"target"    TEXT
	);`
	statement, err := db.Prepare(SQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}
func insertWorkout(db *sql.DB, workout *fitness.FitnessExercise) {
	SQL := `INSERT INTO workouts(id, bodypart, equipment, gif, name, target) VALUES(?,?,?,?,?,?)`
	statement, err := db.Prepare(SQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec(
		workout.Id,
		workout.BodyPart,
		workout.Equipment,
		fmt.Sprintf("/api/gifs/%v.gif", workout.Id),
		workout.Name,
		workout.Target,
	)
}
func main() {

	workouts := fitness.ParseFE(FE_CSV)
	os.Remove(DBNAME)
	file, err := os.Create(DBNAME)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	sqlitedb, err := sql.Open("sqlite3", DBNAME)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sqlitedb.Close()
	createWorkoutTable(sqlitedb)
	for _, workouts := range workouts {
		insertWorkout(sqlitedb, &workouts)
	}
}
