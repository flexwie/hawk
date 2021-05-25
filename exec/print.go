package exec

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/jedib0t/go-pretty/v6/table"
	"hawk.wie.gg/models"
)

func PrintTable(entries []models.Entry) (string, error) {
	t := table.NewWriter()
	typ := reflect.TypeOf(models.Entry{})

	header := table.Row{}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Tag.Get("hawk") == "ignore" {
			continue
		}
		header = append(header, f.Name)
	}

	t.AppendHeader(header)
	t.AppendSeparator()

	for _, j := range entries {
		row := table.Row{}
		s := reflect.ValueOf(j)
		for i := 0; i < s.NumField(); i++ {
			f := typ.Field(i)
			if f.Tag.Get("hawk") == "ignore" {
				continue
			}

			row = append(row, fmt.Sprintf("%v", s.Field(i).String()))
		}
		t.AppendRow(row)
	}

	return t.Render(), nil
}

func PrintJSON(entries []models.Entry, pretty bool) (string, error) {
	var j []byte
	var err error

	if pretty {
		j, err = json.MarshalIndent(entries, "", "  ")
	} else {
		j, err = json.Marshal(entries)
	}

	return string(j), err
}
