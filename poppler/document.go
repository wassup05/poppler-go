// Go Bindings for the poppler library (glib)
package poppler

/*
#cgo pkg-config: poppler-glib
#include <stdlib.h>
#include <poppler.h>
#include <glib.h>
*/
import "C"

import (
	"errors"
	"time"
	"unsafe"
)

type PopplerDocument struct {
	document popplerDocument
}

type popplerDocument *C.struct__PopplerDocument

type PopplerFontInfo struct {
	info popplerFontInfo
}

type popplerFontInfo *C.struct__PopplerFontInfo

type PopplerFontsIter struct {
	fontsIter popplerFontsIter
}

type popplerFontsIter *C.struct__PopplerFontsIter

type PopplerIndexIter struct {
	indexIter popplerIndexIter
}

type popplerIndexIter *C.struct__PopplerIndexIter

type PopplerLayersIter struct {
	layersIter popplerLayersIter
}

type popplerLayersIter *C.struct__PopplerLayersIter

type PopplerPSFile struct {
	psfile popplerPSFile
}

type popplerPSFile *C.struct__PopplerPSFile

type PopplerPageRange struct {
	prange popplerPageRange
}

type popplerPageRange *C.struct__PopplerPageRange

type PopplerDest struct {
	dest popplerDest
}

type popplerDest *C.struct__PopplerDest

func (pd *PopplerDocument) FindDest(name string) *PopplerDest {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	res := C.poppler_document_find_dest(pd.document, cname)

	if res == nil {
		return nil
	}

	return &PopplerDest{
		dest: res,
	}
}

// NewDocFromFilename returns a new PopplerDocument on which many further methods can be used.
// path should be absolute.
// returns a PopplerError in case of any error. Use PopplerError.Message to check the exact error.
func NewDocFromFilename(path string) (*PopplerDocument, *PopplerError) {
	cfilename := C.CString(path)
	defer C.free(unsafe.Pointer(cfilename))

	filenameUri := C.g_filename_to_uri(cfilename, nil, nil)

	defer C.free(unsafe.Pointer(filenameUri))

	var cerr *C.GError = nil

	var cdoc popplerDocument

	cdoc = C.poppler_document_new_from_file(filenameUri, nil, &cerr)

	if cdoc != nil {
		return &PopplerDocument{
			document: cdoc,
		}, nil
	}

	return nil, gErrorToPopplerError(*cerr)
}

// NewDocFromFilenameWithPass same as NewDocFromFilename but for PDF files that are encrypted
func NewDocFromFilenameWithPass(filename string, passwd string) (*PopplerDocument, *PopplerError) {
	cfilename := C.CString(filename)
	cpasswd := C.CString(passwd)
	defer C.free(unsafe.Pointer(cfilename))
	defer C.free(unsafe.Pointer(cpasswd))

	filenameUri := C.g_filename_to_uri(cfilename, nil, nil)

	defer C.free(unsafe.Pointer(filenameUri))

	var cerr *C.GError = nil

	var cdoc popplerDocument

	cdoc = C.poppler_document_new_from_file(filenameUri, cpasswd, &cerr)

	if cdoc != nil {
		return &PopplerDocument{
			document: cdoc,
		}, nil
	}

	return nil, gErrorToPopplerError(*cerr)
}

// NewDocFromBytes same as NewDocFromFilename but with bytes as input
func NewDocFromBytes(b []byte) (*PopplerDocument, *PopplerError) {
	cbytes := C.CBytes(b)
	defer C.free(cbytes)

	gbytes := C.g_bytes_new(C.gconstpointer(cbytes), C.size_t(len(b)))
	var cerr *C.GError = nil

	var cdoc popplerDocument

	cdoc = C.poppler_document_new_from_bytes(gbytes, nil, &cerr)

	if cdoc != nil {
		return &PopplerDocument{
			document: cdoc,
		}, nil
	}

	return nil, gErrorToPopplerError(*cerr)
}

