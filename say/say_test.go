package say

import "testing"

func TestSay(t *testing.T) {

	s, _ := Say("Hello ")

	if s != "Hello Hello Gopher xx" {
		t.Error("Expected 'Hello GopherHello Gopher' but got ", s)
	}
}
