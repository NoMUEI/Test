package arradnslice


func Sum(a []int) int {
	var sum int
	for _,i := range a {
		sum += i
	}
	return sum
}


func SumAll(numberToSum...[]int) []int {
	lenthofnumbers := len(numberToSum)
	sums := make([]int, lenthofnumbers)
	for i, numbers := range numberToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}