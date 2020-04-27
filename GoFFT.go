package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/mjibson/go-dsp/fft"
)

func main() {
	data := make([]float64, 0)
	strfftvalabsreal := make([]string, 0)
	var datanew float64
	var datastring string

	// Open the file
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	y := 0
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//
		datanew, err = strconv.ParseFloat(record[1], 64)
		data = append(data, datanew)
		y++
	}

	fftval := fft.FFTReal(data)

	// Take absolute value of real part of compex number

	for i := 0; i < len(fftval); i++ {
		datastring = fmt.Sprintf("%.2f", math.Abs(real(fftval[i])))
		strfftvalabsreal = append(strfftvalabsreal, datastring)
	}

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(strfftvalabsreal)

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
