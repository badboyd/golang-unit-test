package main

import (
	"bytes"
	"log"
	"testing"
	"text/template"

	"github.com/fteem/go-playground/golden-files/books"
)

const (
	header string = `
| Title         | Author        | Publisher |  Pages  |  ISBN  |  Price  |
| ------------- | ------------- | --------- | ------- | ------ | ------- |
`
	rowTemplate string = "|  {{ .Title }}  |  {{ .Author }}  |  {{ .Publisher }}  |  {{ .Pages }}  |  {{ .ISBN }}  |  {{ .Price }}  |\n"
)

type Book struct {
	ISBN      string
	Title     string
	Author    string
	Pages     int
	Publisher string
	Price     int
}

var Books []Book = []Book{
	Book{
		ISBN:      "978-1591847786",
		Title:     "Hooked",
		Author:    "Nir Eyal",
		Pages:     256,
		Publisher: "Portfolio",
		Price:     19,
	},
	Book{
		ISBN:      "978-1434442017",
		Title:     "The Great Gatsby",
		Author:    "F. Scott Fitzgerald",
		Pages:     140,
		Publisher: "Wildside Press",
		Price:     12,
	},
	Book{
		ISBN:      "978-1784756260",
		Title:     "Then She Was Gone: A Novel",
		Author:    "Lisa Jewell",
		Pages:     448,
		Publisher: "Arrow",
		Price:     29,
	},
	Book{
		ISBN:      "978-1094400648",
		Title:     "Think Like a Billionaire",
		Author:    "James Altucher",
		Pages:     852,
		Publisher: "Scribd, Inc.",
		Price:     9,
	},
}

func Generate(books []books.Book) string {
	buf := bytes.NewBufferString(header)

	t := template.Must(template.New("table").Parse(rowTemplate))

	for _, book := range books {
		err := t.Execute(buf, book)
		if err != nil {
			log.Println("Error executing template:", err)
		}
	}

	return buf.String()
}

// START OMIT

func TestGenerate(t *testing.T) {
	testcases := []struct {
		name  string
		books []books.Book
		want  string
	}{
		{
			name: "WithInventory",
			books: []books.Book{
				books.Book{
					Title:  "The Da Vinci Code",
					Author: "Dan Brown",
					Pages:  592,
					ISBN:   "978-0552161275",
					Price:  7,
				},
				books.Book{
					Title:  "American on Purpose",
					Author: "Craig Ferguson",
					Pages:  288,
					ISBN:   "978-0061959158",
					Price:  19,
				},
			},
			want: `
| Title         | Author        | Pages  | ISBN  | Price  |
| ------------- | ------------- | ------ | ----- | ------ |
| The Da Vinci Code | Dan Brown | 592 | 978-0552161275 | 7 |
| American on Purpose | Craig Ferguson | 288 | 978-0061959158 | 19 |
`,
		},
		{
			name:  "EmptyInventory",
			books: []books.Book{},
			want: `
| Title         | Author        | Pages  | ISBN  | Price  |
| ------------- | ------------- | ------ | ----- | ------ |
`,
		},
	}

	for _, testcase := range testcases {
		got := Generate(testcase.books)
		if got != testcase.want {
			t.Errorf("Want:\n%s\nGot:%s", testcase.want, got)
		}
	}
}

// END OMIT

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestGenerate", F: TestGenerate})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
