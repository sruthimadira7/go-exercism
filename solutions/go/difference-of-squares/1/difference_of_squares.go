package differenceofsquares

func SquareOfSum(n int) int {
    Sum:= 0
	for i:=1; i<=n; i++{
        Sum += i
    }
    return Sum * Sum
}

func SumOfSquares(n int) int {
    SumOfSquare := 0
    
	for i:=1; i<=n; i++{
        SumOfSquare += i * i
    }

    return SumOfSquare
}

func Difference(n int) int {
    return SquareOfSum(n) - SumOfSquares(n)
}
