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
- [DI](https://github.com/riacataquian/learn-go-with-tests/tree/master/di)
  - Testing printed output (`fmt.Printf`).
  - Demonstrates a testable, decoupled and flexible code with dependency injection.
- [Countdown](https://github.com/riacataquian/learn-go-with-tests/tree/master/countdown)
  - Demonstrates _mocking_ and _spies_ to test against _how_ a dependency is used.
  - Test against behavior over implementation.
- [Concurrency](https://github.com/riacataquian/learn-go-with-tests/tree/master/concurrency)
  - Demonstrates use of `goroutines` as the basic unit of concurrency in Go,
    for non-blocking operations.
  - Anonymous functions maintain access to the lexical scope they are defined in.
  - Map is bad for concurrent writes.
  - Effective use of `channel` for better communication and control between goroutines,
    prevents race condition bugs when making use of native Go data structures manually.
  - Anonymous struct fields: see result struct.
  - `go test -race` as a tool to detect race conditions.
- [Racer](https://github.com/riacataquian/learn-go-with-tests/tree/master/racer)
  - Demonstrates use of `select` construct: helps you wait on multiple channels.
  - Use of `net/http/httptest` for a convenient and controlled http tests: demonstrates a mock implementation of httpserver.
  - (Sometimes) Use `time.After` to prevent your system blocking forever.
- [HttpServer](https://github.com/riacataquian/learn-go-with-tests/tree/master/httpserver)
  - Implement `Handler` interface to create web servers.
  - Use `go build` which takes in all `.go` files in the directory to build the program as an executable file via `./myprogram`.
- [JSON](https://github.com/riacataquian/learn-go-with-tests/tree/master/httpserver) (HttpServer extension)
  - Usage of `ServeMux`, Go's request multiplexer.
  - Encapsulate `http.Handler` to a struct so setting up routes can only be done once.
  - When types are embedded, its methods are promoted. Only expose what is necessary to your public API.
  - JSON decoding and encoding.
