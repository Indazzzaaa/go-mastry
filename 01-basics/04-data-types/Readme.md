# Chapter 4 – Data Types

## Numeric Types

- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64
  Note : int, unint, size depends on OS (32 or 64bit)
- float32, float64
- complex64, complex128
- byte ( alias for uint8)
- rune (alias for int32)

## Boolean

- Only true or false
- No implicit conversion ( like int -> bool conversion)

## String & Rune

- Strings are immutable
- UTF-8 encoded ( 1-4 bytes)
- rune = int32 (Unicode character)
- string indexing returns bytes not character

## Type Conversion Vs Casting

- Go does NOT allow implicit conversion
  ```go
  var a int = 10
  var b float64 = a // ERROR
  var c float64 = float64(a) // Casting
  ```
- Must explicitly convert: float64(x)

## Type Alias vs Defined Type

Alias:
type A = int => (Same type A and int is same we can assign a type to int)

Defined:
type B int => New type, not assinable to int directly

Defined types are NOT assignable without conversion.

## Rule

Prefer explicit types when modeling domain data.
