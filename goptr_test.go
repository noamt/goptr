package goptr_test

import (
	"github.com/noamt/goptr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	b := true
	assert.False(t, goptr.Bool(nil))
	assert.True(t, goptr.Bool(&b))

	assert.True(t, goptr.BoolOrDefault(nil, true))
	assert.True(t, goptr.BoolOrDefault(&b, false))
}

func TestFloat32(t *testing.T) {
	f := float32(2.2)
	assert.Equal(t, float32(0), goptr.Float32(nil))
	assert.Equal(t, float32(2.2), goptr.Float32(&f))

	assert.Equal(t, float32(1.1), goptr.Float32OrDefault(nil, 1.1))
	assert.Equal(t, float32(2.2), goptr.Float32OrDefault(&f, 1.1))
}

func TestFloat64(t *testing.T) {
	f := 2.2
	assert.Equal(t, 0.0, goptr.Float64(nil))
	assert.Equal(t, 2.2, goptr.Float64(&f))

	assert.Equal(t, 1.1, goptr.Float64OrDefault(nil, 1.1))
	assert.Equal(t, 2.2, goptr.Float64OrDefault(&f, 1.1))
}

func TestInt(t *testing.T) {
	i := 2
	assert.Equal(t, 0, goptr.Int(nil))
	assert.Equal(t, 2, goptr.Int(&i))

	assert.Equal(t, 1, goptr.IntOrDefault(nil, 1))
	assert.Equal(t, 2, goptr.IntOrDefault(&i, 1))
}

func TestInt16(t *testing.T) {
	i := int16(2)
	assert.Equal(t, int16(0), goptr.Int16(nil))
	assert.Equal(t, int16(2), goptr.Int16(&i))

	assert.Equal(t, int16(1), goptr.Int16OrDefault(nil, 1))
	assert.Equal(t, int16(2), goptr.Int16OrDefault(&i, 1))
}

func TestInt32(t *testing.T) {
	i := int32(2)
	assert.Equal(t, int32(0), goptr.Int32(nil))
	assert.Equal(t, int32(2), goptr.Int32(&i))

	assert.Equal(t, int32(1), goptr.Int32OrDefault(nil, 1))
	assert.Equal(t, int32(2), goptr.Int32OrDefault(&i, 1))
}

func TestInt64(t *testing.T) {
	i := int64(2)
	assert.Equal(t, int64(0), goptr.Int64(nil))
	assert.Equal(t, int64(2), goptr.Int64(&i))

	assert.Equal(t, int64(1), goptr.Int64OrDefault(nil, 1))
	assert.Equal(t, int64(2), goptr.Int64OrDefault(&i, 1))
}

func TestInt8(t *testing.T) {
	i := int8(2)
	assert.Equal(t, int8(0), goptr.Int8(nil))
	assert.Equal(t, int8(2), goptr.Int8(&i))

	assert.Equal(t, int8(1), goptr.Int8OrDefault(nil, 1))
	assert.Equal(t, int8(2), goptr.Int8OrDefault(&i, 1))
}

func TestString(t *testing.T) {
	s := "bar"
	assert.Equal(t, "", goptr.String(nil))
	assert.Equal(t, "bar", goptr.String(&s))

	assert.Equal(t, "foo", goptr.StringOrDefault(nil, "foo"))
	assert.Equal(t, "bar", goptr.StringOrDefault(&s, "foo"))
}

func TestUint(t *testing.T) {
	i := uint(2)
	assert.Equal(t, uint(0), goptr.Uint(nil))
	assert.Equal(t, uint(2), goptr.Uint(&i))

	assert.Equal(t, uint(1), goptr.UintOrDefault(nil, 1))
	assert.Equal(t, uint(2), goptr.UintOrDefault(&i, 1))
}

func TestUint16(t *testing.T) {
	i := uint16(2)
	assert.Equal(t, uint16(0), goptr.Uint16(nil))
	assert.Equal(t, uint16(2), goptr.Uint16(&i))

	assert.Equal(t, uint16(1), goptr.Uint16OrDefault(nil, 1))
	assert.Equal(t, uint16(2), goptr.Uint16OrDefault(&i, 1))
}

func TestUint32(t *testing.T) {
	i := uint32(2)
	assert.Equal(t, uint32(0), goptr.Uint32(nil))
	assert.Equal(t, uint32(2), goptr.Uint32(&i))

	assert.Equal(t, uint32(1), goptr.Uint32OrDefault(nil, 1))
	assert.Equal(t, uint32(2), goptr.Uint32OrDefault(&i, 1))
}

func TestUint64(t *testing.T) {
	i := uint64(2)
	assert.Equal(t, uint64(0), goptr.Uint64(nil))
	assert.Equal(t, uint64(2), goptr.Uint64(&i))

	assert.Equal(t, uint64(1), goptr.Uint64OrDefault(nil, 1))
	assert.Equal(t, uint64(2), goptr.Uint64OrDefault(&i, 1))
}

func TestUint8(t *testing.T) {
	i := uint8(2)
	assert.Equal(t, uint8(0), goptr.Uint8(nil))
	assert.Equal(t, uint8(2), goptr.Uint8(&i))

	assert.Equal(t, uint8(1), goptr.Uint8OrDefault(nil, 1))
	assert.Equal(t, uint8(2), goptr.Uint8OrDefault(&i, 1))
}
