package poppler

/*
#cgo pkg-config: poppler-glib
#include <stdlib.h>
#include <poppler.h>
#include <cairo.h>
*/
import "C"

import (
	"errors"
	"unsafe"

	"github.com/ungerik/go-cairo"
)

type PopplerAnnotMapping struct {
	mapping popplerAnnotMapping
}

type popplerAnnotMapping *C.struct__PopplerAnnotMapping

type PopplerFormFieldMapping struct {
	mapping popplerFormFieldMapping
}

type popplerFormFieldMapping *C.struct__PopplerFormFieldMapping

func (pffm *PopplerFormFieldMapping) Area() *PopplerRectangle {
	return &PopplerRectangle{
		rect: &pffm.mapping.area,
	}
}

func (pffm *PopplerFormFieldMapping) Form() *PopplerFormField {
	return &PopplerFormField{
		form: pffm.mapping.field,
	}
}

type PopplerImageMapping struct {
	mapping popplerImageMapping
}

type popplerImageMapping *C.struct__PopplerImageMapping

func (pim *PopplerImageMapping) Area() *PopplerRectangle {
	return &PopplerRectangle{
		rect: &pim.mapping.area,
	}
}

func (pim *PopplerImageMapping) Id() int {
	return int(pim.mapping.image_id)
}

type PopplerLinkMapping struct {
	mapping popplerLinkMapping
}

type popplerLinkMapping *C.struct__PopplerLinkMapping

func (plm *PopplerLinkMapping) Area() *PopplerRectangle {
	return &PopplerRectangle{
		rect: &plm.mapping.area,
	}
}

func (plm *PopplerLinkMapping) Action() *PopplerAction {
	return &PopplerAction{
		action: plm.mapping.action,
	}
}

type PopplerAnnot struct {
	annot popplerAnnot
}

type popplerAnnot *C.struct__PopplerAnnot

func (pam *PopplerAnnotMapping) Area() *PopplerRectangle {
	return &PopplerRectangle{
		rect: &pam.mapping.area,
	}
}

func (pam *PopplerAnnotMapping) Annot() *PopplerAnnot {
	return &PopplerAnnot{
		annot: pam.mapping.annot,
	}
}

type PopplerRectangle struct {
	rect popplerRectangle
}

type popplerRectangle *C.struct__PopplerRectangle

type PopplerPageTransition struct {
	transition popplerPageTransition
}

type popplerPageTransition *C.struct__PopplerPageTransition

type PopplerPoint struct {
	point popplerPoint
}

type popplerPoint *C.struct__PopplerPoint

type PopplerQuadrilateral struct {
	quad popplerQuadrilateral
}

type popplerQuadrilateral *C.struct__PopplerQuadrilateral

type PopplerTextAttributes struct {
	attrs popplerTextAttributes
}

type popplerTextAttributes *C.struct__PopplerTextAttributes

func (pta *PopplerTextAttributes) Start() int {
	return int(pta.attrs.start_index)
}

func (pta *PopplerTextAttributes) End() int {
	return int(pta.attrs.end_index)
}

func (pta *PopplerTextAttributes) Font() string {
	return toGoString(pta.attrs.font_name)
}

func (pta *PopplerTextAttributes) IsUnderlined() bool {
	return toGoBool(pta.attrs.is_underlined)
}

func (pta *PopplerTextAttributes) FontSize() float64 {
	return float64(pta.attrs.font_size)
}

func (pta *PopplerTextAttributes) Color() *PopplerColor {
	return &PopplerColor{
		color: &pta.attrs.color,
	}
}

type PopplerPage struct {
	page popplerPage
}

type popplerPage *C.struct__PopplerPage

// Annot stuff

func (pa *PopplerAnnotMapping) Copy() *PopplerAnnotMapping {
	return &PopplerAnnotMapping{
		mapping: C.poppler_annot_mapping_copy(pa.mapping),
	}
}

func (pa *PopplerAnnotMapping) Free() {
	C.poppler_annot_mapping_free(pa.mapping)
	pa.mapping = nil
}

func NewAnnotMapping() *PopplerAnnotMapping {
	return &PopplerAnnotMapping{
		mapping: C.poppler_annot_mapping_new(),
	}
}

func (pp *PopplerPage) AddAnnot(pa *PopplerAnnot) {
	C.poppler_page_add_annot(pp.page, pa.annot)
}

// Form Field stuff

