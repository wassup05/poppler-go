package poppler

// custom struct for knowing which free_func to call in GoList and GoArray

type PopplerType uint8

const (
	gobject PopplerType = iota
	rectangle
)

// Defined in PopplerDocument

type PopplerFontType uint8

const (
	PopplerFontTypeUnknown PopplerFontType = iota
	PopplerFontTypeType1
	PopplerFontTypeType1C
	PopplerFontTypeType1COT
	PopplerFontTypeType3
	PopplerFontTypeTrueType
	PopplerFontTypeTrueTypeOT
	PopplerFontTypeTypeO
	PopplerFontTypeTypeOT
	PopplerFontTypeType0
	PopplerFontTypeType0C
	PopplerFontTypeType0COT
	PopplerFontTypeType2
	PopplerFontTypeType2OT
)

type PopplerPDFConformance uint8

const (
	PopplerPDFSubtypeConfUnset PopplerPDFConformance = iota
	PopplerPDFSubtypeConfA
	PopplerPDFSubtypeConfB
	PopplerPDFSubtypeConfG
	PopplerPDFSubtypeConfN
	PopplerPDFSubtypeConfP
	PopplerPDFSubtypeConfPG
	PopplerPDFSubtypeConfU
	PopplerPDFSubtypeConfNone
)

type PopplerPDFPart uint8

const (
	PopplerPDfSubtypePartUnset PopplerPDFPart = iota
	PopplerPDfSubtypePart1
	PopplerPDfSubtypePart2
	PopplerPDfSubtypePart3
	PopplerPDfSubtypePart4
	PopplerPDfSubtypePart5
	PopplerPDfSubtypePart6
	PopplerPDfSubtypePart7
	PopplerPDfSubtypePart8
	PopplerPDfSubtypePartNone
)

type PopplerPDFSubtype uint8

const (
	PopplerPDFSubtypeUnset PopplerPDFSubtype = iota
	PopplerPDFSubtypePDFA
	PopplerPDFSubtypePDFE
	PopplerPDFSubtypePDFUA
	PopplerPDFSubtypePDFVT
	PopplerPDFSubtypePDFX
	PopplerPDFSubtypePDFNone
)

type PopplerPageLayout uint8

const (
	PopplerPageLayoutUnset PopplerPageLayout = iota
	PopplerPageLayoutSinglePage
	PopplerPageLayoutOneColumn
	PopplerPageLayoutTwoColumnLeft
	PopplerPageLayoutTwoColumnRight
	PopplerPageLayoutTwoPageLeft
	PopplerPageLayouttwoPageRight
)

type PopplerPageMode uint8

const (
	PopplerPageModeUnset PopplerPageMode = iota
	PopplerPageModeNone
	PopplerPageModeUseOutlines
	PopplerPageModeUseThumbs
	PopplerPageModeFullScreen
	PopplerPageModeUseOC
	PopplerPageModeUseAttachments
)

type PopplerPermissions uint8

const (
	PopplerPermissionsCanPrint PopplerPermissions = iota
	PopplerPermissionsCanModify
	PopplerPermissionsCanCopy
	PopplerPermissionsCanAddNotes
	PopplerPermissionsCanFillForm
	PopplerPermissionsCanExtractContents
	PopplerPermissionsCanAssemble
	PopplerPermissionsCanPrintInHighRes
	PopplerPermissionsFull
)

type PopplerPrintDuplex uint8

const (
	PopplerPrintDuplexNone PopplerPrintDuplex = iota
	PopplerPrintDuplexSimplex
	PopplerPrintDuplexDuplexShortEdge
	PopplerPrintDuplexDuplexLongEdge
)

type PopplerPrintScaling uint8

const (
	PopplerPrintScalingAppDefault PopplerPrintScaling = iota
	PopplerPrintScalingNone
)

type PopplerViewerPreferences uint8

const (
	PopplerViewerPreferencesUnset PopplerViewerPreferences = iota
	PopplerViewerPreferencesHideToolbar
	PopplerViewerPreferencesHideMenubar
	PopplerViewerPreferencesHideWindowUI
	PopplerViewerPreferencesFitWindow
	PopplerViewerPreferencesCenterWindow
	PopplerViewerPreferencesDisplayDocTitle
	PopplerViewerPreferencesDirectionRTL
)

type PopplerPrintFlags uint8

const (
	PopplerPrintDocument PopplerPrintFlags = iota
	PopplerPrintMarkupAnnots
	PopplerPrintStampAnnotsOnly
	PopplerPrintAll
)

// Defined in PopplerPage

type PopplerPageTransitionType uint8

const (
	PopplerPageTransitionReplace PopplerPageTransitionType = iota
	PopplerPageTransitionSplit
	PopplerPageTransitionBlinds
	PopplerPageTransitionBox
	PopplerPageTransitionWipe
	PopplerPageTransitionDissolve
	PopplerPageTransitionGlitter
	PopplerPageTransitionFly
	PopplerPageTransitionPush
	PopplerPageTransitionCover
	PopplerPageTransitionUncover
	PopplerPageTransitionFade
)

