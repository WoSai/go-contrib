package yaml

import (
	"github.com/wosai/go-contrib/encoding"
	"gopkg.in/yaml.v2"
)

type yamlCodec struct {
	name string
}

var _ encoding.Codec = (*yamlCodec)(nil)

func (c *yamlCodec) Name() string {
	return c.name
}

func (c *yamlCodec) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (c *yamlCodec) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func init() {
	encoding.RegisterCodec(&yamlCodec{name: "yaml"})
	encoding.RegisterCodec(&yamlCodec{"yml"})
}
