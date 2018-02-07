package time

var year = 1

func GetCurrentYear() int {
	return year
}

func EndThisYear() {
	year = year + 1
}
