package printing

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	canvasSize   = 28 // Size of the canvas
	screenWidth  = 48 // Screen width
	screenHeight = 28 // Screen height
	buttonWidth  = 20 // Button width
	buttonHeight = 5  // Button height
)

// Array to store the state of each pixel
var pixels [canvasSize][canvasSize]float64

// Function to get a one-dimensional array from a two-dimensional array
func getOneDimensionalArray() []float64 {
	var array []float64
	for y := 0; y < canvasSize; y++ {
		for x := 0; x < canvasSize; x++ {
			array = append(array, pixels[y][x])
		}
	}
	return array
}

// Game implements ebiten.Game interface
type Game struct {
	Array chan []float64
}

// Update updates the game by one tick
func (g *Game) Update() error {

	// Get the cursor position
	mouseX, mouseY := ebiten.CursorPosition()

	// If the left mouse button is pressed
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

		// Determine which pixel was clicked (taking into account the scale)
		x := mouseX
		y := mouseY

		// If within the canvas
		if x >= 0 && x < canvasSize && y >= 0 && y < canvasSize {

			// Set the pixel to black
			pixels[y][x] = 1.0
		}

		// Check if the button was pressed (in the right bottom corner)
		if mouseX >= canvasSize && mouseX <= screenWidth && mouseY >= buttonWidth && mouseY <= screenHeight {

			// Get the one-dimensional array and send it to the channel
			g.Array <- getOneDimensionalArray()
		}

		// Check if the button was pressed (in the right top corner)
		if mouseX >= canvasSize && mouseX <= screenWidth && mouseY >= 0 && mouseY <= buttonHeight {

			// Clear the canvas
			pixels = [canvasSize][canvasSize]float64{}
		}
	}

	return nil
}

// Draw renders the canvas
func (g *Game) Draw(screen *ebiten.Image) {

	// Fill the screen with white color
	screen.Fill(color.White)

	// Draw each pixel on the screen
	for y := 0; y < canvasSize; y++ {
		for x := 0; x < canvasSize; x++ {

			// If the pixel is black
			if pixels[y][x] == 1.0 {

				// Draw a black pixel
				screen.Set(x, y, color.Black)
			}
		}
	}

	// Draw the button for getting the array
	drawButton(screen, canvasSize, screenHeight-buttonHeight)

	// Draw the button for clearing the canvas
	drawButton(screen, canvasSize, 0)
}

// Function to draw a button
func drawButton(screen *ebiten.Image, x, y float64) {

	// Button rectangle
	buttonRect := ebiten.NewImage(buttonWidth, buttonHeight)
	buttonRect.Fill(color.RGBA{R: 0, G: 0, B: 255, A: 255}) // Blue color for the button

	// Button position
	geo := ebiten.GeoM{}
	geo.Translate(x, y)

	// Draw the button
	screen.DrawImage(buttonRect, &ebiten.DrawImageOptions{GeoM: geo})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set the screen size
	return screenWidth, screenHeight
}
