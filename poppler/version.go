package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
*/
import "C"

// GetBackend returns the backend poppler was compiled with
func GetBackend() PopplerBackend {
	back := C.poppler_get_backend()

	return PopplerBackend(back)
}

// GetVersion returns poppler version
func GetVersion() string {
	return toGoString(C.poppler_get_version())
}