func (pa *PopplerFormFieldMapping) Copy() *PopplerFormFieldMapping {
	return &PopplerFormFieldMapping{
		mapping: C.poppler_form_field_mapping_copy(pa.mapping),
	}
}

func (pa *PopplerFormFieldMapping) Free() {
	C.poppler_form_field_mapping_free(pa.mapping)
	pa.mapping = nil
}

func NewFormFieldMapping() *PopplerFormFieldMapping {
	return &PopplerFormFieldMapping{
		mapping: C.poppler_form_field_mapping_new(),
	}
}

// Image mapping stuff

func (pim *PopplerImageMapping) Copy() *PopplerImageMapping {
	return &PopplerImageMapping{
		mapping: C.poppler_image_mapping_copy(pim.mapping),
	}
}

func (pim *PopplerImageMapping) Free() {
	C.poppler_image_mapping_free(pim.mapping)
	pim.mapping = nil
}

func NewImageMapping() *PopplerImageMapping {
	return &PopplerImageMapping{
		mapping: C.poppler_image_mapping_new(),
	}
}

// link mapping stuff

func (plm *PopplerLinkMapping) Copy() *PopplerLinkMapping {
	return &PopplerLinkMapping{
		mapping: C.poppler_link_mapping_copy(plm.mapping),
	}
}

func (plm *PopplerLinkMapping) Free() {
	C.poppler_link_mapping_free(plm.mapping)
	plm.mapping = nil
}

func NewLinkMapping() *PopplerLinkMapping {
	return &PopplerLinkMapping{
		mapping: C.poppler_link_mapping_new(),
	}
}

// Getting the pages of a document

func (pd *PopplerDocument) GetPage(index int) *PopplerPage {
	var page PopplerPage
	cpage := C.poppler_document_get_page(pd.document, C.int(index))

	page.page = cpage

	return &page
}

func (pd *PopplerDocument) GetPageByLabel(label string) *PopplerPage {
	var page PopplerPage

	clabel := C.CString(label)
	defer C.free(unsafe.Pointer(clabel))
	cpage := C.poppler_document_get_page_by_label(pd.document, clabel)

	page.page = cpage

	return &page
}

func (pp *PopplerPage) Close() {
	C.g_object_unref(C.gpointer(pp.page))
}

// Getting the properties of pages

func (pp *PopplerPage) GetDuration() float64 {
	return float64(C.poppler_page_get_duration(pp.page))
}

func (pp *PopplerPage) GetIndex() int {
	return int(C.poppler_page_get_index(pp.page))
}

func (pp *PopplerPage) GetLabel() string {
	return toGoString(C.poppler_page_get_label(pp.page))
}

func (pp *PopplerPage) GetSize() (float64, float64) {
	var h C.double
	var w C.double

	C.poppler_page_get_size(pp.page, &w, &h)

	return float64(h), float64(w)
}

// GetCropBox returns a PopplerRectangle covering the entire page
func (pp *PopplerPage) GetCropBox() *PopplerRectangle {
	box := NewRectangle()

	C.poppler_page_get_crop_box(pp.page, box.rect)

	return box
}

// GetBoundingBox returns a PopplerRectangle covering all the vector graphics, attachments in a page.
// also returns if the page has vector graphics.
func (pp *PopplerPage) GetBoundingBox() (*PopplerRectangle, bool) {
	box := NewRectangle()

	gbool := C.poppler_page_get_bounding_box(pp.page, box.rect)

	return box, toGoBool(gbool)
}

func (pp *PopplerPage) GetText() string {
	return toGoString(C.poppler_page_get_text(pp.page))
}

func (pp *PopplerPage) GetTextForArea(rect *PopplerRectangle) string {
	return toGoString(C.poppler_page_get_text_for_area(pp.page, rect.rect))
}

func (pp *PopplerPage) GetSelectedText(style PopplerSelectionStyle, selection *PopplerRectangle) string {
	return toGoString(C.poppler_page_get_selected_text(pp.page, C.PopplerSelectionStyle(style), selection.rect))
}

func (pp *PopplerPage) GetTextAttrs() []*PopplerTextAttributes {
	golist := make([]*PopplerTextAttributes, 0)

	res := C.poppler_page_get_text_attributes(pp.page)

	toGoList(res, func(gp C.gpointer) {
		golist = append(golist, &PopplerTextAttributes{
			attrs: (popplerTextAttributes)(gp),
		})
	})

	return golist
}

