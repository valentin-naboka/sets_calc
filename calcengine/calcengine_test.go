package calcengine_test

import (
	"reflect"
	"runtime"
	"sets_calc/calcengine"
	"sets_calc/lexer"
	"sets_calc/parser"
	"strings"
	"testing"
)

//TODO: extract to test utils
func printCaller(t *testing.T, depth int) {
	function, file, line, _ := runtime.Caller(depth)
	trimName := func(n string) string {
		i := strings.LastIndex(n, "/")
		if i == -1 {
			return n
		}
		return n[i+1:]
	}
	t.Logf("%s: line %d, function: %s\n", trimName(file), line, trimName(runtime.FuncForPC(function).Name()))
}

func checkValue(t *testing.T, reader calcengine.IFileReader, input string, expected []int) {
	r := strings.NewReader(input)
	scanner := lexer.NewScanner(r)
	parser := parser.Parser{Scanner: scanner}
	expr, err := parser.BuildExpression()

	if err != nil {
		printCaller(t, 2)
		t.Fatal("unable to parse en expression: ", err)
	}

	engine := calcengine.NewCalculator(reader)
	result, err := engine.Execute(expr)

	if err != nil {
		printCaller(t, 2)
		t.Fatal("unable to calulate en expression: ", err)
	}

	if !reflect.DeepEqual(expected, result) {
		printCaller(t, 2)
		t.Errorf("\nexpected: %v\nactual:   %v", expected, result)
	}
}

type FakeFileReader struct {
	Data map[string][]int
}

func NewFakeFileReader() *FakeFileReader {
	return &FakeFileReader{make(map[string][]int)}
}

func (r *FakeFileReader) Read(filename string) ([]int, error) {
	return r.Data[filename], nil
}

func TestCalcEngine(t *testing.T) {
	{
		fakeReader := NewFakeFileReader()
		fakeReader.Data["a.txt"] = []int{1, 2, 3}
		fakeReader.Data["b.txt"] = []int{2, 3, 4}
		fakeReader.Data["c.txt"] = []int{1, 2, 3, 4, 5}

		checkValue(t, fakeReader, "[ EQ 1 c.txt ]", []int{1, 2, 3, 4, 5})
		checkValue(t, fakeReader, "[ EQ 1 a.txt ]", []int{1, 2, 3})

		checkValue(t, fakeReader, "[ GR 1 c.txt [ EQ 3 a.txt a.txt b.txt ] ]", []int{2, 3})
		checkValue(t, fakeReader, "[ LE 2 a.txt [ GR 1 b.txt c.txt ] ]", []int{1, 4})
	}

	{
		fakeReader := NewFakeFileReader()

		fakeReader.Data["a.txt"] = []int{1}
		fakeReader.Data["b.txt"] = []int{5, 6, 7, 8, 9}
		fakeReader.Data["c.txt"] = []int{1, 2, 3, 4}
		fakeReader.Data["d.txt"] = []int{8, 9}
		fakeReader.Data["e.txt"] = []int{6, 7, 8, 9}
		fakeReader.Data["f.txt"] = []int{1, 2}
		fakeReader.Data["g.txt"] = []int{2, 3, 4}

		checkValue(t, fakeReader, "[ EQ 1 a.txt b.txt c.txt d.txt e.txt ]", []int{2, 3, 4, 5})
		checkValue(t, fakeReader, "[ EQ 1 a.txt b.txt c.txt d.txt e.txt f.txt g.txt ]", []int{5})
		checkValue(t, fakeReader, "[ EQ 2 a.txt b.txt c.txt d.txt e.txt f.txt g.txt ]", []int{3, 4, 6, 7})
		checkValue(t, fakeReader, "[ EQ 3 a.txt b.txt c.txt d.txt e.txt f.txt g.txt ]", []int{1, 2, 8, 9})

		fakeReader.Data["h.txt"] = []int{10}
		checkValue(t, fakeReader, "[ LE 2 a.txt b.txt c.txt d.txt e.txt f.txt g.txt h.txt ]", []int{5, 10})

		checkValue(t, fakeReader, "[ GR 1 a.txt b.txt c.txt d.txt e.txt f.txt g.txt h.txt ]", []int{1, 2, 3, 4, 6, 7, 8, 9})
		checkValue(t, fakeReader, "[ GR 2 a.txt b.txt c.txt d.txt e.txt f.txt g.txt h.txt ]", []int{1, 2, 8, 9})
	}
}
