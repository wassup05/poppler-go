package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
#include <stdlib.h>
*/
import "C"

type PopplerLayer struct {
	layer popplerLayer
}

type popplerLayer *C.struct__PopplerLayer

func (pl *PopplerLayer) GetRadioButtonGroupId() int {
	return int(C.poppler_layer_get_radio_button_group_id(pl.layer))
}

func (pl *PopplerLayer) GetTitle() string {
	return toGoString(C.poppler_layer_get_title(pl.layer))
}

func (pl *PopplerLayer) Hide() {
	C.poppler_layer_hide(pl.layer)
}

func (pl *PopplerLayer) IsParent() bool {
	return toGoBool(C.poppler_layer_is_parent(pl.layer))
}

func (pl *PopplerLayer) IsVisible() bool {
	return toGoBool(C.poppler_layer_is_visible(pl.layer))
}

func (pl *PopplerLayer) Show() {
	C.poppler_layer_show(pl.layer)
}
