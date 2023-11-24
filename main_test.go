package main

import (
	"fmt"
	"github.com/pmezard/go-difflib/difflib"
	"testing"
	"time"
)

func TestBuildUsingDisplayModeByDate(t *testing.T) {
	documents := getCDSDocuments()
	builder := NewTreeNodeBuilder(documents, ByDate)
	builder.Build()
	got := builder.ToString()

	want := `root (10)
  2023 (3)
    - Date: 2023-11-23, Specialty: Neurology, Type: Progress Note
    - Date: 2023-11-01, Specialty: Cardiology, Type: Discharge Summary
    - Date: 2023-05-23, Specialty: Cardiology, Type: Lab Result
  2022 (2)
    - Date: 2022-11-23, Specialty: Cardiology, Type: Discharge Summary
    - Date: 2022-09-15, Specialty: Neurology, Type: Progress Note
  2021 (1)
    - Date: 2021-06-21, Specialty: Neurology, Type: Progress Note
  2019 (1)
    - Date: 2019-11-23, Specialty: Cardiology, Type: Lab Result
  2017 (1)
    - Date: 2017-10-27, Specialty: Cardiology, Type: Discharge Summary
  2014 (1)
    - Date: 2014-11-23, Specialty: Cardiology, Type: Lab Result
  2010 (1)
    - Date: 2010-01-02, Specialty: Cardiology, Type: Discharge Summary
`

	if want != got {
		//t.Errorf("ToString output is incorrect.\nExpected:\n%s\nGot:\n%s", want, got)
		// Compute the differences
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(want),
			B:        difflib.SplitLines(got),
			FromFile: "want",
			ToFile:   "got",
			Context:  3,
		}

		text, err := difflib.GetUnifiedDiffString(diff)
		if err != nil {
			fmt.Println("Error computing diff:", err)
			return
		}

		// Print the diff
		//fmt.Println(text)
		t.Errorf("ToString output is incorrect.\nDiff:\n%s\n", text)
	} else {
		fmt.Println(got)
	}
}

func TestBuildUsingDisplayModeBySpecialty(t *testing.T) {
	documents := getCDSDocuments()
	builder := NewTreeNodeBuilder(documents, BySpecialty)
	builder.Build()
	got := builder.ToString()

	want := `root (10)
  Cardiology (7)
    Discharge Summary (4)
      - Date: 2023-11-01, Specialty: Cardiology, Type: Discharge Summary
      - Date: 2022-11-23, Specialty: Cardiology, Type: Discharge Summary
      - Date: 2017-10-27, Specialty: Cardiology, Type: Discharge Summary
      - Date: 2010-01-02, Specialty: Cardiology, Type: Discharge Summary
    Lab Result (3)
      - Date: 2023-05-23, Specialty: Cardiology, Type: Lab Result
      - Date: 2019-11-23, Specialty: Cardiology, Type: Lab Result
      - Date: 2014-11-23, Specialty: Cardiology, Type: Lab Result
  Neurology (3)
    Progress Note (3)
      - Date: 2023-11-23, Specialty: Neurology, Type: Progress Note
      - Date: 2022-09-15, Specialty: Neurology, Type: Progress Note
      - Date: 2021-06-21, Specialty: Neurology, Type: Progress Note
`

	if want != got {
		//t.Errorf("ToString output is incorrect.\nExpected:\n%s\nGot:\n%s", want, got)
		// Compute the differences
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(want),
			B:        difflib.SplitLines(got),
			FromFile: "want",
			ToFile:   "got",
			Context:  3,
		}

		text, err := difflib.GetUnifiedDiffString(diff)
		if err != nil {
			fmt.Println("Error computing diff:", err)
			return
		}

		// Print the diff
		//fmt.Println(text)
		t.Errorf("ToString output is incorrect.\nDiff:\n%s\n", text)
	} else {
		fmt.Println(got)
	}
}

func TestBuildUsingDisplayModeByType(t *testing.T) {
	documents := getCDSDocuments()
	builder := NewTreeNodeBuilder(documents, ByType)
	builder.Build()
	got := builder.ToString()

	want := `root (10)
  Discharge Summary (4)
    Cardiology (4)
      - Date: 2023-11-01, Specialty: Cardiology, Type: Discharge Summary
      - Date: 2022-11-23, Specialty: Cardiology, Type: Discharge Summary
      - Date: 2017-10-27, Specialty: Cardiology, Type: Discharge Summary
      - Date: 2010-01-02, Specialty: Cardiology, Type: Discharge Summary
  Lab Result (3)
    Cardiology (3)
      - Date: 2023-05-23, Specialty: Cardiology, Type: Lab Result
      - Date: 2019-11-23, Specialty: Cardiology, Type: Lab Result
      - Date: 2014-11-23, Specialty: Cardiology, Type: Lab Result
  Progress Note (3)
    Neurology (3)
      - Date: 2023-11-23, Specialty: Neurology, Type: Progress Note
      - Date: 2022-09-15, Specialty: Neurology, Type: Progress Note
      - Date: 2021-06-21, Specialty: Neurology, Type: Progress Note
`

	if want != got {
		//t.Errorf("ToString output is incorrect.\nExpected:\n%s\nGot:\n%s", want, got)
		// Compute the differences
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(want),
			B:        difflib.SplitLines(got),
			FromFile: "want",
			ToFile:   "got",
			Context:  3,
		}

		text, err := difflib.GetUnifiedDiffString(diff)
		if err != nil {
			fmt.Println("Error computing diff:", err)
			return
		}

		// Print the diff
		//fmt.Println(text)
		t.Errorf("ToString output is incorrect.\nDiff:\n%s\n", text)
	} else {
		fmt.Println(got)
	}
}

func getCDSDocuments() []Document {

	documents := []Document{
		{
			Date:      time.Date(2023, 11, 01, 14, 30, 45, 100, time.Local),
			Specialty: "Cardiology",
			Type:      "Discharge Summary",
		},
		{
			Date:      time.Date(2023, 11, 23, 14, 30, 45, 100, time.Local),
			Specialty: "Neurology",
			Type:      "Progress Note",
		},
		{
			Date:      time.Date(2022, 9, 15, 14, 30, 45, 100, time.Local),
			Specialty: "Neurology",
			Type:      "Progress Note",
		},
		{
			Date:      time.Date(2021, 6, 21, 14, 30, 45, 100, time.Local),
			Specialty: "Neurology",
			Type:      "Progress Note",
		},
		{
			Date:      time.Date(2010, 01, 02, 14, 30, 45, 100, time.Local),
			Specialty: "Cardiology",
			Type:      "Discharge Summary",
		},
		{
			Date:      time.Date(2017, 10, 27, 14, 30, 45, 100, time.Local),
			Specialty: "Cardiology",
			Type:      "Discharge Summary",
		},
		{
			Date:      time.Date(2022, 11, 23, 14, 30, 45, 100, time.Local),
			Specialty: "Cardiology",
			Type:      "Discharge Summary",
		},
		{
			Date:      time.Date(2023, 5, 23, 14, 30, 45, 100, time.Local),
			Specialty: "Cardiology",
			Type:      "Lab Result",
		},
		{
			Date:      time.Date(2019, 11, 23, 14, 30, 45, 100, time.Local),
			Specialty: "Cardiology",
			Type:      "Lab Result",
		},
		{
			Date:      time.Date(2014, 11, 23, 14, 30, 45, 100, time.Local),
			Specialty: "Cardiology",
			Type:      "Lab Result",
		},
	}
	return documents
}
