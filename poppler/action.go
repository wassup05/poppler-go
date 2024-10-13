package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
*/
import "C"

type PopplerAction struct {
	action popplerAction
}

type popplerAction *C.PopplerAction

func (pa *PopplerAction) Copy() *PopplerAction {
	return &PopplerAction{
		action: C.poppler_action_copy(pa.action),
	}
}

func (pa *PopplerAction) Free() {
	C.poppler_action_free(pa.action)
	pa.action = nil
}

func (pd *PopplerDest) Copy() *PopplerDest {
	return &PopplerDest{
		dest: C.poppler_dest_copy(pd.dest),
	}
}

func (pd *PopplerDest) Free() {
	C.poppler_dest_free(pd.dest)
	pd.dest = nil
}
