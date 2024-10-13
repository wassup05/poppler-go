package poppler

/*
#cgo  pkg-config : poppler-glib
#include <poppler.h>
#include <stdlib.h>
*/
import "C"

type PopplerColor struct {
	color popplerColor
}

type popplerColor *C.struct__PopplerColor

func (pc *PopplerColor) Red() uint16 {
	return uint16(pc.color.red)
}

func (pc *PopplerColor) Green() uint16 {
	return uint16(pc.color.green)
}

func (pc *PopplerColor) Blue() uint16 {
	return uint16(pc.color.blue)
}

func NewColor() *PopplerColor {
	value := C.poppler_color_new()

	return &PopplerColor{
		color: value,
	}
}

func (pc *PopplerColor) Free() {
	C.poppler_color_free(pc.color)
	pc.color = nil
}

func (pc *PopplerColor) Copy() *PopplerColor {
	res := C.poppler_color_copy(pc.color)

	return &PopplerColor{
		color: res,
	}
}
