package gsheetsreader


import (
	"fmt"
	"io/ioutil"
	"log"
)


type Producto struct {
    NSO string `json:"RegSan/NSO"`
    NombreProducto string `json:"NombreProducto"`
    MarcaProducto string `json:"MarcaProducto"`
    Titular string `json:"Titular"`
    FechaEmision string `json:"FechaEmision"`
    FechaVigencia string `json:"FechaVigencia"`
}


func GetData(spreadsheetId string) map[string]Producto {

	// Use google credentials to access Sheets api
	// -------------------------------------------

	creds, err := ioutil.ReadFile("keys/verifynso-dd9a44c0a761.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	srv := getService(creds)


	// Get data from spreadsheet
	// -------------------------

	// Initialize list
	productos := make(map[string]Producto)

	// Read spreadsheet
	readRange := "Sheet1"
	sheetdata, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Parse data to data producto object
	if len(sheetdata.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		
		for _, row := range sheetdata.Values {

			// Parse values to strings
			record := make([]string, 6)

			for j := 0; j < len(row); j++ {
				record[j] = row[j].(string)

			}

			// save strings to map
			productos[record[0]] = Producto{
				NSO: record[0],
				NombreProducto: record[1],
				MarcaProducto: record[2],
				Titular: record[3],
				FechaEmision: record[4],
				FechaVigencia: record[5],
			} 

			// use only first rows as example
			// if i==20 {
			// 	break
			// }
		}
		//fmt.Print(productos)
	}
	return productos
}
