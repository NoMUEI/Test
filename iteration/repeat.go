package iteration


func Repeat(s string, times int) string {
	var t string
	for i := 0; i < times; i++ {
		t += s
	}
	return t
}