func (pp *PopplerPage) GetTextAttrsForArea(rect *PopplerRectangle) []*PopplerTextAttributes {
	golist := make([]*PopplerTextAttributes, 0)

	res := C.poppler_page_get_text_attributes_for_area(pp.page, rect.rect)

	toGoList(res, func(gp C.gpointer) {
		golist = append(golist, &PopplerTextAttributes{
			attrs: (popplerTextAttributes)(gp),
		})
	})

	return golist
}

func (pp *PopplerPage) GetThumbnailSize() (int, int) {
	var h C.int
	var w C.int

	result := C.poppler_page_get_thumbnail_size(pp.page, &w, &h)

	if result == 0 {
		return 0, 0
	}

	return int(h), int(w)
}

func (pp *PopplerPage) GetTransition() *PopplerPageTransition {
	return &PopplerPageTransition{
		transition: C.poppler_page_get_transition(pp.page),
	}
}

func (pp *PopplerPage) FindText(text string) []*PopplerRectangle {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	golist := make([]*PopplerRectangle, 0)

	res := C.poppler_page_find_text(pp.page, ctext)

	toGoList(res, func(gp C.gpointer) {
		golist = append(golist, &PopplerRectangle{
			rect: (popplerRectangle)(gp),
		})
	})

	return golist
}

// GetAnnotMapping returns a list of mapping of a area on a page to a Popplerannot
func (pp *PopplerPage) GetAnnotMapping() []*PopplerAnnotMapping {
	golist := make([]*PopplerAnnotMapping, 0)

	res := C.poppler_page_get_annot_mapping(pp.page)

	toGoList(res, func(gp C.gpointer) {
		golist = append(golist, &PopplerAnnotMapping{
			mapping: (popplerAnnotMapping)(gp),
		})
	})

	return golist
}

// GetFormFieldMapping returns a list of mapping of a area on a page to a Popplerformfield
func (pp *PopplerPage) GetFormFieldMapping() []*PopplerFormFieldMapping {
	golist := make([]*PopplerFormFieldMapping, 0)

	res := C.poppler_page_get_annot_mapping(pp.page)

	toGoList(res, func(gp C.gpointer) {
		golist = append(golist, &PopplerFormFieldMapping{
			mapping: (popplerFormFieldMapping)(gp),
		})
	})

	return golist
}

// GetImageMapping returns a list of mapping of a area on a page to a unique identifier to an image
func (pp *PopplerPage) GetImageMapping() []*PopplerImageMapping {
	golist := make([]*PopplerImageMapping, 0)

	res := C.poppler_page_get_annot_mapping(pp.page)

	toGoList(res, func(gp C.gpointer) {
		golist = append(golist, &PopplerImageMapping{
			mapping: (popplerImageMapping)(gp),
		})
	})

	return golist
}

// GetLinkMapping returns a list of mapping of a area on a page to a
func (pp *PopplerPage) GetLinkMapping() []*PopplerLinkMapping {
	golist := make([]*PopplerLinkMapping, 0)

	res := C.poppler_page_get_annot_mapping(pp.page)

	toGoList(res, func(gp C.gpointer) {
		golist = append(golist, &PopplerLinkMapping{
			mapping: (popplerLinkMapping)(gp),
		})
	})

	return golist
}

func (pp *PopplerPage) RemoveAnnot(annot *PopplerAnnot) {
	C.poppler_page_remove_annot(pp.page, annot.annot)
}

// Rendering stuff

func (pp *PopplerPage) GetImageSurface(index int) *cairo.Surface {
	cs := C.poppler_page_get_image(pp.page, C.int(index))
	ctx := C.cairo_create(cs)

	return cairo.NewSurfaceFromC(cairo.Cairo_surface(unsafe.Pointer(cs)), cairo.Cairo_context(unsafe.Pointer(ctx)))
}

func (pp *PopplerPage) GetThumbnailSurface() (*cairo.Surface, error) {
	cs := C.poppler_page_get_thumbnail(pp.page)

	if cs == nil {
		return nil, errors.New("No thumbnail for this page")
	}

	c := C.cairo_create(cs)

	return cairo.NewSurfaceFromC(cairo.Cairo_surface(unsafe.Pointer(cs)), cairo.Cairo_context(unsafe.Pointer(c))), nil
}

func (pp *PopplerPage) RenderToSurface(surface *cairo.Surface) {
	_, ctx := surface.Native()

	C.poppler_page_render(pp.page, (*C.cairo_t)(unsafe.Pointer(ctx)))
}

