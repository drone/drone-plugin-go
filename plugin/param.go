package plugin

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

var Stdin *ParamSet

func init() {
	if len(os.Args) > 1 {
		buf := bytes.NewBufferString(os.Args[1])
		Stdin = NewParamSet(buf)
	} else {
		Stdin = NewParamSet(os.Stdin)
	}
}

type ParamSet struct {
	reader io.Reader
	params map[string]interface{}
}

func NewParamSet(reader io.Reader) *ParamSet {
	var p = new(ParamSet)
	p.reader = reader
	p.params = map[string]interface{}{}
	return p
}

// Param defines a parameter with the specified name.
func (p ParamSet) Param(name string, value interface{}) {
	p.params[name] = value
}

// Parse parses parameter definitions from the map.
func (p ParamSet) Parse() error {
	raw := map[string]json.RawMessage{}
	err := json.NewDecoder(p.reader).Decode(&raw)
	if err != nil {
		return err
	}

	for key, val := range p.params {
		err := json.Unmarshal(raw[key], val)
		if err != nil {
			return err
		}
	}

	return nil
}

// Param defines a parameter with the specified name.
func Param(name string, value interface{}) {
	Stdin.Param(name, value)
}

// Parse parses parameter definitions from the map.
func Parse() error {
	return Stdin.Parse()
}
