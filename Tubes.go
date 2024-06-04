package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Hewan adalah tipe data struct untuk binatang
type Hewan struct {
	Species string
	Count   int
}

// Peserta struct for participant data
type Peserta struct {
	Nama     string
	Alamat   string
	Kelompok int
}

var pesertas []Peserta

func SelectHewan() {
	hewans := koutaSapi()
	for _, hewan := range hewans {
		fmt.Printf("Jumlah %s: %d\n", hewan.Species, hewan.Count)
	}
}

func SelectPeserta() {
	sort.Slice(pesertas, func(i, j int) bool {
		return pesertas[i].Nama < pesertas[j].Nama
	})
	fmt.Println("Daftar Peserta:")
	for _, peserta := range pesertas {
		fmt.Printf("Nama: %s\n", peserta.Nama)
		fmt.Printf("Alamat: %s\n", peserta.Alamat)
		fmt.Printf("Kelompok: %d\n\n", peserta.Kelompok)
	}
}

func SelectKelompok(kelompok int) {
	for _, peserta := range pesertas {
		if peserta.Kelompok == kelompok {
			fmt.Println("Nama Peserta:", peserta.Nama)
		}
	}
}

func PesertabyKelompok(kelompok int) []Peserta {
	var peserta []Peserta
	for _, p := range pesertas {
		if p.Kelompok == kelompok {
			peserta = append(peserta, p)
		}
	}
	return peserta
}

func SelectAdmin() {
	fmt.Println("Menu Admin:")
	fmt.Println("1. Tambah Peserta")
	fmt.Println("2. Ubah Peserta")
	fmt.Println("3. Hapus Peserta")
}

// TambahPeserta adds a new participant to the list
func TambahPeserta(p Peserta) {
	pesertas = append(pesertas, p)
	fmt.Println("Peserta berhasil ditambahkan:", p.Nama)
}

// UbahPeserta updates a participant's data by Nama
func UbahPeserta(nama string, p Peserta) {
	for i, peserta := range pesertas {
		if peserta.Nama == nama {
			pesertas[i].Nama = p.Nama
			pesertas[i].Alamat = p.Alamat
			pesertas[i].Kelompok = p.Kelompok
			fmt.Println("Peserta berhasil diubah:", p.Nama)
			return
		}
	}
	fmt.Println("Peserta tidak ditemukan")
}

func CariPeserta(nama string) int {
	for i, peserta := range pesertas {
		if peserta.Nama == nama {
			return i
		}
	}
	return -1
}

func HapusPesertaByName(nama string) {
	i := CariPeserta(nama)
	if i != -1 {
		pesertas = append(pesertas[:i], pesertas[i+1:]...)
		fmt.Println("Peserta berhasil dihapus")
	} else {
		fmt.Println("Peserta tidak ditemukan")
	}
}

// HapusPeserta deletes a participant by Name
func HapusPeserta(nama string) {
	HapusPesertaByName(nama)
}

// koutaSapi mengembalikan slice Hewan yang berisi jumlah sapi, kambing, dan domba
func koutaSapi() []Hewan {
	hewans := []Hewan{
		{"Sapi", 10},
		{"Kambing", 5},
		{"Domba", 6},
	}
	return hewans
}

func main() {
	for {
		var namaPeserta, alamatPeserta string
		var nomorKelompok int

		// Membuat generator angka acak lokal
		src := rand.NewSource(time.Now().UnixNano())
		r := rand.New(src)

		fmt.Println("*** ---------------------------- ***")
		fmt.Println("***                              ***")
		fmt.Println("***       APLIKASI KURBAN        ***")
		fmt.Println("***                              ***")
		fmt.Println("*** ---------------------------- ***")
		fmt.Println("*** Menu Utama ***")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Kouta Sapi")
		fmt.Println("3. Admin")
		fmt.Println("4. Kelompok")
		fmt.Println("5. Daftar Peserta")
		fmt.Println("6. Quit")
		fmt.Println("------------------")
		fmt.Print("Select: ")
		
		var Select int
		fmt.Scanln(&Select)

		switch Select {
		case 1:
			fmt.Println("*** Halaman Registrasi ***")
			fmt.Print("Nama Peserta: ")
			fmt.Scan(&namaPeserta)

			fmt.Print("Alamat Peserta: ")
			fmt.Scan(&alamatPeserta)

			// Menghasilkan nomor kelompok secara acak antara 1 sampai 10 (misalnya)
			nomorKelompok = r.Intn(10) + 1
			fmt.Printf("Nomor Kelompok: %d\n", nomorKelompok)

			newPeserta := Peserta{Nama: namaPeserta, Alamat: alamatPeserta, Kelompok: nomorKelompok}
			TambahPeserta(newPeserta)
		case 2:
			fmt.Println("Anda memilih Kouta Sapi")
			SelectHewan()
		case 3:
			fmt.Println("*** Menu Admin ***")
			fmt.Println("1. Tambah Peserta")
			fmt.Println("2. Ubah Peserta")
			fmt.Println("3. Hapus Peserta")
			fmt.Println("*** Menu Admin ***")
			fmt.Print("Select: ")
			fmt.Scanln(&Select)
			switch Select { // untuk masuk ke halaman admin
			case 1:
				var p Peserta
				fmt.Print("Masukkan Nama: ")
				fmt.Scan(&p.Nama)
				fmt.Print("Masukkan Alamat: ")
				fmt.Scan(&p.Alamat)
				fmt.Print("Masukkan Kelompok: ")
				fmt.Scan(&p.Kelompok)
				TambahPeserta(p)
			case 2:
				var p Peserta
				fmt.Print("Masukkan Nama yang akan diubah: ")
				var nama string
				fmt.Scan(&nama)
				fmt.Print("Masukkan Nama baru: ")
				fmt.Scan(&p.Nama)
				fmt.Print("Masukkan Alamat baru: ")
				fmt.Scan(&p.Alamat)
				fmt.Print("Masukkan Kelompok baru: ")
				fmt.Scan(&p.Kelompok)
				UbahPeserta(nama, p)
			case 3:
				var nama string
				fmt.Print("Masukkan nama peserta yang akan dihapus: ")
				fmt.Scan(&nama)
				HapusPeserta(nama)
			}
		case 4:
			fmt.Println("*** Halaman Kelompok ***")
			fmt.Println("1. Kelompok 1")
			fmt.Println("2. Kelompok 2")
			fmt.Println("3. Kelompok 3")
			fmt.Println("4. Kelompok 4")
			fmt.Println("5. Kelompok 5")
			fmt.Println("6. Kelompok 6")
			fmt.Println("7. Kelompok 7")
			fmt.Println("8. Kelompok 8")
			fmt.Println("9. Kelompok 9")
			fmt.Println("10. Kelompok 10")
			fmt.Println("-------------------")
			fmt.Print("Pilih Kelompok: ")
			var SelectKelompok int
			fmt.Scanln(&SelectKelompok)
			peserta := PesertabyKelompok(SelectKelompok)
			for _, p := range peserta {
				fmt.Printf("Nama: %s, Alamat: %s, Kelompok: %d\n", p.Nama, p.Alamat, p.Kelompok)
			}
		case 5:
			fmt.Print()
			SelectPeserta()
		case 6:
			fmt.Println("Arigatou gozaimasu!") // "Thank you" in Japanese
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
