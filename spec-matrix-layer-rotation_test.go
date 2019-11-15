package matrixLayerRotation

import (
	"fmt"
	"testing"
)

// result should be:
// 3 4 8 12
// 2 11 10 16
// 1 7 6 15
// 5 9 13 14

func TestMatrixLayerRotation(t *testing.T) {
	fmt.Println("4 x 4")
	testArray := [][]int32{
		[]int32{1, 2, 3, 4},
		[]int32{5, 6, 7, 8},
		[]int32{9, 10, 11, 12},
		[]int32{13, 14, 15, 16},
	}

	matrixRotation(testArray, int32(2))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("5 x 5")
	testArray = [][]int32{
		[]int32{1, 2, 3, 4, 5},
		[]int32{6, 7, 8, 9, 10},
		[]int32{11, 12, 13, 14, 15},
		[]int32{16, 17, 18, 19, 20},
		[]int32{21, 22, 23, 24, 25},
	}

	matrixRotation(testArray, int32(2))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("5 x 4")
	testArray = [][]int32{
		[]int32{1, 2, 3, 4},
		[]int32{5, 6, 7, 8},
		[]int32{9, 10, 11, 12},
		[]int32{13, 14, 15, 16},
		[]int32{17, 18, 19, 20},
	}

	matrixRotation(testArray, int32(2))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("4 x 5")
	testArray = [][]int32{
		[]int32{1, 2, 3, 4, 5},
		[]int32{6, 7, 8, 9, 10},
		[]int32{11, 12, 13, 14, 15},
		[]int32{16, 17, 18, 19, 20},
	}

	matrixRotation(testArray, int32(2))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("5 x 6")
	testArray = [][]int32{
		[]int32{1, 2, 3, 4, 5, 6},
		[]int32{7, 8, 9, 10, 11, 12},
		[]int32{13, 14, 15, 16, 17, 18},
		[]int32{19, 20, 21, 22, 23, 24},
		[]int32{25, 26, 27, 28, 29, 30},
	}

	matrixRotation(testArray, int32(2))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("6 x 5")
	testArray = [][]int32{
		[]int32{1, 2, 3, 4, 5},
		[]int32{6, 7, 8, 9, 10},
		[]int32{11, 12, 13, 14, 15},
		[]int32{16, 17, 18, 19, 20},
		[]int32{21, 22, 23, 24, 25},
		[]int32{26, 27, 28, 29, 30},
	}

	matrixRotation(testArray, int32(2))
}
