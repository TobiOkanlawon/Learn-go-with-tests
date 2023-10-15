package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumEach(collection ...[]int) []int {
	var sums []int

	for _, numbers := range collection {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(collection ...[]int) []int {
	var sums []int

	for _, numbers := range collection {
		if len(numbers) > 1 {
			sums = append(sums, Sum(numbers[1:]))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}

func main() {

}
