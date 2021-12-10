package helper

import "strconv"

//struct
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

//interface
type Description interface {
	GetDescription() string
}

//func
func (ph Phone) GetDescription() string {
	var desc = "Name : " + ph.Name + "\n" +
		"Brand : " + ph.Brand + "\n" +
		"Year : " + strconv.Itoa(ph.Year) + "\n" +
		"Colors : "
	for index, item := range ph.Colors {
		if index == 0 {
			desc += item
		} else {
			desc += ", " + item
		}
	}
	return desc
}
