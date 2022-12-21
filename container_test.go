package gocontainer

import (
	"testing"
)

func TestContainer(t *testing.T) {
	Register("string_example", "123")
	t.Run("register values", func(t *testing.T) {
		var value = Get("string_example")
		if value != "123" {
			t.Error("was expected 123 and return", value)
			return
		}
	})
	t.Run("replace register", func(t *testing.T) {
		Register("string_example", "321")
		var value = Get("string_example")
		if value != "321" {
			t.Error("was expected 321 and return", value)
			return
		}
	})
	t.Run("has register", func(t *testing.T) {
		if !Has("string_example") {
			t.Error("was expected true")
			return
		}
		if Has("string_example_2") {
			t.Error("was expected false")
			return
		}
	})
	t.Run("delete register", func(t *testing.T) {
		Register("int_example", 100)
		Delete("int_example")
		if Has("int_example") {
			t.Error("was expected false")
		}
	})
}
