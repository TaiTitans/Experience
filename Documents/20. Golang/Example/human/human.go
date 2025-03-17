package main

import "fmt"

type Human struct {
	Name     string
	YearOld  int
	Hometown string
	Job      string
}

func (h Human) Live(NoiCuTru string) {
	fmt.Printf("Ho ten: %s\n Noi cu tru: %s\n", h.Name, NoiCuTru)
}

func (h Human) CompanyAddress(DiaChiCQ string) {
	fmt.Printf("Ho ten: %s\n Co Quan: %s\n", h.Name, DiaChiCQ)
}

func main() {
	nguoi1 := Human{
		Name:     "Nguyen Van A",
		YearOld:  2002,
		Hometown: "Can Tho",
		Job:      "IT",
	}

	nguoi1.Live("Dai hoc Can Tho")
	nguoi1.CompanyAddress("Duong 3/2, Ninh Kieu, Can Tho")
}
