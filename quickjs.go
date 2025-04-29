/*
Package quickjs Go bindings to QuickJS: a fast, small, and embeddable ES2020 JavaScript interpreter
*/
package quickjs

/*
#cgo CFLAGS: -I./deps/include
#cgo darwin,amd64 LDFLAGS:  -lquickjs -lm -ldl
#cgo darwin,arm64 LDFLAGS:  -lquickjs -lm -ldl
#cgo linux,amd64 LDFLAGS:  -lquickjs -lm -ldl
#cgo linux,arm64 LDFLAGS:  -lquickjs -lm -ldl
#cgo windows,amd64 LDFLAGS: -lquickjs -lm -ldl
#cgo windows,386 LDFLAGS: -lquickjs -lm -ldl
*/
import "C"
