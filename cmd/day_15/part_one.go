package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	grid := NewGrid()

	regex, err := regexp.Compile(`Sensor at x=(-?\d*), y=(-?\d*): closest beacon is at x=(-?\d*), y=(-?\d*)`)
	if err != nil {
		return err
	}
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	for _, line := range file {
		matches := regex.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			continue
		}
		match := matches[0][1:]
		sensorX, _ := strconv.Atoi(match[0])
		sensorY, _ := strconv.Atoi(match[1])
		beaconX, _ := strconv.Atoi(match[2])
		beaconY, _ := strconv.Atoi(match[3])

		beacon := &Item{
			Symbol: Beacon,
			X:      beaconX,
			Y:      beaconY,
		}

		sensor := &Item{
			Symbol: Sensor,
			X:      sensorX,
			Y:      sensorY,
		}

		sensor.MinBeaconDistance = grid.GetDistance(sensor, beacon)

		grid.AddItem(sensor)
		grid.AddItem(beacon)
	}

	y := 2000000
	cantBe := 0
	for x := grid.MinX; x <= grid.MaxX; x++ {
		itemAt := grid.GetAt(x, y)
		if itemAt != nil {
			if itemAt.Symbol == Beacon || itemAt.Symbol == Sensor {
				continue
			}
		}
		cross := &Item{Symbol: Cross, X: x, Y: y}
		for _, item := range grid.ItemsBySymbols[Sensor] {
			distanceFromCrossToItem := grid.GetDistance(cross, item)
			if item.MinBeaconDistance >= distanceFromCrossToItem {
				coverage := &Item{
					Symbol: Coverage,
					X:      x,
					Y:      y,
				}
				grid.AddItem(coverage)
				cantBe++
				break
			}
		}
	}

	fmt.Println(cantBe)

	return nil
}
