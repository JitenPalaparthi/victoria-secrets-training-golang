// package slices is to create a new user defined package called slice
// this is a int slice with additional methods

package slices

import (
	"errors"
	"math/rand"
)

type SliceInt []int // SliceInt is a user defined type

func (si SliceInt) Fill(nums ...int) SliceInt {

	//if si == nil {
	si = make(SliceInt, len(nums))
	//}

	for i := range si {
		si[i] = nums[i]
	}

	return si

}

// Fill fills the slice with random numbers
func (si SliceInt) RandFill(len int) SliceInt {

	if si == nil {
		si = make(SliceInt, len)
	}

	for i := range si {
		si[i] = rand.Int()
	}
	return si

}

// InferDefault func infers the slice with a given value.
// normal slice by default it is infered with 0, provided by slice is int.
func (si SliceInt) InferDefaults(val int) SliceInt {

	if si == nil {
		return nil
	}

	for i := range si {
		si[i] = val
	}

	return si

}

// Max func returns the biggest value of the slice.
// if slice is nil then infer returns 0 and error
// if length of the slice is 1 then it returns the 0th element
func (si SliceInt) Max() (int, error) {

	if si == nil {
		return 0, errors.New("nil slice")
	}

	if len(si) == 1 {
		return si[0], nil
	}

	max := si[0]

	for i := 1; i < len(si); i++ {
		if max < si[i] {
			max = si[i]
		}

	}

	return max, nil
}

func (si SliceInt) SecondBiggest() (int, error) {

	if len(si) == 1 {
		return si[0], nil
	} else if len(si) == 2 {
		if si[0] > si[1] {
			return si[1], nil
		}
		return si[0], nil
	}

	max, err := si.Max() // get the maximum number
	if err != nil {
		return 0, err
	}

	slice := make([]int, 0) // create a another slice

	for _, v := range si {
		if v != max {
			slice = append(slice, v) // add eleents from si to slice exept the max number
		}
	}

	max, err = SliceInt(slice).Max() // the the max number from the second slice

	return max, err
}

func (si SliceInt) SecondBiggest2() (int, error) {

	if si == nil {
		return 0, errors.New("nil slice")
	}

	if len(si) == 1 {
		return si[0], nil
	}

	if len(si) == 2 {
		if si[0] > si[1] {
			return si[1], nil
		}
		return si[0], nil
	}

	var max, secondMax int
	if si[0] > si[1] {
		max = si[0]
		secondMax = si[1]
	} else {
		max = si[1]
		secondMax = si[0]
	}
	for i := 2; i < len(si); i++ {

		if si[i] > max {
			secondMax = max
			max = si[i]
		} else {
			if si[i] > secondMax {
				secondMax = si[i]
			}
		}

	}

	return secondMax, nil
}
