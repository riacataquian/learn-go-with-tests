# Learn Go with tests

- [Hello](https://github.com/riacataquian/learn-go-with-tests/tree/master/hello)
    - Subtests, which groups test suites together.
- [Adder](https://github.com/riacataquian/learn-go-with-tests/tree/master/adder)
    - Example test suite for documentation.
    - Run `godoc -http :8000` to view local workspace packages and their documentation.
- [Iteration](https://github.com/riacataquian/learn-go-with-tests/tree/master/iteration)
    - Benchmarking performance.
    - Run `go test -bench=.` to execute benchmarks.
    - `Square` vs `SquareX` results:

      ```
      BenchmarkSquare-4    	50000000	        33.9 ns/op
      BenchmarkSquareX-4   	10000000	       138 ns/op
      ```

      In `Square`, we initialized an array with N length compared to `SquareX` in which
      we append item as we iterate through the slice.
- [Arrays](https://github.com/riacataquian/learn-go-with-tests/tree/master/arrays)
    - Use of `range` over `for`.
    - `go test -cover` to see test coverage percentage.
    - `reflect.DeepEqual` to check equality of two variables (hmm, though it is not "type-safe").
    - Creating a variadic function.
- [Shapes](https://github.com/riacataquian/learn-go-with-tests/tree/master/shapes)
    - Demonstrates a table-driven tests.
    - Interfaces, structs and its methods.
    - Structure test for better readability and maintainability.
- [Wallet](https://github.com/riacataquian/learn-go-with-tests/tree/master/wallet)
    - Use https://github.com/kisielk/errcheck to check error handling.
    - Use pointers, to pass something by reference.
- [Greeter](https://github.com/riacataquian/learn-go-with-tests/tree/master/greeter)
    - Demonstrates a testable, decoupled and flexible code with dependency injection.
- [Countdown](https://github.com/riacataquian/learn-go-with-tests/tree/master/countdown)
    - Demonstrates _mocking_ in tests.
    - Favor testing behaviour over implementation.
- [Concurrency](https://github.com/riacataquian/learn-go-with-tests/tree/master/concurrency)
    - Performing mocks in tests for a better controlled and faster execution.
    - `goroutines` as the basic unit of concurrency in Go.
    - Use of anonymous functions.
    - `go` keyword which accepts a function then execute it in another process.
    - How map is bad for concurrent writes.
    - Effective use of `channel` as the communication between different processes.
    - `go test -race` as a tool to detect race conditions.
