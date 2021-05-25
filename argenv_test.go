package argenv

import (
	"testing"
)

type MyTest struct {
    One string `default:"One" description:"Description of One"`
    SecondVariable string `default:"Second" description:"Description of SecondVariable"`
    ThirdVariableS int `default:"33" description:"Description of ThirdVariableS"`
}

var test *MyTest

func TestLoad(t *testing.T) {
    e := &ArgEnv{}
    test = &MyTest{}
    e.Load(test)
}
