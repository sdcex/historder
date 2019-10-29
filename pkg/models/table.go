package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// Table ...
type Table struct {
	Titles    []string
	TitleDict map[string]int
	Rows      [][]string
}

// NewTable ...
func NewTable(titles []string) *Table {
	tb := &Table{
		Titles:    titles,
		TitleDict: make(map[string]int),
		Rows:      [][]string{},
	}
	for i, title := range titles {
		tb.TitleDict[title] = i
	}
	return tb
}

// AddRowMap ...
func (tb *Table) AddRowMap(values map[string]string) error {
	row := make([]string, len(tb.Titles))
	for k, v := range values {
		if pos, ok := tb.TitleDict[k]; ok {
			row[pos] = v
		}
	}
	tb.Rows = append(tb.Rows, row)
	return nil
}

// AddRowList ...
func (tb *Table) AddRowList(values []string) error {
	if len(values) != len(tb.Titles) {
		return fmt.Errorf("row %v: col size is %v, expect %v", len(tb.Rows), len(values), len(tb.Titles))
	}
	tb.Rows = append(tb.Rows, values)
	return nil
}

// DumpData ...
func (tb *Table) DumpData() [][]string {
	return tb.Rows
}

// DumpTitles ...
func (tb *Table) DumpTitles() []string {
	return tb.Titles
}

// Save ...
func (tb *Table) Save() error {
	fileName := time.Now().Format(time.RFC3339) + ".csv"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	wr := csv.NewWriter(file)
	err = wr.Write(tb.Titles)
	if err != nil {
		return err
	}
	return wr.WriteAll(tb.Rows)
}
