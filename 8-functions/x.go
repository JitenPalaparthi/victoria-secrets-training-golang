package main

// the return type is towards the end of the function signature
func areaOfRect(l, b float32) float32 { // float32 is an unnamed return type
	return l * b
}
func Rect(l, b float32) (area float32, perimeter float32) { // named return types
	//func Rect(l, b float32) (float32, float32) { // unnamed return types.
	area = l * b // No need to redeclare area becase it has been declared in the function return type
	perimeter = 2 * (l + b)
	return area, perimeter
}
