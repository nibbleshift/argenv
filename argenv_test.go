package argenv

import (
	"testing"
	"reflect"
	"os"
)

type Test1 struct {
    One string `default:"One" description:"Description of One"`
    SecondVariable string `default:"Second" description:"Description of SecondVariable"`
    ThirdVariableS int `default:"33" description:"Description of ThirdVariableS"`
}

type Test2 struct {
    Apple string `default:"One" description:"Description of One"`
    Orange string `default:"Second" description:"Description of SecondVariable"`
    Banana int `default:"33" description:"Description of ThirdVariableS"`
    AppleSauce int `default:"Random" description:"Description of ThirdVariableS"`
    AppleSauceS int `default:"MoreRandom" description:"Description of ThirdVariableS"`
}

var test1 *Test1
var test2 *Test2

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

	os.Setenv("APPLE", "Two")
	os.Setenv("ORANGE", "First")
	os.Setenv("BANANA", "66")
	os.Setenv("APPLE_SAUCE", "Toast")
	os.Setenv("APPLE_SAUCE_S", "Toasted")

    e := &ArgEnv{}
    test2 = &Test2{}
    e.Load(test2)


	assertEqual(t, test2.Apple, "Two")
	assertEqual(t, test2.Orange, "First")
	assertEqual(t, test2.Banana, 66)
}