// Save saves a PopplerDocument to a path.
// path must be absolute.
// returns a *PopplerError in case of any error.
func (pd *PopplerDocument) Save(filename string) *PopplerError {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	filenameUri := C.g_filename_to_uri(cfilename, nil, nil)

	defer C.free(unsafe.Pointer(filenameUri))

	var cerr *C.GError = nil

	result := C.poppler_document_save(pd.document, filenameUri, &cerr)

	if result == 0 {
		return gErrorToPopplerError(*cerr)
	}

	return nil
}

// SaveACopy saves a copy of the original document. Any change made in the document such as form fields filled will mot be reflected.
// path must be absolute.
// returns a *PopplerError in case of any error.
func (pd *PopplerDocument) SaveACopy(filename string) *PopplerError {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	filenameUri := C.g_filename_to_uri(cfilename, nil, nil)

	defer C.free(unsafe.Pointer(filenameUri))

	var cerr *C.GError = nil

	result := C.poppler_document_save_a_copy(pd.document, filenameUri, &cerr)

	if result == 0 {
		return gErrorToPopplerError(*cerr)
	}

	return nil
}

// Close frees the memory associated with a Document.
// Use it after freeing all the opened pages individually.
func (pd *PopplerDocument) Close() {
	C.g_object_unref(C.gpointer(pd.document))
	pd.document = nil
}

// Getting the properties of a document

// GetAuthor returns the author or an empty string in case the field is not present
func (pd *PopplerDocument) GetAuthor() string {
	return toGoString(C.poppler_document_get_author(pd.document))
}

// GetCreator returns the creator or an empty string in case the field is not present
func (pd *PopplerDocument) GetCreator() string {
	return toGoString(C.poppler_document_get_creator(pd.document))
}

// GetCreationDate returns the creation date in time.Time.
// returns an empty instance of time.Time (which can be checked with isZero() method provided by the `time` stdlib) and an error in case no field available
func (pd *PopplerDocument) GetCreationDate() (time.Time, error) {
	ctime := C.poppler_document_get_creation_date(pd.document)

	if ctime == C.long(-1) {
		return time.Time{}, errors.New("No date available")
	}

	return time.Unix(int64(ctime), 0), nil
}

// GetCreationDateTime same as GetCreationDate.
// included separately because poppler provides a separate function for it which uses GDateTime.
func (pd *PopplerDocument) GetCreationDateTime() (time.Time, error) {
	gdtime := C.poppler_document_get_creation_date_time(pd.document)

	if gdtime == nil {
		return time.Time{}, errors.New("No datetime available")
	}

	ctime := C.g_date_time_to_unix(gdtime)
	C.g_date_time_unref(gdtime)

	return time.Unix(int64(ctime), 0), nil
}

// GetModDate same as GetCreationDate but returns Modified date instead
func (pd *PopplerDocument) GetModDate() (time.Time, error) {
	ctime := C.poppler_document_get_modification_date(pd.document)

	if ctime == C.long(-1) {
		return time.Time{}, errors.New("No date available")
	}

	return time.Unix(int64(ctime), 0), nil
}

// GetModDateTime same as GetCreationDateTime but returns Modified date time instead
func (pd *PopplerDocument) GetModDateTime() (time.Time, error) {
	gdtime := C.poppler_document_get_modification_date_time(pd.document)

	if gdtime == nil {
		return time.Time{}, errors.New("No datetime available")
	}

	ctime := C.g_date_time_to_unix(gdtime)
	C.g_date_time_unref(gdtime)

	return time.Unix(int64(ctime), 0), nil
}

// GetKeywords returns keywords or empty string in case field not present
func (pd *PopplerDocument) GetKeywords() string {
	return toGoString(C.poppler_document_get_keywords(pd.document))
}

// GetMetadata returns metadata or empty string in case field not present
func (pd *PopplerDocument) GetMetadata() string {
	return toGoString(C.poppler_document_get_metadata(pd.document))
}

// GetAttachmentCount returns attachment count in Document
func (pd *PopplerDocument) GetAttachmentCount() uint {
	return uint(C.poppler_document_get_n_attachments(pd.document))
}

// GetPageCount returns page count in Document
func (pd *PopplerDocument) GetPageCount() uint {
	return uint(C.poppler_document_get_n_pages(pd.document))
}

