
@title[Intro]
# Jack Diver
### Go presentation
#### SDAGE 2nd year 

---

@title[The go programming language]
# Go overview

- Go was built with simplcity in mind, for general purpose and systems programming in mind. 
- It is a compiled, statically and strongly typed language. 
- It supports concurrency and garbage collection out of the box. 
- Go does not support inheritance in the traditional sense.

---

@title[The history of go]
# History of Go

- Go devlopment was started in 2007 and announced in 2009.
- It was developed by Robert Griesemer, Rob Pike and Ken Thompson at google.
- A shared dislike between the three was of C++'s complexity. 
- They wanted to create a language that removed common gripes with languages such as C++, Java, python etc... but keep the strong points such as being readable ("light on the page"), being scalable to large systems, and supporting networking.

---

@title[Basic types in go]
### Basic types in Go

- Go has automatic type deduction for variables, but also allows them to be explicitly typed. This is important in some cases.
- Types are specified after the variable name.
- Go is a strictly typed language, with no implicit conversions.

+++?code=types.go&lang=golang&title=types.go

@[8-14](Global variable definitions)
@[17-21](We can delcare multiple variables of different type through one statement)
@[23-27](The same is done for constants, where const replaces var)
@[30-36](Here we compare three different syntax used for delcaring variables)
@[31](Automatically determined type)
@[33-34](Declaration requires type and will be given a zero value. Then we can assign later.)
@[36](The var keyword can be omitted completely for local variables if we use the := syntax.)
@[49-51](Using type deduction with floating point numbers yeilds a double precision variable. There are no float32 literals.)
@[53-58](Go has very strict typing with no implict conversions, even between integral types. All conversions must be explicit.)

---

@title[Functions in go]
### Functions in Go

- Go functions return types are again specified at the end of the declaration before the function body.
- There is no void in Go, we can simply ommit the return type.
- Go supports:
    * Returning multiple values
    * Variadic arguments
    * Named return values
    * Parameter lists that share type
- Go doesn't support function overloading or generics.

+++?code=functions.go&lang=golang&title=functions.go

@[5-12](No overloading.)
@[14-23](Variadic arguments.)
@[25-27](Multiple return values.)
@[29-33](Named return value.)

---

@title[Collections in go]
### Collections in Go

- Go has several built-in collection types:
    * Arrays (string acts similarly)
    * Slices
    * Maps
- A slice is the ```std::vector``` of Go

+++?code=collections.go&lang=golang&title=collections.go

