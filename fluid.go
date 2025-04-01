package fluid

import (
	"errors"
	"fmt"
	"github.com/kamora/mathanol"
	"math/rand/v2"
	"reflect"
)

func convert(value uint16) string {
	if value <= 25 {
		return fmt.Sprintf("%c", 'a'+value)
	}

	if value <= 51 {
		return fmt.Sprintf("%c", 'A'+(value-26))
	}

	if value <= 61 {
		return fmt.Sprintf("%d", value-52)
	}

	panic(errors.New("value out of range"))
}

func Encode[T uint16 | uint32 | uint64](value T) string {
	output := ""
	value = (^T(0)) - value

	limit := 3

	if reflect.TypeFor[T]().Kind() == reflect.Uint32 {
		limit = 6
	}

	if reflect.TypeFor[T]().Kind() == reflect.Uint64 {
		limit = 11
	}

	if value > 61 {
		for v := range mathanol.Rem[T](value, 62) {
			output += convert(uint16(v))
		}
	} else {
		output += convert(uint16(value))
	}

	output = fmt.Sprintf("%d%s", len(output), output)

	for len(output) < limit {
		output += convert(rand.N[uint16](62))
	}

	return output
}
