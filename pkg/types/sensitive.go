package types

import (
	"encoding"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

var (
	_ fmt.Formatter          = (*SensitiveString)(nil)
	_ json.Marshaler         = (*SensitiveString)(nil)
	_ yaml.Marshaler         = (*SensitiveString)(nil)
	_ encoding.TextMarshaler = (*SensitiveString)(nil)
)

type SensitiveString string

func (s SensitiveString) Format(f fmt.State, c rune) {
	f.Write([]byte("***"))
}

func (s SensitiveString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string("***"))
}

func (s SensitiveString) MarshalYAML() (any, error) {
	return json.Marshal(string("***"))
}

func (s SensitiveString) MarshalText() (text []byte, err error) {
	return []byte("***"), nil
}
