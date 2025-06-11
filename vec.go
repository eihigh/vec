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

// NewAs2 creates a 2D vector from x, y components of type In, returning type Out.
func NewAs2[Out, In Scalar](x, y In) Vec2g[Out] { return Vec2g[Out]{Out(x), Out(y)} }

// NewAs3 creates a 3D vector from x, y, z components of type In, returning type Out.
func NewAs3[Out, In Scalar](x, y, z In) Vec3g[Out] {
	return Vec3g[Out]{Out(x), Out(y), Out(z)}
}

// NewAs4 creates a 4D vector from x, y, z, w components of type In, returning type Out.
func NewAs4[Out, In Scalar](x, y, z, w In) Vec4g[Out] {
	return Vec4g[Out]{Out(x), Out(y), Out(z), Out(w)}
}

// Splat2 creates a 2D vector with all components set to x.
func Splat2[S Scalar](x S) Vec2g[S] { return Vec2g[S]{x, x} }

// Splat3 creates a 3D vector with all components set to x.
func Splat3[S Scalar](x S) Vec3g[S] { return Vec3g[S]{x, x, x} }

// Splat4 creates a 4D vector with all components set to x.
func Splat4[S Scalar](x S) Vec4g[S] { return Vec4g[S]{x, x, x, x} }

// As2 converts a 2D vector from type In to type Out.
func As2[Out, In Scalar, V Vec2like[In]](v V) Vec2g[Out] {
	vv := Vec2g[In](v)
	return Vec2g[Out]{Out(vv.X), Out(vv.Y)}
}

// As3 converts a 3D vector from type In to type Out.
func As3[Out, In Scalar, V Vec3like[In]](v V) Vec3g[Out] {
	vv := Vec3g[In](v)
	return Vec3g[Out]{Out(vv.X), Out(vv.Y), Out(vv.Z)}
}

// As4 converts a 4D vector from type In to type Out.
func As4[Out, In Scalar, V Vec4like[In]](v V) Vec4g[Out] {
	vv := Vec4g[In](v)
	return Vec4g[Out]{Out(vv.X), Out(vv.Y), Out(vv.Z), Out(vv.W)}
}

// ToArray2 returns components as array.
func ToArray2[V Vec2like[S], S Scalar](v V) [2]S {
	vv := Vec2g[S](v)
	return [2]S{vv.X, vv.Y}
}

// ToArray3 returns components as array.
func ToArray3[V Vec3like[S], S Scalar](v V) [3]S {
	vv := Vec3g[S](v)
	return [3]S{vv.X, vv.Y, vv.Z}
}

// ToArray4 returns components as array.
func ToArray4[V Vec4like[S], S Scalar](v V) [4]S {
	vv := Vec4g[S](v)
	return [4]S{vv.X, vv.Y, vv.Z, vv.W}
}

// ToSlice2 returns components as slice.
func ToSlice2[V Vec2like[S], S Scalar](v V) []S {
	vv := Vec2g[S](v)
	return []S{vv.X, vv.Y}
}

// ToSlice3 returns components as slice.
func ToSlice3[V Vec3like[S], S Scalar](v V) []S {
	vv := Vec3g[S](v)
	return []S{vv.X, vv.Y, vv.Z}
}

// ToSlice4 returns components as slice.
func ToSlice4[V Vec4like[S], S Scalar](v V) []S {
	vv := Vec4g[S](v)
	return []S{vv.X, vv.Y, vv.Z, vv.W}
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

// ===================
// Math API (package functions)
// Multiple vector operations are defined as global functions.
// ===================

// Dot2 returns the dot product of two 2D vectors.
func Dot2[V1, V2 Vec2like[S], S Scalar](a V1, b V2) S {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return va.X*vb.X + va.Y*vb.Y
}

// Dot3 returns the dot product of two 3D vectors.
func Dot3[V1, V2 Vec3like[S], S Scalar](a V1, b V2) S {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return va.X*vb.X + va.Y*vb.Y + va.Z*vb.Z
}

// Dot4 returns the dot product of two 4D vectors.
func Dot4[V1, V2 Vec4like[S], S Scalar](a V1, b V2) S {
	va := Vec4g[S](a)
	vb := Vec4g[S](b)
	return va.X*vb.X + va.Y*vb.Y + va.Z*vb.Z + va.W*vb.W
}

// Lerp2 linearly interpolates between a and b by t.
func Lerp2[V1, V2 Vec2like[S], S Scalar](a V1, b V2, t float64) V1 {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return V1(Vec2g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
	})
}

