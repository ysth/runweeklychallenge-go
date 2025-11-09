Go library to facilitate running a solution to [the Weekly Challenge](https://theweeklychallenge.org) using one or more sets of inputs provided as JSON command line arguments

Example usage running a "solution" to sum integers:

    package main
    
    import run "github.com/ysth/runweeklychallenge-go"
    
    func SumOfInts(ints []int) (int) {
        sum := int(0)
        for _, num := range ints {
            sum += num
        }
        return sum
    }
    
    func main() {
        // runSolution runs the solution for a single set of inputs; it may
        // reformat the output if desired and can use helper functions from
        // this module to extract parts of the input to pass to the solution
        // and pass through returned errors if provided by the solution
        runSolution := func(inputs any) (any, error) {
            return SumOfInts(run.AsIntSlice(inputs, "ints")), nil
        }

        // inputs example for use in error message if incorrectly formatted
        // inputs are given
        inputsExample := `{"ints":[1,2,3]}`

        // jsonschema (draft2020-12) for a set of inputs; it should validate
        // types such that any runtime conversion from the unmarshalled json
        // will not fail
        inputsSchemaJSON := `{
            "type": "object",
            "properties": {
                "ints": {
                    "type": "array",
                    "items": { "type": "integer" }
                }
            },
            "required": ["ints"],
            "additionalProperties": false
        }`
    
        run.RunWeeklyChallenge(runSolution, inputsExample, inputsSchemaJSON)
    }

Example output:

    $ go run example.go '{"ints":[1,2,3]}' '{"ints":[]}'
    Inputs: {"ints":[1,2,3]}
    Output: 6
    Inputs: {"ints":[]}
    Output: 0

You must provide an example JSON inputs string (used in error messages), a JSON schema for inputs, and a shim function to run the solution given the decoded JSON inputs and reformat the output if desired.  Helper functions are provided to extract properties of the JSON input as various types to pass to the solution in the shim function.
