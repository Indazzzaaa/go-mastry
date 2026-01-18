# Chapter 2 – Program Structure

## Packages

- Folder = package
- A Compilation Unit
- All files in folder must use same package name.

## main package

- Entry point of program
- Must contain func main()

## Imports

- Single import: import "fmt"
- Multiple imports using () => impot ( "fmt" "math")
  - Aliases => import m "math"
- Unused imports cause compile error

## Visibility

- Capitalized = exported => Public
- lowercase = unexported => Private

## init() function

- Runs before main()
- Runs once per package
- Used for setup , config loading, validation.
- Allowd one per file, and execution order will be same as file order.

## Observations

- init() runs before main()
- Multiple init() allowed
- Order depends on file compilation order
