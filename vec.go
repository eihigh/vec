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
// Helper API
// ====================

func Splat2[S Scalar](x S) Vec2g[S] { return Vec2g[S]{x, x} }
func Splat3[S Scalar](x S) Vec3g[S] { return Vec3g[S]{x, x, x} }
func Splat4[S Scalar](x S) Vec4g[S] { return Vec4g[S]{x, x, x, x} }

func (a Vec2g[S]) XY() (x, y S)         { return a.X, a.Y }
func (a Vec3g[S]) XYZ() (x, y, z S)     { return a.X, a.Y, a.Z }
func (a Vec4g[S]) XYZW() (x, y, z, w S) { return a.X, a.Y, a.Z, a.W }

func (a Vec2g[S]) Vec3(z S) Vec3g[S]    { return Vec3g[S]{a.X, a.Y, z} }
func (a Vec2g[S]) Vec4(z, w S) Vec4g[S] { return Vec4g[S]{a.X, a.Y, z, w} }

func (a Vec3g[S]) Vec2() Vec2g[S]    { return Vec2g[S]{a.X, a.Y} }
func (a Vec3g[S]) Vec4(w S) Vec4g[S] { return Vec4g[S]{a.X, a.Y, a.Z, w} }

func (a Vec4g[S]) Vec2() Vec2g[S] { return Vec2g[S]{a.X, a.Y} }
func (a Vec4g[S]) Vec3() Vec3g[S] { return Vec3g[S]{a.X, a.Y, a.Z} }

func (a Vec2g[S]) Int() Vec2g[int] { return Vec2g[int]{int(a.X), int(a.Y)} }
func (a Vec3g[S]) Int() Vec3g[int] { return Vec3g[int]{int(a.X), int(a.Y), int(a.Z)} }
func (a Vec4g[S]) Int() Vec4g[int] { return Vec4g[int]{int(a.X), int(a.Y), int(a.Z), int(a.W)} }

func (a Vec2g[S]) Uint() Vec2g[uint] { return Vec2g[uint]{uint(a.X), uint(a.Y)} }
func (a Vec3g[S]) Uint() Vec3g[uint] { return Vec3g[uint]{uint(a.X), uint(a.Y), uint(a.Z)} }
func (a Vec4g[S]) Uint() Vec4g[uint] { return Vec4g[uint]{uint(a.X), uint(a.Y), uint(a.Z), uint(a.W)} }

func (a Vec2g[S]) Float64() Vec2g[float64] {
	return Vec2g[float64]{float64(a.X), float64(a.Y)}
}
func (a Vec3g[S]) Float64() Vec3g[float64] {
	return Vec3g[float64]{float64(a.X), float64(a.Y), float64(a.Z)}
}
func (a Vec4g[S]) Float64() Vec4g[float64] {
	return Vec4g[float64]{float64(a.X), float64(a.Y), float64(a.Z), float64(a.W)}
}

func (a Vec2g[S]) Float32() Vec2g[float32] {
	return Vec2g[float32]{float32(a.X), float32(a.Y)}
}
func (a Vec3g[S]) Float32() Vec3g[float32] {
	return Vec3g[float32]{float32(a.X), float32(a.Y), float32(a.Z)}
}
func (a Vec4g[S]) Float32() Vec4g[float32] {
	return Vec4g[float32]{float32(a.X), float32(a.Y), float32(a.Z), float32(a.W)}
}

func (a Vec2g[S]) Array() [2]S { return [2]S{a.X, a.Y} }
func (a Vec3g[S]) Array() [3]S { return [3]S{a.X, a.Y, a.Z} }
func (a Vec4g[S]) Array() [4]S { return [4]S{a.X, a.Y, a.Z, a.W} }

func (a Vec2g[S]) Slice() []S { return []S{a.X, a.Y} }
func (a Vec3g[S]) Slice() []S { return []S{a.X, a.Y, a.Z} }
func (a Vec4g[S]) Slice() []S { return []S{a.X, a.Y, a.Z, a.W} }

