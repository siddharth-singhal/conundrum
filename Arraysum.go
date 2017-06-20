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
		fmt.Println(arrayLen)
		numArrayString := strings.Split(arrayString, " ")
		var numArray []int = make([]int, arrayLen)
		var counter int = 0
		for _, numString := range numArrayString {
			numArray[counter], _ = strconv.Atoi(numString)
			counter++
		}
	}
}
