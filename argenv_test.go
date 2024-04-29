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
	AppleSauceS int    `default:"0" description:"Description of AppleSauceS"`
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

type Test6 struct {
	DefaultOne   string `default:"One" description:"Description of DefaultOne"`
	DefaultTwo   string `default:"Second" description:"Description of DefaultTwo"`
	DefaultThree int    `default:"0" description:"Description of DefaultThree"`
	DefaultFour  int    `default:"1" description:"Description of DefaultFour"`
	DefaultFive  int    `default:"2" description:"Description of DefaultFiveAndLast"`
}

type Test7 struct {
	BoolOne   bool `default:"true" description:"Description of BoolOne"`
	BoolTwo   bool `default:"false" description:"Description of BoolTwo"`
	BoolThree bool `default:"true" description:"Description of BoolThree"`
	BoolFour  bool `default:"false" description:"Description of BoolFour"`
}

type Test8 struct {
	ArgBoolOne   bool `default:"true" description:"Description of ArgBoolOne"`
	ArgBoolTwo   bool `default:"false" description:"Description of ArgBoolTwo"`
	ArgBoolThree bool `default:"true" description:"Description of ArgBoolThree"`
	ArgBoolFour  bool `default:"false" description:"Description of ArgBoolFour"`
}

type Test9 struct {
	EnvBoolOne   bool `default:"true" description:"Description of EnvBoolOne"`
	EnvBoolTwo   bool `default:"true" description:"Description of EnvBoolTwo"`
	EnvBoolThree bool `default:"false" description:"Description of EnvBoolThree"`
	EnvBoolFour  bool `default:"false" description:"Description of EnvBoolFour"`
}

type Test10 struct {
	Float32One   float64 `default:"1.1" description:"Description of EnvFloatOne"`
	Float32Two   float64 `default:"1.2" description:"Description of EnvFloatTwo"`
	Float32Three float64 `default:"0.1" description:"Description of EnvFloatThree"`
	Float32Four  float64 `default:"0.001" description:"Description of EnvFloatFour"`
}

type Test11 struct {
	InitOne   float64 `default:"1.1" description:"Description of EnvInitOne"`
	InitTwo   string  `default:"Test" description:"Description of EnvInitTwo"`
	InitThree bool    `default:"true" description:"Description of EnvInitThree"`
	InitFour  int     `default:"5" description:"Description of EnvInitFour"`
}

var test1 *Test1
var test2 *Test2
var test3 *Test3
var test4 *Test4
var test5 *Test5
var test6 *Test6
var test7 *Test7
var test8 *Test8
var test9 *Test9
var test10 *Test10

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

func TestLoadDefaults(t *testing.T) {
	e := &ArgEnv{}
	test6 = &Test6{}
	e.Load(test6)

	assertEqual(t, test6.DefaultOne, "One")
	assertEqual(t, test6.DefaultTwo, "Second")
	assertEqual(t, test6.DefaultThree, 0)
	assertEqual(t, test6.DefaultFour, 1)
	assertEqual(t, test6.DefaultFive, 2)
}

func TestLoadDefaultBool(t *testing.T) {
	e := &ArgEnv{}
	test7 = &Test7{}
	e.Load(test7)

	assertEqual(t, test7.BoolOne, true)
	assertEqual(t, test7.BoolTwo, false)
	assertEqual(t, test7.BoolThree, true)
	assertEqual(t, test7.BoolFour, false)
}

func TestLoadArgBool(t *testing.T) {
	os.Args = append(os.Args, "-arg-bool-one=false")
	os.Args = append(os.Args, "-arg-bool-two=true")
	os.Args = append(os.Args, "-arg-bool-three=true")
	os.Args = append(os.Args, "-arg-bool-four=true")

	e := &ArgEnv{}
	test8 = &Test8{}
	e.Load(test8)

	assertEqual(t, test8.ArgBoolOne, false)
	assertEqual(t, test8.ArgBoolTwo, true)
	assertEqual(t, test8.ArgBoolThree, true)
	assertEqual(t, test8.ArgBoolFour, true)
}

func TestLoadEnvBool(t *testing.T) {
	os.Setenv("ENV_BOOL_ONE", "false")
	os.Setenv("ENV_BOOL_TWO", "false")
	os.Setenv("ENV_BOOL_THREE", "true")
	os.Setenv("ENV_BOOL_FOUR", "true")

	e := &ArgEnv{}
	test9 = &Test9{}
	e.Load(test9)

	assertEqual(t, test9.EnvBoolOne, false)
	assertEqual(t, test9.EnvBoolTwo, false)
	assertEqual(t, test9.EnvBoolThree, true)
	assertEqual(t, test9.EnvBoolFour, true)
}

func TestLoadDefaultFloat32(t *testing.T) {
	e := &ArgEnv{}
	test10 = &Test10{}
	e.Load(test10)

	assertEqual(t, test10.Float32One, 1.1)
	assertEqual(t, test10.Float32Two, 1.2)
	assertEqual(t, test10.Float32Three, 0.1)
	assertEqual(t, test10.Float32Four, 0.001)
}

func TestInitVariety(t *testing.T) {
	test11 := &Test11{}
	Init(test11)
	assertEqual(t, test11.InitOne, 1.1)
	assertEqual(t, test11.InitTwo, "Test")
	assertEqual(t, test11.InitThree, true)
	assertEqual(t, test11.InitFour, 5)
}
