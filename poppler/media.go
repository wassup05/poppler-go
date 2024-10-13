package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type PopplerMedia struct {
	media popplerMedia
}

type popplerMedia *C.struct__PopplerMedia

type PopplerMediaSaveFunc func() bool

func (pm *PopplerMedia) GetAutoPlay() bool {
	return toGoBool(C.poppler_media_get_auto_play(pm.media))
}

func (pm *PopplerMedia) GetFilename() string {
	return toGoString(C.poppler_media_get_filename(pm.media))
}

func (pm *PopplerMedia) GetMimeType() string {
	return toGoString(C.poppler_media_get_mime_type(pm.media))
}

func (pm *PopplerMedia) GetRepeatCount() float64 {
	return float64(C.poppler_media_get_repeat_count(pm.media))
}

func (pm *PopplerMedia) GetShowControls() bool {
	return toGoBool(C.poppler_media_get_show_controls(pm.media))
}

func (pm *PopplerMedia) IsEmbedded() bool {
	return toGoBool(C.poppler_media_is_embedded(pm.media))
}

// Save saves the media to an absolute path.
// returns an error if any.
func (pm *PopplerMedia) Save(path string) *PopplerError {
	cfilename := C.CString(path)
	defer C.free(unsafe.Pointer(cfilename))

	var cerr *C.GError = nil

	result := C.poppler_media_save(pm.media, cfilename, &cerr)

	if result == 0 {
		return gErrorToPopplerError(*cerr)
	}

	return nil
}
