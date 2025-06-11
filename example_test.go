package vec_test

import (
	"fmt"
	"image/color"
	"math"

	"github.com/eihigh/vec"
)

func Example_arithmetic() {
	// Vector-vector arithmetic operations
	a := vec.Vec2{10, 20}
	b := vec.Vec2{3, 4}
	fmt.Println("Vector operations:")
	fmt.Println("a + b =", a.Add(b)) // vector addition
	fmt.Println("a - b =", a.Sub(b)) // vector subtraction
	fmt.Println("a * b =", a.Mul(b)) // component-wise multiplication
	fmt.Println("a / b =", a.Div(b)) // component-wise division

	// Vector-scalar arithmetic operations
	fmt.Println("\nScalar operations:")
	fmt.Println("a + 2 =", a.Adds(2)) // add scalar to each component
	fmt.Println("a - 2 =", a.Subs(2)) // subtract scalar from each component
	fmt.Println("a * 2 =", a.Muls(2)) // multiply each component by scalar
	fmt.Println("a / 2 =", a.Divs(2)) // divide each component by scalar

	// Other operations
	fmt.Println("\nOther operations:")
	fmt.Println("-a =", a.Neg())     // negation
	fmt.Println("a == b =", a.Eq(b)) // equality check

	// Output:
	// Vector operations:
	// a + b = {13 24}
	// a - b = {7 16}
	// a * b = {30 80}
	// a / b = {3.3333333333333335 5}
	//
	// Scalar operations:
	// a + 2 = {12 22}
	// a - 2 = {8 18}
	// a * 2 = {20 40}
	// a / 2 = {5 10}
	//
	// Other operations:
	// -a = {-10 -20}
	// a == b = false
}

func Example_geometryAndTransformations() {
	// Geometric operations
	v1 := vec.Vec2{3, 4}
	v2 := vec.Vec2{1, 0}

	fmt.Println("Geometric operations:")
	fmt.Printf("Length of %v = %.2f\n", v1, vec.Len2(v1))
	fmt.Printf("Squared length = %.0f\n", vec.LenSq2(v1))
	fmt.Printf("Normalized = %v\n", vec.Normalize2(v1))
	fmt.Printf("Angle = %.2f radians\n", vec.Angle2(v1))

	// Dot and cross products
	fmt.Println("\nDot and cross products:")
	fmt.Printf("Dot product = %.0f\n", vec.Dot2(v1, v2))
	fmt.Printf("Cross product (2D) = %.0f\n", vec.Cross2(v1, v2))

	// 3D operations
	v3 := vec.Vec3{1, 2, 3}
	v4 := vec.Vec3{4, 5, 6}
	fmt.Printf("Cross product (3D) = %v\n", vec.Cross3(v3, v4))

	// Projections and reflections
	fmt.Println("\nProjections and reflections:")
	normal := vec.Vec2{0, 1}
	incident := vec.Vec2{1, -1}
	fmt.Printf("Project %v onto %v = %v\n", incident, normal, vec.Project2(incident, normal))
	fmt.Printf("Reflect %v off %v = %v\n", incident, normal, vec.Reflect2(incident, normal))

	// Transformations
	fmt.Println("\nTransformations:")
	a := vec.Vec2{1, 0}
	b := vec.Vec2{0, 1}
	fmt.Printf("Lerp(a, b, 0.5) = %v\n", vec.Lerp2(a, b, 0.5))
	fmt.Printf("Rotate %v by π/2 = %v\n", a, vec.Rotate2(a, math.Pi/2))

	// Spherical interpolation
	s1 := vec.Vec3{1, 0, 0}
	s2 := vec.Vec3{0, 1, 0}
	fmt.Printf("Slerp(s1, s2, 0.5) = %v\n", vec.Slerp3(s1, s2, 0.5))

	// Output:
	// Geometric operations:
	// Length of {3 4} = 5.00
	// Squared length = 25
	// Normalized = {0.6 0.8}
	// Angle = 0.93 radians
	//
	// Dot and cross products:
	// Dot product = 3
	// Cross product (2D) = -4
	// Cross product (3D) = {-3 6 -3}
	//
	// Projections and reflections:
	// Project {1 -1} onto {0 1} = {-0 -1}
	// Reflect {1 -1} off {0 1} = {1 1}
	//
	// Transformations:
	// Lerp(a, b, 0.5) = {0.5 0.5}
	// Rotate {1 0} by π/2 = {6.123233995736757e-17 1}
	// Slerp(s1, s2, 0.5) = {0.7071067811865475 0.7071067811865475 0}
}

func Example_constructorsAndUtilities() {
	// Constructors
	fmt.Println("Constructors:")
	fmt.Printf("New2(3, 4) = %v\n", vec.New2(3, 4))
	fmt.Printf("Splat2(5) = %v\n", vec.Splat2(5))

	// Type conversions
	v := vec.Vec2{-3.7, 4.2}
	fmt.Println("\nType conversions:")
	fmt.Printf("Original: %v\n", v)
	fmt.Printf("To int: %v\n", v.Int())
	fmt.Printf("To float32: %v\n", v.Float32())
	fmt.Printf("To uint8: %v\n", vec.Cast2[uint8](v))

	// Dimension conversions
	v2 := vec.Vec2{1, 2}
	v3 := vec.Vec3{3, 4, 5}
	fmt.Println("\nDimension conversions:")
	fmt.Printf("2D to 3D: %v -> %v\n", v2, v2.Vec3(3))
	fmt.Printf("3D to 2D: %v -> %v\n", v3, v3.Vec2())

	// Array/slice conversions
	fmt.Println("\nArray/slice conversions:")
	fmt.Printf("To array: %v\n", v2.Array())
	fmt.Printf("To slice: %v\n", v2.Slice())
	x, y := v2.XY()
	fmt.Printf("Components: x=%v, y=%v\n", x, y)

	// Functional operations
	fmt.Println("\nFunctional operations:")
	v4 := vec.Vec2{-1.5, 2.7}
	fmt.Printf("Map(Abs): %v -> %v\n", v4, vec.Map2(v4, math.Abs))

	v5 := vec.Vec2{10, 20}
	v6 := vec.Vec2{3, 7}
	fmt.Printf("Zip(Max): %v, %v -> %v\n", v5, v6, vec.Zip2(v5, v6, math.Max))

	// Apply a function to the vector
	// that takes multiple scalars (e.g., color conversion)
	rgb := vec.Vec3g[uint8]{255, 128, 64}
	ycbcr := vec.Apply3(rgb, color.RGBToYCbCr)
	fmt.Printf("Apply(RGBToYCbCr): %v -> %v\n", rgb, ycbcr)

	// Output:
	// Constructors:
	// New2(3, 4) = {3 4}
	// Splat2(5) = {5 5}
	//
	// Type conversions:
	// Original: {-3.7 4.2}
	// To int: {-3 4}
	// To float32: {-3.7 4.2}
	// To uint8: {253 4}
	//
	// Dimension conversions:
	// 2D to 3D: {1 2} -> {1 2 3}
	// 3D to 2D: {3 4 5} -> {3 4}
	//
	// Array/slice conversions:
	// To array: [1 2]
	// To slice: [1 2]
	// Components: x=1, y=2
	//
	// Functional operations:
	// Map(Abs): {-1.5 2.7} -> {1.5 2.7}
	// Zip(Max): {10 20}, {3 7} -> {10 20}
	// Apply(RGBToYCbCr): {255 128 64} -> {159 75 197}
}
