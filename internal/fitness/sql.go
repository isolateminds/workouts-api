package fitness

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db = NewDB()

func scanQuery(exercises *sql.Rows) []FitnessExercise {

	fe := []FitnessExercise{}
	for exercises.Next() {
		var e FitnessExercise
		err := exercises.Scan(&e.Id, &e.BodyPart, &e.Equipment, &e.GifURL, &e.Name, &e.Target)
		if err != nil {
			panic(err)
		}
		fe = append(fe, e)
	}
	return fe
}
func QueryWorkouts(search string) []FitnessExercise {

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

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "workouts.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
