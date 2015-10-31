package pg_query

// Note(LukasFittl): This needs Go 1.5 for $SRCDIR support, see
// https://github.com/golang/go/commit/131758183f7dc2610af489da3a7fcc4d30c6bc48

/*
#cgo CFLAGS: -I${SRCDIR}/tmp/libpg_query-master
#cgo LDFLAGS: -L${SRCDIR}/tmp/libpg_query-master -lpg_query -fstack-protector
#include <pg_query.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func init() {
	C.pg_query_init()
}

func Parse(input string) string {
	input_c := C.CString(input)
	defer C.free(unsafe.Pointer(input_c))

	result_c := C.pg_query_parse(input_c)
	defer C.free(unsafe.Pointer(result_c.parse_tree))
	defer C.free(unsafe.Pointer(result_c.stderr_buffer))

	result := C.GoString(result_c.parse_tree)

	return result
}