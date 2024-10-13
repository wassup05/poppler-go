package poppler

/*
#cgo pkg-config: poppler-glib
#include <stdlib.h>
#include <poppler.h>
#include <glib.h>
*/
import "C"
import "time"

type PopplerFormField struct {
	form popplerFormField
}

type popplerFormField *C.struct__PopplerFormField

type PopplerCertInfo struct {
	info popplerCertInfo
}

type popplerCertInfo *C.struct__PopplerCertificateInfo

type PopplerSigInfo struct {
	info popplerSigInfo
}

type popplerSigInfo *C.struct__PopplerSigInfo

func (pci *PopplerCertInfo) GetExpTime() time.Time {
	gdtime := C.poppler_certificate_info_get_expiration_time(pci.info)

	if gdtime == nil {
		return time.Time{}
	}
	ctime := C.g_date_time_to_unix(gdtime)

	return time.Unix(int64(ctime), 0)
}

func (pci *PopplerCertInfo) GetIssuanceTime() time.Time {
	gdtime := C.poppler_certificate_info_get_issuance_time(pci.info)

	if gdtime == nil {
		return time.Time{}
	}
	ctime := C.g_date_time_to_unix(gdtime)

	return time.Unix(int64(ctime), 0)
}

func (pci *PopplerCertInfo) GetIssuerCommonName() string {
	return toGoString(C.poppler_certificate_info_get_issuer_common_name(pci.info))
}

func (pci *PopplerCertInfo) GetIssuerEmail() string {
	return toGoString(C.poppler_certificate_info_get_issuer_email(pci.info))
}

func (pci *PopplerCertInfo) GetIssuerOrg() string {
	return toGoString(C.poppler_certificate_info_get_issuer_organization(pci.info))
}

func (pci *PopplerCertInfo) GetSubCommonName() string {
	return toGoString(C.poppler_certificate_info_get_subject_common_name(pci.info))
}

func (pci *PopplerCertInfo) GetSubEmail() string {
	return toGoString(C.poppler_certificate_info_get_subject_email(pci.info))
}

func (pci *PopplerCertInfo) GetSubOrg() string {
	return toGoString(C.poppler_certificate_info_get_subject_organization(pci.info))
}

// FormField stuff

func (pff *PopplerFormField) Free() {
	C.g_object_unref(C.gpointer(pff.form))
	pff.form = nil
}

func (pff *PopplerFormField) GetBtnType() PopplerFormButtonType {
	return PopplerFormButtonType(C.poppler_form_field_button_get_button_type(pff.form))
}

func (pff *PopplerFormField) GetChoiceType() PopplerFormChoiceType {
	return PopplerFormChoiceType(C.poppler_form_field_choice_get_choice_type(pff.form))
}

func (pff *PopplerFormField) GetChoiceItem(id int) string {
	return toGoString(C.poppler_form_field_choice_get_item(pff.form, C.int(id)))
}

func (pff *PopplerFormField) GetChoiceItemCount() int {
	return int(C.poppler_form_field_choice_get_n_items(pff.form))
}

func (pff *PopplerFormField) GetChoiceText() string {
	return toGoString(C.poppler_form_field_choice_get_text(pff.form))
}

func (pff *PopplerFormField) IsChoiceEditable() bool {
	return toGoBool(C.poppler_form_field_choice_is_editable(pff.form))
}

func (pff *PopplerFormField) GetBtnState() bool {
	return toGoBool(C.poppler_form_field_button_get_state(pff.form))
}

func (pff *PopplerFormField) SetBtnState(state bool) {
	C.poppler_form_field_button_set_state(pff.form, togboolean(state))
}

func (pff *PopplerFormField) CanSelectMultiple() bool {
	return toGoBool(C.poppler_form_field_choice_can_select_multiple(pff.form))
}

func (pff *PopplerFormField) CommitOnChange() bool {
	return toGoBool(C.poppler_form_field_choice_commit_on_change(pff.form))
}

func (pff *PopplerFormField) DoSpellCheck() bool {
	return toGoBool(C.poppler_form_field_choice_do_spell_check(pff.form))
}
