# I/O

Similar to other languages, such as Java, Go models data input and output as a stream that flows from sources to targets.
Data resources, such as files, networked connections, or even some in-memory objects, can be modeled as streams of bytes from which data can be *read* or *written* to.

The stream of data is represented as **slice of bytes** (**[]byte**) that can be accessed for reading or writing.
The `io` package makes available the `io.Reader` interface to implement code that *reads* and transfer data from a source into a stream of bytes.
Conversely, the `io.Writer` interface lets implementer create code that reads data from a provided stream of bytes and *writes* it as output to a target resource.
Both interfaces are used extensively in Go as a standard idiom to express IO operations.
This makes it possible to interchange readers and writers of different implementations and contexts with predictable results.

## The io.Reader interface

The `io.Reader` interface is simple.
It consists of a single method `Reaa([]byte)(int, error)`, intended to let programmers implement code that *reads* data, from an arbitrary source, and transfers it into the provided slice of bytes.
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```
The `Read` method returns the total number of bytes transferred into the provided slice and an error value (if necessary).
As a guidance, implementations of the `io.Reader` should return an error value of `io.EOF` when the reader has no more data to transfer into stream `p`.
