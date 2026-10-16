package main
import "fmt"

func main() {


type dataBuku struct {
judul string
genre string
tahun int
rating float64
}


var buku1, buku2 dataBuku

// INPUT
fmt.Print("Masukan Judul buku : ")
fmt.Scan(&buku1.judul, &buku2.judul)

fmt.Print("Masukan genre buku : ")
fmt.Scan(&buku1.genre, &buku2.genre)

fmt.Print("Masukan tahun buku : ")
fmt.Scan(&buku1.tahun, &buku2.tahun)

fmt.Print("Masukan rating buku : ")
fmt.Scan(&buku1.rating, &buku2.rating)




// OUTPUT
fmt.Println("--- Book 1 ---")
fmt.Println("Judul Buku 1 : ", buku1.judul)
fmt.Println("Genre Buku 1 : ", buku1.genre)
fmt.Println("Tahun Buku 1 : ", buku1.tahun)
fmt.Println("Rating Buku 1 : ", buku1.rating)

fmt.Println("--- Book 2 ---")
fmt.Println("Judul Buku 2 : ", buku2.judul)
fmt.Println("Genre Buku 2 : ", buku2.genre)
fmt.Println("Tahun Buku 2 : ", buku2.tahun)
fmt.Println("Rating Buku 2 : ", buku2.rating)
}


