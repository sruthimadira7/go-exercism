package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, errors.New("Only positive numbers are allowed")
    }
    
    numSteps := 0
    
    for n != 1{
         if n % 2 == 0{
        	n = n / 2
         }else{
             n = n * 3 + 1
         }
        numSteps++
    }
    return numSteps, nil
}
