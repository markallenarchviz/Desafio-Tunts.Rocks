package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Data struct {
	Matricula string
	Aluno     string
	Faltas    string
	P1        string
	P2        string
	P3        string
}

func main() {
	// API key for accessing the Google Sheets API
	apiKey := "AIzaSyAfjKrcpc9uOXflJOJLMyc5FfH2np7Ards"

	// ID of the Google Sheet you want to access
	spreadsheetID := "1jJCYe__gprZznPm-iecd4x5p0W9oXxXyhlHOqhF7IAY"

	// Connect to Google Sheets API with API key
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	// Get values from the spreadsheet
	readRange := "A1:F27"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var dataArray []Data

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			if len(row) >= 6 {
				data := Data{
					Matricula: row[0].(string),
					Aluno:     row[1].(string),
					Faltas:    row[2].(string),
					P1:        row[3].(string),
					P2:        row[4].(string),
					P3:        row[5].(string),
				}
				dataArray = append(dataArray, data)
			}
		}
	}

	fmt.Println(dataArray[4].Aluno)

}
