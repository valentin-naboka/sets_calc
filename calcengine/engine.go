package calcengine

import (
	"bufio"
	"log"
	"os"
	"sets_calc/ast"
	"sort"
	"strconv"

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

type Calculator struct {
}

func (c *Calculator) Execute(expression *ast.Expression) ([]int, error) {
	return c.calcExpr(expression)
}

func (c *Calculator) calcExpr(e *ast.Expression) ([]int, error) {
	var sets [][]int
	for _, set := range e.Sets {
		file, ok := set.(ast.File)
		if ok {
			numbers, err := c.readFile(&file)
			_ = err

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

//TODO: Filereader interface to make unittests
func (c *Calculator) readFile(file *ast.File) ([]int, error) {

	f, err := os.Open(file.Name)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
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
	sort.Sort(&setsWrapper)

	for currentSetIdx, currentSet := range setsWrapper.data {
		for len(setsWrapper.data[currentSetIdx]) != 0 {

			//Remove currentItem (which is the last item in the set) from the slice
			LastItemIdx := len(currentSet) - 1
			lastItem := currentSet[LastItemIdx]
			setsWrapper.data[currentSetIdx] = currentSet[:LastItemIdx]
			occurenceCount := 1

			for otherSetIdx := currentSetIdx + 1; otherSetIdx < len(setsWrapper.data); otherSetIdx++ {

				//Skip the slice if it's already empty
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

// //TODO: don;t use whole array. Use only ranges???
// func (c *Calculator) calculate(n int, predicate func(int, int) bool, sets ...[]int) []int {
// 	var result []int
// 	for _, set := range sets[:len(sets)-1] {
// 		for _, item := range set {
// 			count := 1
// 			for j := 1; j < len(sets); j++ {

// 				//TODO: remove a set
// 				if len(sets[j]) == 0 {
// 					continue
// 				}

// 				if item >= sets[j][0] && item <= sets[j][len(sets[j])-1] {
// 					//TODO: remove element
// 					for i := 0; i < len(sets[j]); i++ {
// 						if sets[j][i] == item {
// 							copy(sets[j][i:], sets[j][i+1:]) // Shift a[i+1:] left one index.
// 							sets[j][len(sets[j])-1] = 0      // Erase last element (write zero value).

// 							sets[j] = sets[j][:len(sets[j])-1] // Truncate slice.
// 							break
// 						}
// 					}
// 					count++
// 				}
// 			}
// 			if predicate(n, count) {
// 				result = append(result, item)
// 			}
// 		}
// 	}

// 	for _, item := range sets[len(sets)-1] {
// 		if predicate(n, 1) {
// 			result = append(result, item)
// 		}
// 	}
// 	return result
// }
