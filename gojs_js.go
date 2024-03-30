package gojs

import (
	"syscall/js"
	"unsafe"
)

const nanHead = 0x7FF80000
const typeFlagObject = 1

var jsGo = predefValue(6, typeFlagObject) // instance of the Go class in JavaScript

var JSGo = *(*js.Value)(unsafe.Pointer(&jsGo))

type ref uint64
type Value struct {
	_     [0]func() // uncomparable; to make == not compile
	ref   ref       // identifies a JavaScript value, see ref type
	gcPtr *ref      // used to trigger the finalizer when the Value is not referenced any more
}

func predefValue(id uint32, typeFlag byte) Value {
	return Value{ref: (nanHead|ref(typeFlag))<<32 | ref(id)}
}
