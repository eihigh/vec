package vec

import "math"

// ====================
// Types
// ====================

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Scalar interface {
	Integer | Float
}

type (
	Vec2g[S Scalar] struct{ X, Y S }
	Vec3g[S Scalar] struct{ X, Y, Z S }
	Vec4g[S Scalar] struct{ X, Y, Z, W S }

	Vec2  = Vec2g[float64]
	Vec3  = Vec3g[float64]
	Vec4  = Vec4g[float64]
	Vec2i = Vec2g[int]
	Vec3i = Vec3g[int]
	Vec4i = Vec4g[int]
	Vec2u = Vec2g[uint]
	Vec3u = Vec3g[uint]
	Vec4u = Vec4g[uint]

	Vec2like[S Scalar] interface {
		~struct{ X, Y S }
	}
	Vec3like[S Scalar] interface {
		~struct{ X, Y, Z S }
	}
	Vec4like[S Scalar] interface {
		~struct{ X, Y, Z, W S }
	}
)

// ====================
// Utility API
// Convenience functions for creating vectors, converting between types,
// and applying functions to vector components.
// ====================

// New2 creates a 2D vector from x, y components.
func New2[S Scalar](x, y S) Vec2g[S] { return Vec2g[S]{x, y} }

// New3 creates a 3D vector from x, y, z components.
func New3[S Scalar](x, y, z S) Vec3g[S] { return Vec3g[S]{x, y, z} }

// New4 creates a 4D vector from x, y, z, w components.
func New4[S Scalar](x, y, z, w S) Vec4g[S] { return Vec4g[S]{x, y, z, w} }

// Splat2 creates a 2D vector with all components set to x.
func Splat2[S Scalar](x S) Vec2g[S] { return Vec2g[S]{x, x} }

// Splat3 creates a 3D vector with all components set to x.
func Splat3[S Scalar](x S) Vec3g[S] { return Vec3g[S]{x, x, x} }

// Splat4 creates a 4D vector with all components set to x.
func Splat4[S Scalar](x S) Vec4g[S] { return Vec4g[S]{x, x, x, x} }

// Cast2 converts a 2D vector from type S to type T.
func Cast2[T, S Scalar, V Vec2like[S]](v V) Vec2g[T] {
	vv := Vec2g[S](v)
	return Vec2g[T]{T(vv.X), T(vv.Y)}
}

// Cast3 converts a 3D vector from type S to type T.
func Cast3[T, S Scalar, V Vec3like[S]](v V) Vec3g[T] {
	vv := Vec3g[S](v)
	return Vec3g[T]{T(vv.X), T(vv.Y), T(vv.Z)}
}

// Cast4 converts a 4D vector from type S to type T.
func Cast4[T, S Scalar, V Vec4like[S]](v V) Vec4g[T] {
	vv := Vec4g[S](v)
	return Vec4g[T]{T(vv.X), T(vv.Y), T(vv.Z), T(vv.W)}
}

// XY returns the x, y components.
func (a Vec2g[S]) XY() (x, y S) { return a.X, a.Y }

// XYZ returns the x, y, z components.
func (a Vec3g[S]) XYZ() (x, y, z S) { return a.X, a.Y, a.Z }

// XYZW returns the x, y, z, w components.
func (a Vec4g[S]) XYZW() (x, y, z, w S) { return a.X, a.Y, a.Z, a.W }

// Vec3 extends to 3D by appending z.
func (a Vec2g[S]) Vec3(z S) Vec3g[S] { return Vec3g[S]{a.X, a.Y, z} }

// Vec4 extends to 4D by appending z, w.
func (a Vec2g[S]) Vec4(z, w S) Vec4g[S] { return Vec4g[S]{a.X, a.Y, z, w} }

// Vec2 truncates to 2D.
func (a Vec3g[S]) Vec2() Vec2g[S] { return Vec2g[S]{a.X, a.Y} }

// Vec4 extends to 4D by appending w.
func (a Vec3g[S]) Vec4(w S) Vec4g[S] { return Vec4g[S]{a.X, a.Y, a.Z, w} }

// Vec2 truncates to 2D.
func (a Vec4g[S]) Vec2() Vec2g[S] { return Vec2g[S]{a.X, a.Y} }

// Vec3 truncates to 3D.
func (a Vec4g[S]) Vec3() Vec3g[S] { return Vec3g[S]{a.X, a.Y, a.Z} }

// Int converts to int vector.
func (a Vec2g[S]) Int() Vec2g[int] { return Vec2g[int]{int(a.X), int(a.Y)} }

