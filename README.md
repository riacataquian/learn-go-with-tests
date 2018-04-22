# Learn Go with tests

1. Hello World
    - Subtests, which groups test suites together.
2. Adder
    - Example test suite for documentation.
    - Run `godoc -http :8000` to view local workspace packages and their documentation.
3. Iteration
    - Benchmarking performance.
    - Run `go test -bench=.` to execute benchmarks.
    - `Square` vs `SquareX` results:

      ```
      BenchmarkSquare-4    	50000000	        33.9 ns/op
      BenchmarkSquareX-4   	10000000	       138 ns/op
      ```

      In `Square`, we initialized an array with N length compared to `SquareX` in which
      we append item as we iterate through the slice.
4. Arrays
    - Use of `range` over `for`.
    - `go test -cover` to see test coverage percentage.
    - `reflect.DeepEqual` to check equality of two variables (hmm, though it is not "type-safe").
    - Creating a variadic function.
5. Shapes
    - Demonstrates a table-driven tests.
    - Interfaces, structs and its methods.
    - Structure test for better readability and maintainability.
6. Wallet
    - Use https://github.com/kisielk/errcheck to check error handling.
    - Use pointers, to pass something by reference.
