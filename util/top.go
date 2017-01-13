package util

import (
	"strings"
)

type TopRecord struct {
	Tel   string  `json:"tel"`
	Names []string  `json:"names"`
}

type NameIndex struct {
	Row     int      `json:"tel"`
	Columns []int    `json:"columns"`
}

func ParseRecords(records [][]string) ([]*TopRecord, int) {
	topRecords := []*TopRecord{}
	if records == nil || len(records) == 0 {
		return topRecords, 0
	}
	nameIndex := getNameTitleIndex(records)
	if nameIndex.Row == -1 {
		return topRecords, -1
	}

	telIndex := getTelIndex(records)
	for rIndex, line := range records {
		// skip title row
		if rIndex <= nameIndex.Row {
			continue
		}
		names := []string{}
		for _, column := range nameIndex.Columns {
			names = append(names, line[column])
		}
		topRecord := &TopRecord{
			Names: names,
		}
		if telIndex != -1 {
			topRecord.Tel = line[telIndex]
		}
		topRecords = append(topRecords, topRecord)

	}
	return topRecords, 0
}

func getNameTitleIndex(records [][]string) *NameIndex {
	nameIndex := &NameIndex{
		Row: -1,
		Columns: []int{},
	}
	for rIndex, record := range records {
		find := false
		for cIndex, val := range record {
			if isName(val) {
				nameIndex.Row = rIndex
				nameIndex.Columns = append(nameIndex.Columns, cIndex)
				find = true
			}
		}
		if find {
			break
		}
	}
	return nameIndex
}

func getTelIndex(records [][]string) int {
	minRatio := 0.8
	telsCount := make([]int, len(records[0]))
	for _, line := range records {
		for index, cell := range line {
			if isTel(cell) {
				telsCount[index] += 1;
			}

		}
	}
	telIndex := maxIndex(telsCount)
	maxCount := telsCount[telIndex]
	if float64(maxCount) / float64(len(records)) > minRatio {
		return telIndex
	}
	return -1
}

func isName(val string) bool {
	if strings.Contains(val, "姓名") || strings.Contains(val, "名字") ||
	strings.Contains(val, "收件人") || strings.Contains(val, "收件方") ||
	strings.Contains(val, "收货人") || strings.Contains(val, "收货方") {
		return true
	}
	return false
}

func isTel(str string) bool {
	if strings.HasPrefix(str, "1") && len(str) == 11 {
		return true
	}
	return false
}

func maxIndex(numArray []int) int {
	max := numArray[0]
	idx := 0
	for index, v := range numArray {
		if max < v {
			max = v
			idx = index
		}
	}
	return idx
}

func getColumn(records [][]string, column int) []string {
	list := []string{}
	for _, record := range records {
		if (isTel(record[column])) {
			list = append(list, record[column])
		}
	}
	return list
}