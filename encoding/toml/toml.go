package toml

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/wosai/go-contrib/encoding"
)

type tomlCodec struct {
}

func (c tomlCodec) Name() string {
	return "toml"
}

func (c tomlCodec) Marshal(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := toml.NewEncoder(buf)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c tomlCodec) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

func init() {
	encoding.RegisterCodec(tomlCodec{})
}
