package main

import (
	"fmt"

	"github.com/jcasado94/nats-points/mongo/drivers"
	"github.com/jcasado94/nats-points/mongo/entity"
)

func main() {
	driver, _ := drivers.NewMongoDriver()
	a := entity.NewArticle("a", []string{"tag"}, 2)
	fmt.Println(driver.InsertArticle(&a))
}
