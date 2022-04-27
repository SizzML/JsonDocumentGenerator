package documentbuilder

import (
	. "JSONDocumentGenerator/document"
)

type DocumentBuilder interface {
	AppendView(View) DocumentBuilder
	SetMetadata(Metadata) DocumentBuilder
	Build() *DocumentDto
}

type viewBuilder struct {
}

type documentBuilder struct {
	Views    []View
	Metadata Metadata
}

func (dB *documentBuilder) Build() *DocumentDto {
	return &DocumentDto{
		Document: Document{
			Views:    dB.Views,
			Metadata: dB.Metadata,
		},
	}
}

func (dB *documentBuilder) AppendView(view View) DocumentBuilder {
	dB.Views = append(dB.Views, view)
	return dB
}
func (dB *documentBuilder) SetMetadata(metadata Metadata) DocumentBuilder {
	dB.Metadata = metadata
	return dB
}

func NewDocument() DocumentBuilder {
	return &documentBuilder{}
}
