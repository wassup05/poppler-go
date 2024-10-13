package poppler

import "os"

func NewDocumentFromOSFile(file *os.File) (*PopplerDocument, *PopplerError) {
	return NewDocFromFilename(file.Name())
}

type DocumentInfo struct {
	Title, Subject, Author, Creator, Producer, Keywords, Metadata, Version string
	PageCount, AttachmentCount, MajorVersion, MinorVersion                 uint
	IsLinearized, HasJS, HasAttachments                                    bool
}

func (pd *PopplerDocument) GetInfo() *DocumentInfo {
	major, minor := pd.GetPdfVersion()
	return &DocumentInfo{
		Title:           pd.GetTitle(),
		Subject:         pd.GetSubject(),
		Author:          pd.GetAuthor(),
		Creator:         pd.GetCreator(),
		Producer:        pd.GetProducer(),
		Keywords:        pd.GetKeywords(),
		Metadata:        pd.GetMetadata(),
		Version:         pd.GetPdfVersionString(),
		PageCount:       pd.GetPageCount(),
		AttachmentCount: pd.GetAttachmentCount(),
		MajorVersion:    major,
		MinorVersion:    minor,
		IsLinearized:    pd.isLinearized(),
		HasJS:           pd.hasJavascript(),
		HasAttachments:  pd.hasAttachments(),
	}
}

type PageInfo struct {
	Duration, Height, Width float64
	Text, Label             string
	Index                   int
}

func (pp *PopplerPage) GetInfo() *PageInfo {
	h, w := pp.GetSize()
	return &PageInfo{
		Duration: pp.GetDuration(),
		Height:   h,
		Width:    w,
		Index:    pp.GetIndex(),
		Text:     pp.GetText(),
	}
}
