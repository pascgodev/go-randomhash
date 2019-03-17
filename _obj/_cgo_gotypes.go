// Code generated by cmd/cgo; DO NOT EDIT.

package randomhash

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_U8 = _Ctype_uint8_t

type _Ctype___uint8_t = _Ctype_uchar

type _Ctype_intgo = _Ctype_ptrdiff_t

type _Ctype_long int64

type _Ctype_ptrdiff_t = _Ctype_long

type _Ctype_size_t = _Ctype_ulong

type _Ctype_uchar uint8

type _Ctype_uint8_t = _Ctype___uint8_t

type _Ctype_ulong uint64

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


func _Cfunc_CBytes(b []byte) unsafe.Pointer {
	p := _cgo_cmalloc(uint64(len(b)))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], b)
	return p
}
//go:cgo_import_static _cgo_218d9a5f6027_Cfunc_RandomHash_Simple
//go:linkname __cgofn__cgo_218d9a5f6027_Cfunc_RandomHash_Simple _cgo_218d9a5f6027_Cfunc_RandomHash_Simple
var __cgofn__cgo_218d9a5f6027_Cfunc_RandomHash_Simple byte
var _cgo_218d9a5f6027_Cfunc_RandomHash_Simple = unsafe.Pointer(&__cgofn__cgo_218d9a5f6027_Cfunc_RandomHash_Simple)

//go:cgo_unsafe_args
func _Cfunc_RandomHash_Simple(p0 *_Ctype_U8, p1 *_Ctype_U8) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_218d9a5f6027_Cfunc_RandomHash_Simple, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}

func _Cfunc__CMalloc(n _Ctype_size_t) unsafe.Pointer {
	return _cgo_cmalloc(uint64(n))
}
//go:cgo_import_static _cgo_218d9a5f6027_Cfunc_free
//go:linkname __cgofn__cgo_218d9a5f6027_Cfunc_free _cgo_218d9a5f6027_Cfunc_free
var __cgofn__cgo_218d9a5f6027_Cfunc_free byte
var _cgo_218d9a5f6027_Cfunc_free = unsafe.Pointer(&__cgofn__cgo_218d9a5f6027_Cfunc_free)

//go:cgo_unsafe_args
func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_218d9a5f6027_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}

//go:cgo_import_static _cgo_218d9a5f6027_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_218d9a5f6027_Cfunc__Cmalloc _cgo_218d9a5f6027_Cfunc__Cmalloc
var __cgofn__cgo_218d9a5f6027_Cfunc__Cmalloc byte
var _cgo_218d9a5f6027_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_218d9a5f6027_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_218d9a5f6027_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
