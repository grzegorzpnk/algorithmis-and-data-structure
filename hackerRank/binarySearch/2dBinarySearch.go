package binarySearch

func searchMatrix(matrix [][]int, target int) bool {

	//matrix is m x n -> m is number of rows, n is number of columns
	m := len(matrix)
	n := len(matrix[0])

	left := 0
	right := m*n - 1

	for left <= right {
		mid_index := left + (right-left)/2
		mid_val := matrix[mid_index/n][mid_index%n]

		if mid_val == target {
			return true
		} else if mid_val < target {
			left = mid_index + 1
		} else if mid_val > target {
			right = mid_index - 1
		}

	}

	return false
}

// func searchMatrix(matrix [][]int, target int) bool {

//     m := len(matrix) // ilosc "paczek"
//     n := len(matrix[0]) //ilosc w paczce
//     left := 0
//     right := m*n-1 //ostatni element numer ilosc paczek x ilosc w paczce, -1 bo to tablica

//     for left <= right {

//         mid := left + (right - left)/2
//         mid_val := matrix[mid/n][mid%n] //tego nie rozumiem

//         if mid_val == traget {
//             return true
//         } else if mid_val < target{
//             left := mid + 1
//         } else {
//             right:= mid -1
//         }

//     }

//     return false

// }
