package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Human struct {
	HoTen        string
	NamSinh      int
	NgheNghiep   string
	QueQuan      string
	DiaChiCoQuan string
}

func (h Human) Work() {
	fmt.Printf("%s lam viec tai %s\n", h.HoTen, h.DiaChiCoQuan)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nhap so luong nguoi: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	soLuongNguoi, _ := strconv.Atoi(input)

	danhSachNguoi := make([]Human, soLuongNguoi)
	for i := 0; i < soLuongNguoi; i++ {
		fmt.Printf("\nNhap thong tin nguoi thu %d:\n", i+1)

		fmt.Printf("Nhap ho ten: ")
		hoTen, _ := reader.ReadString('\n')
		danhSachNguoi[i].HoTen = strings.TrimSpace(hoTen)

		fmt.Print("Nhap nam sinh: ")
		namSinhStr, _ := reader.ReadString('\n')
		namSinh, _ := strconv.Atoi(strings.TrimSpace(namSinhStr))
		danhSachNguoi[i].NamSinh = namSinh

		fmt.Print("Nhap que quan: ")
		queQuan, _ := reader.ReadString('n')
		danhSachNguoi[i].QueQuan = strings.TrimSpace(queQuan)

		fmt.Print("Nhap nghe nghiep: ")
		ngheNghiep, _ := reader.ReadString('\n')
		danhSachNguoi[i].NgheNghiep = strings.TrimSpace(ngheNghiep)

		fmt.Print("Nhap dia chi co quan: ")
		diaChiaCoQuan, _ := reader.ReadString('\n')
		danhSachNguoi[i].DiaChiCoQuan = strings.TrimSpace(diaChiaCoQuan)
	}
	fmt.Println("\nThong tin nhung nguoi da nhap: ")
	for _, nguoi := range danhSachNguoi {
		fmt.Printf("Họ tên: %s\n", nguoi.HoTen)
		fmt.Printf("Năm sinh: %d\n", nguoi.NamSinh)
		fmt.Printf("Quê quán: %s\n", nguoi.QueQuan)
		fmt.Printf("Nghề nghiệp: %s\n", nguoi.NgheNghiep)
		nguoi.Work()
		fmt.Println("--------------------")
	}
}
