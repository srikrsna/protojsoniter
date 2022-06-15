package protojsoniter

import (
	jsoniter "github.com/json-iterator/go"
)

// Writer is the interface implemented by types that
// can write valid JSON represenation of themselves via *jsoniter.Stream
type Writer interface {
	WriteJSON(w *jsoniter.Stream)
}

// Reader is the interface implemented by types
// that can read themselves from a *jsoniter.Iterator.
type Reader interface {
	ReadJSON(iter *jsoniter.Iterator)
}
