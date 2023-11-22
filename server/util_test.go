package server

import (
	"bytes"
	"testing"
)

var topic = []byte("This is a sample byte slice to be converted to a string.")

func BenchmarkSliceConversion(b *testing.B) {
	b.Run("Option1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = string(topic[:])
		}
	})

	b.Run("Option2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = string(topic)
		}
	})

	b.Run("Option3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = bytesToString(topic)
		}
	})

	b.Run("Option4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = bytes.NewBuffer(topic).String()
		}
	})
}
