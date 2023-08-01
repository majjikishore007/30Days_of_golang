# Arrays and Slices

### Arrays

Arrays have a *fixed capacity* which you define when you declare the variable.
We can initialize an array in two ways:

- [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

It is sometimes useful to also print the inputs to the function in the error
message and we are using the `%v` placeholder which is the "default" format,
which works well for arrays.

[Read more about the format strings](https://golang.org/pkg/fmt/)

Basics types of string formating

- %v : default
- %d : base10 (decimal)
- %s : stirng
- %p : address of the 0th element
- %t : boolean

```
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
```

writing the sum function 

```
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
```

To get the value out of an array at a particular index, just use `array[index]`
syntax. In this case, we are using `for` to iterate 5 times to work through the
array and add each item onto `sum`.

Refactor

Let's introduce `[range](https://gobyexample.com/range)` to help clean up our code

```
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

```
range
```

lets you iterate over an array. Every time it is called it returns two
values, the index and the value. We are choosing to ignore the index value by
using

```
_
```

[blank identifier](https://golang.org/doc/effective_go.html#blank)

.

### Arrays and their type

An interesting property of arrays is that the size is encoded in its type. If you try
to pass an `[4]int` into a function that expects `[5]int`, it won't compile.
They are different types so it's just the same as trying to pass a `string` into
a function that wants an `int`.

You may be thinking it's quite cumbersome that arrays have a fixed length, and most
of the time you probably won't be using them!

Go has *slices* which do not encode the size of the collection and instead can
have any size.

The next requirement will be to sum collections of varying sizes.

### Slices

We will now use the [slice type](https://golang.org/doc/effective_go.html#slices) which allows us to have collections of
any size. The syntax is very similar to arrays, you just omit the size when
declaring them

`mySlice := []int{1,2,3}` rather than `myArray := [3]int{1,2,3}`

## Write the minimal amount of code for the test to run and check the failing test output

We need to define SumAll according to what our test wants.

Go can let you write *[variadic functions](https://gobyexample.com/variadic-functions)* that can take a variable number of arguments.

```go
func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}
```

Try to compile but our tests still don't compile!

`./sum_test.go:26:9: invalid operation: got != want (slice can only be compared to nil)`

Go does not let you use equality operators with slices. You *could* write
a function to iterate over each `got` and `want` slice and check their values
but for convenience sake, we can use `[reflect.DeepEqual](https://golang.org/pkg/reflect/#DeepEqual)` which is
useful for seeing if *any* two variables are the same.

```
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

(make sure you `import reflect` in the top of your file to have access to `DeepEqual`)

It's important to note that `reflect.DeepEqual` is not "type safe", the code
will compile even if you did something a bit silly. To see this in action,
temporarily change the test to:

```
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := "bob"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

What we have done here is try to compare a `slice` with a `string`. Which makes
no sense, but the test compiles! So while using `reflect.DeepEqual` is
a convenient way of comparing slices (and other things) you must be careful
when using it.

Change the test back again and run it, you should have test output looking like this

`sum_test.go:30: got [] want [3 9]`

## Write enough code to make it pass

What we need to do is iterate over the varargs, calculate the sum using our
`Sum` function from before and then add it to the slice we will return

```
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
```

Lots of new things to learn!

There's a new way to create a slice. `make` allows you to create a slice with
a starting capacity of the `len` of the `numbersToSum` we need to work through.

You can index slices like arrays with `mySlice[N]` to get the value out or
assign it a new value with `=`

The tests should now pass