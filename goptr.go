// Package goptr provides utility functions for pointers to primitive types
package goptr

// Bool resolves the given pointer b. If pointer b is nil, false returned
func Bool(b *bool) bool {
	return BoolOrDefault(b, false)
}

// BoolOrDefault resolves the given pointer b. If pointer b is nil, the default value of d is returned
func BoolOrDefault(b *bool, d bool) bool {
	if b != nil {
		return *b
	}
	return d
}

// String resolves the given pointer s. If pointer s is nil, an empty string is returned
func String(s *string) string {
	return StringOrDefault(s, "")
}

// StringOrDefault resolves the given pointer s. If pointer s is nil, the default value of d is returned
func StringOrDefault(s *string, d string) string {
	if s != nil {
		return *s
	}
	return d
}

// Float32 resolves the given pointer f. If pointer f is nil, 0 is returned
func Float32(f *float32) float32 {
	return Float32OrDefault(f, 0)
}

// Float32OrDefault resolves the given pointer f. If pointer f is nil, the default value of d is returned
func Float32OrDefault(f *float32, d float32) float32 {
	if f != nil {
		return *f
	}
	return d
}

// Float64 resolves the given pointer f. If pointer f is nil, 0 is returned
func Float64(f *float64) float64 {
	return Float64OrDefault(f, 0)
}

// Float64OrDefault resolves the given pointer f. If pointer f is nil, the default value of d is returned
func Float64OrDefault(f *float64, d float64) float64 {
	if f != nil {
		return *f
	}
	return d
}

// Int resolves the given pointer i. If pointer i is nil, 0 is returned
func Int(i *int) int {
	return IntOrDefault(i, 0)
}

// IntOrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func IntOrDefault(i *int, d int) int {
	if i != nil {
		return *i
	}
	return d
}

// Int8 resolves the given pointer i. If pointer i is nil, 0 is returned
func Int8(i *int8) int8 {
	return Int8OrDefault(i, 0)
}

// Int8OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Int8OrDefault(i *int8, d int8) int8 {
	if i != nil {
		return *i
	}
	return d
}

// Int16 resolves the given pointer i. If pointer i is nil, 0 is returned
func Int16(i *int16) int16 {
	return Int16OrDefault(i, 0)
}

// Int16OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Int16OrDefault(i *int16, d int16) int16 {
	if i != nil {
		return *i
	}
	return d
}

// Int32 resolves the given pointer i. If pointer i is nil, 0 is returned
func Int32(i *int32) int32 {
	return Int32OrDefault(i, 0)
}

// Int32OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Int32OrDefault(i *int32, d int32) int32 {
	if i != nil {
		return *i
	}
	return d
}

// Int64 resolves the given pointer i. If pointer i is nil, 0 is returned
func Int64(i *int64) int64 {
	return Int64OrDefault(i, 0)
}

// Int64OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Int64OrDefault(i *int64, d int64) int64 {
	if i != nil {
		return *i
	}
	return d
}

// Uint resolves the given pointer i. If pointer i is nil, 0 is returned
func Uint(i *uint) uint {
	return UintOrDefault(i, 0)
}

// UintOrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func UintOrDefault(i *uint, d uint) uint {
	if i != nil {
		return *i
	}
	return d
}

// Uint8 resolves the given pointer i. If pointer i is nil, 0 is returned
func Uint8(i *uint8) uint8 {
	return Uint8OrDefault(i, 0)
}

// Uint8OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Uint8OrDefault(i *uint8, d uint8) uint8 {
	if i != nil {
		return *i
	}
	return d
}

// Uint16 resolves the given pointer i. If pointer i is nil, 0 is returned
func Uint16(i *uint16) uint16 {
	return Uint16OrDefault(i, 0)
}

// Uint16OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Uint16OrDefault(i *uint16, d uint16) uint16 {
	if i != nil {
		return *i
	}
	return d
}

// Uint32 resolves the given pointer i. If pointer i is nil, 0 is returned
func Uint32(i *uint32) uint32 {
	return Uint32OrDefault(i, 0)
}

// Uint32OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Uint32OrDefault(i *uint32, d uint32) uint32 {
	if i != nil {
		return *i
	}
	return d
}

// Uint64 resolves the given pointer i. If pointer i is nil, 0 is returned
func Uint64(i *uint64) uint64 {
	return Uint64OrDefault(i, 0)
}

// Uint64OrDefault resolves the given pointer i. If pointer i is nil, the default value of d is returned
func Uint64OrDefault(i *uint64, d uint64) uint64 {
	if i != nil {
		return *i
	}
	return d
}