type PopplerPageTransitionAlignment uint8

const (
	PopplerPageTransitionVertical PopplerPageTransitionAlignment = iota
	PopplerPageTransitionHorizontal
)

type PopplerPageTransitionDirection uint8

const (
	PopplerPageTransitionInward PopplerPageTransitionDirection = iota
	PopplerPageTransitionOutward
)

type PopplerFindFlags uint8

const (
	PopplerFindDefault PopplerFindFlags = iota
	PopplerFindCaseSensitive
	PopplerFindBackwards
	PopplerFindWholeWordsOnly
	PopplerFindIgnoreDiacritics
	PopplerFindMultiline
)

type PopplerMoviePlayMode uint8

const (
	PopplerMoviePlayModeOnce PopplerMoviePlayMode = iota
	PopplerMoviePlayModeOpen
	PopplerMoviePlayModeRepeat
	PopplerMoviePlayModePalindrome
	PopplerMoviePlayModeUnknown
)

type PopplerBackend uint8

const (
	PopplerBackendUnkown PopplerBackend = iota
	PopplerBackendSplash
	PopplerBackendCairo
)

func (pb PopplerBackend) String() string {
	switch pb {
	case PopplerBackendUnkown:
		return "Unknown"
	case PopplerBackendSplash:
		return "Splash"
	case PopplerBackendCairo:
		return "Cairo"
	default:
		return "Not sure"
	}
}

type PopplerActionType uint8

const (
	PopplerActionUnkown PopplerActionType = iota
	PopplerActionNone
	PopplerActionGoToDest
	PopplerActionGoToRemote
	PopplerActionLaunch
	PopplerActionURI
	PopplerActionNamed
	PopplerActionMovie
	PopplerActionRendition
	PopplerActionOCGState
	PopplerActionJavaScript
	PopplerActionResetForm
)

type PopplerDestType uint8

const (
	PopplerDestUnkown PopplerDestType = iota
	PopplerDestXYZ
	PopplerDestFit
	PopplerDestFitH
	PopplerDestFitV
	PopplerDestFitR
	PopplerDestFitB
	PopplerDestFitBH
	PopplerDestFitBV
	PopplerDestNamed
)

type PopplerActionMovieOperation uint8

const (
	PopplerActionMoviePlay PopplerActionMovieOperation = iota
	PopplerActionMoviePause
	PopplerActionMovieResume
	PopplerActionMovieStop
)

type PopplerActionLayerAction uint8

const (
	PopplerActionLayerOn PopplerActionLayerAction = iota
	PopplerActionLayerOff
	PopplerActionLayerToggle
)

type PopplerSelectionStyle uint8

const (
	PopplerSelectionGlyph PopplerSelectionStyle = iota
	PopplerSelectionWord
	PopplerSelectionLine
)

// Defined in PopplerStructure

type PopplerStructureBlockAlign uint8

const (
	PopplerStructureBlockAlignBefore PopplerStructureBlockAlign = iota
	PopplerStructureBlockAlignMiddle
	PopplerStructureBlockAlignAfter
	PopplerStructureBlockAlignJustify
)

// Defined in PopplerFormField

type PopplerAdditionalActionType uint8

const (
	PopplerAdditionalActionFieldModified PopplerAdditionalActionType = iota
	PopplerAdditionalActionFormatField
	PopplerAdditionalActionValidateField
	PopplerAdditionalActionCalculateField
)

type PopplerCertStatus uint8

const (
	PopplerCertTrusted PopplerCertStatus = iota
	PopplerCertUntrustedIssuer
	PopplerCertUnkwonIssuer
	PopplerCertRevoked
	PopplerCertExpired
	PopplerCertGenericError
	PopplerCertNotVerified
)

type PopplerFormFieldType uint8

const (
	PopplerFormFieldUnkwon PopplerFormFieldType = iota
	PopplerFormFieldButton
	PopplerFormFieldText
	PopplerFormFieldChoice
	PopplerFormFieldSignature
)

type PopplerFormButtonType uint8

const (
	PopplerformButtonPush PopplerFormButtonType = iota
	PopplerformButtonCheck
	PopplerformButtonRadio
)

type PopplerFormChoiceType uint8

const (
	PopplerFormChoiceCombo PopplerFormChoiceType = iota
	PopplerFormChoiceList
)

type PopplerFormTextType uint8

const (
	PopplerFormTextNormal PopplerFormTextType = iota
	PopplerFormTextMultiline
	PopplerFormTextFileSelect
)

type PopplerSigStatus uint8

const (
	PopplerSigValid PopplerSigStatus = iota
	PopplerSigInvalid
	PopplerSigDigestMismatch
	PopplerSigDecodingErr
	PopplerSigGenericErr
	PopplerSigNotFound
	PopplerSigNotVerified
)

type PopplerSigValidationFlags uint8

const (
	PopplerSigValidationFlagValidateCert PopplerSigValidationFlags = iota
	PopplerSigValidationFlagWithoutOSCPRevocationCheck
	PopplerSigValidationFlagUseAIACertFetch
)
