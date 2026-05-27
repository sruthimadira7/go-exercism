package hamming

import ("errors")

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("The given two DNA strands are of unequal length")
    }

    hamming_distance := 0
    
    for i := range a {
    	if a[i] != b[i] {
            hamming_distance++
        }
    }

    return hamming_distance, nil
}