@[6-13](Arrays have a fixed size, given in the declaration.)
@[15-17](Slices have a similar syntax but omit the length.)
@[19-21](Slices can easily be appended to. If the slice grows too large for it's currently allocated memory, a new one is returned, hence the assignment.)
@[23-27](Deleting from a slice is strangely complicated. We must cut the slice down to before the element we want to remove, then append all the elements after it.)
@[29-32](If the order is unimportant we can assign from back and trim the slice by one.)
@[34-35](A pitfall is that slices containing pointers won't get GC'd by this, as the original slice may still exist and reference it. We need to delete the references.)
@[37-39](We are provided with a make function to construct collections of a given size. These elements will have a zero value.)
@[41-44](Maps can be declared, but will give you an error if you attempt to assign to them before using the make function.)
@[47](The make function requires the type anyway so I always use this syntax.)
@[47-54](Maps can have fairly complex types as their key, here an array is used.)
@[56-59](Maps can be created with data like other collections.)
@[62-63](When accessing a value from a map, you can also test for the key.)
@[66](Deleting from a Map is far simpler than from a slice.)

---

@title[Structs and user defined types in go]
### Structs and user defined types in Go

- Go defines types with the type keyword similiar to typedef.
- These can be copies of exisiting types or newly defined structs.
- Constructors and destructors don't exist, we use free functions to create new objects if setup is non-trivial.

+++?code=structs_types.go&lang=golang&title=structs_types.go

@[6](An example of type copying. Not the same as an alias.)
@[10-12](This function is an exampe of strict types in Go. myInt is the same as a regular int but won't be accepted.)
@[33-34](Call)
@[15-18](This is how structs are defined as types, no functions are written here.)
@[21-25](We can emulate constructors using a function that returns an object.)
@[24](Note that it is legal to return a pointer to a local variable. Go performs pointer escape analysis.)
@[30-31](Here we create an instance of the basic class, all structs members can be brace initialised.)
@[37-38,42](Go supports the creation of anonymous structs.)
@[46-49](Structs can be intialised out of order by naming the members.)

---

@title[Object oriented code in go]
### Object oriented code (ish) in Go

- Go is an object oriented language in that it allows the creation of structs that "inherit" from one and other, and also have member functions.
- Go also allows the creation of interfaces.
- Inheritence in Go does not work like other languages, types that implement an interface automatically "inherit" from it.
- Usually inheritance is described by an is-a relationship, but in Go it is an acts-like relationship.

+++?code=oo_style_1.go&lang=golang&title=oo_style_1.go

@[5-9](Interfaces in Go are declared with the interface keyword, and only contain function declarations.)
@[16-18](Here we can see a dog member function. Defined with the special syntax before the function name.)
@[16](There are two flavours of this as we will soon see.)
@[11-18](By adding this member function, Dog now implements the Animal interface.)
@[34-39](This means we can store a Dog within an Animal interface variable.)
@[20-29](Member functions can either be defined with a pointer contract or with a value contract.)
@[20-24](This member function makes a copy of the object and acts on that data.)
@[26-29](This member function acts on the original object through a pointer.)
@[41-47](Hence the object is unchanged when calling the first, but modified when calling the second.)

---

@title[Interface contracts in go]
### Interface contracts in Go

- When passing an object to an interface parameter, a contract is formed.
- The function that recieves the object checks how the member function calls have been bound.
- If the member functions have pointer contracts, a pointer must be passed to the interface.
- If the member functions have been bound by value, the interface can recieve either a pointer or a value.


+++?code=oo_style_2.go&lang=golang&title=oo_style_2.go

@[11-17](Here we can see the same function being bound with two different contracts.)
@[25-27](This function takes an interface and will determine the contract from the call to Noise.)
@[30-31](Here we call the function, this would not work without the &.)

--- 

@title[Composition and struct embedding in go]
### Composition and struct embedding in Go

- Go handles composition in a similar way to most languages.
- You can add a member to a struct, by supplying a name and type.
- Go also offers us struct embedding. Which is used by only giving a type.
- Embedding can be used to emmulate inheritence from other languages, but is still composition.

+++?code=oo_style_3.go&lang=golang&title=oo_style_3.go

@[9-16](A simple struct with one member function.)
@[19-21](Driveway has a member of type Car, standard composition.)
@[39-42](Access to the members data and functions is through the member.)
@[24-28](Now we use struct embedding by omiting the members name.)
@[44-47](We can access Cars data and functions directly through the NewCar type.)
@[30-33](We can even embed a pointer, which can be used to interesting effect.)
@[49-53](If we pass a pointer to create a ProxyCar, changes made to the proxy will affect the original car, through the pointer. Similar to a reference.)
@[55-57](We can still access the embedded part of a struct, like we would with a non-embedded one. This is similar to casting to the base class in other OO languages.)

---

@title[Closures in go]
### Closures in Go

- Go has closures, similar to lambda functions in c++.
- Closures automatically capture all local variables, and go performs analysis on this so that they aren't destroyed when we exit that scope. 

+++?code=closures.go&lang=golang&title=closures.go

@[5-13](This function returns a closure.)
@[6-7](The closure automatically captures these two variables.)
@[8-12](Using the variables here changes their scope, they won't get destroyed when this closure is returned.)
@[17-21](We get the returned closure and can call it successivly to mutate the captured variables.)


---

@title[Generics in go]
### Generics in Go

- Go doesn't have generics like Java or C++.
- We can emulate them in a very useful and safe way using interfaces.
- As Go automatically makes types "inherit" from interfaces that they implement,
we can create interfaces to specify the types that should be passed to a function.
- The key is that we can request that passed types do something, without having to modify any types that could be passed.

+++?code=generics.go&lang=golang&title=generics.go

@[5-9](Here we define the interface that variables must satisfy to be passed into our function. In this case they must convert to a float32.)
@[26-32](Our function signature requires two Convertible types.)
@[12-23](Here I defined two structs which implement the Convertible interface.)
@[35-38](We can pass both of these objects to the function as they satisfy the interface.)
@[39-42](Can't pass a regular int as it doesn't satisfy the interface.)

+++

- Another example would be that all types passed to a sorting algorithm implement the Sortable interface, containing a less than function and are copyable.

---

@title[Packages in go]
### Packages in Go

- Packages are Go's way of organising and encapsulating code.
- The main package must contain a file with a main function to run.
- Other packages can contain types, variables and functions to be used outside of that package.
- Every directory should contain at most one package that matches the directory name.
- Things can only be accessed from outside the package if they are exported, which is done by starting that data with a capital letter.

+++?code=pkg_demos/a_pkg.go&lang=golang&title=pkg_demos/a_pkg.go

Here we can see some types and functions defined in a package.

+++?code=pkg_demo.go&lang=golang&title=pkg_demo.go

@[5](We can use an alias to refer to an imported package.)
@[11-12](Only call functions with captial letter.)
@[15](We can only used named intialisation for members that are exported.)
@[19,22](We can't intialise a member if it an instance of a non-exported type, even if that member itself is exported. However we can still access the member.)

---

@title[Concurrency in go]
### Concurrency in Go

- Go gives us the keyword go to create subroutines.
- Placing go infront of a statement will execute that concurrently.
- We are also given channels which can be used to syncronise routines and also send messages between them.

+++?code=concurrency.go&lang=golang&title=concurrency.go

@[9-13](Simple function to be executed concurrently.)
@[16](Logger is thread safe.)
@[17-19](Here we create 10 threads to print out consecutive integers.)

+++?code=channels.go&lang=golang&title=channels.go

@[23-26](Channels can be created using the chan keyword, the type specifier is for the message that will be sent and recieved.)
@[8-20](This function uses two channels, one for sending messages and one for receiving messages.)
@[8](Channels can be limited to only sending or only receiving messages, note the placement of <-. )
@[12](This line waits here until a message is received and then assigns it to m.)
@[16](This sends our message through the sender channel.)
@[28-30](By setting up to routines with reversed channels we can get them to ping pong agains eachother.)
@[32-33](This line sends the intial signal to start the pinging.)


