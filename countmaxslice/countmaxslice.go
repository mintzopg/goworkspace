package countmaxslice

// func main() {
// 	v := []int{-4, 2, 2}
// 	fmt.Println(countmax(v))

// }

// Countmax function(v []int) int; counts the number of elements in slice 'v' that are equal to the max element in the slice
func Countmax(u []int) int {
	count := 0

	var fx func(u []int, max int)

	fx = func(u []int, max int) {
		// recursive function to used for counting
		if len(u) <= 0 {
			return
		}
		// algo is in here
		if curmax := u[0]; curmax == max {
			count++
		} else if curmax > max {
			max = curmax
			count = 1
		}
		fx(u[1:], max)
	}
	// initial call
	fx(u, 0)
	return count
}
