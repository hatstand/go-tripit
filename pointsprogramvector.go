package tripit

// This file was generated by a tool. Do not edit.

import (
	"os"
	"json"
	"container/vector"
)

// A specialization of Vector for PointsProgram objects
type PointsProgramVector struct {
	vector.Vector
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *PointsProgramVector) AppendVector(x *PointsProgramVector) {
	p.Vector.AppendVector(&x.Vector)
}

// At returns the i'th element of the vector.
func (p *PointsProgramVector) At(i int) PointsProgram {
	return p.Vector.At(i).(PointsProgram)
}

// Copy makes a copy of the vector and returns it.
func (p *PointsProgramVector) Copy() PointsProgramVector {
	return PointsProgramVector{p.Vector.Copy()}
}

// Do calls function f for each element of the vector, in order. The behavior of Do is undefined if f changes *p.
func (p *PointsProgramVector) Do(f func(elem PointsProgram)) {
	p.Vector.Do(func(e interface{}) { f(e.(PointsProgram)) })
}

// Insert inserts into the vector an element of value x before the current element at index i.
func (p *PointsProgramVector) Insert(i int, x PointsProgram) {
	p.Vector.Insert(i, x)
}

// InsertVector inserts into the vector the contents of the vector x such that the 0th element of x appears at
// index i after insertion.
func (p *PointsProgramVector) InsertVector(i int, x *PointsProgramVector) {
	p.Vector.InsertVector(i, &x.Vector)
}

// Last returns the element in the vector of highest index.
func (p *PointsProgramVector) Last() PointsProgram {
	return p.Vector.Last().(PointsProgram)
}

// Pop deletes the last element of the vector.
func (p *PointsProgramVector) Pop() PointsProgram {
	return p.Vector.Pop().(PointsProgram)
}

// Push appends x to the end of the vector.
func (p *PointsProgramVector) Push(x PointsProgram) {
	p.Vector.Push(x)
}

// Resize changes the length and capacity of a vector. If the new length is shorter than the current length,
// Resize discards trailing elements. If the new length is longer than the current length, Resize adds the
// respective zero values for the additional elements. The capacity parameter is ignored unless the new length
// or capacity is longer than the current capacity. The resized vector's capacity may be larger than the
// requested capacity.
func (p *PointsProgramVector) Resize(length, capacity int) *PointsProgramVector {
	p.Vector = *p.Vector.Resize(length, capacity)
	return p
}

// Set sets the i'th element of the vector to value x.
func (p *PointsProgramVector) Set(i int, x PointsProgram) {
	p.Vector.Set(i, x)
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j]. The elements are copied.
// The original vector is unchanged.
func (p *PointsProgramVector) Slice(i, j int) *PointsProgramVector {
	v := p.Vector.Slice(i, j)
	return &PointsProgramVector{*v}
}

// UnmarshalJSON customizes the JSON unmarshalling by accepting single elements or arrays of elements.
func (p *PointsProgramVector) UnmarshalJSON(b []byte) os.Error {
	var arr []PointsProgram
	err := json.Unmarshal(b, &arr)
	if err != nil {
		arr = make([]PointsProgram, 1)
		err := json.Unmarshal(b, &arr[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				arr = arr[0:0]
			} else {
				return err
			}
		}
		
	}
	p.Cut(0, p.Len())
	for _, v := range arr {
		p.Push(v)
	}
	return nil
}

// MarshalJSON customizes the JSON output for Vectors.
func (p *PointsProgramVector) MarshalJSON() ([]byte, os.Error) {
	var a []PointsProgram
	if p == nil {
		a = make([]PointsProgram, 0)
	} else {
		a = make([]PointsProgram, p.Len())
		for i := 0; i < p.Len(); i++ {
			a[i] = p.At(i)
		}
	}
	return json.Marshal(a)
}

// Data returns all the elements as a slice.
func (p *PointsProgramVector) Data() []PointsProgram {
	arr := make([]PointsProgram, p.Len())
	var i int
	i = 0
	p.Do(func(v PointsProgram) {
		arr[i] = v
		i++
	})
	return arr
}
