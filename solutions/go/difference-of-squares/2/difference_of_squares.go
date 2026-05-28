package differenceofsquares

func SquareOfSum(n int) int {
    Sum:= 0
	for i:= range n {
        Sum += i + 1
    }
    return Sum * Sum
}

func SumOfSquares(n int) int {
    SumOfSquare := 0
    
	for i:= range n{
        SumOfSquare += (i + 1) * (i + 1)
    }

    return SumOfSquare
}

func Difference(n int) int {
    return SquareOfSum(n) - SumOfSquares(n)
}
