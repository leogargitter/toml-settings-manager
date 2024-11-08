package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/charmbracelet/huh"
)

// createFormFromStruct dynamically generates a form with groups representing each struct
func createFormFromStruct(s interface{}) *huh.Form {
	var groups []*huh.Group
	val := reflect.ValueOf(s)

	if val.Kind() != reflect.Struct {
		log.Fatalf("Error: input is not a struct. %v", val.Kind())
		return nil
	}

	// Iterate over each field in the main struct
	for i := 0; i < val.NumField(); i++ {
		// field := val.Type().Field(i)
		fieldValue := val.Field(i)

		// Create a slice of fields for the current group
		var fields []huh.Field

		// Process fields within each sub-struct
		for j := 0; j < fieldValue.NumField(); j++ {
			subField := fieldValue.Type().Field(j)
			subFieldValue := fieldValue.Field(j)

			tag := subField.Tag.Get("toml")
			if tag == "" {
				tag = subField.Name
			}

			// Create inputs based on type
			switch subFieldValue.Kind() {
			case reflect.String:
				fields = append(fields, huh.NewInput().Title(tag))
			case reflect.Int:
				fields = append(fields, huh.NewInput().Title(tag))
			case reflect.Bool:
				fields = append(fields, huh.NewSelect[bool]().Title(tag))
			}

		}
		// Add a new group with the fields and group name
		group := huh.NewGroup(fields...)
		groups = append(groups, group)
	}
	return huh.NewForm(groups...)
}

func main() {
	// Create a new instance of Config
	config, err := GetConfig("./example.toml")
	if err != nil {
		return
	}

	// Dynamically create a form based on the Config struct
	form := createFormFromStruct(config)

	// Run the form and get the result
	err = form.Run()
	if err != nil {
		fmt.Println("Error running form:", err)
		return
	}

	// Populate the config struct with values from the form
	// populateStructFromForm(&config, _)

	// Print the populated config struct
	fmt.Printf("Config: %+v\n", config)
}