// GetSignatureCount returns signature count in Document
func (pd *PopplerDocument) GetSignatureCount() uint {
	return uint(C.poppler_document_get_n_signatures(pd.document))
}

// GetPermanentId returns the permanent id of the PDF file that which is a permanent identifier of the file that was made at creation.
// returns an empty string in case field not present
func (pd *PopplerDocument) GetPermanentId() string {
	var id *C.char
	defer C.free(unsafe.Pointer(id))

	result := C.poppler_document_get_id(pd.document, &id, nil)

	if result == 0 {
		return ""
	}

	return toGoString(id)
}

// GetSubTypeString returns the PDF subtype version of document as a string.
// returns an empty string in case field is not present
func (pd *PopplerDocument) GetSubTypeString() string {
	return toGoString(C.poppler_document_get_pdf_subtype_string(pd.document))
}

// GetPdfVersion returns the major, minor version of a PDF Document
func (pd *PopplerDocument) GetPdfVersion() (uint, uint) {
	var cmajor C.uint
	var cminor C.uint

	C.poppler_document_get_pdf_version(pd.document, &cmajor, &cminor)

	major := uint(cmajor)
	minor := uint(cminor)

	return major, minor
}

// GetPdfVersionString returns the version as a string
func (pd *PopplerDocument) GetPdfVersionString() string {
	return toGoString(C.poppler_document_get_pdf_version_string(pd.document))
}

// GetSuggestedPrintCount returns the suggested no of copies of the PDF to be printed
// returns 0 in case document does not specify it
func (pd *PopplerDocument) GetSuggestedPrintCount() uint {
	value := C.poppler_document_get_print_n_copies(pd.document)

	if value == 1 {
		return 0
	}

	return uint(value)
}

// GetProducer returns the producer or an empty string in case field not present
func (pd *PopplerDocument) GetProducer() string {
	return toGoString(C.poppler_document_get_producer(pd.document))
}

// GetSubject returns the subject or an empty string in case field not present
func (pd *PopplerDocument) GetSubject() string {
	return toGoString(C.poppler_document_get_subject(pd.document))
}

// GetTitle returns the title or an empty string in case field not present
func (pd *PopplerDocument) GetTitle() string {
	return toGoString(C.poppler_document_get_title(pd.document))
}

// hasAttachments
func (pd *PopplerDocument) hasAttachments() bool {
	return toGoBool(C.poppler_document_has_attachments(pd.document))
}

// hasJavascript
func (pd *PopplerDocument) hasJavascript() bool {
	return toGoBool(C.poppler_document_has_javascript(pd.document))
}

// isLinearized
func (pd *PopplerDocument) isLinearized() bool {
	return toGoBool(C.poppler_document_is_linearized(pd.document))
}

// GetFormField returns the form field corresponding to id
func (pd *PopplerDocument) GetFormField(id int) *PopplerFormField {
	return &PopplerFormField{
		form: C.poppler_document_get_form_field(pd.document, C.gint(id)),
	}
}

// GetPageLayout
func (pd *PopplerDocument) GetPageLayout() PopplerPageLayout {
	return PopplerPageLayout(C.poppler_document_get_page_layout(pd.document))
}

// GetPageMode
func (pd *PopplerDocument) GetPageMode() PopplerPageMode {
	return PopplerPageMode(C.poppler_document_get_page_mode(pd.document))
}

// GetPDFConformance
func (pd *PopplerDocument) GetPDFConformance() PopplerPDFConformance {
	return PopplerPDFConformance(C.poppler_document_get_pdf_conformance(pd.document))
}

// GetPDFPart
func (pd *PopplerDocument) GetPDFPart() PopplerPDFPart {
	return PopplerPDFPart(C.poppler_document_get_pdf_part(pd.document))
}

// GetPDFSubtype
func (pd *PopplerDocument) GetPDFSubtype() PopplerPDFSubtype {
	return PopplerPDFSubtype(C.poppler_document_get_pdf_subtype(pd.document))
}

// GetPermissions
func (pd *PopplerDocument) GetPermissions() PopplerPermissions {
	return PopplerPermissions(C.poppler_document_get_permissions(pd.document))
}

