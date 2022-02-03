# Sets calculator (test task)

**To build the sources:**

go build -o sets_calc

**To run tests:**

  go test ./...

**To run tests with a coverage report:**

  go test ./...  -coverprofile=coverage.out && go tool cover -html=coverage.out

Note:
despite the fact that the solution is based on handwritten lexer and parser it's also possible to consider generation lexer and parser with Flex and Bison or similar tools.