// Lerp3 linearly interpolates between a and b by t.
func Lerp3[V1, V2 Vec3like[S], S Scalar](a V1, b V2, t float64) V1 {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return V1(Vec3g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
		Z: S(float64(va.Z) + (float64(vb.Z)-float64(va.Z))*t),
	})
}

// Lerp4 linearly interpolates between a and b by t.
func Lerp4[V1, V2 Vec4like[S], S Scalar](a V1, b V2, t float64) V1 {
	va := Vec4g[S](a)
	vb := Vec4g[S](b)
	return V1(Vec4g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
		Z: S(float64(va.Z) + (float64(vb.Z)-float64(va.Z))*t),
		W: S(float64(va.W) + (float64(vb.W)-float64(va.W))*t),
	})
}

// Project2 projects v onto onNormal.
func Project2[V1, V2 Vec2like[S], S Scalar](v V1, onNormal V2) V1 {
	va := Vec2g[S](v)
	vn := Vec2g[S](Normalize2(onNormal))
	return V1(vn.Scale(Dot2(va, vn)))
}

// Project3 projects v onto onNormal.
func Project3[V1, V2 Vec3like[S], S Scalar](v V1, onNormal V2) V1 {
	va := Vec3g[S](v)
	vn := Vec3g[S](Normalize3(onNormal))
	return V1(vn.Scale(Dot3(va, vn)))
}

// Project4 projects v onto onNormal.
func Project4[V1, V2 Vec4like[S], S Scalar](v V1, onNormal V2) V1 {
	va := Vec4g[S](v)
	vn := Vec4g[S](Normalize4(onNormal))
	return V1(vn.Scale(Dot4(va, vn)))
}

// Reflect2 reflects v off the surface with normal n.
func Reflect2[V1, V2 Vec2like[S], S Scalar](v V1, normal V2) V1 {
	va := Vec2g[S](v)
	vn := Vec2g[S](Normalize2(normal))
	dot := Dot2(va, vn)
	return V1(Vec2g[S]{
		X: va.X - S(2*float64(dot)*float64(vn.X)),
		Y: va.Y - S(2*float64(dot)*float64(vn.Y)),
	})
}

// Reflect3 reflects v off the surface with normal n.
func Reflect3[V1, V2 Vec3like[S], S Scalar](v V1, normal V2) V1 {
	va := Vec3g[S](v)
	vn := Vec3g[S](Normalize3(normal))
	dot := Dot3(va, vn)
	return V1(Vec3g[S]{
		X: va.X - S(2*float64(dot)*float64(vn.X)),
		Y: va.Y - S(2*float64(dot)*float64(vn.Y)),
		Z: va.Z - S(2*float64(dot)*float64(vn.Z)),
	})
}

// Reflect4 reflects v off the surface with normal n.
func Reflect4[V1, V2 Vec4like[S], S Scalar](v V1, normal V2) V1 {
	va := Vec4g[S](v)
	vn := Vec4g[S](Normalize4(normal))
	dot := Dot4(va, vn)
	return V1(Vec4g[S]{
		X: va.X - S(2*float64(dot)*float64(vn.X)),
		Y: va.Y - S(2*float64(dot)*float64(vn.Y)),
		Z: va.Z - S(2*float64(dot)*float64(vn.Z)),
		W: va.W - S(2*float64(dot)*float64(vn.W)),
	})
}

// Cross2 returns the 2D cross product (determinant) of two vectors.
func Cross2[V1, V2 Vec2like[S], S Scalar](a V1, b V2) S {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	// 2D cross product is a scalar (determinant)
	return va.X*vb.Y - va.Y*vb.X
}

