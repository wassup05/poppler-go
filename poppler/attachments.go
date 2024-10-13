package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

type PopplerAttachment struct {
	attachment popplerAttachment
}

type popplerAttachment *C.struct__PopplerAttachment

// TODO: Add custom save function support
type PopplerAttachmentSaveFunc func() (bool, *PopplerError)

func (pa *PopplerAttachment) GetCheckSum() string {
	res := C.poppler_attachment_get_checksum(pa.attachment)

	return toGoString(res.str)
}

func (pa *PopplerAttachment) GetDescription() string {
	return toGoString(C.poppler_attachment_get_description(pa.attachment))
}

func (pa *PopplerAttachment) GetName() string {
	return toGoString(C.poppler_attachment_get_name(pa.attachment))
}

func (pa *PopplerAttachment) GetSize() uint64 {
	res := C.poppler_attachment_get_size(pa.attachment)

	return uint64(res)
}

// Save saves the attachment to an absolute path.
// returns an error if any.
func (pa *PopplerAttachment) Save(path string) *PopplerError {
	cfilename := C.CString(path)
	defer C.free(unsafe.Pointer(cfilename))

	var cerr *C.GError = nil

	result := C.poppler_attachment_save(pa.attachment, cfilename, &cerr)

	if result == 0 {
		return gErrorToPopplerError(*cerr)
	}

	return nil
}

// WARN: Don't use this function as it results in a double free error although poppler docs claim that
// they must be freed
func (pa *PopplerAttachment) Free() {
	C.g_object_unref(C.gpointer(pa.attachment))
	pa.attachment = nil
}
