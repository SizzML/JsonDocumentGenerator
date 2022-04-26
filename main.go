package main

import (
	. "JSONDocumentGenerator/document"
	documentbuilder "JSONDocumentGenerator/document-builder"
	viewBuilder "JSONDocumentGenerator/document-builder/view-builder"
	jsonfilegenerator "JSONDocumentGenerator/json-file-generator"

	"github.com/dranikpg/dto-mapper"
)

func main() {
	var _document DocumentDto
	instancedDoc := DocumentDto{
		Document: Document{
			Metadata: Metadata{
				Rows: 12,
			},
			Views: []View{
				{},
			},
		},
	}
	dto.Map(&_document, instancedDoc)
	jsonfilegenerator.GenerateFileByStruct("test", _document)

	otherDOc := documentbuilder.NewDocument()
	view := viewBuilder.NewViewsBuilder()
	view.
		CreateCompany().
		CreateColumns(Company_Columns{BusinessName: "Busines perrito, funciona zorro"}).
		CreateRecords(Company_Records{BusinessName: []string{"Si sirve :33"}}).
		BuildCompany().
		CreateDateRange().
		CreateColumns(DateRange_Columns{From: "Pruebaaaa xd"}).
		CreateRecords(DateRange_Records{From: "Ora si papa"}).
		BuildDateRange()

	// view := View{
	// 	Balance: View_Balance{
	// 		Columns: Balance_Columns{
	// 			InitialBalance: "initial b",
	// 			Deposits:       "deposits",
	// 		},
	// 		Records: Balance_Records{
	// 			InitialBalance: "records initial balance",
	// 		},
	// 	},
	// }
	// view2 := View{
	// 	Balance: View_Balance{
	// 		Columns: Balance_Columns{
	// 			InitialBalance: "initial b2",
	// 			Deposits:       "deposits2",
	// 		},
	// 		Records: Balance_Records{
	// 			InitialBalance: "records initial balance2",
	// 		},
	// 	},
	// }
	otherDOc.AppendView(view.BuildView())
	// otherDOc.AppendView(view2)

	jsonfilegenerator.GenerateFileByStruct("test2", otherDOc.Build())

}