// Cross3 returns the 3D cross product.
func Cross3[V1, V2 Vec3like[S], S Scalar](a V1, b V2) V1 {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return V1(Vec3g[S]{
		X: va.Y*vb.Z - va.Z*vb.Y,
		Y: va.Z*vb.X - va.X*vb.Z,
		Z: va.X*vb.Y - va.Y*vb.X,
	})
}

// Slerp3 spherically interpolates between a and b by t.
func Slerp3[V1, V2 Vec3like[S], S Scalar](a V1, b V2, t float64) V1 {
	va := Vec3g[S](Normalize3(a))
	vb := Vec3g[S](Normalize3(b))

	dot := float64(Dot3(va, vb))

	// Clamp dot product to avoid numerical errors
	if dot > 0.9995 {
		// Vectors are very close, use linear interpolation
		return Normalize3(Lerp3(a, b, t))
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
		return Normalize3(Lerp3(a, b, t))
	}

	a1 := math.Sin((1-t)*theta) / sinTheta
	b1 := math.Sin(t*theta) / sinTheta

	return V1(Vec3g[S]{
		X: S(float64(va.X)*a1 + float64(vb.X)*b1),
		Y: S(float64(va.Y)*a1 + float64(vb.Y)*b1),
		Z: S(float64(va.Z)*a1 + float64(vb.Z)*b1),
	})
}

// Rotate2 rotates v by angle radians.
func Rotate2[V Vec2like[S], S Scalar](v V, angle float64) V {
	va := Vec2g[S](v)
	sin, cos := math.Sincos(angle)
	return V(Vec2g[S]{
		X: S(float64(va.X)*cos - float64(va.Y)*sin),
		Y: S(float64(va.X)*sin + float64(va.Y)*cos),
	})
}

// Map2 applies f to each component of a 2D vector.
func Map2[V Vec2like[S], S Scalar](v V, f func(S) S) V {
	va := Vec2g[S](v)
	return V(Vec2g[S]{f(va.X), f(va.Y)})
}

// Map3 applies f to each component of a 3D vector.
func Map3[V Vec3like[S], S Scalar](v V, f func(S) S) V {
	va := Vec3g[S](v)
	return V(Vec3g[S]{f(va.X), f(va.Y), f(va.Z)})
}

// Map4 applies f to each component of a 4D vector.
func Map4[V Vec4like[S], S Scalar](v V, f func(S) S) V {
	va := Vec4g[S](v)
	return V(Vec4g[S]{f(va.X), f(va.Y), f(va.Z), f(va.W)})
}

// Zip2 applies f to corresponding components of two 2D vectors.
func Zip2[V1, V2 Vec2like[S], S Scalar](a V1, b V2, f func(S, S) S) V1 {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return V1(Vec2g[S]{f(va.X, vb.X), f(va.Y, vb.Y)})
}

// Zip3 applies f to corresponding components of two 3D vectors.
func Zip3[V1, V2 Vec3like[S], S Scalar](a V1, b V2, f func(S, S) S) V1 {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return V1(Vec3g[S]{f(va.X, vb.X), f(va.Y, vb.Y), f(va.Z, vb.Z)})
}

// Zip4 applies f to corresponding components of two 4D vectors.
func Zip4[V1, V2 Vec4like[S], S Scalar](a V1, b V2, f func(S, S) S) V1 {
	va := Vec4g[S](a)
	vb := Vec4g[S](b)
	return V1(Vec4g[S]{f(va.X, vb.X), f(va.Y, vb.Y), f(va.Z, vb.Z), f(va.W, vb.W)})
}

// Apply2 transforms all components of a 2D vector at once.
func Apply2[V Vec2like[S], S Scalar](v V, f func(S, S) (S, S)) V {
	va := Vec2g[S](v)
	x, y := f(va.X, va.Y)
	return V(Vec2g[S]{x, y})
}

// Apply3 transforms all components of a 3D vector at once.
func Apply3[V Vec3like[S], S Scalar](v V, f func(S, S, S) (S, S, S)) V {
	va := Vec3g[S](v)
	x, y, z := f(va.X, va.Y, va.Z)
	return V(Vec3g[S]{x, y, z})
}

