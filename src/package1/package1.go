package package1

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type jsonobject struct {
	Object ObjectType
}

type ObjectType struct {
	Buffer_size int
	Databases   []DatabasesType
}

type DatabasesType struct {
	Host   string
	User   string
	Pass   string
	Type   string
	Name   string
	Tables []TablesType
}

type TablesType struct {
	Name     string
	Statment string
	Regex    string
	Types    []TypesType
}

type TypesType struct {
	Id    string
	Value string
}

type Configuration struct {
	Users  []string
	Groups []string
}

func Test() {

	fmt.Println("FTH7")
}

func ReadJson() {

	file, e := ioutil.ReadFile("../config/config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(file))

	var jsontype jsonobject
	json.Unmarshal(file, &jsontype)

	fmt.Printf("START print object");

	fmt.Printf("Results: %v\n", jsontype)

	fmt.Printf("END: print object");

	fmt.Printf("END: print object element");

	fmt.Println(jsontype)

}

func ReadJson2() {

	file, _ := os.Open("../config/config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Start conf user");

	fmt.Println(configuration.Users)

}
