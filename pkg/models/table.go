package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// Table ...
type Table struct {
	titles    []string
	titleDict map[string]int
	rows      [][]string
	statitics [][]string
}

// NewTable ...
func NewTable(titles []string) *Table {
	tb := &Table{
		titles:    titles,
		titleDict: make(map[string]int),
		rows:      [][]string{},
		statitics: [][]string{},
	}
	for i, title := range titles {
		tb.titleDict[title] = i
	}
	return tb
}

// AddRowMap ...
func (tb *Table) AddRowMap(values map[string]string) error {
	row := make([]string, len(tb.titles))
	for k, v := range values {
		if pos, ok := tb.titleDict[k]; ok {
			row[pos] = v
		}
	}
	tb.rows = append(tb.rows, row)
	return nil
}

// AddRowList ...
func (tb *Table) AddRowList(values []string) error {
	if len(values) != len(tb.titles) {
		return fmt.Errorf("row %v: col size is %v, expect %v", len(tb.rows), len(values), len(tb.titles))
	}
	tb.rows = append(tb.rows, values)
	return nil
}

// AddStatistics ...
func (tb *Table) AddStatistics(stt []string) error {
	tb.statitics = append(tb.statitics, stt)
	return nil
}

// GetTitleIndex ...
func (tb *Table) GetTitleIndex(title string) (int, bool) {
	index, ok := tb.titleDict[title]
	return index, ok
}

// DumpData ...
func (tb *Table) DumpData() [][]string {
	return tb.rows
}

// DumpStatistics ...
func (tb *Table) DumpStatistics() [][]string {
	return tb.statitics
}

// Dumptitles ...
func (tb *Table) Dumptitles() []string {
	return tb.titles
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
	err = wr.Write(tb.titles)
	if err != nil {
		return err
	}
	err = wr.WriteAll(tb.rows)
	if err != nil {
		return err
	}
	err = wr.WriteAll(tb.statitics)
	return err
}
