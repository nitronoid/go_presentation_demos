
@title[The go programming language]
# Go 

A programming language

---

@title[Types in go]
### Basic types in go

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

### Functions in go

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

### Collections in go

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

### Structs and user defined types in Go

- Go defines types with the type keyword similiar to typedef.
- These can be copies of exisiting types or newly defined structs.
- Constructors and destructors don't exist, we use free functions to create new objects if setup is non-trivial.

+++?code=structs_types.go&lang=golang&title=structs_types.go

@[6](An example of type copying. Not the same as an alias.)
@[10-12, 33-34](This function is an exampe of strict types in Go. myInt is the same as a regular int but won't be accepted.)
@[15-18](This is how structs are defined as types, no functions are written here.)
@[21-25, 52](We can emulate constructors using a function that returns an object.)
@[24](Note that it is legal to return a pointer to a local variable. Go performs pointer escape analysis.)
@[30-31](Here we create an instance of the basic class, all structs members can be brace initialised.)
@[37-38,42](Go supports the creation of anonymous structs.)
@[46-49](Anonymous structs can be brace initialised like named ones.)