// GetPrintDuplex
func (pd *PopplerDocument) GetPrintDuplex() PopplerPrintDuplex {
	return PopplerPrintDuplex(C.poppler_document_get_print_duplex(pd.document))
}

// GetPrintScaling
func (pd *PopplerDocument) GetPrintScaling() PopplerPrintScaling {
	return PopplerPrintScaling(C.poppler_document_get_print_scaling(pd.document))
}

// GetAttachments returns all the attachments in the document.
// each member of the list must be freed individually with the Free method
func (pd *PopplerDocument) GetAttachments() []*PopplerAttachment {
	count := pd.GetAttachmentCount()
	golist := make([]*PopplerAttachment, 0, count)
	clist := C.poppler_document_get_attachments(pd.document)

	toGoList(clist, func(gp C.gpointer) {
		golist = append(golist, &PopplerAttachment{
			attachment: (popplerAttachment)(gp),
		})
	})

	return golist
}

// GetSigFields returns all the signature fields in a document.
// each member must be freed with the Free method.
func (pd *PopplerDocument) GetSigFields() []*PopplerFormField {
	golist := make([]*PopplerFormField, 0)
	clist := C.poppler_document_get_signature_fields(pd.document)

	toGoList(clist, func(gp C.gpointer) {
		golist = append(golist, &PopplerFormField{
			form: (popplerFormField)(gp),
		})
	})

	return golist
}

// ResetForm exclude indicates if these should be reset or every other form except these
func (pd *PopplerDocument) ResetForm(forms []*PopplerFormField, exclude bool) {
	var clist *C.GList = nil

	for _, i := range forms {
		clist = C.g_list_prepend(clist, C.gpointer(i.form))
	}

	clist = C.g_list_reverse(clist)

	C.poppler_document_reset_form(pd.document, clist, togboolean(exclude))

	C.g_list_free(clist)
}

// setting properties of a pdf

func (pd *PopplerDocument) SetAuthor(author string) {
	cs := C.CString(author)
	defer C.free(unsafe.Pointer(cs))
	C.poppler_document_set_author(pd.document, cs)
}

// ClearAuthor clears author field
func (pd *PopplerDocument) ClearAuthor() {
	C.poppler_document_set_author(pd.document, nil)
}

func (pd *PopplerDocument) SetCreator(creator string) {
	cs := C.CString(creator)
	defer C.free(unsafe.Pointer(cs))
	C.poppler_document_set_creator(pd.document, cs)
}

func (pd *PopplerDocument) ClearCreator() {
	C.poppler_document_set_creator(pd.document, nil)
}

func (pd *PopplerDocument) SetKeywords(keywords string) {
	cs := C.CString(keywords)
	defer C.free(unsafe.Pointer(cs))
	C.poppler_document_set_keywords(pd.document, cs)
}

func (pd *PopplerDocument) ClearKeywords() {
	C.poppler_document_set_keywords(pd.document, nil)
}

func (pd *PopplerDocument) SetProducer(producer string) {
	cs := C.CString(producer)
	defer C.free(unsafe.Pointer(cs))
	C.poppler_document_set_producer(pd.document, cs)
}

func (pd *PopplerDocument) ClearProducer() {
	C.poppler_document_set_producer(pd.document, nil)
}

func (pd *PopplerDocument) SetSubject(subject string) {
	cs := C.CString(subject)
	defer C.free(unsafe.Pointer(cs))
	C.poppler_document_set_subject(pd.document, cs)
}

func (pd *PopplerDocument) ClearSubject() {
	C.poppler_document_set_subject(pd.document, nil)
}

func (pd *PopplerDocument) SetTitle(title string) {
	cs := C.CString(title)
	defer C.free(unsafe.Pointer(cs))
	C.poppler_document_set_title(pd.document, cs)
}

func (pd *PopplerDocument) ClearTitle() {
	C.poppler_document_set_title(pd.document, nil)
}

func (pd *PopplerDocument) SetCreationDate(t time.Time) {
	if t.IsZero() {
		C.poppler_document_set_creation_date(pd.document, C.time_t(-1))
		return
	}

	C.poppler_document_set_creation_date(pd.document, C.time_t(t.Unix()))
}

