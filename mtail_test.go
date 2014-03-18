package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_pad(t *testing.T) {
	Convey("Returns the same value when length matches", t, func() {
		So(pad("hithere", 7), ShouldEqual, "hithere")
	})

	Convey("Returns the padded value when length less than", t, func() {
		So(pad("short", 10), ShouldEqual, "short     ")
	})

	Convey("Returns the stripped value when length more than", t, func() {
		So(pad("thisistoolong", 10), ShouldEqual, "...toolong")
	})
}
