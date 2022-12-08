package model

import "strconv"

type File struct {
	Name string
	Size int
}

func (f File) String() string {
	return strconv.Itoa(f.Size) + " " + f.Name
}
