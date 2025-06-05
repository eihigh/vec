package vec_test

import (
	"fmt"
	"math"

	"github.com/eihigh/vec"
)

func Example_map() {
	v := vec.Vec2{3.14, 2.71}
	v = v.Map(math.Floor)
	fmt.Println(v)

	// Output:
	// {3 2}
}
