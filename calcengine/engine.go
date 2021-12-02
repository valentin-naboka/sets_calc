package calcengine

import (
	"sets_calc/ast"
	"sort"

	"github.com/pkg/errors"
)

func makePredicate(op ast.Operator) func(int, int) bool {
	switch op {
	case ast.Equal:
		return func(n int, count int) bool {
			return n == count
		}

	case ast.GreaterThan:
		return func(n int, count int) bool {
			return n < count
		}
	case ast.LessThan:
		return func(n int, count int) bool {
			return n > count
		}
	}
	return nil
}

type IFileReader interface {
	Read(filename string) ([]int, error)
}

type Calculator struct {
	fileReader IFileReader
}

func NewCalculator(reader IFileReader) *Calculator {
	return &Calculator{fileReader: reader}
}

func (c *Calculator) Execute(expression *ast.Expression) ([]int, error) {
	return c.calcExpr(expression)
}

func (c *Calculator) calcExpr(e *ast.Expression) ([]int, error) {
	var sets [][]int
	for _, set := range e.Sets {
		file, ok := set.(ast.File)
		if ok {
			numbers, err := c.fileReader.Read(file.Name)

			if err != nil {
				return nil, errors.Wrapf(err, "unable to read numbers from the file: %s", file.Name)
			}

			sets = append(sets, numbers)
			continue
		}

		expr, ok := set.(*ast.Expression)
		if !ok {
			return nil, errors.Errorf("unexpected type of set")
		}
		res, err := c.calcExpr(expr)
		if err != nil {
			return nil, err
		}

		if res != nil {
			sets = append(sets, res)
		}
	}

	return c.calculate(int(e.N), makePredicate(e.OpType), sets...), nil
}

type Sets struct {
	data [][]int
}

func (s *Sets) Len() int {
	return len(s.data)
}

func (s *Sets) Less(i, j int) bool {
	return s.data[i][len(s.data[i])-1] > s.data[j][len(s.data[j])-1]
}

func (s *Sets) Swap(i, j int) {
	tmp := s.data[i]
	s.data[i] = s.data[j]
	s.data[j] = tmp
}

func (c *Calculator) calculate(n int, predicate func(int, int) bool, sets ...[]int) []int {

	var result []int

	setsWrapper := Sets{data: sets}
	// NOTE: Sort sets in descending order based on the last element of the set.
	// I.e, [{1,2,3}, {4,5,6}, {3,4,5}] will be rearanged to [{4,5,6}, {3,4,5}, {1,2,3}]
	sort.Sort(&setsWrapper)

	for currentSetIdx, currentSet := range setsWrapper.data {
		for len(setsWrapper.data[currentSetIdx]) != 0 {
			// Remove currentItem (which is the last item in the set) from the slice
			LastItemIdx := len(currentSet) - 1
			lastItem := currentSet[LastItemIdx]
			setsWrapper.data[currentSetIdx] = currentSet[:LastItemIdx]
			currentSet = setsWrapper.data[currentSetIdx]
			occurenceCount := 1

			for otherSetIdx := currentSetIdx + 1; otherSetIdx < len(setsWrapper.data); otherSetIdx++ {

				// Skip the slice if it's already empty
				if len(setsWrapper.data[otherSetIdx]) == 0 {
					continue
				}

				if lastItem == setsWrapper.data[otherSetIdx][len(setsWrapper.data[otherSetIdx])-1] {
					occurenceCount++
					setsWrapper.data[otherSetIdx] = setsWrapper.data[otherSetIdx][:len(setsWrapper.data[otherSetIdx])-1]
				}
			}

			if predicate(n, occurenceCount) {
				result = append(result, lastItem)
			}
		}
	}

	// NOTE: in a production code, it's better to find a robust library for such things,
	// in order to avoid "reinventing the wheel"
	for i := 0; i < len(result)/2; i++ {
		tmp := result[i]
		result[i] = result[len(result)-i-1]
		result[len(result)-i-1] = tmp
	}
	return result
}
