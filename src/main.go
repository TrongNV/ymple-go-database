package main

import (
	//"os"
	/*"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"*/

	//"net/http"
	"./package1"
	"./package_redis"
	"./package_gin"
	"./package_mongodb"
	"./package_elast"
	//   "./package_mysql"
)

func main() {

	package_mongodb.InsertMongo()

	package1.Test();

	package1.ReadJson2();

	package_redis.Read("db to read")

	package_elast.Read()

	package_gin.StartRouter()
	// package_mysql.Read()
}

