package main

import (
	"fmt"

	"github.com/adamelfsborg-code/learning-go/tutorial_22/util"
)

func main() {
	shapes := []util.Shape{
		util.Square{Length: 15.2},
		util.Circle{Radius: 7.5},
		util.Circle{Radius: 12.3},
		util.Square{Length: 4.9},
	}

	for _, v := range shapes {
		util.PrintShapeInfo(v)
		fmt.Println("---")
	}

	adapters := []util.Database{
		util.PostgreSQLAdapter{Host: "test", Port: 5432, User: "admin", Database: "newsletter", Password: "123"},
		util.MongoDBAdapter{ConnectionString: "testing.mongodb.com"},
	}

	for _, value := range adapters {
		util.ConnectToDb(value)
	}

}
