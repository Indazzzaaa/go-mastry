# Chapter 9 – Methods & Receivers

## Methods

- Functions attached to a type (struct).
- Go support two types of methods (Value Receivers (Copy Passed), Pointer Receiver(Reference of Original passed ))

## Value Receivers

- Receive a copy
- Cannot modify original
- Good for small immutable types

  ```go
  func (u User) SetName(name string) {
      u.Name = name   // modifies copy original data remain same
  }

  ```

## Pointer Receivers

- Receive address
- Can modify original
- Avoid copying large structs

  ```go
  func (u *User) SetName(name string) {
      u.Name = name   // modifies original ✅
  }

  ```

## Method Sets

- T has value methods
- \*T has both value + pointer methods

  ```go
  //Meaning
  var u User  // u can call value methods
  var p *User  // p can call both

  ```

## Embedded Methods

- Methods of embedded struct are promoted
- Enables composition

  ```go
  type Animal struct {}

  func (a Animal) Speak() {}

  type Dog struct {
      Animal
  }

  // to call
  var d Dog
  d.speak()

  ```

## Best Practice

- Use pointer receivers by default
- Be consistent

## Mental Model

Struct = data
Method = behavior
Receiver = owner
Pointer receiver = real object
Value receiver = copy
Embedding = reuse behavior
