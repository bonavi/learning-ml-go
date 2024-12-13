# Neural Network Learning Project

## Overview

This repository is an educational project aimed at understanding and demonstrating how neural networks work. It is not intended to be an out-of-the-box solution for immediate use. Training a neural network is a time-consuming process, and reducing the training time is not the goal of this repository.

## Project Structure

- `main.go`: The main entry point of the application.
- `model/`: Contains the implementation of the neural network.
- `mnist/`: Contains functions to load the MNIST dataset.
- `utlis/`: Contains utility functions used by the neural network.

## Getting Started

### Prerequisites

- Go programming language installed on your machine.

### Installation

1. Clone the repository:
```sh
git clone https://github.com/bonavi/learning-ml-go.git
cd learning-ml-go
```

2. Install dependencies:`go mod tidy`

### Running the Project
1. Run the application:  `go run main.go`
2. Wait for the neural network to train.
3. Watch accuracy improve over time.
4. Draw a digit on the canvas and see the neural network predict the digit.
## Important Notes
- Training Time: Training the neural network takes a significant amount of time. This project is designed to help you understand the process, not to optimize training time.
- Educational Purpose: The primary goal of this project is to provide a clear and understandable example of how neural networks function. It is not optimized for performance or efficiency.
