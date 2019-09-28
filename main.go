package main

import (
	"log"

	"github.com/jcasado94/nats-points/invalidator"
)

func main() {
	inv, _ := invalidator.NewInvalidator()
	err := inv.InvalidateAllArticles("Spain")
	log.Print(err)
}
