# vec

[![Go Reference](https://pkg.go.dev/badge/github.com/eihigh/vec.svg)](https://pkg.go.dev/github.com/eihigh/vec)

A Go vector math library with generic types.

## Quick Start

`Vec2` is an alias for `Vec2g[float64]`. Other aliases include `Vec2i`, `Vec2u`.

Basic arithmetic operations, comparisons, and type/size conversions are provided as methods, while other operations are available as global functions.

As will be seen later, global functions accept vector types defined in any package, and type conversions are flexible.

```go
import "github.com/eihigh/vec"

// Basic usage
a := vec.Vec2{10, 20}
b := vec.Vec2{3, 4}

c := a.Add(b)        // {13, 24}
d := a.Scale(2)      // {20, 40} - scalar multiplication
e := vec.Normalize2(a)   // unit vector

// Different types
vi := vec.Vec2i{10, 20}  // int vectors
vf := vec.Vec2g[float32]{10.5, 20.5} // generic vectors
```

## Interoperability

### Any Vector Types

Package functions work seamlessly with any vector types from any package:

```go
import "image"

p := image.Point{X: 10, Y: 20}
q := image.Point{X: 3, Y: 4}

// Use image.Point directly with vec functions
d := vec.Dot2(p, q) // 110
```

### Flexible Type Conversions

Create and convert vectors between different numeric types:

```go
// Basic constructors
v1 := vec.New2(3.14, 2.71)           // vec.Vec2{3.14, 2.71}
v2 := vec.New3(1, 2, 3)              // vec.Vec3i{1, 2, 3}

// Type conversion during construction
vi := vec.NewAs2[int](3.14, 2.71)    // vec.Vec2i{3, 2}
vf := vec.NewAs3[float32](1, 2, 3)   // vec.Vec3g[float32]{1, 2, 3}

// Convert existing vectors
p := image.Point{X: 10, Y: 20}
vec.As2[float32](p) // vec.Vec2g[float32]{10, 20}
```

### Component Access and Integration

Unpack vector components for use with any function:

```go
v := vec.Vec2{3, 4}

// Unpack for multi-argument functions
hypot := math.Hypot(v.XY())  // 5.0

// Works with any function signature
func drawPoint(x, y float64) { /* ... */ }
drawPoint(v.XY())
```

### Scalar Operations

Perform arithmetic operations with scalars on all components:

```go
v := vec.Vec2{3, 4}

v.Adds(10)  // {13, 14} - add scalar to all components
v.Muls(2)   // {6, 8}   - multiply all components by scalar
v.Divs(2)   // {1.5, 2} - divide all components by scalar
```

## Quick Reference

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
a.Eqs(10)  // false
```

### Geometry and Transformations

```go
// Length and normalization
v := vec.Vec2{3, 4}
vec.Len2(v)       // 5.0
vec.LenSq2(v)     // 25
vec.Normalize2(v) // {0.6, 0.8}

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
vec.As2[int](v)     // vec.Vec2i{3, 4}
vec.As2[float32](v) // vec.Vec2g[float32]{3.7, 4.2}
vec.As2[uint8](v)   // vec.Vec2g[uint8]{3, 4}

// Dimension conversions
v2 := vec.Vec2{1, 2}
v2.Vec3(3)  // vec.Vec3{1, 2, 3}

v3 := vec.Vec3{3, 4, 5}
v3.Vec2()   // vec.Vec2{3, 4}

// Array/slice conversions
vec.ToArray2(v)  // [2]float64{3.7, 4.2}
vec.ToSlice2(v)  // []float64{3.7, 4.2}
x, y := v.XY()   // 3.7, 4.2
```

### Functional Operations

```go
// Map: apply function to each component
v := vec.Vec2{-1.5, 2.7}
vec.Map2(v, math.Abs)    // {1.5, 2.7}
vec.Map2(v, math.Floor)  // {-2, 2}

// Zip: combine two vectors with a function
a := vec.Vec2{10, 20}
b := vec.Vec2{3, 7}
vec.Zip2(a, b, math.Max)  // {10, 20}

// Apply: transform all components at once
rgb := vec.Vec3g[uint8]{255, 128, 64}
vec.Apply3(rgb, color.RGBToYCbCr) // {159 75 197}
```

## Types

- `Vec2`, `Vec3`, `Vec4` - float64 vectors (default)
- `Vec2i`, `Vec3i`, `Vec4i` - int vectors
- `Vec2u`, `Vec3u`, `Vec4u` - uint vectors
- `Vec2g[T]`, `Vec3g[T]`, `Vec4g[T]` - generic vectors for any scalar type

## License

MIT
