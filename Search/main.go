package main

import(
	"fmt"
)
func MergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	leftSide := MergeSort(items[0 : len(items)/2])
	rightSide := MergeSort(items[len(items)/2 : ])
	i := 0
	j := 0
	combined := []int{}
	for i < len(leftSide) || j < len(rightSide) {
		if i >= len(leftSide) {
			combined = append(combined, rightSide[j:]...)
			j = len(rightSide)
			continue
	 	}
		if j >= len(rightSide) {
			combined = append(combined, leftSide[i:]...)
			i = len(leftSide)
			continue
	 	}
		if leftSide[i] < rightSide[j] {
			combined = append(combined, leftSide[i])
			i = i + 1
			continue
	 	}
		combined = append(combined, rightSide[j])
		j = j + 1
	}
	return combined
}

func BubbleSort(items []int) {
	for {
		sortHappened := false
		for i := 0; i < len(items)-1; i++ {
			if items[i] > items[i+1] {
				temp := items[i]
				items[i] = items[i+1]
				items[i+1] = temp
				sortHappened = true
			}
		}
		if !sortHappened {
			break
		}
	}
}

func NormalSearch(vet []int, val int)(int, error){
	for i:=0; i < len(vet); i++{
		if vet[i] == val {
			return i, nil
		}
	}
	return 0, fmt.Errorf("valor não encontrado")
}

func quicksort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	leftSide := []int{}
	rightSide := []int{}
	pivot := values[len(values)-1]
	for _, v := range values[0 : len(values)-1] {
		if v < pivot {
			leftSide = append(leftSide, v)
			continue
			}
			rightSide = append(rightSide, v)
	}
	sortedLeftSide := quicksort(leftSide)
	sortedRightSide := quicksort(rightSide)
	sorted := append(sortedLeftSide, pivot)
	sorted = append(sorted, sortedRightSide...)
	return sorted
}
func BinarySearch(finding int, values []int) bool {
	if len(values) == 0 {
			return false
	}
	if len(values) == 1 {
		if values[0] == finding {
			return values[0] == finding
		}
		return false
		}
		found := false
		leftHalf := values[0 : len(values)/2]
		rightHalf := values[len(values)/2 : ]
		if finding >= rightHalf[0] {
			found = BinarySearch(finding, rightHalf)
		}else {
			found = BinarySearch(finding, leftHalf)
		}
		return found
	}


func main(){
	//fmt.Println(NormalSearch([]int{1,2,3,4},5))
	//values := []int{4, 3, 2, 1}
	//fmt.Println(values)
	//BubbleSort(values)
	//fmt.Println(values)
	
	values := []int{4, 3, 2, 1}
	sorted := MergeSort(values)
	fmt.Println(sorted)
}
