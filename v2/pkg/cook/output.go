package cook

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/config"
	"github.com/glitchedgitz/cook/v2/pkg/parse"
)

func (cook *COOK) AppendMode(values []string) {
	tmp := []string{}
	till := len(cook.Final)
	if len(cook.Final) > len(values) {
		till = len(values)
	}
	for i := 0; i < till; i++ {
		tmp = append(tmp, cook.Final[i]+values[i])
	}
	cook.Final = tmp
}

func (cook *COOK) PermutationMode(values []string) {
	tmp := []string{}
	for _, t := range cook.Final {
		for _, v := range values {
			tmp = append(tmp, t+v)
		}
	}
	cook.Final = tmp
}

func (cook *COOK) CheckParam(p string, array *[]string) bool {
	if val, exists := cook.Params[p]; exists {
		if config.PipeInput(val, array) || config.RawInput(val, array) || RepeatOp(val, array) || cook.Config.ParseFunc(val, array) || cook.Config.ParseFile(p, val, array) || cook.CheckMethods(val, array) {
			return true
		}

		*array = append(*array, parse.SplitValues(val)...)
		return true
	}
	return false
}

func (cook *COOK) Print() {

	if !cook.PrintResult {
		return
	}

	var outputLines []string
	
	if len(cook.MethodsForAll) > 0 {
		tmp := []string{}

		for _, meth := range strings.Split(cook.MethodsForAll, ",") {
			cook.ApplyMethods(cook.Final, parse.SplitMethods(meth), &tmp)
		}
		outputLines = tmp
	} else {
		outputLines = cook.Final
	}
	
	// Write to file if OutputFile is specified
	if cook.OutputFile != "" {
		cook.WriteToFile(outputLines)
	}
	
	// Write to stdout if no file specified or if OutputBoth is true
	if cook.OutputFile == "" || cook.OutputBoth {
		for _, v := range outputLines {
			fmt.Println(v)
		}
	}
}

// WriteToFile writes the output lines to the specified file
func (cook *COOK) WriteToFile(lines []string) {
	f, err := os.OpenFile(cook.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}
	defer f.Close()
	
	for _, line := range lines {
		if _, err := f.WriteString(line + "\n"); err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
	}
}
