# Pointers , Errors 

## Pointers
- A pointer is a variable that stores the memory address of another variable.
- The type *T is a pointer to a T value. Its zero value is nil.



Go copies values when you pass them to functions/methods so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples could be very large data or perhaps things you intend only to have one instance of (like database connection pools).

## nil

Pointers can be nil
When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception, the compiler won't help you here.
Useful for when you want to describe a value that could be missing

## Errors

Errors are the way to signify failure when calling a function/method.
By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.
Don’t just check errors, handle them gracefully

## Create new types from existing ones

Useful for adding more domain specific meaning to values
Can let you implement interfaces

Pointers and errors are a big part of writing Go that you need to get comfortable with. Thankfully the compiler will usually help you out if you do something wrong, just take your time and read the error.


ref: https://github.com/quii/learn-go-with-tests/blob/master/pointers-and-errors.md