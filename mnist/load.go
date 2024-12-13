package mnist

import (
	"archive/zip"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

func UnzipAndLoad(zipFileName, fileName string) (testCases [][]float64, labels []int, err error) {

	// Check if the file exists
	if _, err := os.Stat(fileName); err != nil {

		// If the file does not exist
		if os.IsNotExist(err) {

			log.Println("Unzipping the file...")

			// Open the zip file
			zipReader, err := zip.OpenReader(zipFileName)
			if err != nil {
				return nil, nil, err
			}
			defer zipReader.Close()

			// Get the csv file in the zip file
			file := zipReader.File[0]

			// Open the file
			rc, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			// Create a new file
			f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return nil, nil, err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return nil, nil, err
			}
		} else {
			return nil, nil, err
		}
	}

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Create a new csv reader
	csvReader := csv.NewReader(file)

	// Read the csv file
	for {

		// Read a line from the csv file
		testCase, err := csvReader.Read()
		if err != nil {

			// If we reach the end of the file, break the loop
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, nil, err
		}

		// Parse the label
		label, err := strconv.Atoi(testCase[0])
		if err != nil {
			return nil, nil, err
		}

		labels = append(labels, label)

		// Cut the label from the image
		image := testCase[1:]

		// Make a float array of the pixels of the image
		testCasePixels := make([]float64, 28*28)

		// Read the pixels of the image
		for i, pixelString := range image {

			// Parse the pixel
			pixelFloat, err := strconv.ParseFloat(pixelString, 64)
			if err != nil {
				return nil, nil, err
			}

			// Normalize the pixel
			normalizedPixel := pixelFloat / 255

			if normalizedPixel > 0.5 {
				testCasePixels[i] = 1
			} else {
				testCasePixels[i] = 0
			}
		}

		// Append the image to the test cases
		testCases = append(testCases, testCasePixels)
	}

	return testCases, labels, nil
}
