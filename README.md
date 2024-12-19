# nex
- PRUDP/NEX Server Library.

## Example Usage
```Go
package main

import (
	"github.com/RetendoNetwork/nex"
)

func main() {
	server := nex.NewServer()
	server.SetPrudpVersion(1)
	server.SetKeySize(16)
	server.SetFragmentSize(962)
	server.SetAccessKey("ridfebb9")
	server.Listen(":6000")
}
```