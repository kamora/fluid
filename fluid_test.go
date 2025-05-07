package fluid

import (
	"math/rand/v2"
	"testing"
)

const (
	MaxUint16 = ^uint16(0)
	MaxUint32 = ^uint32(0)
	MaxUint64 = ^uint64(0)
)

func Test(t *testing.T) {
	t.Run("int16", func(t *testing.T) {
		for i := 0; i < 10000; i++ {
			n := rand.N[uint16](MaxUint16)
			s := Encode[uint16](n)

			if len(s) != 4 {
				t.Errorf("failed for %d - expected %d but get %d", n, 4, len(s))
			}
		}
	})

	t.Run("int32", func(t *testing.T) {
		for i := 0; i < 100000; i++ {
			n := rand.N[uint32](MaxUint32)
			s := Encode[uint32](n)

			if len(s) != 7 {
				t.Errorf("failed for %d - expected %d but get %d", n, 7, len(s))
			}
		}
	})

	t.Run("int64", func(t *testing.T) {
		for i := 0; i < 1000000; i++ {
			n := rand.N[uint64](MaxUint64)
			s := Encode[uint64](n)

			if len(s) != 13 {
				t.Errorf("failed for %d - expected %d but get %d", n, 12, len(s))
			}
		}
	})
}
