package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
#include <stdlib.h>
*/
import "C"

type PopplerMovie struct {
	movie popplerMovie
}

type popplerMovie *C.struct__PopplerMovie

func (pmo *PopplerMovie) GetAspect() (int, int) {
	var h C.int
	var w C.int

	C.poppler_movie_get_aspect(pmo.movie, &w, &h)

	return int(h), int(w)
}

func (pmo *PopplerMovie) GetDurationInNs() uint64 {
	return uint64(C.poppler_movie_get_duration(pmo.movie))
}

func (pmo *PopplerMovie) GetFilename() string {
	return toGoString(C.poppler_movie_get_filename(pmo.movie))
}

func (pmo *PopplerMovie) GetPlayMode() PopplerMoviePlayMode {
	res := C.poppler_movie_get_play_mode(pmo.movie)

	return PopplerMoviePlayMode(res)
}

func (pmo *PopplerMovie) GetRate() float64 {
	return float64(C.poppler_movie_get_rate(pmo.movie))
}

func (pmo *PopplerMovie) GetRotationAngle() uint16 {
	return uint16(C.poppler_movie_get_rotation_angle(pmo.movie))
}

func (pmo *PopplerMovie) GetStart() uint64 {
	return uint64(C.poppler_movie_get_start(pmo.movie))
}

func (pmo *PopplerMovie) GetVolume() float64 {
	return float64(C.poppler_movie_get_volume(pmo.movie))
}

func (pmo *PopplerMovie) IsSync() bool {
	return toGoBool(C.poppler_movie_is_synchronous(pmo.movie))
}

func (pmo *PopplerMovie) NeedsPoster() bool {
	return toGoBool(C.poppler_movie_need_poster(pmo.movie))
}

func (pmo *PopplerMovie) ShowControls() bool {
	return toGoBool(C.poppler_movie_show_controls(pmo.movie))
}
