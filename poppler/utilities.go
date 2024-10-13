package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
#include <time.h>
#include <stdlib.h>
*/
import "C"

import (
	"time"
	"unsafe"
)

// CanParseDate checks if date is a valid pdf format date or not and returns the time corresponding to it
// returns an empty instance of time.Time if it is not in a valid format
func CanParseDate(date string) (bool, time.Time) {
	var ctime C.time_t

	cdate := C.CString(date)
	defer C.free(unsafe.Pointer(cdate))

	value := C.poppler_date_parse(cdate, &ctime)

	if value == C.int(0) {
		return false, time.Time{}
	}

	return true, time.Unix(int64(ctime), 0)
}

func togboolean(state bool) C.gboolean {
	var cstate C.gboolean = 0
	if state {
		cstate = 1
	}
	return cstate
}

func toGoString(cstr *C.char) string {
	gostr := C.GoString(cstr)
	C.free(unsafe.Pointer(cstr))

	return gostr
}

func toGoList(clist *C.GList, callback func(gp C.gpointer)) {
	if clist == nil {
		return
	}
	callback(clist.data)
	for node := clist.next; node != nil; node = node.next {
		callback(node.data)
	}
	C.g_list_free(clist)
}

func toGoBool(gbool C.gboolean) bool {
	return gbool != C.int(0)
}
