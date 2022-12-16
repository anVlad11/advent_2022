package main

import (
	"fmt"
	"math"
	"strings"
)

type ItemSymbol string

const (
	Beacon   = ItemSymbol("B")
	Sensor   = ItemSymbol("S")
	Dot      = ItemSymbol(".")
	Coverage = ItemSymbol("#")
	Cross    = ItemSymbol("x")
)

type Item struct {
	Symbol            ItemSymbol
	MinBeaconDistance int
	X                 int
	Y                 int
}

func (m *Item) WithSymbol(symbol ItemSymbol) *Item {
	return &Item{
		Symbol:            symbol,
		MinBeaconDistance: m.MinBeaconDistance,
		X:                 m.X,
		Y:                 m.Y,
	}
}

type Grid struct {
	Grid           map[int]map[int]*Item
	Items          []*Item
	ItemsBySymbols map[ItemSymbol][]*Item
	MinX, MaxX     int
	MinY, MaxY     int
}

func (g *Grid) GetAt(x, y int) *Item {
	if _, exists := g.Grid[x]; !exists {
		return nil
	}

	if item, exists := g.Grid[x][y]; exists {
		return item
	}

	return nil
}

func (g *Grid) GetDistance(item1, item2 *Item) int {
	x1, y1 := item1.X, item1.Y
	x2, y2 := item2.X, item2.Y

	return GetDistance(x1, x2, y1, y2)
}

func GetDistance(x1, x2, y1, y2 int) int {
	distanceX := x1 - x2
	if distanceX < 0 {
		distanceX = -distanceX
	}

	distanceY := y1 - y2
	if distanceY < 0 {
		distanceY = -distanceY
	}

	return distanceX + distanceY
}

func (g *Grid) RemoveItem(item *Item) {
	delete(g.Grid[item.X], item.Y)
	for i := range g.Items {
		if g.Items[i] == item {
			g.Items = append(g.Items[0:i], g.Items[i+1:]...)
		}
	}
	for i := range g.ItemsBySymbols[item.Symbol] {
		if g.ItemsBySymbols[item.Symbol][i] == item {
			g.ItemsBySymbols[item.Symbol] = append(g.ItemsBySymbols[item.Symbol][0:i], g.ItemsBySymbols[item.Symbol][i+1:]...)
		}
	}
}

func (g *Grid) AddItem(item *Item) {
	if _, exists := g.Grid[item.X]; !exists {
		g.Grid[item.X] = map[int]*Item{}
	}
	g.Grid[item.X][item.Y] = item
	g.Items = append(g.Items, item)
	if _, exists := g.ItemsBySymbols[item.Symbol]; !exists {
		g.ItemsBySymbols[item.Symbol] = []*Item{item}
	} else {
		g.ItemsBySymbols[item.Symbol] = append(g.ItemsBySymbols[item.Symbol], item)
	}

	addRadius := item.MinBeaconDistance

	if item.X-addRadius < g.MinX {
		g.MinX = item.X - addRadius
	}
	if item.X+addRadius > g.MaxX {
		g.MaxX = item.X + addRadius
	}
	if item.Y-addRadius < g.MinY {
		g.MinY = item.Y - addRadius
	}
	if item.Y+addRadius > g.MaxY {
		g.MaxY = item.Y + addRadius
	}

}

func (g *Grid) PrintGrid(cursors ...*Item) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("X: %d -> %d\n", g.MinX, g.MaxX))
	sb.WriteString(fmt.Sprintf("Y: %d -> %d\n", g.MinY, g.MaxY))
	/*
		grid := map[int]map[int]*Item{}
		for _, item := range g.Items {
			if _, exists := grid[item.X]; !exists {
				grid[item.X] = map[int]*Item{}
			}
			if _, exists := grid[item.X][item.Y]; !exists {
				grid[item.X][item.Y] = item
			}
		}
		g.Grid = grid

	*/

	for y := g.MinY; y <= g.MaxY; y++ {
		sb.WriteString(fmt.Sprintf("%d\t", y))
		for x := g.MinX; x <= g.MaxX; x++ {
			hasCursor := false
			for _, cursor := range cursors {
				if cursor != nil {
					if cursor.X == x && cursor.Y == y {
						sb.WriteString(string(cursor.Symbol))
						hasCursor = true
						break
					}
				}
			}
			if hasCursor {
				continue
			}

			line, lineExists := g.Grid[x]
			if !lineExists {
				sb.WriteString(string(Dot))
				continue
			}

			item, exists := line[y]
			if exists {
				sb.WriteString(string(item.Symbol))
			} else {
				sb.WriteString(string(Dot))
			}
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func NewGrid() *Grid {
	return &Grid{
		Grid:           map[int]map[int]*Item{},
		MinX:           math.MaxInt,
		MaxX:           math.MinInt,
		MinY:           math.MaxInt,
		MaxY:           math.MinInt,
		Items:          []*Item{},
		ItemsBySymbols: map[ItemSymbol][]*Item{},
	}
}
