1. What features of Go facilitate Object Oriented Programming?
Mainly interface{} and struct methods.

2. What are some of the disadvantages of using Go over other compiled languages such as C++?
Go has an inbuilt garbage collector which can make memory management more difficult as compared to other compiled languages. so it is less suited for system programming and more suited for network code and api servers.

3. What are some of the advantages of using Go over dynamic languages such as Python?
Go is fast. Go also doesn't have any type errors that can occur at runtime in dynamic languages

4. What is the standard way in a Go application to share information about a request among many processing steps?
using the http/requests library where the request is a struct with info including header, body, and cookies and is passed into functions.
Note: not sure if this was the intended answer as there are many ways to share the info about a request depending on how the info needs to be processed. e.g. using a channel and go routines to async process the data.

5. Are the arrays passed by value OR by reference when handled as args in a function?
neither. An array in go is more accurately described as a pointer to the beginning of an array with a length. and that pointer/length is passed by value. so the underlying array itself isn't passed by either.

6. How is inheritance supported in Golang?
It's not. at least not in a traditional sense of inheritance. you can declare a struct within a struct to make a composite that is kinda like inheritance. i.e.
```
type shape struct {
    ...// stuff relating to shape
}

type square struct {
    shape // or....
    ShapeAttributes shape
}

7. Demonstrate how we can interchange the values of two variables without using a third variable?
num1, num2 = num2, num1

8. Do parallelism and concurrency means same in Golang - please explain?
for practical purposes, yes. Behind the scenes the Go scheduler takes care of whether go routines run in parallel or concurrently.

9. What does GOPATH denote? What problem does it solve?
Where dependencies are installed/stored. it keeps from having multiple copies of the same library installed among multiple projects. to be clear, it is not dependency management

10. State difference between static and dynamic declaration types of variable?

11. What are channels and how and why are they used in Go?
channels are like mini event streams. they are used to communicate to and from go routines.

12. Is there a way or provision in Golang that can facilitate cleanup (sort of a garbage collector routine)?
Garbage collection is handled automagically by go itself.

13. What are various ways we can address data synchronization issues within Golang?
One way I have used for syncing data is wait groups. Another is using switch statements that read from channels

14. Can we develop a cross platform Go binary? If so then how?
yes, set the `GOOS` and the `GOARCH` env variables and then run `go build`.

15. How do we define "dynamic structs" in Golang?
By using the `comparable` and `any` type parameters. This feature is more aptly called generics. e.g.
```
type MyStruct[T any] struct { }
```

16. Is Go strict-typed language - please explain?
yes. it enforces typing on all data including variables, arrays, maps, and function parameters

Write a command line utility in Go which takes as arguments a list of directories. The program should output the sizes of each of the individual directories passed as well as a cumulative total. If a --recursive flag is provided, output the sizes of each of the individual directories passed and sub-directories recursively as well as a cumulative total. If a "--human" flag is passed, format the sizes to be human friendly by outputting the size in the most appropriate unit of bytes. For example, 304K for 304,000 bytes and 300M for 300000000 bytes.
 


