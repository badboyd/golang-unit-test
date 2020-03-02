package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
)

type Record struct {
	student string
	subject string
	grade   string
}

type Gradebook []Record

func NewGradebook(csvFile io.Reader) (Gradebook, error) {
	var gradebook Gradebook
	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return gradebook, err
		}

		if len(line) < 3 {
			return gradebook, fmt.Errorf("Invalid file structure")
		}

		gradebook = append(gradebook, Record{
			student: line[0],
			subject: line[1],
			grade:   line[2],
		})
	}

	return gradebook, nil
}

func (gb *Gradebook) FindByStudent(student string) []Record {
	var records []Record
	for _, record := range *gb {
		if student == record.student {
			records = append(records, record)
		}
	}
	return records
}

func buildGradebook(t *testing.T, path string) *Gradebook {
	csvFile, err := os.Open(path)
	if err != nil {
		t.Errorf("error opening file: %v", err)
	}

	gradebook, err := NewGradebook(csvFile)
	if err != nil {
		t.Fatalf("Cannot create Gradebook: %v", err)
	}

	return &gradebook
}

// START OMIT

func TestFindByStudent(t *testing.T) {
	cases := []struct {
		fixture string
		student string
		want    Gradebook
		name    string
	}{
		{
			fixture: "./testdata/empty.csv",
			student: "Jane",
			want:    Gradebook{},
			name:    "EmptyFixture",
		},
		{fixture: "./testdata/valid.csv",
			student: "Jane",
			want: Gradebook{
				Record{
					student: "Jane",
					subject: "Chemistry",
					grade:   "A",
				},
				Record{
					student: "Jane",
					subject: "Algebra",
					grade:   "B",
				},
			},
			name: "ValidFixture",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gradebook := buildGradebook(t, tc.fixture)

			got := gradebook.FindByStudent(tc.student)
			for idx, wantedGrade := range tc.want {
				// wantedGrade := tc.want[idx]
				gotGrade := got[idx]
				if gotGrade != wantedGrade {
					t.Errorf("Expected: %v, got: %v", wantedGrade, gotGrade)
				}
			}

		})
	}
}

// END OMIT

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestFindByStudent", F: TestFindByStudent})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
