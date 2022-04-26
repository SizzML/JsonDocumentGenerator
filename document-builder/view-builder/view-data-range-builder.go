package viewbuilder

import (
	. "JSONDocumentGenerator/document"
)

type DateRangeBuilder interface {
	injectViewInstance(*viewBuilder)
	CreateColumns(DateRange_Columns) DateRangeBuilder
	CreateRecords(DateRange_Records) DateRangeBuilder
	BuildDateRange() ViewBuilder
}

type dateRangeBuilder struct {
	viewBuilderInstance *viewBuilder
	columns             DateRange_Columns
	records             DateRange_Records
}

func (drb *dateRangeBuilder) injectViewInstance(viewBuilder *viewBuilder) {
	drb.viewBuilderInstance = viewBuilder
}

func (drb *dateRangeBuilder) CreateColumns(dateRangeColumns DateRange_Columns) DateRangeBuilder {
	drb.columns = dateRangeColumns
	return drb
}

func (drb *dateRangeBuilder) CreateRecords(dateRangeRecords DateRange_Records) DateRangeBuilder {
	drb.records = dateRangeRecords
	return drb
}

func (drb *dateRangeBuilder) BuildDateRange() ViewBuilder {
	drb.viewBuilderInstance.DateRange = View_DateRange{
		Columns: drb.columns,
		Records: drb.records,
	}
	return drb.viewBuilderInstance
}