func (a Vec2g[S]) Map(f func(S) S) Vec2g[S] { return Vec2g[S]{f(a.X), f(a.Y)} }
func (a Vec3g[S]) Map(f func(S) S) Vec3g[S] { return Vec3g[S]{f(a.X), f(a.Y), f(a.Z)} }
func (a Vec4g[S]) Map(f func(S) S) Vec4g[S] { return Vec4g[S]{f(a.X), f(a.Y), f(a.Z), f(a.W)} }

func (a Vec2g[S]) Map2(b Vec2g[S], f func(S, S) S) Vec2g[S] {
	return Vec2g[S]{f(a.X, b.X), f(a.Y, b.Y)}
}
func (a Vec3g[S]) Map2(b Vec3g[S], f func(S, S) S) Vec3g[S] {
	return Vec3g[S]{f(a.X, b.X), f(a.Y, b.Y), f(a.Z, b.Z)}
}
func (a Vec4g[S]) Map2(b Vec4g[S], f func(S, S) S) Vec4g[S] {
	return Vec4g[S]{f(a.X, b.X), f(a.Y, b.Y), f(a.Z, b.Z), f(a.W, b.W)}
}

func (a Vec2g[S]) Apply(f func(S, S) (S, S)) Vec2g[S] {
	x, y := f(a.X, a.Y)
	return Vec2g[S]{x, y}
}

func (a Vec3g[S]) Apply(f func(S, S, S) (S, S, S)) Vec3g[S] {
	x, y, z := f(a.X, a.Y, a.Z)
	return Vec3g[S]{x, y, z}
}

func (a Vec4g[S]) Apply(f func(S, S, S, S) (S, S, S, S)) Vec4g[S] {
	x, y, z, w := f(a.X, a.Y, a.Z, a.W)
	return Vec4g[S]{x, y, z, w}
}

// ===================
// Math API (functions)
// ===================

func Dot2[S Scalar, V Vec2like[S]](a, b V) S {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return va.X*vb.X + va.Y*vb.Y
}

func Dot3[S Scalar, V Vec3like[S]](a, b V) S {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return va.X*vb.X + va.Y*vb.Y + va.Z*vb.Z
}

func Dot4[S Scalar, V Vec4like[S]](a, b V) S {
	va := Vec4g[S](a)
	vb := Vec4g[S](b)
	return va.X*vb.X + va.Y*vb.Y + va.Z*vb.Z + va.W*vb.W
}

func Lerp2[S Scalar, V Vec2like[S]](a, b V, t float64) Vec2g[S] {
	va := Vec2g[S](a)
	vb := Vec2g[S](b)
	return Vec2g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
	}
}

func Lerp3[S Scalar, V Vec3like[S]](a, b V, t float64) Vec3g[S] {
	va := Vec3g[S](a)
	vb := Vec3g[S](b)
	return Vec3g[S]{
		X: S(float64(va.X) + (float64(vb.X)-float64(va.X))*t),
		Y: S(float64(va.Y) + (float64(vb.Y)-float64(va.Y))*t),
		Z: S(float64(va.Z) + (float64(vb.Z)-float64(va.Z))*t),
	}
}

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

func Project2[S Scalar, V Vec2like[S]](v, onNormal V) Vec2g[S] {
	va := Vec2g[S](v)
	vn := Vec2g[S](onNormal).Normalize()
	return vn.Scale(Dot2(va, vn))
}

func Project3[S Scalar, V Vec3like[S]](v, onNormal V) Vec3g[S] {
	va := Vec3g[S](v)
	vn := Vec3g[S](onNormal).Normalize()
	return vn.Scale(Dot3(va, vn))
}

func Project4[S Scalar, V Vec4like[S]](v, onNormal V) Vec4g[S] {
	va := Vec4g[S](v)
	vn := Vec4g[S](onNormal).Normalize()
	return vn.Scale(Dot4(va, vn))
}

// ===================
// Math API (methods)
// ===================

// Vec2
// ---

