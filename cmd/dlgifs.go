//Downloads workout gifs
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dannyvidal/workout-service/internal/fitness"
)

const FE_CSV = "./fitness_exercises.csv"

func GetGifs(fe *[]fitness.FitnessExercise, dirname string) <-chan interface{} {

	completed := make(chan interface{})
	go func() {
		defer close(completed)
		for _, fe := range *fe {

			res, err := http.Get(fe.GifURL)
			if err != nil {
				fmt.Printf("<%v>: %v\n", res.StatusCode, err.Error())
			}

			file, err := os.Create(fmt.Sprintf("%v/%v.gif", dirname, fe.Id))
			if err != nil {
				fmt.Printf("%v\n", err.Error())
			}
			w, err := io.Copy(file, res.Body)
			if err != nil {
				fmt.Printf("bytes<%v>%v\n", w, err.Error())
			}
			defer res.Body.Close()
			defer file.Close()
		}
	}()
	return completed
}
func main() {
	fe := fitness.ParseFE(FE_CSV)
	fmt.Println("Downloading all gifs from csv url(s)...")
	<-GetGifs(&fe, "./gifs")
	fmt.Println("Done.")

}