func (pd *PopplerDocument) SetCreationDateTime(t time.Time) {
	if t.IsZero() {
		C.poppler_document_set_creation_date_time(pd.document, nil)
		return
	}

	gdtime := C.g_date_time_new_from_unix_utc(C.gint64(t.Unix()))

	C.poppler_document_set_creation_date_time(pd.document, gdtime)

	C.g_date_time_unref(gdtime)
}

func (pd *PopplerDocument) SetModDate(t time.Time) {
	if t.IsZero() {
		C.poppler_document_set_modification_date(pd.document, C.time_t(-1))
		return
	}

	C.poppler_document_set_modification_date(pd.document, C.time_t(t.Unix()))
}

func (pd *PopplerDocument) SetModDateTime(t time.Time) {
	if t.IsZero() {
		C.poppler_document_set_modification_date_time(pd.document, nil)
		return
	}

	gdtime := C.g_date_time_new_from_unix_utc(C.gint64(t.Unix()))

	C.poppler_document_set_modification_date_time(pd.document, gdtime)

	C.g_date_time_unref(gdtime)
}

// NewPSFile returns a new PopplerPSFile instance to render the document to.
// path must be absolute.
// first indicates the first page you want to render and pages indicates the number of pages you want to render (with first as reference)
func (pd *PopplerDocument) NewPSFile(path string, first uint, pages uint) *PopplerPSFile {
	cfilename := C.CString(path)
	defer C.free(unsafe.Pointer(cfilename))

	psf := C.poppler_ps_file_new(pd.document, cfilename, C.int(first), C.int(pages))

	return &PopplerPSFile{
		psfile: psf,
	}
}

// Free
func (psf *PopplerPSFile) Free() {
	C.poppler_ps_file_free(psf.psfile)
	psf.psfile = nil
}

func (psf *PopplerPSFile) SetDuplex(duplex bool) {
	C.poppler_ps_file_set_duplex(psf.psfile, togboolean(duplex))
}

func (psf *PopplerPSFile) SetPaperSize(height float64, width float64) {
	C.poppler_ps_file_set_paper_size(psf.psfile, C.double(width), C.double(height))
}

// Font stuff

func (pd *PopplerDocument) NewFontInfo() *PopplerFontInfo {
	return &PopplerFontInfo{
		info: C.poppler_font_info_new(pd.document),
	}
}

func (pfi *PopplerFontInfo) Free() {
	C.poppler_font_info_free(pfi.info)
	pfi.info = nil
}

// Scan returns a PopplerFontsIter which can be used to iterate over the different fonts in a document.
// pages indicates the number of pages to scan
func (pfi *PopplerFontInfo) Scan(pages uint) *PopplerFontsIter {
	var iter *C.struct__PopplerFontsIter
	res := C.poppler_font_info_scan(pfi.info, C.int(pages), &iter)

	if res == C.int(0) {
		return nil
	}

	return &PopplerFontsIter{
		fontsIter: iter,
	}
}

func (pfi *PopplerFontsIter) Copy() *PopplerFontsIter {
	return &PopplerFontsIter{
		fontsIter: C.poppler_fonts_iter_copy(pfi.fontsIter),
	}
}

func (pfi *PopplerFontsIter) Free() {
	C.poppler_fonts_iter_free(pfi.fontsIter)
	pfi.fontsIter = nil
}

func (pfi *PopplerFontsIter) GetEncoding() string {
	return toGoString(C.poppler_fonts_iter_get_encoding(pfi.fontsIter))
}

func (pfi *PopplerFontsIter) GetFilename() string {
	return toGoString(C.poppler_fonts_iter_get_file_name(pfi.fontsIter))
}

func (pfi *PopplerFontsIter) GetFontType() PopplerFontType {
	return PopplerFontType(C.poppler_fonts_iter_get_font_type(pfi.fontsIter))
}

func (pfi *PopplerFontsIter) GetFullname() string {
	return toGoString(C.poppler_fonts_iter_get_full_name(pfi.fontsIter))
}

