package calcengine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sets_calc/ast"
	"strconv"
)

type Calculator struct {
}

func (c *Calculator) Execute(expression *ast.Expression) []int {
	return c.calcExpr(expression)
}

//TODO: don no use whole array. Use only ranges
func (c *Calculator) calcExpr(e *ast.Expression) []int {
	var sets [][]int
	for _, set := range e.Sets {
		file, ok := set.(ast.File)
		if ok {
			numbers, err := c.readFile(&file)
			_ = err

			sets = append(sets, numbers)
			//TODO: remove
			fmt.Printf("%v", numbers)
		} else {
			expr, ok := set.(*ast.Expression)
			if !ok {
				return nil
			}
			res := c.calcExpr(expr)
			if res != nil {
				sets = append(sets, res)
			}
		}
	}

	var predicate func(int, int) bool

	switch e.OpType {
	case ast.Equal:
		{
			predicate = func(n int, count int) bool {
				return n == count
			}
			break
		}
	case ast.GreaterThan:
		{
			predicate = func(n int, count int) bool {
				return n < count
			}
			break
		}
	case ast.LessThan:
		{
			predicate = func(n int, count int) bool {
				return n > count
			}
			break
		}
	}

	return c.calculate(int(e.N), predicate, sets...)
}

//TODO: review.
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

func (c *Calculator) calculate(n int, predicate func(int, int) bool, sets ...[]int) []int {
	var result []int
	for _, set := range sets[:len(sets)-1] {
		for _, item := range set {
			count := 1
			for j := 1; j < len(sets); j++ {

				//TODO: remove a set
				if len(sets[j]) == 0 {
					continue
				}

				if item >= sets[j][0] && item <= sets[j][len(sets[j])-1] {
					//TODO: remove element
					for i := 0; i < len(sets[j]); i++ {
						if sets[j][i] == item {
							copy(sets[j][i:], sets[j][i+1:]) // Shift a[i+1:] left one index.
							sets[j][len(sets[j])-1] = 0      // Erase last element (write zero value).

							sets[j] = sets[j][:len(sets[j])-1] // Truncate slice.
							break
						}
					}
					count++
				}
			}
			if predicate(n, count) {
				result = append(result, item)
			}
		}
	}

	for _, item := range sets[len(sets)-1] {
		if predicate(n, 1) {
			result = append(result, item)
		}
	}
	return result
}
