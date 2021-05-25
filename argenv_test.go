package argenv

import (
	"testing"
	"reflect"
)

type MyTest struct {
    One string `default:"One" description:"Description of One"`
    SecondVariable string `default:"Second" description:"Description of SecondVariable"`
    ThirdVariableS int `default:"33" description:"Description of ThirdVariableS"`
}

var test *MyTest

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}


func TestLoadValues(t *testing.T) {
    e := &ArgEnv{}
    test = &MyTest{}
    e.Load(test)

	assertEqual(t, test.One, "One")
	assertEqual(t, test.SecondVariable, "Second")
	assertEqual(t, test.ThirdVariableS, 33)
}
