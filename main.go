package main

import (
	. "JSONDocumentGenerator/document"
	documentbuilder "JSONDocumentGenerator/document-builder"
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
	view := View{
		Balance: Balance{
			Columns: BalanceType{
				InitialBalance: "initial b",
				Deposits:       "deposits",
			},
			Records: BalanceType{
				InitialBalance: "records initial balance",
			},
		},
	}
	view2 := View{
		Balance: Balance{
			Columns: BalanceType{
				InitialBalance: "initial b2",
				Deposits:       "deposits2",
			},
			Records: BalanceType{
				InitialBalance: "records initial balance2",
			},
		},
	}
	otherDOc.AppendView(view)
	otherDOc.AppendView(view2)

	jsonfilegenerator.GenerateFileByStruct("test2", otherDOc.Build())
}
