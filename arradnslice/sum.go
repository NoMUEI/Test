package arradnslice


func Sum(a []int) int {
	var sum int
	for _,i := range a {
		sum += i
	}
	return sum
}


func SumAll(a, b []int) []int {
	var sumarr []int
	var suma int
	var sumb int
	for _, i := range a {
		suma += i
	}
	for _, j := range b {
		sumb += j
	}
	sumarr = append(sumarr,suma)
	sumarr = append(sumarr,suma)
	return  sumarr
}