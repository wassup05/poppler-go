package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
*/
import "C"

type PopplerTextSpan struct {
	span popplerTextSpan
}

type popplerTextSpan *C.struct__PopplerTextSpan

func (pts *PopplerTextSpan) Copy() *PopplerTextSpan {
	return &PopplerTextSpan{
		span: C.poppler_text_span_copy(pts.span),
	}
}

func (pts *PopplerTextSpan) Free() {
	C.poppler_text_span_free(pts.span)
	pts.span = nil
}

func (pts *PopplerTextSpan) GetColor() *PopplerColor {
	var ccolor popplerColor
	C.poppler_text_span_get_color(pts.span, ccolor)

	return &PopplerColor{
		color: ccolor,
	}
}

func (pts *PopplerTextSpan) GetFontName() string {
	return toGoString(C.poppler_text_span_get_font_name(pts.span))
}

func (pts *PopplerTextSpan) GetText() string {
	return toGoString(C.poppler_text_span_get_text(pts.span))
}

func (pts *PopplerTextSpan) IsBoldFont() bool {
	return toGoBool(C.poppler_text_span_is_bold_font(pts.span))
}

func (pts *PopplerTextSpan) IsFixedWidthFont() bool {
	return toGoBool(C.poppler_text_span_is_fixed_width_font(pts.span))
}

func (pts *PopplerTextSpan) IsSerifFont() bool {
	return toGoBool(C.poppler_text_span_is_serif_font(pts.span))
}
