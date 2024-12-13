package main

import (
	"fmt"
	"log"

	"neuralNetwork/mnist"
	"neuralNetwork/model"
	"neuralNetwork/utlis"
)

const (
	learningRate = 0.1
	epoch        = 5
)

const (
	trainDataFileName    = "mnist/mnist_train.csv"
	testDataFileName     = "mnist/mnist_test.csv"
	trainZipDataFileName = "mnist/mnist_train.csv.zip"
	testZipDataFileName  = "mnist/mnist_test.csv.zip"
)

func main() {

	// Initialize the network with 784 input neurons, 32 hidden neurons, and 10 output neurons
	network := model.NewNetwork(
		[]int{
			784,
			28,
			10,
		},
		learningRate,
		utlis.Sigmoid,
		utlis.SigmoidDerivative,
	)

	// Load the training data
	log.Println("Loading training data...")
	testCases, labels, err := mnist.UnzipAndLoad(trainZipDataFileName, trainDataFileName)
	if err != nil {
		log.Fatal(err)
	}

	// Train the network
	log.Println("Training the network...")
	network.Train(testCases, labels, epoch)

	// Load the test data
	log.Println("Loading test data...")
	testCases, labels, err = mnist.UnzipAndLoad(testZipDataFileName, testDataFileName)
	if err != nil {
		log.Fatal(err)
	}

	successes := 0

	// Loop through the test testCases
	log.Println("Testing the network...")
	for i := 0; i < len(testCases); i++ {

		testCase := testCases[i]

		// Set the input values
		network.SetInput(testCase)

		// Get the predictions of the network
		prediction := network.GetPredictions()

		// Get the best result
		bestPred := utlis.GetBestResult(prediction)

		// Compare expected and predicted values
		if bestPred == labels[i] && bestPred != -1 {
			successes++
		}
	}

	log.Printf("10.000 Test values accuracy: %.2f%%", float64(successes)/float64(len(testCases))*100)

	arrayCh := make(chan []float64)
	//game := &printing.Game{
	//	Array: arrayCh,
	//}

	log.Println("Starting the real-time prediction...")
	go func() {
		for {

			// Listen for the array channel
			select {
			case array := <-arrayCh:

				// Set the input values
				network.SetInput(array)

				// Get the predictions of the network in real time
				prediction := network.GetPredictions()

				// Get the best result
				bestPred := utlis.GetBestResult(prediction)

				// Print the predicted value
				fmt.Printf("Predicted value: %d\n", bestPred)
			}
		}
	}()

	// Run the graphical interface
	log.Println("Running the graphical interface...")
	//if err := ebiten.RunGame(game); err != nil {
	//	log.Fatal(err)
	//}
}
