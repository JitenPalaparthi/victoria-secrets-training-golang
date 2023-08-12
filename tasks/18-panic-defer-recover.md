
Calc(a,b any)(float64,float64,float64,float64)

if a,b or a or b are other than 14 types(we already knew) panic it .

call the Calc function more than once.

the last call must be a successful execution.

Calc("Hi","Vitoria") // error panic

Calc([]int{10,12},10) // error panic

Calc(10,20) // should give 30,-10,200,0