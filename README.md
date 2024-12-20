# nex
- PRUDP/NEX Server Library.

## Informations
- It's NEX Library written in Go for Retendo Network.
- Nintendo used the Library Quazal Rendez-Vous by Ubisoft and Quazal Companie and added some modifications for create "NEX". NEX is library for Online Games Servers for Nintendo 2DS/3DS, Nintendo Wii U and Nintendo Switch Consoles. (But today Nintendo Switch used NPLN).
- For create our NEX Library, we used the NEX documentation from NintendoClients, thanks Kinnay.

## Installation 
```
go get github.com/RetendoNetwork/nex
```

## Example Usage
```Go
package main

import (
	"github.com/RetendoNetwork/nex"
)

func main() {
	server := nex.NewServer()
	
	server.SetPRUDPVersion(1)
	server.SetKeySize(16)
	server.SetFragmentSize(962)
	server.SetAccessKey("ridfebb9")

	server.OnData("Data", func(packet *nex.PacketV0) {
		request := packet.RMCRequest()

		fmt.Printf("Protocol: %#v\n", request.GetProtocol())
		fmt.Printf("Method: %#v\n", request.GetMethod())
	})

	server.Listen(6000)
}
```