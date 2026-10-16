package main
import "fmt"

func main() {
	var harga, diskon1, hargaAkhir float64

	fmt.Scan(&harga)
	diskon1 = harga * 0.9
	hargaAkhir = diskon1 * 0.8

	fmt.Printf("%.2f\n", hargaAkhir)
}