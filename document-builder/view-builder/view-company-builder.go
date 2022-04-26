package viewbuilder

import (
	. "JSONDocumentGenerator/document"
)

type CompanyBuilder interface {
	injectViewInstance(*viewBuilder)
	CreateRecords(Company_Records) CompanyBuilder
	CreateColumns(Company_Columns) CompanyBuilder
	BuildCompany() ViewBuilder
}

type companyBuilder struct {
	columns             Company_Columns
	records             Company_Records
	viewBuilderInstance *viewBuilder
}

func (cb *companyBuilder) injectViewInstance(viewBuilder *viewBuilder) {
	cb.viewBuilderInstance = viewBuilder
}

func (cb *companyBuilder) CreateColumns(companyColunms Company_Columns) CompanyBuilder {
	cb.columns = companyColunms
	return cb
}

func (cb *companyBuilder) CreateRecords(companyRecords Company_Records) CompanyBuilder {
	cb.records = companyRecords
	return cb
}

func (cb *companyBuilder) BuildCompany() ViewBuilder {
	cb.viewBuilderInstance.Company = View_Company{
		Columns: cb.columns,
		Records: cb.records,
	}
	return cb.viewBuilderInstance
}
