# Chapter 13 – Interface Internals

## Interface layout

interface value =

- itab pointer (type + method table)
- data pointer (actual value)

## itab

- Stores concrete type info
- Stores method addresses
- Enables dynamic dispatch

## Dynamic type/value

- Interface stores runtime type and value
- Changes when assigning new concrete value

```go
var s Speaker = Dog{}

itab → (Dog, Speak method address)
data → Dog{}

s.Speak() => internally : itab.method[0](data) // dynamic dispatch
--------------------------------------------
var i interface{}
i = 10
// internally : type  = int value = 10
i= "hello"
// now : type  = string value = "hello"
So interface value is : (dynamic_type, dynamic_value)




```

## Nil pitfalls

Case 1:
type=nil, value=nil → interface == nil => eg. var s Speaker

Case 2:
type=\*Dog, value=nil → interface != nil

```go
var d *Dog = nil
var s Speaker = d

// Interanlly
type = *Dog
value = nil

```

## Rule

Always check concrete value when dealing with interfaces.

## Debug tip

Use:
fmt.Printf(\"%T %v\", i, i)
