package vite

import (
	"errors"
	"fmt"
)

var reservedKeys = map[string]bool{
	"file":           true,
	"src":            true,
	"isEntry":        true,
	"isDynamicEntry": true,
	"dynamicImports": true,
	"css":            true,
	"assets":         true,
}

func (v *Vite) Data() AssetsData {
	return v.data
}

func (v *Vite) SetArgs(vars AssetsData) error {
	for k, val := range vars {
		if !reservedKeys[k] {
			v.data[k] = val
		} else {
			return errors.New(fmt.Sprintf("Field %v is readonly", k))
		}
	}

	return nil
}

func (v *Vite) SetArg(k string, val any) error {
	if !reservedKeys[k] {
		v.data[k] = val
	} else {
		return errors.New(fmt.Sprintf("Field %v is readonly", k))
	}

	return nil
}
