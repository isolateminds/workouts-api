package fitness

import (
	"encoding/csv"
	"os"
)

type FitnessExercise struct {
	BodyPart  string `json:"bodypart"`
	Equipment string `json:"equipment"`
	GifURL    string `json:"gifurl"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
}

// parses the fitness csv file and returns an array
func ParseFE(path string) []FitnessExercise {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	fitnessExercises := make([]FitnessExercise, len(records)-1)
	for idx, record := range records {
		// 0th index is csv headers
		if idx != 0 {
			fitnessExercises[idx-1] = FitnessExercise{
				record[0],
				record[1],
				record[2],
				record[3],
				record[4],
				record[5],
			}
		}

	}
	return fitnessExercises
}