// Apply4 transforms all components of a 4D vector at once.
func Apply4[V Vec4like[S], S Scalar](v V, f func(S, S, S, S) (S, S, S, S)) V {
	va := Vec4g[S](v)
	x, y, z, w := f(va.X, va.Y, va.Z, va.W)
	return V(Vec4g[S]{x, y, z, w})
}

// LenSq2 returns the squared length of a 2D vector.
func LenSq2[V Vec2like[S], S Scalar](v V) S {
	va := Vec2g[S](v)
	return va.X*va.X + va.Y*va.Y
}

// LenSq3 returns the squared length of a 3D vector.
func LenSq3[V Vec3like[S], S Scalar](v V) S {
	va := Vec3g[S](v)
	return va.X*va.X + va.Y*va.Y + va.Z*va.Z
}

// LenSq4 returns the squared length of a 4D vector.
func LenSq4[V Vec4like[S], S Scalar](v V) S {
	va := Vec4g[S](v)
	return va.X*va.X + va.Y*va.Y + va.Z*va.Z + va.W*va.W
}

// Len2 returns the length of a 2D vector.
func Len2[V Vec2like[S], S Scalar](v V) float64 {
	va := Vec2g[S](v)
	return math.Sqrt(float64(va.X*va.X + va.Y*va.Y))
}

// Len3 returns the length of a 3D vector.
func Len3[V Vec3like[S], S Scalar](v V) float64 {
	va := Vec3g[S](v)
	return math.Sqrt(float64(va.X*va.X + va.Y*va.Y + va.Z*va.Z))
}

// Len4 returns the length of a 4D vector.
func Len4[V Vec4like[S], S Scalar](v V) float64 {
	va := Vec4g[S](v)
	return math.Sqrt(float64(va.X*va.X + va.Y*va.Y + va.Z*va.Z + va.W*va.W))
}

// Angle2 returns the angle of a 2D vector in radians.
func Angle2[V Vec2like[S], S Scalar](v V) float64 {
	va := Vec2g[S](v)
	return math.Atan2(float64(va.Y), float64(va.X))
}

// Normalize2 returns the unit vector of a 2D vector.
// Returns zero vector if the input has zero length.
func Normalize2[V Vec2like[S], S Scalar](v V) V {
	va := Vec2g[S](v)
	l := Len2(va)
	if l == 0 {
		return V(Vec2g[S]{0, 0})
	}
	return V(Vec2g[S]{va.X / S(l), va.Y / S(l)})
}

// Normalize3 returns the unit vector of a 3D vector.
// Returns zero vector if the input has zero length.
func Normalize3[V Vec3like[S], S Scalar](v V) V {
	va := Vec3g[S](v)
	l := Len3(va)
	if l == 0 {
		return V(Vec3g[S]{0, 0, 0})
	}
	return V(Vec3g[S]{va.X / S(l), va.Y / S(l), va.Z / S(l)})
}

// Normalize4 returns the unit vector of a 4D vector.
// Returns zero vector if the input has zero length.
func Normalize4[V Vec4like[S], S Scalar](v V) V {
	va := Vec4g[S](v)
	l := Len4(va)
	if l == 0 {
		return V(Vec4g[S]{0, 0, 0, 0})
	}
	return V(Vec4g[S]{va.X / S(l), va.Y / S(l), va.Z / S(l), va.W / S(l)})
}

// ===================
// Math API (methods)
// Arithmetic operations, comparisons, dimension conversions,
// and component unpacking are defined as methods.
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

// Eqs returns whether all components equal s.
func (a Vec2g[S]) Eqs(s S) bool { return a.X == s && a.Y == s }

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

// Eqs returns whether all components equal s.
func (a Vec3g[S]) Eqs(s S) bool {
	return a.X == s && a.Y == s && a.Z == s
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

// Eqs returns whether all components equal s.
func (a Vec4g[S]) Eqs(s S) bool {
	return a.X == s && a.Y == s && a.Z == s && a.W == s
}

// Scale is an alias for Muls.
func (a Vec4g[S]) Scale(s S) Vec4g[S] {
	return Vec4g[S]{a.X * s, a.Y * s, a.Z * s, a.W * s}
}
