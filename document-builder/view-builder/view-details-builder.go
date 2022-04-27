package viewbuilder

import (
	. "JSONDocumentGenerator/document"
)

type DetailsBuilder interface {
	injectViewInstance(*viewBuilder)
	CreateRecords(Details_Records) DetailsBuilder
	CreateColumns(Details_Columns) DetailsBuilder
	BuildDetails() ViewBuilder
}

type detailsBuilder struct {
	viewBuilderInstance *viewBuilder
	columns             Details_Columns
	records             Details_Records
}

func (db *detailsBuilder) injectViewInstance(viewBuilder *viewBuilder) {
	db.viewBuilderInstance = viewBuilder
}

func (db *detailsBuilder) CreateColumns(detailsColunms Details_Columns) DetailsBuilder {
	db.columns = detailsColunms
	return db
}

func (db *detailsBuilder) CreateRecords(detailsRecords Details_Records) DetailsBuilder {
	db.records = detailsRecords
	return db
}

func (db *detailsBuilder) BuildDetails() ViewBuilder {
	db.viewBuilderInstance.Details = View_Details{
		Columns: db.columns,
		Records: db.records,
	}
	return db.viewBuilderInstance
}
