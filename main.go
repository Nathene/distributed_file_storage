package main

import (
	"log"

	"github.com/Nathene/distributed_file_storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.LisenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

}
