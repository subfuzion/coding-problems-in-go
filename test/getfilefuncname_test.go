package test

import "testing"

func TestGetFileFuncName(t *testing.T) {
	name := GetFileFuncName(GetFileFuncName)
	if name != "test/getfilefuncname.go/GetFileFuncName" {
		t.Errorf("Expected: %s; Actual: %s\n", "test/getfilefuncname.go/GetFileFuncName", name)
	}
}
