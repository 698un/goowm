package mymath

import "math"

type MVector struct {
	X float32
	Y float32
	Z float32
} //MVector


func (v1 *MVector) CalcModul{
	return math.Sqrt(v1.x*v1.x+
	                 v1.y*v1.y+
					 v1.z*v1.z)
	
	
}//calcModul