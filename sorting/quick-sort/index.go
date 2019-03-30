package algorithms

import "fmt"

func main() {
    fmt.Println(quick_sort([]int{1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,16,31,30,29,28,27,20, 10, 12, 11}))
}

func partition_func(A []int, start int, end int) int {
    pivot := A[end - 1]
    pIndex := start
    for i := start; i < end; i++ {
        if A[i] <= pivot {
            temp := A[i]
            A[i] = A[pIndex]
            A[pIndex] = temp
            pIndex++
        }
    }
    return pIndex
}

func quick_sort_func(A []int, start int, end int) {
    if start >= end {
        return
    }
    pIndex := partition_func(A, start, end)
    quick_sort_func(A, start, pIndex - 1)
    quick_sort_func(A, pIndex, end)
}

func quick_sort(A []int) []int {
    quick_sort_func(A, 0, len(A))
    return A
}