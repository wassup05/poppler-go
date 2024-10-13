package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
*/
import "C"

type PopplerStrctureElement struct {
	element popplerStructureElement
}

type popplerStructureElement *C.struct__PopplerStructureElement

type PopplerStrctureElementIter struct {
	elementIter popplerStructureElementIter
}

type popplerStructureElementIter *C.struct__PopplerStrctureElementIter

func (pse *PopplerStrctureElement) GetAbbr() string {
	return toGoString(C.poppler_structure_element_get_abbreviation(pse.element))
}

func (pse *PopplerStrctureElement) GetActualText() string {
	return toGoString(C.poppler_structure_element_get_actual_text(pse.element))
}

func (pse *PopplerStrctureElement) GetAltText() string {
	return toGoString(C.poppler_structure_element_get_alt_text(pse.element))
}

func (pse *PopplerStrctureElement) GetBgColor() *PopplerColor {
	var color popplerColor
	res := C.poppler_structure_element_get_background_color(pse.element, color)

	if res == C.int(0) {
		return nil
	}

	return &PopplerColor{
		color: color,
	}
}

func (pse *PopplerStrctureElement) GetBaselineShift() float64 {
	return float64(C.poppler_structure_element_get_baseline_shift(pse.element))
}

func (pse *PopplerStrctureElement) GetBlockAlign() PopplerStructureBlockAlign {
	res := C.poppler_structure_element_get_block_align(pse.element)

	return PopplerStructureBlockAlign(res)
}
