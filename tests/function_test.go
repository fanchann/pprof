package tests

import (
	"learn/memprof/function"
	"testing"
)

func TestGenerateUserJson(t *testing.T) { // this test consume memory 709.14MB
	function.WriteFileJson(1000000)
}

func TestReadUserJson(t *testing.T) {
	_, ttl, err := function.ReadFileJson() // mem consume 397.38MB
	if err != nil {
		t.FailNow()
	}
	if ttl != 1000000 {
		t.Fail()
	}
}
