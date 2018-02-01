
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

# Functions in go

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

