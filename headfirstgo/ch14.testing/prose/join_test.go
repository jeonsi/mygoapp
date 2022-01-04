package prose

import (
	"fmt"
	"testing"
)

func TestOneElement(t *testing.T) {
	// t.Error("no test written yet")
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		// t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
		t.Error(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	// t.Error("no test written yet")
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		// t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
		t.Error(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	// t.Error("no test written yet")
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		// t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
		t.Error(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
