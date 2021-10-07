# Go

- These are notes from: https://www.udemy.com/course/go-the-complete-developers-guide/ 

* https://golang.org/dl/ 
* Visual Studio Code — recommended editor for Go
    * Install Go extension
    * Enable analysis tools

## Hello World Code:
```
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

## Go CLI

- To compile code into an executable:
`go build main.go`

- To compile AND run code:
`go run main.go`

- To format all the code in each file of the current directory:
`go fmt`

- To compile and install a package:
`go install`

- To download the raw source code of someone else's package:
`go get`

- To run all tests in the current project:
`go test`

## Packages

- Package == Project == Workspace
- The purpose of a package is to group together code with a similar purpose.
- Each file much have their package in the first line. 
- Two types of packages:
    1. executables - generates a file that we can run
        - must be named main
        - must also have a function called main to be the entrypoint
    2. reusables - code used as "helpers"; good place to put reusable logic; creating a dependency

## Imports

- Import statements are used to give our code access to code in other packages.
- fmt is a standard library package - "format"

## Functions / File Organization

- Func = Function
- `func function-name (list of arguments) return-type {     function body }`
- Files always have package declaration, then package imports, then declare functions and tell Go to do things.

## Variable Declarations

- Example: `var card string = "Ace of Spades"`
- Go is a **statically typed language**. In other words we care about the types we assign to each variable.
- Basic Go Types:
    - bool
    - string
    - int
    - float64
    - array
    - map
    - etc.
- You can also declare variables like this: `card := "Ace of Spades"`
    - ":=" is syntax for declaring a NEW variable. To change the value, you can just use "="
    - Go infers the type based on the value on the right side
- You can initialize a variable outside of a function, but you cannot assign a value to it.

## Functions and Return Types

- If you want a function to be able to return a value, you must have the return type after the function name and parameters.
    - `func newCard() string {}`
- Note: the order of function definitions in a file does not matter. You can call a function in an earlier line than it is defined.

## Arrays and Slices
- **Array** = fixed length list of things
- **Slice** = an array that can grow or shrink
    - every element must be the same type
    - `fruits := []string{"apples", "pears", "bananas"}`
- You can add items to a slice with `append()` -- note that it returns a new slice. 
    - `fruits = append(fruits, "strawberries")`
- You can use range syntax on slices. `fruits[0:2]` -- takes the first two elements of fruits. (From 0 and up to but not including index 2.)
    - when starting from 0 you can shorthand to `fruits[:2]`
    - or if ending at the end of the slice you can shorthand `fruits[2:]` (from index 2 to the end)
- `len(fruits)` to get the length of a slice
- You can alter the value of an element at an index with `fruit[0] = "orange"`

## Loops
- You can iterate through a closed set with a for-loop. Example:
```
for i, fruit := range fruits {
	fmt.Println(i, fruit)
}
```
- i here is the index -- it does not have to be declared if it's not used in the code (can use underscore)
- don't need to define the item (fruit here) -- if not needed, leave it out

## Object-Oriented vs. Go

- Go is NOT an object-oriented programming language.
    - no classes
- In Go you can instead "extend" types and add some functions meant to work with it. 
    - For example, I can create a deck type that uses a string slice as a base type. `type deck []string`
    - Then I can have functions with deck as a **receiver**. (A function with a receiver is like a method -- a function that belongs to an instance of a class in OO.)
- Here's an example of a receiver function. Any variable of type deck now gets access to the print function.
    - d here is the actual copy of the variable of type deck we're working with (by convention Go doesn't use "this" or "self" but usually an initial abbreviation)
```
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

## More Notes
- You can use an underscore (_) when you want to tell go you know there's a variable there but you don't want to use it. `for _, suit := range cardSuits { ... }`
- You can tell Go you want to return more than one thing by having the return type in parentheses and comma separating the items. `func deal(d deck, handSize int) (deck, deck) { ... }`
- Conversion to byte slice = `[]byte("Greetings!")`
- Can generate a pseudo-random integer with `rand.Intn(r)` where r is a range
    - be careful that by default it always uses the same seed!


## Testing
- Go testing is not RSpec, mocha, jasmine, selenium, etc.!
- To make a test, create a new file ending in `_test.go`
- To run all tests in a package, run the command `go test`
- Tests need to have `t *testing.T` as a parameter -- this is the test handler.

## Structs
- Data structure. Collection of properties that are related together.
- Think of it like a Java plain object. Or like a Python dictionary.
- You can define a struct like:
```
type person struct {
	firstName string
	lastName  string
}
```
- Then you can use it like: `amanda := person{"Amanda", "Grice"}` where it matches the fields based on the order. **not recommended
- You can also create a person like: `amanda := person{firstName: "Amanda", lastName: "Grice"}`
- You can create a person without any properties set like `var alex person` - this will default firstName and lastName to empty strings.
- You can print all the properties of a struct like `fmt.Printf("%+v", amanda)`
- You can update properties of a struct like: `alex.firstName = "Alex"`
- Structs can have properties that are other structs:
```
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo   contactInfo
}

func main() {
	amanda := person{
		firstName: "Amanda",
		lastName:  "Grice",
		contactInfo: contactInfo{
			email:   "fake@gmail.com",
			zipCode: 12345,
		},
	}
	fmt.Printf("%+v\n", amanda)
}
```
- Structs can also have receiver functions.
```
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
...
amanda.print()
```