// Int converts to int vector.
func (a Vec3g[S]) Int() Vec3g[int] { return Vec3g[int]{int(a.X), int(a.Y), int(a.Z)} }

// Int converts to int vector.
func (a Vec4g[S]) Int() Vec4g[int] { return Vec4g[int]{int(a.X), int(a.Y), int(a.Z), int(a.W)} }

// Uint converts to uint vector.
func (a Vec2g[S]) Uint() Vec2g[uint] { return Vec2g[uint]{uint(a.X), uint(a.Y)} }

// Uint converts to uint vector.
func (a Vec3g[S]) Uint() Vec3g[uint] { return Vec3g[uint]{uint(a.X), uint(a.Y), uint(a.Z)} }

// Uint converts to uint vector.
func (a Vec4g[S]) Uint() Vec4g[uint] { return Vec4g[uint]{uint(a.X), uint(a.Y), uint(a.Z), uint(a.W)} }

// Float64 converts to float64 vector.
func (a Vec2g[S]) Float64() Vec2g[float64] {
	return Vec2g[float64]{float64(a.X), float64(a.Y)}
}

// Float64 converts to float64 vector.
func (a Vec3g[S]) Float64() Vec3g[float64] {
	return Vec3g[float64]{float64(a.X), float64(a.Y), float64(a.Z)}
}

// Float64 converts to float64 vector.
func (a Vec4g[S]) Float64() Vec4g[float64] {
	return Vec4g[float64]{float64(a.X), float64(a.Y), float64(a.Z), float64(a.W)}
}

// Float32 converts to float32 vector.
func (a Vec2g[S]) Float32() Vec2g[float32] {
	return Vec2g[float32]{float32(a.X), float32(a.Y)}
}

// Float32 converts to float32 vector.
func (a Vec3g[S]) Float32() Vec3g[float32] {
	return Vec3g[float32]{float32(a.X), float32(a.Y), float32(a.Z)}
}

// Float32 converts to float32 vector.
func (a Vec4g[S]) Float32() Vec4g[float32] {
	return Vec4g[float32]{float32(a.X), float32(a.Y), float32(a.Z), float32(a.W)}
}

// Array returns components as array.
func (a Vec2g[S]) Array() [2]S { return [2]S{a.X, a.Y} }

// Array returns components as array.
func (a Vec3g[S]) Array() [3]S { return [3]S{a.X, a.Y, a.Z} }

// Array returns components as array.
func (a Vec4g[S]) Array() [4]S { return [4]S{a.X, a.Y, a.Z, a.W} }

// Slice returns components as slice.
func (a Vec2g[S]) Slice() []S { return []S{a.X, a.Y} }

// Slice returns components as slice.
func (a Vec3g[S]) Slice() []S { return []S{a.X, a.Y, a.Z} }

// Slice returns components as slice.
func (a Vec4g[S]) Slice() []S { return []S{a.X, a.Y, a.Z, a.W} }

// ===================
// Math API (package functions)
// Multiple vector operations are defined as global functions.
// ===================

// Dot2 returns the dot product of two 2D vectors.
func Dot2[S Scalar, V Vec2like[S]](a, b V) S {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return va.X*vb.X + va.Y*vb.Y
}

// Dot3 returns the dot product of two 3D vectors.
func Dot3[S Scalar, V Vec3like[S]](a, b V) S {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return va.X*vb.X + va.Y*vb.Y + va.Z*vb.Z
}

// Dot4 returns the dot product of two 4D vectors.
func Dot4[S Scalar, V Vec4like[S]](a, b V) S {
	va := Vec4g[S](a)
	vb := Vec4g[S](b)
	return va.X*vb.X + va.Y*vb.Y + va.Z*vb.Z + va.W*vb.W
}

// Lerp2 linearly interpolates between a and b by t.
func Lerp2[S Scalar, V Vec2like[S]](a, b V, t float64) Vec2g[S] {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return Vec2g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
	}
}

// Lerp3 linearly interpolates between a and b by t.
func Lerp3[S Scalar, V Vec3like[S]](a, b V, t float64) Vec3g[S] {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return Vec3g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
		Z: S(float64(va.Z) + (float64(vb.Z)-float64(va.Z))*t),
	}
}

// Lerp4 linearly interpolates between a and b by t.
func Lerp4[S Scalar, V Vec4like[S]](a, b V, t float64) Vec4g[S] {
	va := Vec4g[S](a)
	vb := Vec4g[S](b)
	return Vec4g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
		Z: S(float64(va.Z) + (float64(vb.Z)-float64(va.Z))*t),
		W: S(float64(va.W) + (float64(vb.W)-float64(va.W))*t),
	}
}

