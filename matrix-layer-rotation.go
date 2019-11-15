package matrixLayerRotation

import (
	"fmt"
	"strings"
)

type value struct {
	slice []int32
	index int32
}

func (v *value) nextValue() int32 {
	returnValue := v.slice[v.index]
	v.index++
	return returnValue
}

func (v *value) rotate(r int32) {
	length := int32(len(v.slice))
	r = r % length
	v.slice = append(v.slice[length-r:length:length], v.slice[:length-r]...)
}

func matrixRotation(matrix [][]int32, r int32) {
	rowLength := int32(len(matrix))
	columnLength := int32(len(matrix[0]))
	level := int32(0)
	for {
		done := rotateLevel(matrix, level, rowLength, columnLength, r)
		if done {
			break
		}
		level++
	}

	printArray(matrix)
}

func rotateLevel(matrix [][]int32, level int32, rowLength int32, columnLength int32, r int32) bool {

	values := &value{
		slice: make([]int32, 0),
		index: 0,
	}
	var endingPattern int32
	if rowLength <= columnLength {
		endingPattern = rowLength - (level * 2)
	} else {
		endingPattern = columnLength - (level * 2)
	}

	if endingPattern == 2 && columnLength <= rowLength {
		for i := level; i < rowLength-level; i++ {
			values.slice = append(values.slice, matrix[i][0+level])
		}
		for i := rowLength - 1 - level; i >= level; i-- {
			values.slice = append(values.slice, matrix[i][columnLength-1-level])
		}
	} else {
		// start in lower left corner and go counter clockwise
		for i := level; i < columnLength-level; i++ {
			values.slice = append(values.slice, matrix[rowLength-1-level][i])
		}
		for i := rowLength - 2 - level; i > level; i-- {
			values.slice = append(values.slice, matrix[i][columnLength-1-level])
		}
		for i := columnLength - level - 1; i >= level; i-- {
			values.slice = append(values.slice, matrix[level][i])
		}
		for i := level + 1; i < rowLength-1-level; i++ {
			values.slice = append(values.slice, matrix[i][0+level])
		}
	}

	values.rotate(r)

	if endingPattern == 2 && columnLength <= rowLength {
		for i := level; i < rowLength-level; i++ {
			matrix[i][0+level] = values.nextValue()
		}
		for i := rowLength - 1 - level; i >= level; i-- {
			matrix[i][columnLength-1-level] = values.nextValue()
		}
	} else {
		// start in lower left corner and go counter clockwise
		for i := level; i < columnLength-level; i++ {
			matrix[rowLength-1-level][i] = values.nextValue()
		}
		for i := rowLength - 2 - level; i > level; i-- {
			matrix[i][columnLength-1-level] = values.nextValue()
		}
		for i := columnLength - level - 1; i >= level; i-- {
			matrix[level][i] = values.nextValue()
		}
		for i := level + 1; i < rowLength-1-level; i++ {
			matrix[i][0+level] = values.nextValue()
		}
	}

	if endingPattern == 2 || endingPattern == 3 {
		return true
	}
	return false
}

func printArray(matrix [][]int32) {
	for _, value := range matrix {
		fmt.Println(strings.Trim(fmt.Sprint(value), "[]"))
	}
}
