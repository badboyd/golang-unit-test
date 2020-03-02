package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"text/template"

	"github.com/fteem/go-playground/golden-files/books"
)

var (
	update = flag.Bool("update", false, "update the golden files of this test")
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
		name   string
		books  []books.Book
		golden string
	}{
		{
			name: "WithInventory",
			books: []books.Book{
				books.Book{
					Title:     "The Da Vinci Code",
					Author:    "Dan Brown",
					Publisher: "Corgi",
					Pages:     592,
					ISBN:      "978-0552161275",
					Price:     7,
				},
				books.Book{
					Title:     "American on Purpose",
					Author:    "Craig Ferguson",
					Publisher: "Harper Collins",
					Pages:     288,
					ISBN:      "978-0061959158",
					Price:     19,
				},
			},
			golden: "inventory",
		},
		{
			name:   "EmptyInventory",
			books:  []books.Book{},
			golden: "empty",
		},
	}

	for _, testcase := range testcases {
		got := Generate(testcase.books)
		want := goldenValue(t, testcase.golden, got, *update)

		if got != want {
			t.Errorf("Want:\n%s\nGot:\n%s", want, got)
		}
	}
}

func goldenValue(t *testing.T, goldenFile string, actual string, update bool) string {
	t.Helper()
	goldenPath := "testdata/" + goldenFile + ".golden"

	f, _ := os.OpenFile(goldenPath, os.O_RDWR, 0644)
	defer f.Close()

	if update {
		_, err := f.WriteString(actual)
		if err != nil {
			t.Fatalf("Error writing to file %s: %s", goldenPath, err)
		}

		return actual
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Error opening file %s: %s", goldenPath, err)
	}
	return string(content)
}

// END OMIT

func main() {
	flag.Parse()

	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestGenerate", F: TestGenerate})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
