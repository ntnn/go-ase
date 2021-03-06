// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

type ValueConverter struct{}

var DefaultValueConverter ValueConverter

func (conv ValueConverter) ConvertValue(v interface{}) (driver.Value, error) {
	// Check the default value converter
	if driver.IsValue(v) {
		return v, nil
	}

	// Convert any values that can be handled as another type
	switch value := v.(type) {
	case int:
		return int64(value), nil
	case uint:
		return uint64(value), nil
	}

	// Check the reflect types if the value is handled.
	sv := reflect.TypeOf(v)
	for _, kind := range ReflectTypes {
		if kind == sv {
			return v, nil
		}
	}

	return nil, fmt.Errorf("unsupported type %T, a %s", v, sv.Kind())
}
