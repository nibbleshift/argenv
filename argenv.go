package argenv

import (
	"flag"
	"fmt"
	"reflect"
	"unicode"
	"strings"
	"os"
	"strconv"
)


type ArgEnv struct {
	Entries []Entry
}

type Entry struct {
	Name string
	EnvName string
	FlagName string
	Type reflect.Type
	Value reflect.Value
	Description string
	Default string
}

func normalize(name string) (envName string, flagName string) {
	for i, c := range name {
		// envName
		if i != 0 && unicode.IsUpper(c) {
			envName += fmt.Sprintf("_%s", strings.ToUpper(string(c)))
		} else {
			envName += strings.ToUpper(string(c))
		}

		// flagName
		if i != 0 && unicode.IsUpper(c) {
			flagName += fmt.Sprintf("-%s", strings.ToLower(string(c)))
		} else {
			flagName += strings.ToLower(string(c))
		}
	}

	return
}

func setValue(o reflect.Value, field string, value interface{}) {
	switch v := value.(type) {
		case int:
			intVal,_ := value.(int)
			o.SetInt(int64(intVal))
		case string:
			strVal,_ := value.(string)
			o.SetString(strVal)
		default:
			fmt.Println("Unknown type: ", v)
	}
}

func setInt(o interface{}, field string, value int) {
    reflect.ValueOf(&o).Elem().FieldByName(field).SetInt(7)
}

func (e *ArgEnv) Load(o interface{}) {
    val := reflect.ValueOf(o).Elem()

	e.Entries = make([]Entry, val.NumField())

    for i:=0; i<val.NumField(); i++ {
		entry := Entry{}
		entry.Name = val.Type().Field(i).Name
		entry.EnvName, entry.FlagName = normalize(entry.Name)
		entry.Type = val.Type().Field(i).Type
		entry.Value = val.Field(i)
		entry.Description = reflect.TypeOf(o).Elem().Field(i).Tag.Get("description")
		entry.Default = reflect.TypeOf(o).Elem().Field(i).Tag.Get("default")
		e.Entries[i] = entry
    }

	for _, item := range e.Entries {
		switch item.Type.String() {
			case "string":
				var value string
				flag.StringVar(&value, item.FlagName, item.Default, item.Description)
				setValue(item.Value, item.FlagName, value)
			case "int":
				var value int
				i, _ := strconv.ParseInt(item.Default, 10, 64)
				flag.IntVar(&value, item.FlagName, int(i), item.Description)
				setValue(item.Value, item.FlagName, value)
			default:
				fmt.Printf("Unknown type %s\n", item.Type)
		}
	}
	flag.Parse()
}

var Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
}