// Project2 projects v onto onNormal.
func Project2[S Scalar, V Vec2like[S]](v, onNormal V) Vec2g[S] {
	va := Vec2g[S](v)
	vn := Vec2g[S](onNormal).Normalize()
	return vn.Scale(Dot2(va, vn))
}

// Project3 projects v onto onNormal.
func Project3[S Scalar, V Vec3like[S]](v, onNormal V) Vec3g[S] {
	va := Vec3g[S](v)
	vn := Vec3g[S](onNormal).Normalize()
	return vn.Scale(Dot3(va, vn))
}

// Project4 projects v onto onNormal.
func Project4[S Scalar, V Vec4like[S]](v, onNormal V) Vec4g[S] {
	va := Vec4g[S](v)
	vn := Vec4g[S](onNormal).Normalize()
	return vn.Scale(Dot4(va, vn))
}

// Reflect2 reflects v off the surface with normal n.
func Reflect2[S Scalar, V Vec2like[S]](v, normal V) Vec2g[S] {
	va := Vec2g[S](v)
	vn := Vec2g[S](normal)
	dot := Dot2(va, vn)
	return Vec2g[S]{
		X: va.X - S(2*float64(dot)*float64(vn.X)),
		Y: va.Y - S(2*float64(dot)*float64(vn.Y)),
	}
}

// Reflect3 reflects v off the surface with normal n.
func Reflect3[S Scalar, V Vec3like[S]](v, normal V) Vec3g[S] {
	va := Vec3g[S](v)
	vn := Vec3g[S](normal)
	dot := Dot3(va, vn)
	return Vec3g[S]{
		X: va.X - S(2*float64(dot)*float64(vn.X)),
		Y: va.Y - S(2*float64(dot)*float64(vn.Y)),
		Z: va.Z - S(2*float64(dot)*float64(vn.Z)),
	}
}

// Cross2 returns the 2D cross product (scalar).
func Cross2[S Scalar, V Vec2like[S]](a, b V) float64 {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return float64(va.X*vb.Y - va.Y*vb.X)
}

// Cross3 returns the 3D cross product.
func Cross3[S Scalar, V Vec3like[S]](a, b V) Vec3g[S] {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return Vec3g[S]{
		X: va.Y*vb.Z - va.Z*vb.Y,
		Y: va.Z*vb.X - va.X*vb.Z,
		Z: va.X*vb.Y - va.Y*vb.X,
	}
}

// Slerp3 spherically interpolates between a and b by t.
func Slerp3[S Scalar, V Vec3like[S]](a, b V, t float64) Vec3g[S] {
	va := Vec3g[S](a).Normalize()
	vb := Vec3g[S](b).Normalize()

	dot := float64(Dot3(va, vb))

	// Clamp dot product to avoid numerical errors
	if dot > 0.9995 {
		// Vectors are very close, use linear interpolation
		return Lerp3(va, vb, t).Normalize()
	}

	if dot < -1 {
		dot = -1
	} else if dot > 1 {
		dot = 1
	}

	theta := math.Acos(dot)
	sinTheta := math.Sin(theta)

	if sinTheta < 0.001 {
		// Vectors are parallel, use linear interpolation
		return Lerp3(va, vb, t).Normalize()
	}

	a1 := math.Sin((1-t)*theta) / sinTheta
	b1 := math.Sin(t*theta) / sinTheta

	return Vec3g[S]{
		X: S(float64(va.X)*a1 + float64(vb.X)*b1),
		Y: S(float64(va.Y)*a1 + float64(vb.Y)*b1),
		Z: S(float64(va.Z)*a1 + float64(vb.Z)*b1),
	}
}

// Rotate2 rotates v by angle radians.
func Rotate2[S Scalar, V Vec2like[S]](v V, angle float64) Vec2g[S] {
	va := Vec2g[S](v)
	sin, cos := math.Sincos(angle)
	return Vec2g[S]{
		X: S(float64(va.X)*cos - float64(va.Y)*sin),
		Y: S(float64(va.X)*sin + float64(va.Y)*cos),
	}
}

// Map2 applies f to each component of a 2D vector.
func Map2[S Scalar, V Vec2like[S]](v V, f func(S) S) Vec2g[S] {
	va := Vec2g[S](v)
	return Vec2g[S]{f(va.X), f(va.Y)}
}

// Map3 applies f to each component of a 3D vector.
func Map3[S Scalar, V Vec3like[S]](v V, f func(S) S) Vec3g[S] {
	va := Vec3g[S](v)
	return Vec3g[S]{f(va.X), f(va.Y), f(va.Z)}
}

