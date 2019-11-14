package algorithm

// InsertionSort (https://www.geeksforgeeks.org/insertion-sort/)
// Time Complexity: O(n*2)
// Auxiliary Space: O(1)
// Boundary Cases: Insertion sort takes maximum time to sort if elements are sorted in reverse order. And it takes minimum time (Order of n) when elements are already sorted.
// Algorithmic Paradigm: Incremental Approach
// Sorting In Place: Yes
// Stable: Yes
// Online: Yes
// Uses: Insertion sort is used when number of elements is small. It can also be useful when input array is almost sorted, only few elements are misplaced in complete big array.
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && tmp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
	return arr
}

// MergeSort (https://www.geeksforgeeks.org/merge-sort/)
// Time Complexity: Sorting arrays on different machines. Merge Sort is a recursive algorithm and time complexity can be expressed as following recurrence relation.
// T(n) = 2T(n/2) + Theta(n)
// The above recurrence can be solved either using Recurrence Tree method or Master method. It falls in case II of Master Method and solution of the recurrence is Theta(n Logn ).
// Time complexity of Merge Sort is Theta(n Log n) in all 3 cases (worst, average and best) as merge sort always divides the array into two halves and take linear time to merge two halves.
// Auxiliary Space: O(n)
// Algorithmic Paradigm: Divide and Conquer
// Sorting In Place: No in a typical implementation
// Stable: Yes
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	leftArr := MergeSort(arr[:middle])
	rightArr := MergeSort(arr[middle:])

	result := []int{}
	var leftIndex, rightIndex int
	for leftIndex < len(leftArr) && rightIndex < len(rightArr) {
		if leftArr[leftIndex] <= rightArr[rightIndex] {
			result = append(result, leftArr[leftIndex])
			leftIndex++
		} else {
			result = append(result, rightArr[rightIndex])
			rightIndex++
		}
	}
	if leftIndex < len(leftArr) {
		result = append(result, leftArr[leftIndex:]...)
	} else {
		result = append(result, rightArr[rightIndex:]...)
	}
	return result
}
