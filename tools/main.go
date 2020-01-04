package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v3"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

func main() {

	// Open our yamlFile
	yamlFile, err := os.Open("quizz.yaml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer yamlFile.Close()

	byteValue, _ := ioutil.ReadAll(yamlFile)

	var data Data
	err = yaml.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	//pp.Println(data.Kahoot)

	database := "Kahoot-Like-App"
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/" + database)
	if err != nil {
		log.Fatalf("cannot connect to mongodb host: %v", err)
	}
	/*
		c := session.DB(database).C(collection)
		query := ""
		err := c.Find(query).One(&result)
		if err != nil {
			log.Fatalf("cannot connect to mongodb host: %v", err)
		}
	*/

	for _, question := range data.Kahoot.Questions {
		pp.Println(question)
	}

	pp.Println(session)

	// Use session as normal
	session.Close()

}
