package conf

import (
	"testing"
)

func TestMakeConf(t *testing.T) {
	c := MakeConf("conf.json")

	if c == nil {
		t.Error("TestMakeConf : MakeConf(conf.json) fail!")
	}
}
