package viewbuilder

import (
	. "JSONDocumentGenerator/document"
)

type BalanceBuilder interface {
	injectViewInstance(*viewBuilder)
	CreateRecords(Balance_Records) BalanceBuilder
	CreateColumns(Balance_Columns) BalanceBuilder
	BuildBalance() ViewBuilder
}

type balanceBuilder struct {
	viewBuilderInstance *viewBuilder
	columns             Balance_Columns
	records             Balance_Records
}

func (bb *balanceBuilder) injectViewInstance(viewBuilder *viewBuilder) {
	bb.viewBuilderInstance = viewBuilder
}

func (bb *balanceBuilder) CreateRecords(records Balance_Records) BalanceBuilder {
	bb.records = records
	return bb
}

func (bb *balanceBuilder) CreateColumns(columns Balance_Columns) BalanceBuilder {
	bb.columns = columns
	return bb
}

func (bb *balanceBuilder) BuildBalance() ViewBuilder {
	bb.viewBuilderInstance.Balance = View_Balance{
		Columns: bb.columns,
		Records: bb.records,
	}
	return bb.viewBuilderInstance
}
