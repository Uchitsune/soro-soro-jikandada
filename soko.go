package main

import "fmt"
func main() {
	var harga float64
	var ongkir float64
	var jumlah int
	fmt.Println("masukan harga barang")
	fmt.Scan(&harga)

	fmt.Println("masukan berapa ongkirnya")
	fmt.Scan(&ongkir)

	fmt.Println("masukan jumlah barang")
	fmt.Scan(&jumlah)

	subtotal := float64(jumlah) * harga
	total := subtotal + ongkir
	fmt.Printf("checkout %.2fRP total", total)
	
}