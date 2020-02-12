package cardhandlers

import (
	"testing"
)

func TestDoWaifu2x(t *testing.T) {
	test := &Waifu2xAble{FileBaseName: "1996true.png"}
	err := test.DoWaifu2x()
	if err != nil {
		t.Error(err)
	}
}
