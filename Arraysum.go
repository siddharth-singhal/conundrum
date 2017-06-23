package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*Given an array of integers, find a combination of four elements in the array whose sum is equal to a given value X.

Input:
First line consists of T test cases. First line of every test case consists of an integer N, denoting the number of elements in an array. Second line consists of N spaced array elements. Third line of every test case X.

Output:
Single line output, print 1 if combination is found else 0.

Constraints:
1<=T<=100
1<=N<=100

Example:
Input:
1
6
1 5 1 0 6 0
7
Output:
1*/

type arrayDetail struct {
	i   int
	j   int
	sum int
}

func main() {
	var numOfCases, arrayLen, sum int
	var arrayString string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanf("%d", &numOfCases)
	for i := 0; i < numOfCases; i++ {
		fmt.Scan(&arrayLen)
		for scanner.Scan() {
			arrayString = scanner.Text()
			if len(arrayString) > 0 {
				break
			}
		}
		fmt.Scanf("%d", &sum)
		numArrayString := strings.Split(arrayString, " ")
		var auxillaryArray []arrayDetail = make([]arrayDetail, 0)
		var numArray []int = make([]int, 0)
		for _, numString := range numArrayString {
			num, _ := strconv.Atoi(numString)
			numArray = append(numArray, num)
		}
		var flag int = 0
		for i := 0; i < arrayLen-1; i++ {
			for j := i + 1; j < arrayLen; j++ {
				if i != j {
					var element arrayDetail
					element.i = i
					element.j = j
					element.sum = numArray[i] + numArray[j]
					auxillaryArray = append(auxillaryArray, element)
				}
			}
		}
		for i := 0; i < len(auxillaryArray)-1; i++ {
			if flag == 1 {
				break
			}
			for j := i + 1; j < len(auxillaryArray); j++ {
				if auxillaryArray[i].sum+auxillaryArray[j].sum == sum && auxillaryArray[i].i != auxillaryArray[j].i &&
					auxillaryArray[i].i != auxillaryArray[j].j &&
					auxillaryArray[i].j != auxillaryArray[j].i &&
					auxillaryArray[i].j != auxillaryArray[j].j {
					flag = 1
				}
			}
		}
		fmt.Println(flag)
	}
}