func (pfi *PopplerFontsIter) GetName() string {
	return toGoString(C.poppler_fonts_iter_get_name(pfi.fontsIter))
}

func (pfi *PopplerFontsIter) GetSubstituteName() string {
	return toGoString(C.poppler_fonts_iter_get_substitute_name(pfi.fontsIter))
}

func (pfi *PopplerFontsIter) IsEmbedded() bool {
	return toGoBool(C.poppler_fonts_iter_is_embedded(pfi.fontsIter))
}

func (pfi *PopplerFontsIter) IsSubset() bool {
	return toGoBool(C.poppler_fonts_iter_is_subset(pfi.fontsIter))
}

// Next sets the iter to point to the next fontsIter.
// returns true if this was successful
func (pfi *PopplerFontsIter) Next() bool {
	return toGoBool(C.poppler_fonts_iter_next(pfi.fontsIter))
}

// Index stuff

// NewIndexIter returns a new IndexIter which can iterated over to transverse the entire index of a document
func (pd *PopplerDocument) NewIndexIter() *PopplerIndexIter {
	return &PopplerIndexIter{
		indexIter: C.poppler_index_iter_new(pd.document),
	}
}

// Next
func (pii *PopplerIndexIter) Next() bool {
	return toGoBool(C.poppler_index_iter_next(pii.indexIter))
}

func (pii *PopplerIndexIter) Copy() *PopplerIndexIter {
	return &PopplerIndexIter{
		indexIter: C.poppler_index_iter_copy(pii.indexIter),
	}
}

func (pii *PopplerIndexIter) Free() {
	C.poppler_index_iter_free(pii.indexIter)
	pii.indexIter = nil
}

// GetAction returns the action corresponding to the IndexIter
func (pii *PopplerIndexIter) GetAction() *PopplerAction {
	return &PopplerAction{
		action: C.poppler_index_iter_get_action(pii.indexIter),
	}
}

// GetChild returns the child (sub index) of a particular node of a IndexIter.
// returns null if there is no child.
func (pii *PopplerIndexIter) GetChild() *PopplerIndexIter {
	child := C.poppler_index_iter_get_child(pii.indexIter)

	if child == nil {
		return nil
	}

	return &PopplerIndexIter{
		indexIter: child,
	}
}

// IsOpen If the nodes of an Index should be opened by default to an user.
func (pii *PopplerIndexIter) IsOpen() bool {
	return toGoBool(C.poppler_index_iter_is_open(pii.indexIter))
}

// layers stuff

// NewLayersIter returns a new LayersIter to transverse the layers of a document
func (pd *PopplerDocument) NewLayersIter() *PopplerLayersIter {
	layer := C.poppler_layers_iter_new(pd.document)

	if layer == nil {
		return nil
	}

	return &PopplerLayersIter{
		layersIter: C.poppler_layers_iter_new(pd.document),
	}
}

func (pli *PopplerLayersIter) Next() bool {
	return toGoBool(C.poppler_layers_iter_next(pli.layersIter))
}

func (pli *PopplerLayersIter) Copy() *PopplerLayersIter {
	return &PopplerLayersIter{
		layersIter: C.poppler_layers_iter_copy(pli.layersIter),
	}
}

func (pli *PopplerLayersIter) Free() {
	C.poppler_layers_iter_free(pli.layersIter)
	pli.layersIter = nil
}

// GetChild returns the child of a particular node of a LayersIter.
// returns null if there is no child.
func (pli *PopplerLayersIter) GetChild() *PopplerLayersIter {
	child := C.poppler_layers_iter_get_child(pli.layersIter)

	if child == nil {
		return nil
	}

	return &PopplerLayersIter{
		layersIter: child,
	}
}

// GetLayer returns the layer corresponding to a LayersIter
func (pli *PopplerLayersIter) GetLayer() *PopplerLayer {
	layer := C.poppler_layers_iter_get_layer(pli.layersIter)

	if layer == nil {
		return nil
	}

	return &PopplerLayer{
		layer: layer,
	}
}

func (pli *PopplerLayersIter) GetTitle() string {
	return toGoString(C.poppler_layers_iter_get_title(pli.layersIter))
}
