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
        runSolution := func(inputs any) (any, error) {
            return SumOfInts(run.AsIntSlice(inputs, "ints")), nil
        }
    
        inputExample := `{"ints":[1,2,3]}`
    
        inputSchema := `{
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
    
        run.RunWeeklyChallenge(runSolution, inputExample, inputSchema)
    }

You must provide an example JSON inputs string (used in error messages), a JSON schema for inputs, and a shim function to call the solution given the decoded JSON.  Helper functions are provided to extract properties of the JSON input as various types to pass to the solution in the shim function.
