package excelize

import "testing"

func TestAxisLowerOrEqualThan(t *testing.T) {
	trueExpectedInputList := [][2]string{
		[2]string{"A", "B"},
		[2]string{"A", "AA"},
		[2]string{"B", "AA"},
		[2]string{"BC", "ABCD"},
		[2]string{"1", "2"},
		[2]string{"2", "11"},
	}

	for _, trueExpectedInput := range trueExpectedInputList {
		isLowerOrEqual := axisLowerOrEqualThan(trueExpectedInput[0], trueExpectedInput[1])
		if !isLowerOrEqual {
			t.Fatalf("Expected %v <= %v = true, got false\n", trueExpectedInput[0], trueExpectedInput[1])
		}
	}

	falseExpectedInputList := [][2]string{
		[2]string{"B", "A"},
		[2]string{"AA", "A"},
		[2]string{"AA", "B"},
		[2]string{"ABCD", "AB"},
		[2]string{"2", "1"},
		[2]string{"11", "2"},
	}

	for _, falseExpectedInput := range falseExpectedInputList {
		isLowerOrEqual := axisLowerOrEqualThan(falseExpectedInput[0], falseExpectedInput[1])
		if isLowerOrEqual {
			t.Fatalf("Expected %v <= %v = false, got true\n", falseExpectedInput[0], falseExpectedInput[1])
		}
	}
}

func TestGetCellColRow(t *testing.T) {
	cellExpectedColRowList := map[string][2]string{
		"C220":    [2]string{"C", "220"},
		"aaef42":  [2]string{"aaef", "42"},
		"bonjour": [2]string{"bonjour", ""},
		"59":      [2]string{"", "59"},
		"":        [2]string{"", ""},
	}

	for cell, expectedColRow := range cellExpectedColRowList {
		col, row := getCellColRow(cell)

		if col != expectedColRow[0] {
			t.Fatalf("Expected cell %v to return col %v, got col %v\n", cell, expectedColRow[0], col)
		}

		if row != expectedColRow[1] {
			t.Fatalf("Expected cell %v to return row %v, got row %v\n", cell, expectedColRow[1], row)
		}
	}
}
