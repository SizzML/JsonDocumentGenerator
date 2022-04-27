package main

import (
	. "JSONDocumentGenerator/document"
	documentbuilder "JSONDocumentGenerator/document-builder"
	viewBuilder "JSONDocumentGenerator/document-builder/view-builder"
	jsonfilegenerator "JSONDocumentGenerator/json-file-generator"
)

func main() {
	document := documentbuilder.NewDocument()

	// Iterate over the quantity of views that you need to create
	// in this case will be created 2 views in the document
	const NumberOfViews = 2

	for i := 0; i < NumberOfViews; i++ {
		view := viewBuilder.NewView()

		view.
			CreateCompany().
			CreateColumns(Company_Columns{
				BusinessName: "DENOMINACIÓN SOCIAL",
				Address:      "DOMICILIO",
				Name:         "NOMBRE",
			}).
			CreateRecords(Company_Records{
				BusinessName: []string{
					"Mercado Pago S.A.",
					"Compañía de Financiamiento",
					"NIT: 901449740-9",
				},
				Address: []string{
					"Carrera 17 Nº 93-09 p.3",
					"Bogotá D.C., Colombia",
					"Teléfono comercial: 6017053050",
				},
				Name: []string{
					"ADIDAS COLOMBIA LIMITADA",
				},
			}).
			BuildCompany()

		view.
			CreateDateRange().
			CreateColumns(DateRange_Columns{
				From:           "DESDE",
				To:             "HASTA",
				GenerationDate: "FECHA DE GENERACIÓN",
				IdUser:         "ID DE USUARIO",
			}).
			CreateRecords(DateRange_Records{
				From:           "01-03-2021 00:00:00",
				To:             "20-04-2021 00:00:00",
				GenerationDate: "11-11-2021 05:03:05",
				IdUser:         "123456789",
			}).
			BuildDateRange()

		view.
			CreateBalance().
			CreateColumns(Balance_Columns{
				InitialBalance: "BALANCE INICIAL",
				Deposits:       "DEPÓSITOS",
				Debits:         "DÉBITOS",
				Commisions:     "COMISIONES",
				Retencion:      "RETENCIONES POR IMPUESTOS",
				AverageBalance: "PROMEDIO DIARIO DE SALDOS",
				Balance:        "BALANCE FINAL",
			}).
			CreateRecords(Balance_Records{
				InitialBalance: "976210.28",
				Deposits:       "356802198.00",
				Debits:         "340201206.00",
				Commisions:     "14918119.40",
				Retencion:      "0.00",
				AverageBalance: "2659082.88",
				Balance:        "2659082.88",
			}).
			BuildBalance()

		view.
			CreateDetails().
			CreateColumns(Details_Columns{
				AccreditationDate: "FECHA DE ACREDITACIÓN",
				MovementType:      "TIPO DE MOVIMIENTO",
				TransactionType:   "TIPO DE TRANSACCIÓN",
				TransactionId:     "ID DE TRANSACCIÓN",
				Currency:          "MONEDA",
				TransactionAmount: "MONTO DE TRANSACCIÓN",
				AccountCommission: "COMISIÓN DE CUENTA",
				BusinessName:      "NOMBRE DEL COMERCIO",
			}).
			CreateRecords(Details_Records{
				[]string{
					"01-03-2021 00:31:00",
					"ABONO",
					"Pago",
					"13838829463#0",
					"COP (pesos colombianos)",
					"56314.27",
					"-2659.73",
					"Nombre del comercio",
				},
				[]string{
					"01-03-2021 00:31:00",
					"ABONO",
					"Pago",
					"13838829463#1",
					"COP (pesos colombianos)",
					"56314.27",
					"-2659.73",
					"Nombre del comercio",
				},
				[]string{
					"01-03-2021 00:31:00",
					"ABONO",
					"Pago",
					"13838829463#2",
					"COP (pesos colombianos)",
					"56314.27",
					"-2659.73",
					"Nombre del comercio",
				},
				[]string{
					"01-03-2021 00:31:00",
					"ABONO",
					"Pago",
					"13838829463#3",
					"COP (pesos colombianos)",
					"56314.27",
					"-2659.73",
					"Nombre del comercio",
				},
				[]string{
					"01-03-2021 00:31:00",
					"ABONO",
					"Pago",
					"13838829463#4",
					"COP (pesos colombianos)",
					"56314.27",
					"-2659.73",
					"Nombre del comercio",
				},
			}).
			BuildDetails()

		document.AppendView(view.BuildView())
	}

	document.SetMetadata(Metadata{
		Rows: 10,
		RangeDate: Metadata_RangeDate{
			Begin: "01-03-2021 00:00:00",
			End:   "20-04-2021 00:00:00",
		},
		User: "usertest",
	})

	jsonfilegenerator.GenerateFileByStruct("output", document.Build())
}