## Pass By Value and Pointers
- Go is a **"pass by value"** language. 
    - Meaning Go will copy all the fields of a value passed to a function to a new slot in memory. So updates to that object in the function will change the copy and not the original. 
- `&` creates a pointer to an object. `amandaPointer := &amanda`
- `&variable` is saying give me the memory address of the value this variable is pointing at
    - `&` here is an operator
- `*pointer` is saying give me the value this memory address is pointing at
    - `*` here is an operator
- `*type` is saying the type is a pointer to something. ex.) `func (pointerToPerson *person) updateName() {...}` is a receiver function called on a pointer to a person.
    - `*` here is not an operator
- Turn `address` into `value` with `*address`.
- Turn `value` into `address` with `&value`.
- Pointer shortcut: if you have a receiver function that expects to be called on a pointer to a struct, you can also call it on the struct.
- Note: slices operate a little bit differently than structs or other values. When passing a slice to a function, the function does not make a copy and instead modifies the original slice.
    - Go is still copying the slice data structure to a new memory address. But under the hood the slice data structure in memory is a pointer to an Array in memory (and a few other properties). So when copying, it's copying a pointer to the values. The new copy is pointing to the same address.
    - **Reference vs Value Types**
    - Slices are reference types!
    - Other reference types include maps, channels, pointers, and functions.
    - ints, floats, strings, bools, structs are all value types.

## Maps
- A map is a collection of key-value pairs. Similar to a dictionary in Python or an object in Javascript.
- Both keys and values are statically typed. So all keys must be the same type and all values must be the same type.
- You can create a map like this:
```
colors := map[string]string{
	"red":   "#ff0000",
	"green": "#4bf745",
}
```
- You can initialize an empty map like this either of two ways:
    - `var colors map[string]string`
	- `colors := make(map[string]string)`
- You can add to or update a map like this: `colors["white"] = "#ffffff"`
- You can remove from a map like this: `delete(colors, "green")`
- Maps vs Structs
    - Maps
        - all keys must be the same type
        - used to represent a collection of related properties
        - all values must be the same type
        - don't need t know all the keys at compile time
        - keys are indexed - we can iterate over them
        - Reference Type!
    - Structs
        - values can be of different type
        - you need to know all the different fields at compile time
        - keys don't support indexing
        - used to represent a "thing" with a lot of different properties
        - Value Type!


## Conventions
- https://betterprogramming.pub/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a 
- Go uses camelCase for multi-word names.
- For public identifiers -- start with a capital letter.
- For private identifiers -- use a lower-case letter.
- MethodName + er = InterfaceName
- Go doesn't really do getters and setters. Better to just export the field and get rid of the setter and getter methods.
    - if you do write getters and setters, you should not bother putting Get in the getters name
- Short variable names
- Acronyms keep consistent casing
- No fixed line lengths (but shorter is better)

## Interfaces
- Interfaces help with making code reusable across different types.
- Go does NOT support ovreloaded functions. 
- Example of creating an interface:
    - This means that all "bots" no matter the type have a function called getGreeting().
    - This is automatic - don't need to specify anything on englishBot or spanishBot, etc. types. If they implement getGreeting and getX, they are automatically bots.
```
type bot interface {
    getGreeting() string
    getX(string, int) (string, error)
}
```
- Interfaces are NOT generic types.
- Interfaces are implicit. (don't have to write things like x implements interface y)
- Interfaces are a contract to help us manage types.

## Channels & Go Routines
- A **goroutine** is a lightweight thread of execution.
    - The cost of creating a Goroutine is tiny compared to a thread.
    - Common for Go applications to have thousands of Goroutines running concurrently. 
- Goroutines are functions or methods that run concurrently with other functions or methods.
- Used so that we can run function calls asynchronously. 
- `go checkLink(link)`
- **Go Scheduler** runs one routine until it finishes or makes a blocking call (like an HTTP request).
    - not truly being executed at the same time when we only have one CPU core
    - if more than one CPU core, the scheduler runs one thread on each "logical" core
    - By default, Go tries to use one core!
- Concurrency is not parallelism.
- **Concurrency** - we can have multiple threads executing code. If one thread blocks, another one is picked up and worked on. 
- **Parallelism** - multiple threads executed at the exact same time. Requires multiple CPUs
- Always have "Main Routine" -- Go command creates "Child routines"
- **Channels** are used to communicate between goroutines.
    - otherwise a child routine could still be executing when the main routine finishes
    - **Channels are typed!** The type is the type of data that can be communicated across the channel.
- `c := make(chan string)`

- A **function literal** in Go is the same thing as an *anonymous function** in JavaScript or a **lambda function** in Python.


