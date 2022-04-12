package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
		"equipment" TEXT,
		"gif"    TEXT,
		"name"      TEXT,
		"target"    TEXT
	);`
	statement, err := db.Prepare(SQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}
func insertWorkout(db *sql.DB, workout *fitness.FitnessExercise, gifAbs string) {
	SQL := `INSERT INTO workouts(id, equipment, gif, name, target) VALUES(?,?,?,?,?)`
	statement, err := db.Prepare(SQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec(
		workout.Id,
		workout.Equipment,
		fmt.Sprintf("%v/%v.gif", gifAbs, workout.Id),
		workout.Name,
		workout.Target,
	)
}
func main() {
	gifAbs, err := filepath.Abs("./gif")
	if err != nil {
		log.Fatal(err.Error())
	}

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
		insertWorkout(sqlitedb, &workouts, gifAbs)
	}
}
