package main

import (
	"fmt"
	"os"
	"strconv"
)

type siswa struct {
	absen int
	nama string
	alamat string
	pekerjaan string
	alasan string
}

var dataSiswa = []siswa {
	{1, "Leon S Kennedy", "Cianjur", "anggota RPD", "ingin bisa ngoding"},
	{2, "Ada Wong", "Ciamis", "agen intelijen Umbrella", "ingin ganti pekerjaan"},
	{3, "Jill Valentine", "Rancaekek", "anggota S.T.A.R.S.", "mencari kerjaan sampingan"},
	{4, "Chris Redfield", "Singaparna", "pendiri BSAA", "ingin membuat superApps"},
}

func tampilkan (noAbsen int){
	for _, dataSiswa := range dataSiswa {
		if dataSiswa.absen == noAbsen {
			fmt.Println("Data Siswa: ")
			fmt.Println("No Absen:", dataSiswa.absen)
			fmt.Println("Nama:", dataSiswa.nama)
			fmt.Println("Alamat:", dataSiswa.alamat)
			fmt.Println("Pekerjaan:", dataSiswa.pekerjaan)
			fmt.Println("Alasan memilih kelas Golang:", dataSiswa.alasan)
			return
		}
	}
	fmt.Println("teman dengan nomor absen", noAbsen, "tidak ditemukan")
}

func main(){
	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run tugas1.go <nomor_absen>")
		return
	}
	
	noAbsen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Nomor absen harus bilangan bulat")
		return
	}


	tampilkan(noAbsen)

}