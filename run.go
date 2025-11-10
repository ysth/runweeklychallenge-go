package runweeklychallenge

import (
    "encoding/json"
    "fmt"
    "os"
    "github.com/kaptinlin/jsonschema"
)

func RunWeeklyChallenge(runSolution func(inputs any) (any, error), inputsExample string, inputsSchemaJSON string) () {

    // compile the schema for json inputs args
    inputsSchema, err := jsonschema.NewCompiler().Compile([]byte(inputsSchemaJSON))
    if err != nil {
        fmt.Fprintln(os.Stderr, "Invalid jsonschema")
        os.Exit(1)
    }

    // run the solution for each command line arg
    inputs_error := false
    for _, inputsJSON := range os.Args[1:] {

        // show the inputs
        fmt.Printf("Inputs: %s\n", inputsJSON)

        // validate the json input contains what we expect
        validationResult := inputsSchema.Validate([]byte(inputsJSON))
        if ! validationResult.IsValid() {
            fmt.Printf("Error: invalid inputs: %v\n", validationResult.Error())
            inputs_error = true
            continue
        }

        // decode the json
        var inputs any
        err := json.Unmarshal([]byte(inputsJSON), &inputs)
        if err != nil {
            fmt.Printf("Error: invalid inputs JSON: %v\n", err)
            inputs_error = true
            continue
        }

        // run it and show the results
        result, err := runSolution(inputs)
        if err != nil {
            fmt.Printf("Exception: %v\n", err)
            continue
        } else {
            fmt.Printf("Output: %v\n", result)
        }
    }

    // if an input was not as expected, show an example
    if inputs_error {
        fmt.Printf("Expected inputs arguments like '%s'\n", inputsExample)
    }
    return;
}

// some helper functions

func AsInt(inputs any, key string) int {
    inputsMap := inputs.(map[string]any)
    return int(inputsMap[key].(float64))
}

func AsIntSlice(inputs any, key string) []int {
    inputsMap := inputs.(map[string]any)
    inputSlice := inputsMap[key].([]any)
    var result = make([]int, len(inputSlice))
    for i, v := range inputSlice {
        result[i] = int(v.(float64))
    }
    return result
}
