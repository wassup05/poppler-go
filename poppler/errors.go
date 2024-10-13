package poppler

/*
#include <glib.h>
*/
import "C"

type PopplerErrorDomain uint8

const (
	PopplerErrorInvalid PopplerErrorDomain = iota
	PopplerErrorEncrypted
	PopplerErrorOpenFile
	PopplerErrorBadCatalog
	PopplerErrorDamaged
	PopplerErrorSigning
	PopplerGLibFileError
)

// PopplerError a struct that holds a Domain, Code, Message
// Message is a human readable string returned by poppler
type PopplerError struct {
	Domain  PopplerErrorDomain
	Code    int
	Message string
}

func gErrorToPopplerError(gerr C.GError) *PopplerError {
	var err PopplerError

	err.Domain = PopplerErrorDomain(gerr.domain)

	err.Code = int(gerr.code)
	err.Message = toGoString(gerr.message)

	return &err
}
