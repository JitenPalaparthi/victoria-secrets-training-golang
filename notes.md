# nil

- There is no null 
- There is not None
- There is nil in Go
- nil is means in general in Go nothing, It is not pointed to any memory

- The default value of any or interface{} variable is nil
- If no value is given to any or interface{} the type is also nil

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


# Misc

- func GreetMe()  Pascal case
- func greetMe()  camel Case
- func greet_me() snake_case