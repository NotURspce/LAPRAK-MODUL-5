package main
import "fmt"

func main() {
	const (
		quiz float64 = 0.15
		tugas float64 = 0.15
		praktikum float64 = 0.2
		UTS float64 = 0.25
		UAS float64 = 0.25
	)

	type mahasiswa struct {
		NIM string
		NAMA string
		nilaiTugas float64
		nilaiQuiz float64
		nilaiPraktikum float64
		nilaiUAS float64
		nilaiUTS float64
		nilaiAkhir float64
	}

	var dataMahasiswa mahasiswa
	
	fmt.Scan(&dataMahasiswa.NAMA, &dataMahasiswa.NIM, &dataMahasiswa.nilaiTugas, 
		&dataMahasiswa.nilaiQuiz, &dataMahasiswa.nilaiPraktikum, &dataMahasiswa.nilaiUAS, 
		&dataMahasiswa.nilaiUTS)
	
	dataMahasiswa.nilaiAkhir = dataMahasiswa.nilaiTugas * tugas + dataMahasiswa.nilaiQuiz * quiz + 
	dataMahasiswa.nilaiPraktikum * praktikum + dataMahasiswa.nilaiUAS * UAS + 
	dataMahasiswa.nilaiUTS * UTS 

	fmt.Println("Data Mahasiswa : ")
	fmt.Println("Nama : ", dataMahasiswa.NAMA)
	fmt.Println("NIM : ", dataMahasiswa.NIM)
	fmt.Println("Nilai Tugas : ", dataMahasiswa.nilaiTugas)
	fmt.Println("Nilai Quiz : ", dataMahasiswa.nilaiQuiz)
	fmt.Println("Nilai Praktikum : ", dataMahasiswa.nilaiPraktikum)
	fmt.Println("Nilai UTS : ", dataMahasiswa.nilaiUTS)
	fmt.Println("Nilai UAS : ", dataMahasiswa.nilaiUAS)
	fmt.Println("Nilai Akhir : ", dataMahasiswa.nilaiAkhir)
}