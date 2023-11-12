package libs

import "testing"

func TestLogin(t *testing.T) {
	Greeting()
	if Login("admin", "234") != true {
		t.Error("error test fail")
	}
}