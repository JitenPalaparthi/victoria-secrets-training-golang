# nil

- There is no null 
- There is not None
- There is nil in Go
- nil is means in general in Go nothing, It is not pointed to any memory

- The default value of any or interface{} variable is nil
- If no value is given to any or interface{} the type is also nil
- can check nil on error. Because error is an interface object

# functions

- To create a function use keywork called func

- Functions avoid repeatetion of code

- The scope of variables of functions are for the function. All other scoping rules also applied.
- There is a different scope and scope rules for referneces or giving or returning references from functions

- By default all varialbes are pass or call by value
- Can also pass references as arguments they are call by reference.

- func can return more than one return value
- func return values are in stack or in heap is decided by go runtime

# Switch

- switch in go is more powerful than other programming switch cases
- can also write conditions in case, but the switch must me empty switch
- break is imposed by default.No need to give break.
- can also have case values comma saperated
- If you dont want to break the flow , similar to removing break in other programming languages, you can impose fallthrough keyword.
- removing break in other programming languages is equal to giving fallthrough
- fallthrough can also lead to false negative.
- switch statement can also work on types.Types are evaluated from any variable


# arrays

- the type of an array contains its length as well. That means arr:=[5]int ; The type is not int,the type if [5]int
- type inference is applied to arrays as well
- to get lengh of array use len(arr1)
- to get cap of array use cap(arr1); But cap is of no use for array since the array is a fixed size data type

# string and loops

- for loop on string gives byte by byte value 
- for range loop on string gives character by character. range loop can cleaverly understand unicode character.



# Misc

- func GreetMe()  Pascal case
- func greetMe()  camel Case
- func greet_me() snake_case

- string literals

    Strings can be gives between `` ` ``"Hello World"`` ` ``. This takes the as it is string, if they are given in back quote symbol ``

- blank identifier

    _ is a blank identifier. This is a blank variable

- back quote

# method receivers

- receives are call by value unless they are references
- receivers can be for any user defined type. Not only confined to struct types