func (a Vec2g[S]) Add(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X + b.X, a.Y + b.Y} }
func (a Vec2g[S]) Adds(s S) Vec2g[S]       { return Vec2g[S]{a.X + s, a.Y + s} }
func (a Vec2g[S]) Sub(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X - b.X, a.Y - b.Y} }
func (a Vec2g[S]) Subs(s S) Vec2g[S]       { return Vec2g[S]{a.X - s, a.Y - s} }
func (a Vec2g[S]) Mul(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X * b.X, a.Y * b.Y} }
func (a Vec2g[S]) Muls(s S) Vec2g[S]       { return Vec2g[S]{a.X * s, a.Y * s} }
func (a Vec2g[S]) Div(b Vec2g[S]) Vec2g[S] { return Vec2g[S]{a.X / b.X, a.Y / b.Y} }
func (a Vec2g[S]) Divs(s S) Vec2g[S]       { return Vec2g[S]{a.X / s, a.Y / s} }
func (a Vec2g[S]) Neg() Vec2g[S]           { return Vec2g[S]{-a.X, -a.Y} }
func (a Vec2g[S]) Eq(b Vec2g[S]) bool      { return a.X == b.X && a.Y == b.Y }
func (a Vec2g[S]) LenSq() S                { return a.X*a.X + a.Y*a.Y }
func (a Vec2g[S]) Len() float64            { return math.Sqrt(float64(a.X*a.X + a.Y*a.Y)) }
func (a Vec2g[S]) Angle() float64          { return math.Atan2(float64(a.Y), float64(a.X)) }

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

func (a Vec3g[S]) Add(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}
func (a Vec3g[S]) Adds(s S) Vec3g[S] {
	return Vec3g[S]{a.X + s, a.Y + s, a.Z + s}
}
func (a Vec3g[S]) Sub(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}
func (a Vec3g[S]) Subs(s S) Vec3g[S] {
	return Vec3g[S]{a.X - s, a.Y - s, a.Z - s}
}
func (a Vec3g[S]) Mul(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}
func (a Vec3g[S]) Muls(s S) Vec3g[S] {
	return Vec3g[S]{a.X * s, a.Y * s, a.Z * s}
}
func (a Vec3g[S]) Div(b Vec3g[S]) Vec3g[S] {
	return Vec3g[S]{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}
func (a Vec3g[S]) Divs(s S) Vec3g[S] {
	return Vec3g[S]{a.X / s, a.Y / s, a.Z / s}
}
func (a Vec3g[S]) Neg() Vec3g[S] {
	return Vec3g[S]{-a.X, -a.Y, -a.Z}
}
func (a Vec3g[S]) Eq(b Vec3g[S]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}
func (a Vec3g[S]) LenSq() S {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}
func (a Vec3g[S]) Len() float64 {
	return math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z))
}
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
func (a Vec4g[S]) Add(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}
func (a Vec4g[S]) Adds(s S) Vec4g[S] {
	return Vec4g[S]{a.X + s, a.Y + s, a.Z + s, a.W + s}
}
func (a Vec4g[S]) Sub(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}
func (a Vec4g[S]) Subs(s S) Vec4g[S] {
	return Vec4g[S]{a.X - s, a.Y - s, a.Z - s, a.W - s}
}
func (a Vec4g[S]) Mul(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W}
}
func (a Vec4g[S]) Muls(s S) Vec4g[S] {
	return Vec4g[S]{a.X * s, a.Y * s, a.Z * s, a.W * s}
}
func (a Vec4g[S]) Div(b Vec4g[S]) Vec4g[S] {
	return Vec4g[S]{a.X / b.X, a.Y / b.Y, a.Z / b.Z, a.W / b.W}
}
func (a Vec4g[S]) Divs(s S) Vec4g[S] {
	return Vec4g[S]{a.X / s, a.Y / s, a.Z / s, a.W / s}
}
func (a Vec4g[S]) Neg() Vec4g[S] {
	return Vec4g[S]{-a.X, -a.Y, -a.Z, -a.W}
}
func (a Vec4g[S]) Eq(b Vec4g[S]) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W
}
func (a Vec4g[S]) LenSq() S {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W
}
func (a Vec4g[S]) Len() float64 {
	return math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W))
}
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
