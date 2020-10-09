package chart

import (
	"errors"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type ValueFile struct {
	Filename string
	Values   map[interface{}]interface{}
}

func NewValueFile(fname string) (*ValueFile, error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	var vals map[interface{}]interface{}
	err = yaml.Unmarshal(buf, &vals)
	if err != nil {
		return nil, err
	}

	c := &ValueFile{
		Filename: fname,
		Values:   vals,
	}

	return c, nil
}

func (c ValueFile) Get(dotk string) string {
	ks := strings.Split(dotk, ".")

	return c.getVal(ks)
}

func (c ValueFile) Set(dotk string, val string) error {
	ks := strings.Split(dotk, ".")

	return c.setVal(ks, val)
}

func (c ValueFile) Save() error {
	b, err := yaml.Marshal(&c.Values)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.Filename, b, 0644)
}

func (c ValueFile) setVal(keys []string, val interface{}) error {
	vals := c.Values
	for _, k := range keys {

		if v, ok := vals[k]; ok {
			switch v.(type) {
			case map[interface{}]interface{}:
				vals = v.(map[interface{}]interface{})
			default:
				vals[k] = val
				return nil
			}
		} else {
			return errors.New("keys not found.")
		}
	}

	return errors.New("keys not found.")
}

func (c ValueFile) getVal(keys []string) string {
	vals := c.Values
	for _, k := range keys {

		if v, ok := vals[k]; ok {
			switch v.(type) {
			case map[interface{}]interface{}:
				vals = v.(map[interface{}]interface{})
			default:
				return v.(string)
			}
		} else {
			return ""
		}
	}

	return ""
}
