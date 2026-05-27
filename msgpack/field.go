package msgpack

import (
	"reflect"
)

type field struct {
	name      string
	omitEmpty bool
	array     bool
	index     []int
	typ       reflect.Type
	empty     reflect.Value
}

func collectFields(fields []*field, t reflect.Type, visited map[reflect.Type]bool, depth map[string]int, index []int) []*field {
	_ = "STUB: not implemented"
	// Break recursion
	return nil
}

// Skip field if not exported and not anonymous

// Skip field when field tag starts with "-"

// Flatten anonymous struct field

// Check for name collisions

// There is another field with same name and same depth
// Remove that field and skip this field

// Parse empty field tag

func fieldsForType(t reflect.Type) ([]*field, bool) { _ = "STUB: not implemented"; return nil, false }

func fieldByIndex(v reflect.Value, index []int) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}
