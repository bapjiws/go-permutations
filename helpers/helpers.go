package helpers

/*
rotate constructs "rotations" of a list. For instance,
rotate [1,2,3,4] gives [[1,2,3,4],[2,3,4,1],[3,4,1,2],[4,1,2,3]]
*/
func Rotate(input []int) [][]int {
	return RotateHelper(input, [][]int{}, 0)
}

func RotateHelper(src []int, dest [][]int, accumulator int) [][]int {
	if accumulator == len(src) {
		return dest
	}
	return RotateHelper(append(src[1:len(src)], []int{src[0]}...),
		append(dest, [][]int{src}...), accumulator+1)
}

/*
permTail gives permutations of the list's tail. For instance,
permTail [1,2,3,4] gives [[1,2,3,4],[1,3,4,2],[1,4,2,3]]
*/
func PermTail(input []int) [][]int {
	result := make([][]int, 0, len(input[1:]))
	for _, tailRotation := range Rotate(input[1:]) {
		result = append(result, append([]int{input[0]}, tailRotation...))

	}
	return result
}