package main

import (
	"fmt"
	"math/rand"
	t "time"
)

func mainBasic() {
	fmt.Println("The time is", t.Now())
	// => The time is 2018-09-11 11:17:00.604665822 -0500 CDT m=+0.000201955

	fmt.Println("Randomness is everywhere", rand.Intn(64))
	// => Randomness is everywhere 35
}
