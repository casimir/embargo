package eval

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type TestObj struct {
	IntAttr int
	StrAttr string
}

func (s TestObj) DumbMethod(a, b string) string {
	return "DumbOK" + s.StrAttr
}

func (s TestObj) DumbMethod2(a string) int {
	return 42
}

func (s TestObj) DumbMethod3(a string, b int) string {
	return "DumbKO"
}

func TestEvalAttr(t *testing.T) {
	testVar := TestObj{StrAttr: "some string", IntAttr: 42}

	Convey("It should give attribute's value as string from its name.", t, func() {
		So(evalAttr(testVar, "intAttr"), ShouldEqual, "42")
		So(evalAttr(&testVar, "strAttr"), ShouldEqual, "some string")
		So(evalAttr(testVar, "StrAttr"), ShouldEqual, "some string")

		So(evalAttr(testVar, "nope"), ShouldEqual, NopeStr)
		So(evalAttr(testVar, ""), ShouldEqual, NopeStr)
		So(evalAttr(nil, ""), ShouldEqual, NopeStr)
	})
}

func TestEvalMethod(t *testing.T) {
	testVar := TestObj{StrAttr: "some string", IntAttr: 42}

	Convey("It should return method result as string from its name.", t, func() {
		So(evalMethod(&testVar, "DumbMethod", "", "1"), ShouldEqual, "DumbOK"+testVar.StrAttr)
		So(evalMethod(&testVar, "dumbMethod", "", "1"), ShouldEqual, "DumbOK"+testVar.StrAttr)

		So(evalMethod(&testVar, "dumbMethod"), ShouldEqual, NopeStr)
		So(evalMethod(&testVar, "dumbMethod2", ""), ShouldEqual, NopeStr)
		So(evalMethod(&testVar, "dumbMethod3", "", "1"), ShouldEqual, NopeStr)
		So(evalMethod(&testVar, "nope"), ShouldEqual, NopeStr)
		So(evalMethod(&testVar, ""), ShouldEqual, NopeStr)
		So(evalMethod(nil, ""), ShouldEqual, NopeStr)
	})
}

func TestStruct(t *testing.T) {
	testVar := TestObj{StrAttr: "some string", IntAttr: 42}
	testStruct := New(testVar)

	Convey("It should return attribute or method result as string from its name.", t, func() {
		So(testStruct.Eval("intAttr"), ShouldEqual, "42")
		So(testStruct.Eval("dumbMethod", "", "1"), ShouldEqual, "DumbOK"+testVar.StrAttr)
	})
}

func TestEval(t *testing.T) {
	testVar := TestObj{StrAttr: "some string", IntAttr: 42}
	Register("test", New(testVar))

	Convey("It should return attribute or method result as string from its name.", t, func() {
		So(Eval("${test.strAttr}"), ShouldEqual, testVar.StrAttr)
		So(Eval("${test.dumbMethod a b}"), ShouldEqual, "DumbOK"+testVar.StrAttr)
		So(Eval("${test.strAttr a}"), ShouldEqual, NopeStr)
		So(Eval("${nope.toto}"), ShouldEqual, "${nope.toto}")
	})
}
