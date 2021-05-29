package argenv

import (
	"errors"
	"flag"
	"fmt"
	_ "github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

// ArgEnv represents the object used to process Environment and command line parameters.
type ArgEnv struct {
	base    interface{}            // the struct that we are processing (the interface that was passed to Load())
	entries []Entry                // Entry objects that we scanned in `base`, i.e. one Entry per field in base
	values  map[string]interface{} // values for the environment variables/command line parameters
}

// Entry represents a single field in a struct.
type Entry struct {
	Name        string        // Name is the name of a field in a struct
	EnvName     string        // EnvName is the Environment variable name
	FlagName    string        // FlagName is the name of the command line parameter
	Type        string        // Type is the reflect.Type of the field
	Value       reflect.Value // Value is the reflect.Value of the field
	Description string        // Description is extracted from the 'description' struct tag for the field
	Default     string        // Default is extracted from the 'default' struct tag for the field
}

// generateEnvName returns a formated Environment variable string
//
// The variable name will be formatted as follows:
//   - The envName will first process all uppercase letters (except the first instance)
//     and insert a preceding underscore '_', the the entire string will be converted to
//     uppercase. e.x. DebugLogging becomes DEBUG_LOGGING
func (e *ArgEnv) generateEnvName(name string) (envName string, err error) {
	for i, c := range name {
		if i != 0 && unicode.IsUpper(c) {
			envName += fmt.Sprintf("_%s", strings.ToUpper(string(c)))
		} else {
			envName += strings.ToUpper(string(c))
		}
	}
	return
}

// generateFlagName returns command line flag string.
//
// The variable name will be formatted as follows:
//   - The flagName will first process all uppercase letters (except the first instance)
//     and insert a preceding dash '-a', then the entire string will be converted to
//     lowercase e.x.  DebugLogging becomes debug-logging
func (e *ArgEnv) generateFlagName(name string) (flagName string, err error) {
	for i, c := range name {
		if i != 0 && unicode.IsUpper(c) {
			flagName += fmt.Sprintf("-%s", strings.ToLower(string(c)))
		} else {
			flagName += strings.ToLower(string(c))
		}
	}
	return
}

// Usage will print the help/usage display for Flags and environment variables.
//
// This is necessary because not only do we want to inform users about
// flag options, but we also want to list the available environment
// variables that can be used.
func (e *ArgEnv) usage() {
	fmt.Fprintf(os.Stdout, "ArgEnv Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "Available Environment Variables:\n")
	for _, e := range e.entries {
		fmt.Fprintf(os.Stdout, "\t%s\n", e.EnvName)
	}
}

// scanStruct will do a scan of the base struct to gather information on the struct
func (e *ArgEnv) scanStruct() (err error) {
	var (
		numberOfFields int
	)

	// load up the base object (struct)
	val := reflect.ValueOf(e.base).Elem()

	numberOfFields = val.NumField()

	// bail out if there are no fields to scan
	if numberOfFields < 1 {
		err = errors.New("Struct has no fields to scan")
		return
	}

	e.entries = make([]Entry, numberOfFields)

	// this pass will just grab, names, types, and struct tags and store them in ArgEnv.entries
	for i := 0; i < numberOfFields; i++ {
		e.entries[i] = Entry{
			Name:        val.Type().Field(i).Name,
			Type:        val.Type().Field(i).Type.String(),
			Value:       val.Field(i),
			Description: reflect.TypeOf(e.base).Elem().Field(i).Tag.Get("description"),
			Default:     reflect.TypeOf(e.base).Elem().Field(i).Tag.Get("default"),
		}
	}

	return
}

// Load will scan the structure provided and populate the structure using parameters passed via Environment
// variables or command line parameters.  If no values are found in either location, then default values specified
// in the 'default' struct tag will be used to populate the structure.
//
//  It is important to note that Environment variables will take precedence over command line parameters.
func (e *ArgEnv) Load(o interface{}) {
	var (
		err error
	)

	// store the object as our base
	e.base = o

	// scan the base struct
	err = e.scanStruct()

	if err != nil {
		log.Fatalln("Scanning the struct failed!")
	}

	// Process the entries from the scan
	err = e.processEntries()

	if err != nil {
		log.Fatalln("Processing the struct failed!")
	}

}

// setupFlags will setup flag command line parsing from the entries collected during the scan
func (e *ArgEnv) setupFlags() (err error) {
	e.values = make(map[string]interface{})
	// setup arguments for flags
	for i := 0; i < len(e.entries); i++ {
		switch e.entries[i].Type {
		case "string":
			var value string
			flag.StringVar(&value, e.entries[i].FlagName, e.entries[i].Default, e.entries[i].Description)
			e.values[e.entries[i].Name] = &value
		case "int":
			var value int
			var defaultValue int64

			defaultValue, _ = strconv.ParseInt(e.entries[i].Default, 10, 64)

			flag.IntVar(&value, e.entries[i].FlagName, int(defaultValue), e.entries[i].Description)
			e.values[e.entries[i].Name] = &value
		case "bool":
			var value bool
			var defaultValue bool

			defaultValue, _ = strconv.ParseBool(e.entries[i].Default)

			flag.BoolVar(&value, e.entries[i].FlagName, defaultValue, e.entries[i].Description)
			e.values[e.entries[i].Name] = &value
		default:
			log.Printf("Unknown type %s\n", e.entries[i].Type)
		}
	}

	// Overwrite default usage with ours and parse flags
	flag.Usage = e.usage

	// parse all of the flags we've setup
	flag.Parse()

	return
}

// generateVariableNames will populate entries with the environment variable names and command line flag names
func (e *ArgEnv) generateVariableNames() (err error) {
	for i := 0; i < len(e.entries); i++ {
		e.entries[i].EnvName, err = e.generateEnvName(e.entries[i].Name)

		if err != nil {
			log.Fatalf("Failed generating Environment variable name\n")
		}

		e.entries[i].FlagName, err = e.generateFlagName(e.entries[i].Name)

		if err != nil {
			log.Fatalf("Failed generating command line flag name\n")
		}
	}
	return
}

// processEntries will scan the environment and command line parameters and populate the struct provided in Load() with the correct values
func (e *ArgEnv) processEntries() (err error) {
	err = e.generateVariableNames()

	if err != nil {
		log.Fatalf("Failed to generate variable names!")
	}

	err = e.setupFlags()

	if err != nil {
		log.Fatalf("Failed to setup command line flags!")
	}

	for i := 0; i < len(e.entries); i++ {
		switch e.entries[i].Type {
		case "string":
			var value string
			var ptrValue *string

			ptrValue, ok := e.values[e.entries[i].Name].(*string)

			if ok {
				value = *ptrValue
			}

			envValue, ok := os.LookupEnv(e.entries[i].EnvName)

			if ok {
				value = envValue
			}
			e.entries[i].Value.SetString(value)
			break
		case "int":
			var value int
			var intVal int64
			var ok bool
			var err error
			var ptrValue *int

			ptrValue, ok = e.values[e.entries[i].Name].(*int)

			if ok {
				value = *ptrValue
			} else {
				intVal, err = strconv.ParseInt(e.entries[i].Default, 10, 64)

				if err == nil {
					value = int(intVal)
				}
			}

			strEnvValue, ok := os.LookupEnv(e.entries[i].EnvName)

			if ok {
				intVal, err = strconv.ParseInt(strEnvValue, 10, 64)

				if err == nil {
					value = int(intVal)
				}
			}
			e.entries[i].Value.SetInt(int64(value))
		case "bool":
			var value bool
			var boolVal bool
			var ok bool
			var err error
			var ptrValue *bool

			ptrValue, ok = e.values[e.entries[i].Name].(*bool)

			if ok {
				value = *ptrValue
			} else {
				boolVal, err = strconv.ParseBool(e.entries[i].Default)

				if err == nil {
					value = boolVal
				}
			}

			strEnvValue, ok := os.LookupEnv(e.entries[i].EnvName)

			if ok {
				boolVal, err = strconv.ParseBool(strEnvValue)

				if err == nil {
					value = boolVal
				}
			}
			e.entries[i].Value.SetBool(value)
		default:
			log.Printf("Unknown type %s\n", e.entries[i].Type)
		}
	}

	return
}
