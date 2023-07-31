# Integers and Iterations
Example function : these provide more clear documentation to the code 

```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```

## iterations

To do stuff repeatedly in Go, you'll need `for`. In Go there are no `while`, `do`, `until` keywords, you can only use `for`. Which is a good thing!

The `for` syntax is very unremarkable and follows most C-like languages.

```
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
```

Unlike other languages like C, Java, or JavaScript there
 are no parentheses surrounding the three components of the for 
statement and the braces

```
{ }
```

are always required. You might wonder what is happening in the row

different varient of loops are 

package main

import "fmt"

func main() {

```
i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}

for j := 7; j <= 9; j++ {
    fmt.Println(j)
}

infinite loop 
for {
    fmt.Println("loop")
    break
}

for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}

```

}

### Benchmarking

Writing [benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks) in Go is another first-class feature of the language and it is very similar to writing tests.

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

You'll see the code is very similar to a test.

The `testing.B` gives you access to the cryptically named `b.N`.

When the benchmark code is executed, it runs `b.N` times and measures how long it takes.

The amount of times the code is run shouldn't matter to 
you, the framework will determine what is a "good" value for that to let
 you have some decent results.

To run the benchmarks do `go test -bench=.` (or if you're in Windows Powershell `go test -bench="."`)

```
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS

```

What `136 ns/op` means is our 
function takes on average 136 nanoseconds to run (on my computer). Which
 is pretty ok! To test this it ran it 10000000 times.

*NOTE* by default Benchmarks are run sequentially.