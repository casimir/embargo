package eval

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	testStr := "${net.wlan0 speed}|${aaa${aaa}}|${net.wlan0 speed}|${bbbb"

	Convey("It should extract meaningful blocks → ${...}", t, func() {
		parts := Parse(testStr)
		keys := keys(parts)

		So(len(parts), ShouldEqual, 2)
		So(len(keys), ShouldEqual, 2)

		So(keys[0], ShouldEqual, "${"+parts[keys[0]].Text+"}")
		So(parts[keys[0]].Module, ShouldEqual, "net")
		So(parts[keys[0]].Text, ShouldEqual, "net.wlan0 speed")
		So(parts[keys[0]].Tokens, ShouldResemble, []string{"wlan0", "speed"})

		So(keys[1], ShouldEqual, "${"+parts[keys[1]].Text+"}")
		So(parts[keys[1]].Module, ShouldBeEmpty)
		So(parts[keys[1]].Text, ShouldEqual, "aaa${aaa")
		So(parts[keys[1]].Tokens, ShouldResemble, []string{"aaa${aaa"})
	})
}

func keys(vl TokenList) (keys []string) {
	for k, _ := range vl {
		keys = append(keys, k)
	}
	return
}
