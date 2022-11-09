// Code generated by cmd/cgo; DO NOT EDIT.

package megasdkgo

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype__GoString_ string

type _Ctype_char int8

type _Ctype_int int32

type _Ctype_int64_t = _Ctype_long

type _Ctype_intgo = _Ctype_ptrdiff_t

type _Ctype_long int64

type _Ctype_ptrdiff_t = _Ctype_long

type _Ctype_struct_AddDownloadResp struct {
	gid		*_Ctype_char
	errorString	*_Ctype_char
	errorCode	_Ctype_int
	_		[4]byte
}

type _Ctype_struct_DownloadInfo struct {
	name		*_Ctype_char
	completedLength	_Ctype_int64_t
	totalLength	_Ctype_int64_t
	speed		_Ctype_int64_t
	state		_Ctype_int
	gid		*_Ctype_char
	errorCode	_Ctype_int
	errorString	*_Ctype_char
}

type _Ctype_struct_LoginResp struct {
	errorCode	_Ctype_int
	errorString	*_Ctype_char
}

//go:notinheap
type _Ctype_void_notinheap struct{}

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


// CString converts the Go string s to a C string.
//
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func _Cfunc_CString(s string) *_Ctype_char {
	if len(s)+1 <= 0 {
		panic("string too large")
	}
	p := _cgo_cmalloc(uint64(len(s)+1))
	sliceHeader := struct {
		p   unsafe.Pointer
		len int
		cap int
	}{p, len(s)+1, len(s)+1}
	b := *(*[]byte)(unsafe.Pointer(&sliceHeader))
	copy(b, s)
	b[len(s)] = 0
	return (*_Ctype_char)(p)
}

//go:linkname _cgo_runtime_gostring runtime.gostring
func _cgo_runtime_gostring(*_Ctype_char) string

// GoString converts the C string p into a Go string.
func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}
//go:cgo_import_static _cgo_898f5d15be17_Cfunc_addDownload
//go:linkname __cgofn__cgo_898f5d15be17_Cfunc_addDownload _cgo_898f5d15be17_Cfunc_addDownload
var __cgofn__cgo_898f5d15be17_Cfunc_addDownload byte
var _cgo_898f5d15be17_Cfunc_addDownload = unsafe.Pointer(&__cgofn__cgo_898f5d15be17_Cfunc_addDownload)

//go:cgo_unsafe_args
func _Cfunc_addDownload(p0 *_Ctype_char, p1 *_Ctype_char) (r1 *_Ctype_struct_AddDownloadResp) {
	_cgo_runtime_cgocall(_cgo_898f5d15be17_Cfunc_addDownload, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_898f5d15be17_Cfunc_cancelDownload
//go:linkname __cgofn__cgo_898f5d15be17_Cfunc_cancelDownload _cgo_898f5d15be17_Cfunc_cancelDownload
var __cgofn__cgo_898f5d15be17_Cfunc_cancelDownload byte
var _cgo_898f5d15be17_Cfunc_cancelDownload = unsafe.Pointer(&__cgofn__cgo_898f5d15be17_Cfunc_cancelDownload)

//go:cgo_unsafe_args
func _Cfunc_cancelDownload(p0 *_Ctype_char) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_898f5d15be17_Cfunc_cancelDownload, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_898f5d15be17_Cfunc_free
//go:linkname __cgofn__cgo_898f5d15be17_Cfunc_free _cgo_898f5d15be17_Cfunc_free
var __cgofn__cgo_898f5d15be17_Cfunc_free byte
var _cgo_898f5d15be17_Cfunc_free = unsafe.Pointer(&__cgofn__cgo_898f5d15be17_Cfunc_free)

//go:cgo_unsafe_args
func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_898f5d15be17_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_898f5d15be17_Cfunc_getDownloadByGid
//go:linkname __cgofn__cgo_898f5d15be17_Cfunc_getDownloadByGid _cgo_898f5d15be17_Cfunc_getDownloadByGid
var __cgofn__cgo_898f5d15be17_Cfunc_getDownloadByGid byte
var _cgo_898f5d15be17_Cfunc_getDownloadByGid = unsafe.Pointer(&__cgofn__cgo_898f5d15be17_Cfunc_getDownloadByGid)

//go:cgo_unsafe_args
func _Cfunc_getDownloadByGid(p0 *_Ctype_char) (r1 *_Ctype_struct_DownloadInfo) {
	_cgo_runtime_cgocall(_cgo_898f5d15be17_Cfunc_getDownloadByGid, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_898f5d15be17_Cfunc_initmega
//go:linkname __cgofn__cgo_898f5d15be17_Cfunc_initmega _cgo_898f5d15be17_Cfunc_initmega
var __cgofn__cgo_898f5d15be17_Cfunc_initmega byte
var _cgo_898f5d15be17_Cfunc_initmega = unsafe.Pointer(&__cgofn__cgo_898f5d15be17_Cfunc_initmega)

//go:cgo_unsafe_args
func _Cfunc_initmega(p0 *_Ctype_char) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_898f5d15be17_Cfunc_initmega, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_898f5d15be17_Cfunc_login
//go:linkname __cgofn__cgo_898f5d15be17_Cfunc_login _cgo_898f5d15be17_Cfunc_login
var __cgofn__cgo_898f5d15be17_Cfunc_login byte
var _cgo_898f5d15be17_Cfunc_login = unsafe.Pointer(&__cgofn__cgo_898f5d15be17_Cfunc_login)

//go:cgo_unsafe_args
func _Cfunc_login(p0 *_Ctype_char, p1 *_Ctype_char) (r1 *_Ctype_struct_LoginResp) {
	_cgo_runtime_cgocall(_cgo_898f5d15be17_Cfunc_login, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}

//go:cgo_import_static _cgo_898f5d15be17_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_898f5d15be17_Cfunc__Cmalloc _cgo_898f5d15be17_Cfunc__Cmalloc
var __cgofn__cgo_898f5d15be17_Cfunc__Cmalloc byte
var _cgo_898f5d15be17_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_898f5d15be17_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_898f5d15be17_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
