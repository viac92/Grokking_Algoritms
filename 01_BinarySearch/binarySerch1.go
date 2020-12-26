package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"os"
)

//foundItem cerca in una slice un elemento utilizzando l'algoritmo binary search.
func foundItem(nS []int, item int) (int, bool) {
	low := 0
	hight := len(nS)-1

	for low <= hight {
		mid := (low + hight) / 2
		guess := nS[mid]
		
		if item == guess {
			return guess, true
		}
		if guess > item {
			hight = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

func main() {
	file,err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b := bufio.NewScanner(file)
	var numSlice []int

	for b.Scan() {
		t := b.Text()
		num,_ := strconv.Atoi(t)
		numSlice = append(numSlice, num)
	}
	sort.Ints(numSlice)
	
	n := 0
	fmt.Print("Cerca un numero nella slice con il binary search: ")
	fmt.Scan(&n)
	fmt.Println(foundItem(numSlice, n))
}