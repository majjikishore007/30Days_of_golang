# Structs , interfaces and methods

## Structs
A struct is just a named collection of fields where you can store data.

Declare a struct like this

```go
type Rectangle struct {
    Width float64
    Height float64
}
```


Some programming languages allow you to do something like this:
```go
func Area(circle Circle) float64 { ... }
func Area(rectangle Rectangle) float64 { ... }
```
But you cannot in Go

We have two choices:

- You can have functions with the same name declared in different packages. So we could create our Area(Circle) in a new package, but that feels overkill here.
- We can define methods on our newly defined types instead.

## Methods

What are methods?

So far we have only been writing functions but we have been using some methods. When we call t.Errorf we are calling the method Errorf on the instance of our t (testing.T).

The syntax for declaring methods is almost the same as functions and that's because they're so similar. The only difference is the syntax of the method receiver `func (receiverName ReceiverType) MethodName(args)`.


example 

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64  {
    return 0
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64  {
    return 0
}

```

Interfaces are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.

Let's introduce this by refactoring our tests.

```go 
func TestArea(t *testing.T) {

    checkArea := func(t *testing.T, shape Shape, want float64) {
        t.Helper()
        got := shape.Area()
        if got != want {
            t.Errorf("got %g want %g", got, want)
        }
    }

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        checkArea(t, rectangle, 72.0)
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        checkArea(t, circle, 314.1592653589793)
    })

}

```

We are creating a helper function like we have in other exercises but this time we are asking for a Shape to be passed in. If we try to call this with something that isn't a shape, then it will not compile.

How does something become a shape? We just tell Go what a Shape is using an interface declaration

```go
type Shape interface {
    Area() float64
}
```

We're creating a new type just like we did with Rectangle and Circle but this time it is an interface rather than a struct.

Once you add this to the code, the tests will pass.
Wait, what?

This is quite different to interfaces in most other programming languages. Normally you have to write code to say My type Foo implements interface Bar.

But in our case

- `Rectangle` has a method called `Area` that returns a `float64` so it satisfies the Shape interface
- `Circle` has a method called `Area` that returns a `float64` so it satisfies the Shape interface
- `string` does not have such a method, so it doesn't satisfy the interface
etc.


## Decoupling

Decoupling

Notice how our helper does not need to concern itself with whether the shape is a Rectangle or a Circle or a Triangle. By declaring an interface the helper is decoupled from the concrete types and just has the method it needs to do its job.

This kind of approach of using interfaces to declare only what you need is very important in software design and will be covered in more detail in later sections.


reference: https://github.com/quii/learn-go-with-tests/blob/master/structs-methods-and-interfaces.md