// Map4 applies f to each component of a 4D vector.
func Map4[S Scalar, V Vec4like[S]](v V, f func(S) S) Vec4g[S] {
	va := Vec4g[S](v)
	return Vec4g[S]{f(va.X), f(va.Y), f(va.Z), f(va.W)}
}

// Zip2 applies f to corresponding components of two 2D vectors.
func Zip2[S Scalar, V Vec2like[S]](a, b V, f func(S, S) S) Vec2g[S] {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return Vec2g[S]{f(va.X, vb.X), f(va.Y, vb.Y)}
}

// Zip3 applies f to corresponding components of two 3D vectors.
func Zip3[S Scalar, V Vec3like[S]](a, b V, f func(S, S) S) Vec3g[S] {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return Vec3g[S]{f(va.X, vb.X), f(va.Y, vb.Y), f(va.Z, vb.Z)}
}

// Zip4 applies f to corresponding components of two 4D vectors.
func Zip4[S Scalar, V Vec4like[S]](a, b V, f func(S, S) S) Vec4g[S] {
	va := Vec4g[S](a)
	vb := Vec4g[S](b)
	return Vec4g[S]{f(va.X, vb.X), f(va.Y, vb.Y), f(va.Z, vb.Z), f(va.W, vb.W)}
}

// Apply2 transforms all components of a 2D vector at once.
func Apply2[S Scalar, V Vec2like[S]](v V, f func(S, S) (S, S)) Vec2g[S] {
	va := Vec2g[S](v)
	x, y := f(va.X, va.Y)
	return Vec2g[S]{x, y}
}

// Apply3 transforms all components of a 3D vector at once.
func Apply3[S Scalar, V Vec3like[S]](v V, f func(S, S, S) (S, S, S)) Vec3g[S] {
	va := Vec3g[S](v)
	x, y, z := f(va.X, va.Y, va.Z)
	return Vec3g[S]{x, y, z}
}

// Apply4 transforms all components of a 4D vector at once.
func Apply4[S Scalar, V Vec4like[S]](v V, f func(S, S, S, S) (S, S, S, S)) Vec4g[S] {
	va := Vec4g[S](v)
	x, y, z, w := f(va.X, va.Y, va.Z, va.W)
	return Vec4g[S]{x, y, z, w}
}

// ===================
// Math API (methods)
// The arithmetic operations, comparisons, and operations on
// individual vectors (length, angle) are defined as methods.
// ===================

// Vec2
// ---

// Add returns the vector a+b.
func (a Vec2g[S]) Add(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X + b.X, a.Y + b.Y} }

// Adds returns the vector a+(s, s).
func (a Vec2g[S]) Adds(s S) Vec2g[S] { return Vec2g[S]{a.X + s, a.Y + s} }

// Sub returns the vector a-b.
func (a Vec2g[S]) Sub(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X - b.X, a.Y - b.Y} }

// Subs returns the vector a-(s, s).
func (a Vec2g[S]) Subs(s S) Vec2g[S] { return Vec2g[S]{a.X - s, a.Y - s} }

// Mul returns the component-wise product a*b.
func (a Vec2g[S]) Mul(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X * b.X, a.Y * b.Y} }

// Muls returns the vector a*s.
func (a Vec2g[S]) Muls(s S) Vec2g[S] { return Vec2g[S]{a.X * s, a.Y * s} }

// Div returns the component-wise quotient a/b.
func (a Vec2g[S]) Div(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X / b.X, a.Y / b.Y} }

// Divs returns the vector a/s.
func (a Vec2g[S]) Divs(s S) Vec2g[S] { return Vec2g[S]{a.X / s, a.Y / s} }

// Neg returns the negated vector -a.
func (a Vec2g[S]) Neg() Vec2g[S] { return Vec2g[S]{-a.X, -a.Y} }

// Eq returns whether a equals b.
func (a Vec2g[S]) Eq(b Vec2g[S]) bool { return a.X == b.X && a.Y == b.Y }

// LenSq returns the squared length.
func (a Vec2g[S]) LenSq() S { return a.X*a.X + a.Y*a.Y }

// Len returns the length.
func (a Vec2g[S]) Len() float64 { return math.Sqrt(float64(a.X*a.X + a.Y*a.Y)) }

// Angle returns the angle in radians.
func (a Vec2g[S]) Angle() float64 { return math.Atan2(float64(a.Y), float64(a.X)) }

// Normalize returns the unit vector.
func (a Vec2g[S]) Normalize() Vec2g[S] {
	l := a.Len()
	if l == 0 {
		return Vec2g[S]{0, 0}
	}
	return Vec2g[S]{a.X / S(l), a.Y / S(l)}
}

