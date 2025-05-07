# Fluid

  A UID generation library. 

## Install

```bash
go get github.com/kamora/fluid
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/kamora/fluid"
	"math/rand/v2"
)

func main() {

	v1 := fluid.Encode[uint16](rand.N[uint16](^uint16(0)))
	fmt.Println(v1)

	v2 := fluid.Encode[uint32](rand.N[uint32](^uint32(0)))
	fmt.Println(v2)

	v3 := fluid.Encode[uint64](rand.N[uint64](^uint64(0)))
	fmt.Println(v3)

	// Output
	//3EUh
	//5Db9XN3
	//11r0f7uh88vVu
}
```

## License

Released under the [MIT License](./LICENSE).
