package entity

import (
	"github.com/tealeg/xlsx"
)

func ReadSpreadsheet(filePath string) ([]Participant, error) {
	var data []Participant

	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	sheet := xlFile.Sheets[0]
	for i := 1; i < len(sheet.Rows); i++ {
		row := sheet.Rows[i]
		if len(row.Cells) < 4 {
			continue
		}
		name := row.Cells[1].String()
		email := row.Cells[3].String()
		data = append(data, Participant{Name: name, Email: email})
	}

	return data, nil
}