// Scale is an alias for Muls.
func (a Vec2g[S]) Scale(s S) Vec2g[S] { return Vec2g[S]{a.X * s, a.Y * s} }

// Vec3
// ---

// Add returns the vector a+b.
func (a Vec3g[S]) Add(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Adds returns the vector a+(s, s, s).
func (a Vec3g[S]) Adds(s S) Vec3g[S] {
	return Vec3g[S]{a.X + s, a.Y + s, a.Z + s}
}

// Sub returns the vector a-b.
func (a Vec3g[S]) Sub(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Subs returns the vector a-(s, s, s).
func (a Vec3g[S]) Subs(s S) Vec3g[S] {
	return Vec3g[S]{a.X - s, a.Y - s, a.Z - s}
}

// Mul returns the component-wise product a*b.
func (a Vec3g[S]) Mul(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

// Muls returns the vector a*s.
func (a Vec3g[S]) Muls(s S) Vec3g[S] {
	return Vec3g[S]{a.X * s, a.Y * s, a.Z * s}
}

// Div returns the component-wise quotient a/b.
func (a Vec3g[S]) Div(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

// Divs returns the vector a/s.
func (a Vec3g[S]) Divs(s S) Vec3g[S] {
	return Vec3g[S]{a.X / s, a.Y / s, a.Z / s}
}

// Neg returns the negated vector -a.
func (a Vec3g[S]) Neg() Vec3g[S] {
	return Vec3g[S]{-a.X, -a.Y, -a.Z}
}

// Eq returns whether a equals b.
func (a Vec3g[S]) Eq(b Vec3g[S]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

// LenSq returns the squared length.
func (a Vec3g[S]) LenSq() S {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

// Len returns the length.
func (a Vec3g[S]) Len() float64 {
	return math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z))
}

// Normalize returns the unit vector.
func (a Vec3g[S]) Normalize() Vec3g[S] {
	l := a.Len()
	if l == 0 {
		return Vec3g[S]{0, 0, 0}
	}
	return Vec3g[S]{a.X / S(l), a.Y / S(l), a.Z / S(l)}
}

// Scale is an alias for Muls.
func (a Vec3g[S]) Scale(s S) Vec3g[S] {
	return Vec3g[S]{a.X * s, a.Y * s, a.Z * s}
}

// Vec4
// ---
// Add returns the vector a+b.
func (a Vec4g[S]) Add(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

// Adds returns the vector a+(s, s, s, s).
func (a Vec4g[S]) Adds(s S) Vec4g[S] {
	return Vec4g[S]{a.X + s, a.Y + s, a.Z + s, a.W + s}
}

// Sub returns the vector a-b.
func (a Vec4g[S]) Sub(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

// Subs returns the vector a-(s, s, s, s).
func (a Vec4g[S]) Subs(s S) Vec4g[S] {
	return Vec4g[S]{a.X - s, a.Y - s, a.Z - s, a.W - s}
}

// Mul returns the component-wise product a*b.
func (a Vec4g[S]) Mul(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W}
}

// Muls returns the vector a*s.
func (a Vec4g[S]) Muls(s S) Vec4g[S] {
	return Vec4g[S]{a.X * s, a.Y * s, a.Z * s, a.W * s}
}

// Div returns the component-wise quotient a/b.
func (a Vec4g[S]) Div(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X / b.X, a.Y / b.Y, a.Z / b.Z, a.W / b.W}
}

// Divs returns the vector a/s.
func (a Vec4g[S]) Divs(s S) Vec4g[S] {
	return Vec4g[S]{a.X / s, a.Y / s, a.Z / s, a.W / s}
}

// Neg returns the negated vector -a.
func (a Vec4g[S]) Neg() Vec4g[S] {
	return Vec4g[S]{-a.X, -a.Y, -a.Z, -a.W}
}

// Eq returns whether a equals b.
func (a Vec4g[S]) Eq(b Vec4g[S]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}

// LenSq returns the squared length.
func (a Vec4g[S]) LenSq() S {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W
}

// Len returns the length.
func (a Vec4g[S]) Len() float64 {
	return math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W))
}

// Normalize returns the unit vector.
func (a Vec4g[S]) Normalize() Vec4g[S] {
	l := a.Len()
	if l == 0 {
		return Vec4g[S]{0, 0, 0, 0}
	}
	return Vec4g[S]{a.X / S(l), a.Y / S(l), a.Z / S(l), a.W / S(l)}
}

// Scale is an alias for Muls.
func (a Vec4g[S]) Scale(s S) Vec4g[S] {
	return Vec4g[S]{a.X * s, a.Y * s, a.Z * s, a.W * s}
}
