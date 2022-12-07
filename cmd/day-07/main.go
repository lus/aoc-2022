package main

import (
	"fmt"
	"github.com/lus/aoc-2022/internal/x"
	"strings"
)

type ElementType int

const (
	TypeDir ElementType = iota
	TypeFile
)

type Element struct {
	Type    ElementType
	Name    string
	Content []*Element
	Size    int
	Top     *Element
}

func main() {
	input := x.ReadInput(true)
	lines := strings.Split(input, "\n")

	root := &Element{
		Type:    TypeDir,
		Name:    "/",
		Content: nil,
		Size:    0,
		Top:     nil,
	}
	cur := root

	isList := false
	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			isList = false
			line = strings.TrimPrefix(line, "$ ")
			command := strings.Split(line, " ")
			switch command[0] {
			case "cd":
				switch command[1] {
				case "/":
					cur = root
					break
				case "..":
					cur = cur.Top
					break
				default:
					for _, element := range cur.Content {
						if element.Type == TypeDir && element.Name == command[1] {
							cur = element
						}
					}
					break
				}
			case "ls":
				isList = true
				break
			default:
				break
			}
			continue
		}

		if isList {
			raw := strings.Split(line, " ")
			if raw[0] == "dir" {
				cur.Content = append(cur.Content, &Element{
					Type:    TypeDir,
					Name:    raw[1],
					Content: nil,
					Size:    0,
					Top:     cur,
				})
			} else {
				cur.Content = append(cur.Content, &Element{
					Type:    TypeFile,
					Name:    raw[1],
					Content: nil,
					Size:    x.MustInt(raw[0]),
					Top:     cur,
				})
			}
			continue
		}
	}

	directorySizes := flatExtractDirectorySizes(root)
	rootSize := calculateElementSize(root)
	directorySizes = append(directorySizes, rootSize)

	required := 30000000 - (70000000 - rootSize)

	sum := 0
	min := -1
	for _, size := range directorySizes {
		if size <= 100000 {
			sum += size
		}
		if min == -1 || (size >= required && size < min) {
			min = size
		}
	}
	fmt.Printf("The sum of the sizes of all directories with an individual maximum size of 100000 is %d.\n", sum)
	fmt.Printf("The size of the smallest directory that would free up enough space when deleted is %d.\n", min)
}

func flatExtractDirectorySizes(root *Element) []int {
	if root == nil || len(root.Content) == 0 {
		return nil
	}
	var sizes []int
	for _, element := range root.Content {
		if element.Type != TypeDir {
			continue
		}
		sizes = append(sizes, calculateElementSize(element))
		subSizes := flatExtractDirectorySizes(element)
		sizes = append(sizes, subSizes...)
	}
	return sizes
}

func calculateElementSize(element *Element) int {
	if element == nil || len(element.Content) == 0 {
		return 0
	}
	size := 0
	for _, sub := range element.Content {
		if sub.Type == TypeDir {
			size += calculateElementSize(sub)
		} else {
			size += sub.Size
		}
	}
	return size
}
