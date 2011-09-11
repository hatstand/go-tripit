package tripit

// This file was generated by a tool. Do not edit.

import (
	"os"
	"json"
	"container/vector"
)

// A specialization of Vector for *TransportSegment objects
type TransportSegmentPtrVector struct {
	vector.Vector
}

// AppendVector appends the entire vector x to the end of this vector.
func (p *TransportSegmentPtrVector) AppendVector(x *TransportSegmentPtrVector) {
	p.Vector.AppendVector(&x.Vector)
}

// At returns the i'th element of the vector.
func (p *TransportSegmentPtrVector) At(i int) *TransportSegment {
	return p.Vector.At(i).(*TransportSegment)
}

// Copy makes a copy of the vector and returns it.
func (p *TransportSegmentPtrVector) Copy() TransportSegmentPtrVector {
	return TransportSegmentPtrVector{p.Vector.Copy()}
}

// Do calls function f for each element of the vector, in order. The behavior of Do is undefined if f changes *p.
func (p *TransportSegmentPtrVector) Do(f func(elem *TransportSegment)) {
	p.Vector.Do(func(e interface{}) { f(e.(*TransportSegment)) })
}

// Insert inserts into the vector an element of value x before the current element at index i.
func (p *TransportSegmentPtrVector) Insert(i int, x *TransportSegment) {
	p.Vector.Insert(i, x)
}

// InsertVector inserts into the vector the contents of the vector x such that the 0th element of x appears at
// index i after insertion.
func (p *TransportSegmentPtrVector) InsertVector(i int, x *TransportSegmentPtrVector) {
	p.Vector.InsertVector(i, &x.Vector)
}

// Last returns the element in the vector of highest index.
func (p *TransportSegmentPtrVector) Last() *TransportSegment {
	return p.Vector.Last().(*TransportSegment)
}

// Pop deletes the last element of the vector.
func (p *TransportSegmentPtrVector) Pop() *TransportSegment {
	return p.Vector.Pop().(*TransportSegment)
}

// Push appends x to the end of the vector.
func (p *TransportSegmentPtrVector) Push(x *TransportSegment) {
	p.Vector.Push(x)
}

// Resize changes the length and capacity of a vector. If the new length is shorter than the current length,
// Resize discards trailing elements. If the new length is longer than the current length, Resize adds the
// respective zero values for the additional elements. The capacity parameter is ignored unless the new length
// or capacity is longer than the current capacity. The resized vector's capacity may be larger than the
// requested capacity.
func (p *TransportSegmentPtrVector) Resize(length, capacity int) *TransportSegmentPtrVector {
	p.Vector = *p.Vector.Resize(length, capacity)
	return p
}

// Set sets the i'th element of the vector to value x.
func (p *TransportSegmentPtrVector) Set(i int, x *TransportSegment) {
	p.Vector.Set(i, x)
}

// Slice returns a new sub-vector by slicing the old one to extract slice [i:j]. The elements are copied.
// The original vector is unchanged.
func (p *TransportSegmentPtrVector) Slice(i, j int) *TransportSegmentPtrVector {
	v := p.Vector.Slice(i, j)
	return &TransportSegmentPtrVector{*v}
}

// UnmarshalJSON customizes the JSON unmarshalling by accepting single elements or arrays of elements.
func (p *TransportSegmentPtrVector) UnmarshalJSON(b []byte) os.Error {
	var arr []*TransportSegment
	err := json.Unmarshal(b, &arr)
	if err != nil {
		arr = make([]*TransportSegment, 1)
		err := json.Unmarshal(b, &arr[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				arr = arr[0:0]
			} else {
				return err
			}
		}
		
		if arr[0] == nil {
			arr = arr[0:0]
		}
		
	}
	p.Cut(0, p.Len())
	for _, v := range arr {
		p.Push(v)
	}
	return nil
}

// MarshalJSON customizes the JSON output for Vectors.
func (p *TransportSegmentPtrVector) MarshalJSON() ([]byte, os.Error) {
	var a []*TransportSegment
	if p == nil {
		a = make([]*TransportSegment, 0)
	} else {
		a = make([]*TransportSegment, p.Len())
		for i := 0; i < p.Len(); i++ {
			a[i] = p.At(i)
		}
	}
	return json.Marshal(a)
}

// Data returns all the elements as a slice.
func (p *TransportSegmentPtrVector) Data() []*TransportSegment {
	arr := make([]*TransportSegment, p.Len())
	var i int
	i = 0
	p.Do(func(v *TransportSegment) {
		arr[i] = v
		i++
	})
	return arr
}
