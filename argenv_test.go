package argenv

import (
	"os"
	"reflect"
	"testing"
)

type Test1 struct {
	One            string `default:"One" description:"Description of One"`
	SecondVariable string `default:"Second" description:"Description of SecondVariable"`
	ThirdVariableS int    `default:"33" description:"Description of ThirdVariableS"`
}

type Test2 struct {
	Apple       string `default:"One" description:"Description of One"`
	Orange      string `default:"Second" description:"Description of SecondVariable"`
	Banana      int    `default:"33" description:"Description of ThirdVariableS"`
	AppleSauce  int    `default:"Random" description:"Description of ThirdVariableS"`
	AppleSauceS int    `default:"MoreRandom" description:"Description of ThirdVariableS"`
}

type Test3 struct {
	ApiHost     string `default:"One" description:"Description of ApiHost"`
	DbHost      string `default:"Second" description:"Description of DbHost"`
	EnableGlob  int    `default:"1" description:"Description of EnableGlob"`
	AllowRoot   int    `default:"0" description:"Description of AllowRoot"`
	SecureCopyR int    `default:"1" description:"Description of SecureCopyR"`
}

var test1 *Test1
var test2 *Test2
var test3 *Test3

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

func TestLoad(t *testing.T) {
	e := &ArgEnv{}
	test1 = &Test1{}
	e.Load(test1)

	assertEqual(t, test1.One, "One")
	assertEqual(t, test1.SecondVariable, "Second")
	assertEqual(t, test1.ThirdVariableS, 33)
}

func TestLoadEnv(t *testing.T) {
	e := &ArgEnv{}
	test2 = &Test2{}
	e.Load(test2)

	os.Setenv("APPLE", "Two")
	os.Setenv("ORANGE", "First")
	os.Setenv("BANANA", "66")
	os.Setenv("APPLE_SAUCE", "Toast")
	os.Setenv("APPLE_SAUCE_S", "Toasted")

}

func TestUsage(t *testing.T) {
	//os.Args = append(os.Args, "-h")
	os.Setenv("APPLE", "Two")
	os.Setenv("ORANGE", "First")
	os.Setenv("BANANA", "66")
	os.Setenv("APPLE_SAUCE", "Toast")
	os.Setenv("APPLE_SAUCE_S", "Toasted")

	e := &ArgEnv{}
	test3 = &Test3{}
	e.Load(test3)
}
