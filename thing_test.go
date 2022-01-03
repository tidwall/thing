package thing

import "testing"

func TestThing(t *testing.T) {
	v := MakeThing(int(1))
	if v.Value() != 1 {
		t.Fatal("bad news")
	}
}
