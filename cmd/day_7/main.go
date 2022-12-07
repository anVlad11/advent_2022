package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anVlad11/advent_2022/internal/helpers"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println(err)
	}
}

func do() error {
	file, err := helpers.ReadFile("./data/input/day_7_input.txt")
	if err != nil {
		return err
	}

	err = part1(file)
	if err != nil {
		return err
	}

	err = part2(file)
	if err != nil {
		return err
	}

	return nil
}

type FSItem struct {
	ID       int
	Name     string
	Parent   *FSItem
	Children map[string]*FSItem
	IsDir    bool
	Size     int
}

func part1(file []string) error {
	if len(file) == 0 {
		return nil
	}

	root := &FSItem{
		ID:       0,
		Name:     "/",
		Parent:   nil,
		Children: map[string]*FSItem{},
		IsDir:    true,
		Size:     0,
	}

	id := 1

	maxSize := 100000
	smallerThanMaxSize := map[int]*FSItem{root.ID: root}

	var cwd *FSItem
	var command string

	for i := 0; i < len(file); i++ {
		line := file[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			command = parts[1]
			switch command {
			case "cd":
				{
					target := parts[2]
					switch target {
					case "/":
						{
							cwd = root
						}
					case "..":
						{
							if cwd.Parent != nil {
								cwd = cwd.Parent
							}
						}
					default:
						{
							if targetDir, exists := cwd.Children[target]; exists {
								cwd = targetDir
							}
						}
					}
				}
			case "ls":
				{
					continue
				}
			}
		} else {
			switch command {
			case "ls":
				{
					if parts[0] == "dir" {
						if _, exists := cwd.Children[parts[1]]; !exists {
							item := &FSItem{
								ID:       id,
								Name:     parts[1],
								Parent:   cwd,
								Children: map[string]*FSItem{},
								IsDir:    true,
								Size:     0,
							}
							id++
							cwd.Children[parts[1]] = item
							smallerThanMaxSize[item.ID] = item
						}
						continue
					}
					if _, exists := cwd.Children[parts[1]]; !exists {
						size, err := strconv.Atoi(parts[0])
						if err != nil {
							return err
						}

						item := &FSItem{
							ID:       id,
							Name:     parts[1],
							Parent:   cwd,
							Children: nil,
							IsDir:    false,
							Size:     size,
						}
						id++

						cwd.Children[parts[1]] = item
						for parent := item.Parent; parent != nil; parent = parent.Parent {
							parent.Size += size
							if parent.Size > maxSize {
								delete(smallerThanMaxSize, parent.ID)
							}
						}
					}
				}
			}
		}
	}

	totalSize := 0
	for _, item := range smallerThanMaxSize {
		totalSize += item.Size
	}

	fmt.Println(totalSize)

	return nil
}

func part2(file []string) error {
	if len(file) == 0 {
		return nil
	}

	root := &FSItem{
		ID:       0,
		Name:     "/",
		Parent:   nil,
		Children: map[string]*FSItem{},
		IsDir:    true,
		Size:     0,
	}

	id := 1

	var cwd *FSItem
	var command string

	directoriesList := map[int]*FSItem{root.ID: root}

	for i := 0; i < len(file); i++ {
		line := file[i]
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			command = parts[1]
			switch command {
			case "cd":
				{
					target := parts[2]
					switch target {
					case "/":
						{
							cwd = root
						}
					case "..":
						{
							if cwd.Parent != nil {
								cwd = cwd.Parent
							}
						}
					default:
						{
							if targetDir, exists := cwd.Children[target]; exists {
								cwd = targetDir
							}
						}
					}
				}
			case "ls":
				{
					continue
				}
			}
		} else {
			switch command {
			case "ls":
				{
					if parts[0] == "dir" {
						if _, exists := cwd.Children[parts[1]]; !exists {
							item := &FSItem{
								ID:       id,
								Name:     parts[1],
								Parent:   cwd,
								Children: map[string]*FSItem{},
								IsDir:    true,
								Size:     0,
							}
							id++
							cwd.Children[parts[1]] = item
							directoriesList[item.ID] = item
						}
						continue
					}
					if _, exists := cwd.Children[parts[1]]; !exists {
						size, err := strconv.Atoi(parts[0])
						if err != nil {
							return err
						}

						item := &FSItem{
							ID:       id,
							Name:     parts[1],
							Parent:   cwd,
							Children: nil,
							IsDir:    false,
							Size:     size,
						}
						id++

						cwd.Children[parts[1]] = item
						for parent := item.Parent; parent != nil; parent = parent.Parent {
							parent.Size += size
						}
					}
				}
			}
		}
	}

	spaceTotal := 70000000
	spaceTaken := root.Size
	spaceNeeded := 30000000
	spaceToFree := spaceNeeded - (spaceTotal - spaceTaken)
	directorySizeToDelete := root.Size
	for _, directory := range directoriesList {
		if directory.Size < directorySizeToDelete && directory.Size > spaceToFree {
			directorySizeToDelete = directory.Size
		}
	}

	fmt.Println(directorySizeToDelete)

	return nil
}
