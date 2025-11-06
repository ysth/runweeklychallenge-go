package runweeklychallenge

import (
    "encoding/json"
    "fmt"
    "os"
    "github.com/kaptinlin/jsonschema"
)

func RunWeeklyChallenge(solution func(inputs any) (any, error), exampleInput string, schemaJson string) () {

    // compile the schema for json inputs args
    schema, err := jsonschema.NewCompiler().Compile([]byte(schemaJson))
    if err != nil {
        fmt.Fprintln(os.Stderr, "Invalid jsonschema")
        os.Exit(1)
    }

    // run the solution for each command line arg
    errors := false
    for _, inputString := range os.Args[1:] {

        // show the inputs
        fmt.Printf("Input: %s\n", inputString)

        // validate the json input contains what we expect
        validationResult := schema.Validate([]byte(inputString))
        if ! validationResult.IsValid() {
            fmt.Printf("Error: invalid input: %v\n", validationResult.Error())
            errors = true
            continue
        }

        // decode the json
        var inputs any
        err := json.Unmarshal([]byte(inputString), &inputs)
        if err != nil {
            fmt.Printf("Error: invalid json input: %v\n", err)
            errors = true
            continue
        }

        // run it and show the results
        result, err := solution(inputs)
        if err != nil {
            fmt.Printf("Error: %v", err)
            continue
        } else {
            fmt.Printf("Output: %v\n", result)
        }
	}

    // if an input was not as expected, show an example
    if errors {
        fmt.Printf("Expected arguments like '%s'\n", exampleInput)
    }
    return;
}

// some helper functions

func AsInt(inputs any, key string) int {
    inputMap := inputs.(map[string]any)
    return int(inputMap[key].(float64))
}

func AsIntSlice(inputs any, key string) []int {
    inputMap := inputs.(map[string]any)
    inputSlice := inputMap[key].([]any)
    var result = make([]int, len(inputSlice))
    for i, v := range inputSlice {
        result[i] = int(v.(float64))
    }
    return result
}
