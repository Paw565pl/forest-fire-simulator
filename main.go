package main

import (
	"flag"
	"fmt"
	"forest-fire-simulator/forest"
	"sync"
)

func findOptimalForestCoverPercentage(forestDimensionX, forestDimensionY int, maxBurnRatePercentage uint8) (uint8, uint8) {
	if maxBurnRatePercentage <= 0 || maxBurnRatePercentage > 100 {
		panic("Max burn rate must be between 1 and 100")
	}

	totalTrees := forestDimensionX * forestDimensionY
	results := make(map[uint8]uint8, 100)

	const Samples = 1000
	var wg sync.WaitGroup
	var mu sync.Mutex

	for forestCoverPercentage := uint8(1); forestCoverPercentage <= 100; forestCoverPercentage++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			averageBurnRatePercentage := 0

			for range Samples {
				burntTrees, _ := forest.BurnForest(forestDimensionX, forestDimensionY, forestCoverPercentage)
				burnRate := 100 * int(burntTrees) / totalTrees
				averageBurnRatePercentage += burnRate
			}

			averageBurnRatePercentage /= Samples

			mu.Lock()
			results[forestCoverPercentage] = uint8(averageBurnRatePercentage)
			mu.Unlock()
		}()
	}

	wg.Wait()

	optimalForestCoverPercentage := uint8(0)
	optimalAverageBurnRate := uint8(0)

	for forestCoverPercentage, averageBurnRatePercentage := range results {
		if averageBurnRatePercentage <= maxBurnRatePercentage && forestCoverPercentage > optimalForestCoverPercentage {
			optimalForestCoverPercentage = forestCoverPercentage
			optimalAverageBurnRate = averageBurnRatePercentage
		}
	}

	for {
		burntTrees, burntForest := forest.BurnForest(forestDimensionX, forestDimensionY, optimalForestCoverPercentage)

		if burntTrees != 0 {
			forest.SaveForestToFile(burntForest)
			break
		}
	}

	return optimalForestCoverPercentage, optimalAverageBurnRate
}

func main() {
	forestDimensionX := flag.Int("x", 10, "forest dimension X")
	forestDimensionY := flag.Int("y", 10, "forest dimension Y")
	maxBurnRatePercentage := flag.Uint("maxBurn", 30, "max acceptable burn rate percentage of the forest [1-100]")

	flag.Parse()

	fmt.Printf("started calculations for forest with dimensions %dx%d and max acceptable burn rate %d%%...\n", *forestDimensionX, *forestDimensionY, *maxBurnRatePercentage)
	optimalForestCoverPercentage, averageBurnRate := findOptimalForestCoverPercentage(*forestDimensionX, *forestDimensionY, uint8(*maxBurnRatePercentage))
	fmt.Printf("optimal forest cover percentage is: %d%% with average burn rate of: %d%%\n", optimalForestCoverPercentage, averageBurnRate)
	fmt.Println("saved to file a sample burnt forest with calculated optimal cover percentage")
}
