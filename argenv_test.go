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
	AppleSauce  int    `default:"0" description:"Description of ThirdVariableS"`
	AppleSauceS int    `default:"0" description:"Description of AppleSauceS`
	StrawBerry  string `default:"Apples" description:"Description of StrawBerry"`
}

type Test3 struct {
	ApiHost     string `default:"One" description:"Description of ApiHost"`
	DbHost      string `default:"Second" description:"Description of DbHost"`
	EnableGlob  int    `default:"0" description:"Description of EnableGlob"`
	AllowRoot   int    `default:"1" description:"Description of AllowRoot"`
	SecureCopyR int    `default:"2" description:"Description of SecureCopyR"`
}

type Test4 struct {
	FieldOne         string `default:"One" description:"Description of FieldOne"`
	FieldTwo         string `default:"Second" description:"Description of FieldTwo"`
	FieldThree       int    `default:"0" description:"Description of FieldThree"`
	FieldFour        int    `default:"1" description:"Description of FieldFour"`
	FieldFiveAndLast int    `default:"2" description:"Description of FieldFiveAndLast"`
}

type Test5 struct {
	ArgOne   string `default:"One" description:"Description of ArgOne"`
	ArgTwo   string `default:"Second" description:"Description of ArgTwo"`
	ArgThree int    `default:"0" description:"Description of ArgThree"`
	ArgFour  int    `default:"1" description:"Description of ArgFour"`
	ArgFive  int    `default:"2" description:"Description of ArgFiveAndLast"`
}

var test1 *Test1
var test2 *Test2
var test3 *Test3
var test4 *Test4
var test5 *Test5

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
	os.Setenv("STRAW_BERRY", "Fruit")
	os.Setenv("APPLE_SAUCE", "100")
	os.Setenv("APPLE_SAUCE_S", "101")

	e := &ArgEnv{}
	test2 = &Test2{}
	e.Load(test2)

	assertEqual(t, test2.Apple, "Two")
	assertEqual(t, test2.Orange, "First")
	assertEqual(t, test2.Banana, 66)
	assertEqual(t, test2.AppleSauce, 100)
	assertEqual(t, test2.AppleSauceS, 101)
	assertEqual(t, test2.StrawBerry, "Fruit")
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

func TestLoadArg(t *testing.T) {

	os.Args = append(os.Args, "-field-one=first field")
	os.Args = append(os.Args, "-field-two=second")
	os.Args = append(os.Args, "-field-three=155")
	os.Args = append(os.Args, "-field-four=56789")
	os.Args = append(os.Args, "-field-five-and-last=-1")

	e := &ArgEnv{}
	test4 = &Test4{}
	e.Load(test4)

	assertEqual(t, test4.FieldOne, "first field")
	assertEqual(t, test4.FieldTwo, "second")
	assertEqual(t, test4.FieldThree, 155)
	assertEqual(t, test4.FieldFour, 56789)
	assertEqual(t, test4.FieldFiveAndLast, -1)

}

func TestLoadEnvAndOverrideFlags(t *testing.T) {

	os.Args = append(os.Args, "-arg-one=first field")
	os.Args = append(os.Args, "-arg-two=second")
	os.Args = append(os.Args, "-arg-three=155")
	os.Args = append(os.Args, "-arg-four=56789")
	os.Args = append(os.Args, "-arg-five=-1")

	os.Setenv("ARG_ONE", "One")
	os.Setenv("ARG_TWO", "Two")
	os.Setenv("ARG_THREE", "3")
	os.Setenv("ARG_FOUR", "4")
	os.Setenv("ARG_FIVE", "5")

	e := &ArgEnv{}
	test5 = &Test5{}
	e.Load(test5)

	assertEqual(t, test5.ArgOne, "One")
	assertEqual(t, test5.ArgTwo, "Two")
	assertEqual(t, test5.ArgThree, 3)
	assertEqual(t, test5.ArgFour, 4)
	assertEqual(t, test5.ArgFive, 5)
}
