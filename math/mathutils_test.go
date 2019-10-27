package mathutils

import (
	"math"
	"testing"

	"github.com/alessiosavi/GoGPUtils/helper"
)

const total int = 1000

var randomizer helper.RandomGenerator = helper.InitRandomizer()

func TestConvertSize(t *testing.T) {
	t.Log(ConvertSize(1024, "KB"))
	t.Log(ConvertSize(1000000, "MB"))
	t.Log(ConvertSize(1024, "GB"))
}

func BenchmarkSumIntArray(t *testing.B) {
	array := randomizer.RandomIntArray(0, math.MaxInt8, total)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		SumIntArray(array)
	}
}

func BenchmarkSumInt32Array(t *testing.B) {
	array := randomizer.RandomInt32Array(0, math.MaxInt32, total)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		SumInt32Array(array)
	}
}

func BenchmarkSumInt64Array(t *testing.B) {
	array := randomizer.RandomInt64Array(0, math.MaxInt64, total)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		SumInt64Array(array)
	}
}

func BenchmarkSumFloat32Array(t *testing.B) {
	array := randomizer.RandomFloat32Array(0, math.MaxInt64, total)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		SumFloat32Array(array)
	}
}

func BenchmarkSumFloat64Array(t *testing.B) {
	array := randomizer.RandomFloat64Array(0, math.MaxInt64, total)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		SumFloat64Array(array)
	}
}

func BenchmarkMaxIntIndex(t *testing.B) {
	array := randomizer.RandomIntArray(0, math.MaxInt8, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		MaxIntIndex(array)
	}
}

func BenchmarkMaxInt32Index(t *testing.B) {
	array := randomizer.RandomInt32Array(0, math.MaxInt32, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		MaxInt32Index(array)
	}
}

func BenchmarkMaxInt64Index(t *testing.B) {
	array := randomizer.RandomInt64Array(0, math.MaxInt64, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		MaxInt64Index(array)
	}
}

func BenchmarkMaxFloat32Index(t *testing.B) {
	array := randomizer.RandomFloat32Array(0, math.MaxFloat32, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		MaxFloat32Index(array)
	}
}

func BenchmarkMaxFloat64Index(t *testing.B) {
	array := randomizer.RandomFloat64Array(0, math.MaxFloat64, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		MaxFloat64Index(array)
	}
}

func BenchmarkAverageInt(t *testing.B) {
	array := randomizer.RandomIntArray(0, math.MaxInt8, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		AverageInt(array)
	}
}

func BenchmarkAverageInt32(t *testing.B) {
	array := randomizer.RandomInt32Array(0, math.MaxInt32, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		AverageInt32(array)
	}
}

func BenchmarkAverageInt64(t *testing.B) {
	array := randomizer.RandomInt64Array(0, math.MaxInt64, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		AverageInt64(array)
	}
}

func BenchmarkAverageFloat32(t *testing.B) {
	array := randomizer.RandomFloat32Array(0, math.MaxFloat32, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		AverageFloat32(array)
	}
}
func BenchmarkAverageFloat64(t *testing.B) {
	array := randomizer.RandomFloat64Array(0, math.MaxFloat64, total)
	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		AverageFloat64(array)
	}
}
