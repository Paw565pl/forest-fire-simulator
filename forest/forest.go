package forest

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

// createForest creates forest with given dimensions and returns it.
// It panics if forestCoverPercentage is not between 1 and 100
// or if either of the dimensions is less than 1.
func createForest(dimensionX, dimensionY int, forestCoverPercentage uint8) [][]*Tree {
	if dimensionX < 1 || dimensionY < 1 {
		panic("Forest dimensions must be at least 1x1.")
	}
	if forestCoverPercentage <= 0 || forestCoverPercentage > 100 {
		panic("Percentage of forest cover must be between 1 and 100.")
	}

	forest := make([][]*Tree, dimensionY)
	for i := range forest {
		forest[i] = make([]*Tree, dimensionX)
	}

	// create trees
	treesToPlant := dimensionX * dimensionY * int(forestCoverPercentage) / 100
	for treesToPlant != 0 {
		for y, row := range forest {
			for x, block := range row {
				if block != nil {
					continue
				}

				isTreeCreated := rand.IntN(2)
				if isTreeCreated == 0 {
					continue
				}

				forest[y][x] = NewTree()
				treesToPlant--

				if treesToPlant == 0 {
					break
				}
			}
			if treesToPlant == 0 {
				break
			}
		}
	}

	return forest
}

// BurnForest makes simulation of fire spreading in forest with given parameters by randomly hitting it with lightning
// and returns number of burnt trees and burnt forest object.
func BurnForest(dimensionX, dimensionY int, forestCoverPercentage uint8) (uint, [][]*Tree) {
	forest := createForest(dimensionX, dimensionY, forestCoverPercentage)

	lightningStrikeX := rand.IntN(dimensionX)
	lightningStrikeY := rand.IntN(dimensionY)

	// start spreading fire
	burntTrees := burnTree(&forest, lightningStrikeX, lightningStrikeY)
	return burntTrees, forest
}

func burnTree(forestPointer *[][]*Tree, x, y int) uint {
	forest := *forestPointer
	areCoordinatesValid := 0 <= y && y < len(forest) && 0 <= x && x < len(forest[y])

	if !areCoordinatesValid {
		return 0
	}

	tree := forest[y][x]
	if tree == nil || !tree.isAlive {
		return 0
	}

	tree.isAlive = false
	burntTrees := uint(1)

	// spread fire for 8 directions
	const Directions = 8
	directionsX := [Directions]int{-1, 0, 1, -1, 1, -1, 0, 1}
	directionsY := [Directions]int{-1, -1, -1, 0, 0, 1, 1, 1}
	for i := 0; i < 8; i++ {
		burntTrees += burnTree(forestPointer, x+directionsX[i], y+directionsY[i])
	}

	return burntTrees
}

// SaveForestToFile saves representation of a forest in a text file.
func SaveForestToFile(forest [][]*Tree) {
	forestStringSlice := make([]string, len(forest))

	for rowIndex, row := range forest {
		rowStrings := make([]string, len(row))
		for elementIndex, element := range row {
			if element != nil {
				rowStrings[elementIndex] = element.String()
			} else {
				rowStrings[elementIndex] = "💩"
			}
		}
		forestStringSlice[rowIndex] = strings.Join(rowStrings, " ")
	}

	output := strings.Join(forestStringSlice, "\n")

	file, openFileError := os.OpenFile("forest_visualisation.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if openFileError != nil {
		fmt.Printf("Failed to open file: %v\n", openFileError)
	}
	defer func(file *os.File) {
		closeFileError := file.Close()
		if closeFileError != nil {
			fmt.Printf("Failed to close file: %v\n", closeFileError)
		}
	}(file)

	if _, saveError := file.WriteString(output); saveError != nil {
		fmt.Printf("Failed to save file: %v\n", saveError)
	}
}
