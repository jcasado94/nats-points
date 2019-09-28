package main

import (
	"fmt"

	"github.com/jcasado94/nats-points/scraping"
)

func main() {
	scr, _ := scraping.NewTheGuardianScraper()
	fmt.Println(scr.GetAllNews("Spain"))
}