func (pp *PopplerPage) RenderToSurfaceForPrinting(surface *cairo.Surface) {
	_, ctx := surface.Native()

	C.poppler_page_render_for_printing(pp.page, (*C.cairo_t)(unsafe.Pointer(ctx)))
}

func (pp *PopplerPage) RenderToSurfaceForPrintingWithOpts(surface *cairo.Surface, opts PopplerPrintFlags) {
	_, ctx := surface.Native()

	C.poppler_page_render_for_printing_with_options(pp.page, (*C.cairo_t)(unsafe.Pointer(ctx)), C.PopplerPrintFlags(opts))
}

func (pp *PopplerPage) RenderToPS(psfile *PopplerPSFile) {
	C.poppler_page_render_to_ps(pp.page, psfile.psfile)
}

// transition stuff

func (ppt *PopplerPageTransition) Copy() *PopplerPageTransition {
	return &PopplerPageTransition{
		transition: C.poppler_page_transition_copy(ppt.transition),
	}
}

func (ppt *PopplerPageTransition) Free() {
	C.poppler_page_transition_free(ppt.transition)
	ppt.transition = nil
}

func NewPageTransition() *PopplerPageTransition {
	return &PopplerPageTransition{
		transition: C.poppler_page_transition_new(),
	}
}

// point stuff

func (pp *PopplerPoint) X() float64 {
	return float64(pp.point.x)
}

func (pp *PopplerPoint) Y() float64 {
	return float64(pp.point.y)
}

func (pp *PopplerPoint) Copy() *PopplerPoint {
	return &PopplerPoint{
		point: C.poppler_point_copy(pp.point),
	}
}

func (pp *PopplerPoint) Free() {
	C.poppler_point_free(pp.point)
	pp.point = nil
}

func NewPoint() *PopplerPoint {
	return &PopplerPoint{
		point: C.poppler_point_new(),
	}
}

// quadrilateral stuff

func (pq *PopplerQuadrilateral) P1() PopplerPoint {
	return PopplerPoint{
		point: &pq.quad.p1,
	}
}

func (pq *PopplerQuadrilateral) P2() PopplerPoint {
	return PopplerPoint{
		point: &pq.quad.p2,
	}
}

func (pq *PopplerQuadrilateral) P3() PopplerPoint {
	return PopplerPoint{
		point: &pq.quad.p3,
	}
}

func (pq *PopplerQuadrilateral) P4() PopplerPoint {
	return PopplerPoint{
		point: &pq.quad.p4,
	}
}

func (pq *PopplerQuadrilateral) Copy() *PopplerQuadrilateral {
	return &PopplerQuadrilateral{
		quad: C.poppler_quadrilateral_copy(pq.quad),
	}
}

func (pq *PopplerQuadrilateral) Free() {
	C.poppler_quadrilateral_free(pq.quad)
	pq.quad = nil
}

func NewQuadrilateral() *PopplerQuadrilateral {
	return &PopplerQuadrilateral{
		quad: C.poppler_quadrilateral_new(),
	}
}

// rectangle stuff

func (pr *PopplerRectangle) X1() float64 {
	return float64(pr.rect.x1)
}

func (pr *PopplerRectangle) X2() float64 {
	return float64(pr.rect.x2)
}

func (pr *PopplerRectangle) Y1() float64 {
	return float64(pr.rect.y1)
}

func (pr *PopplerRectangle) Y2() float64 {
	return float64(pr.rect.y2)
}

func (pr *PopplerRectangle) Copy() *PopplerRectangle {
	return &PopplerRectangle{
		rect: C.poppler_rectangle_copy(pr.rect),
	}
}

func (pr *PopplerRectangle) Free() {
	C.poppler_rectangle_free(pr.rect)
	pr.rect = nil
}

func NewRectangle() *PopplerRectangle {
	return &PopplerRectangle{
		rect: C.poppler_rectangle_new(),
	}
}

// Text attributes
func (pta *PopplerTextAttributes) Copy() *PopplerTextAttributes {
	return &PopplerTextAttributes{
		attrs: C.poppler_text_attributes_copy(pta.attrs),
	}
}

func (pta *PopplerTextAttributes) Free() {
	C.poppler_text_attributes_free(pta.attrs)
	pta.attrs = nil
}

func NewTextAttributes() *PopplerTextAttributes {
	return &PopplerTextAttributes{
		attrs: C.poppler_text_attributes_new(),
	}
}

// some convinience functions defined in poppler
