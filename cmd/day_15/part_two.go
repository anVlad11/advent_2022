package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/anVlad11/advent_2022/internal/helpers"
)

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	regex, err := regexp.Compile(`Sensor at x=(-?\d*), y=(-?\d*): closest beacon is at x=(-?\d*), y=(-?\d*)`)
	if err != nil {
		return err
	}

	sensors := make([][3]int, 0, len(file))
	sensorsEdges := map[[3]int][][2]int{}
	minCoord := 0
	maxCoord := 4000000

	for _, line := range file {
		if strings.TrimSpace(line) == "" {
			continue
		}
		matches := regex.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			continue
		}
		match := matches[0][1:]
		sensorX := helpers.Atoi(match[0])
		sensorY := helpers.Atoi(match[1])
		beaconX := helpers.Atoi(match[2])
		beaconY := helpers.Atoi(match[3])

		beaconDistance := GetDistance(sensorX, beaconX, sensorY, beaconY)
		sensor := [3]int{sensorX, sensorY, beaconDistance}

		edgeDistance := beaconDistance + 1
		edgePoints := map[[2]int]bool{}
		for i := 0; i <= edgeDistance; i++ {
			notI := edgeDistance - i
			edgePoints[[2]int{sensorX + i, sensorY + notI}] = true
			edgePoints[[2]int{sensorX + i, sensorY - notI}] = true
			edgePoints[[2]int{sensorX - i, sensorY + notI}] = true
			edgePoints[[2]int{sensorX - i, sensorY - notI}] = true
		}

		sensorEdgePoints := [][2]int{}
		for edgePoint := range edgePoints {
			if edgePoint[0] < minCoord || edgePoint[0] > maxCoord || edgePoint[1] < minCoord || edgePoint[1] > maxCoord {
				continue
			}

			sensorEdgePoints = append(sensorEdgePoints, edgePoint)
		}
		sensorsEdges[sensor] = sensorEdgePoints
		sensors = append(sensors, sensor)
	}

	for i, sensor := range sensors {
		edgePoints := sensorsEdges[sensor]
		for _, edgePoint := range edgePoints {
			inRange := false
			for j, anotherSensor := range sensors {
				if i == j {
					continue
				}
				anotherSensorDistanceToEdge := GetDistance(edgePoint[0], anotherSensor[0], edgePoint[1], anotherSensor[1])
				if anotherSensor[2] >= anotherSensorDistanceToEdge {
					inRange = true
					break
				}
			}
			if !inRange {
				fmt.Println(edgePoint[0]*4000000 + edgePoint[1])
				return nil
			}
		}
	}

	return nil
}
