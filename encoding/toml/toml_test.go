package toml

import (
	"fmt"
	"testing"
)

type data struct {
	Name string
	Age  int
}

func TestMarshal(t *testing.T) {
	d1 := &data{"Luca", 18}
	d2, err := tomlCodec{}.Marshal(d1)
	if err != nil {
		t.FailNow()
	}

	fmt.Println(string(d2))
}
