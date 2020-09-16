package main

import (
	"encoding/json"
	"fmt"
	"github.com/filbot3/golang_genesys_api/pkg/structures"
	"io/ioutil"
)

func main() {
	// Read in the JSON File blob.
	characterJSONBlob, err := ioutil.ReadFile("character_sheet_example.json")
	if err != nil {
		// Check for errors, and print if there is an error.
		fmt.Println("error: ", err)
	}

	// Create a blank CharacterSheet structure to put our JSON data into
	var philbo structures.CharacterSheet

	// Unmarshal, or take JSON Byte Strings, and convert it into Golang Structs
	// and other data types.
	err = json.Unmarshal(characterJSONBlob, &philbo)
	if err != nil {
		// Again, check for errors.
		fmt.Println("error: ", err)
	}

	// Print the specific parts of the philbo CharacterSheet Struct we want.
	fmt.Printf("%+v", philbo.Characteristics)
}
