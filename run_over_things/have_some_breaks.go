package runoverthings

import "fmt"

func demoBreak() {
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is Prime \n", i)
		}
	}
}
