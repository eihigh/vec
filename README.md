# vec

[![Go Reference](https://pkg.go.dev/badge/github.com/eihigh/vec.svg)](https://pkg.go.dev/github.com/eihigh/vec)

A Go vector math library with generic types.

## Quick Start

```go
import "github.com/eihigh/vec"

// Basic usage
a := vec.Vec2{10, 20}
b := vec.Vec2{3, 4}

c := a.Add(b)        // {13, 24}
d := a.Scale(2)      // {20, 40} - scalar multiplication
e := a.Normalize()   // unit vector

// Different types
vi := vec.Vec2i{10, 20}  // int vectors
vf := vec.Vec2g[float32]{10.5, 20.5} // generic vectors
```

## Features

### Functional Composition

Combine vectors with any function using `Map`, `Map2`, and `Apply`:

```go
v := vec.Vec2{-1.5, 2.7}
v.Map(math.Abs)    // {1.5, 2.7}
v.Map(math.Floor)  // {-2, 2}

a := vec.Vec2{10, 20}
b := vec.Vec2{3, 7}
a.Map2(b, math.Max)  // {10, 20}

rgb := vec.Vec3g[uint8]{255, 128, 64}
rgb.Apply(color.RGBToYCbCr) // {159 75 197}
```

### Component Unpacking

Pass vector components directly to multi-argument functions:

```go
v := vec.Vec2{3, 4}
hypot := math.Hypot(v.XY())  // 5.0

// Works with any function expecting x, y
func drawPoint(x, y float64) { /* ... */ }
drawPoint(v.XY())
```

### Type Interoperability

Package functions accept any struct with matching components:

```go
import "image"

p := image.Point{X: 10, Y: 20}
q := image.Point{X: 3, Y: 4}

// Use standard library types with vec functions
d := vec.Dot2(p, q)  // 110
```

## Notes

### Methods vs Package Functions

Basic operations are methods, while geometric operations are package functions:

```go
// Methods: arithmetic and vector properties
v.Add(w)      // vector arithmetic
v.Len()       // length
v.Normalize() // normalization

// Package functions: geometric operations
vec.Dot2(a, b)     // dot product
vec.Cross3(a, b)   // cross product
vec.Lerp2(a, b, t) // interpolation
```

### Scalar Operations with "s" Suffix

Operations with scalars require an "s" suffix:

```go
v.Add(w)   // vector + vector
v.Adds(5)  // vector + scalar (adds to each component)

v.Mul(w)   // component-wise multiplication (Hadamard product)
v.Muls(5)  // scalar multiplication (or use v.Scale(5))
```

### Mul is Component-wise

`Mul` performs component-wise multiplication, not scalar multiplication:

```go
a := vec.Vec2{2, 3}
b := vec.Vec2{4, 5}
a.Mul(b)  // {8, 15} - NOT scalar multiplication
a.Muls(2) // {4, 6}  - scalar multiplication
```

### Cross2 Returns Scalar

The 2D cross product returns a scalar (signed area):

```go
a := vec.Vec2{3, 0}
b := vec.Vec2{0, 4}
vec.Cross2(a, b)  // returns 12.0 (float64)
```

## Examples

### Arithmetic Operations

```go
a := vec.Vec2{10, 20}
b := vec.Vec2{3, 4}

// Vector operations
a.Add(b)  // {13, 24}
a.Sub(b)  // {7, 16}
a.Mul(b)  // {30, 80} - component-wise
a.Div(b)  // {3.33.., 5}

// Scalar operations
a.Adds(2)  // {12, 22}
a.Subs(2)  // {8, 18}
a.Muls(2)  // {20, 40}
a.Divs(2)  // {5, 10}

// Other
a.Neg()    // {-10, -20}
a.Eq(b)    // false
```

### Geometry and Transformations

```go
// Length and normalization
v := vec.Vec2{3, 4}
v.Len()       // 5.0
v.LenSq()     // 25
v.Normalize() // {0.6, 0.8}

// Dot and cross products
vec.Dot2(a, b)   // dot product
vec.Cross2(a, b) // 2D cross product (scalar)
vec.Cross3(a, b) // 3D cross product (vector)

// Transformations
vec.Lerp2(a, b, 0.5)          // linear interpolation
vec.Rotate2(v, math.Pi/2)     // rotate by 90 degrees
vec.Slerp3(a, b, 0.5)         // spherical interpolation

// Projections and reflections
vec.Project2(v, normal)  // project v onto normal
vec.Reflect2(v, normal)  // reflect v off surface
```

### Type Conversions

```go
// Type conversions
v := vec.Vec2{3.7, 4.2}
v.Int()     // vec.Vec2i{3, 4}
v.Float32() // vec.Vec2g[float32]{3.7, 4.2}

// Generic vector conversions
vec.Cast2[uint8](v) // vec.Vec2g[uint8]{3, 4}

// Dimension conversions
v2 := vec.Vec2{1, 2}
v2.Vec3(3)  // vec.Vec3{1, 2, 3}

v3 := vec.Vec3{3, 4, 5}
v3.Vec2()   // vec.Vec2{3, 4}

// Array/slice conversions
v.Array()  // [2]float64{3.7, 4.2}
v.Slice()  // []float64{3.7, 4.2}
x, y := v.XY()  // 3.7, 4.2
```

## Types

- `Vec2`, `Vec3`, `Vec4` - float64 vectors (default)
- `Vec2i`, `Vec3i`, `Vec4i` - int vectors
- `Vec2u`, `Vec3u`, `Vec4u` - uint vectors
- `Vec2g[T]`, `Vec3g[T]`, `Vec4g[T]` - generic vectors for any scalar type

## License

MIT
