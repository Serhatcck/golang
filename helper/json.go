package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//exampleOne()
	//exampleTwo()
	exampleUnstructed()
}

func exampleOne() {
	type Bird struct {
		Species     string
		Description string
	}
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`

	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf(bird.Species)
}

func exampleTwo() {
	type Bird struct {
		Species     string
		Description string
	}
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds : %+v", birds)
}

func exampleUnstructed() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"empty animal"}`
	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)

	birds := result["birds"].(map[string]interface{})
	animals := result["animals"]

	for key, value := range birds {
		fmt.Println(key, value.(string))
	}
	log.Print("end birds")
	fmt.Println(animals)
}
