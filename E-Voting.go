package main

import "fmt"

type Kandidat struct {
	nourut   int
	nama     string
	visimisi string
	suara    int
}

var data [100]Kandidat
var n int

func tambah() {
	if n >= 100 {
		fmt.Println("Data penuh")
		return
	}

	fmt.Println("\nMASUKKAN KANDIDAT")

	fmt.Print("Nomor urut : ")
	fmt.Scan(&data[n].nourut)

	fmt.Print("Nama : ")
	fmt.Scan(&data[n].nama)

	fmt.Print("Visi misi : ")
	fmt.Scan(&data[n].visimisi)

	data[n].suara = 0
	n++

	fmt.Println("Data berhasil ditambah")
}

func tampil() {
	if n == 0 {
		fmt.Println("Data kosong")
		return
	}

	fmt.Println("\nDAFTAR KANDIDAT")

	for i := 0; i < n; i++ {
		fmt.Println("-------------------")
		fmt.Println("No Urut :", data[i].nourut)
		fmt.Println("Nama :", data[i].nama)
		fmt.Println("Visi :", data[i].visimisi)
		fmt.Println("Suara :", data[i].suara)
	}
}

func ubah() {
	var cari int

	fmt.Print("Nomor urut yang akan diubah : ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {
		if data[i].nourut == cari {

			fmt.Print("Nama baru : ")
			fmt.Scan(&data[i].nama)

			fmt.Print("Visi misi baru : ")
			fmt.Scan(&data[i].visimisi)

			fmt.Println("Data berhasil diubah")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func hapus() {
	var cari int

	fmt.Print("Nomor urut yang akan dihapus : ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {

		if data[i].nourut == cari {

			for j := i; j < n-1; j++ {
				data[j] = data[j+1]
			}

			n--
			fmt.Println("Data berhasil dihapus")
			return
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func voting() {
	var pilih int

	fmt.Print("Masukkan nomor urut kandidat : ")
	fmt.Scan(&pilih)

	for i := 0; i < n; i++ {

		if data[i].nourut == pilih {
			data[i].suara++
			fmt.Println("Vote berhasil")
			return
		}
	}

	fmt.Println("Kandidat tidak ditemukan")
}

func sequentialSearch() {
	var cari int
	var ketemu bool

	fmt.Print("Nomor urut yang dicari : ")
	fmt.Scan(&cari)

	for i := 0; i < n; i++ {

		if data[i].nourut == cari {

			fmt.Println("Data ditemukan")
			fmt.Println("Nama :", data[i].nama)
			fmt.Println("Visi :", data[i].visimisi)
			fmt.Println("Suara :", data[i].suara)

			ketemu = true
			break
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}
func binarySearch() {

	var temp [100]Kandidat

	for i := 0; i < n; i++ {
		temp[i] = data[i]
	}

	for i := 1; i < n; i++ {

		key := temp[i]
		j := i - 1

		for j >= 0 && temp[j].nourut > key.nourut {
			temp[j+1] = temp[j]
			j--
		}

		temp[j+1] = key
	}

	var cari int

	fmt.Print("Nomor urut yang dicari : ")
	fmt.Scan(&cari)

	kiri := 0
	kanan := n - 1

	for kiri <= kanan {

		tengah := (kiri + kanan) / 2

		if temp[tengah].nourut == cari {

			fmt.Println("Data ditemukan")
			fmt.Println("No Urut :", temp[tengah].nourut)
			fmt.Println("Nama :", temp[tengah].nama)
			fmt.Println("Visi :", temp[tengah].visimisi)
			fmt.Println("Suara :", temp[tengah].suara)

			return
		}

		if cari < temp[tengah].nourut {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Data tidak ditemukan")
}

func insertionSortNomorKecil() {

	var temp [100]Kandidat

	for i := 0; i < n; i++ {
		temp[i] = data[i]
	}

	for i := 1; i < n; i++ {

		key := temp[i]
		j := i - 1

		for j >= 0 && temp[j].nourut > key.nourut {
			temp[j+1] = temp[j]
			j--
		}

		temp[j+1] = key
	}

	fmt.Println("\nNomor urut dari yang Terkecil")

	tampilArray(temp)

}
func selectionSortNomorBesar() {

	var temp [100]Kandidat

	for i := 0; i < n; i++ {
		temp[i] = data[i]
	}

	for i := 0; i < n-1; i++ {

		max := i

		for j := i + 1; j < n; j++ {

			if temp[j].nourut > temp[max].nourut {
				max = j
			}
		}

		temp[i], temp[max] = temp[max], temp[i]
	}

	fmt.Println("\nNomor urut dari yang Terbesar")
	tampilArray(temp)
}
func selectionSortSuaraBesar() {

	var temp [100]Kandidat

	for i := 0; i < n; i++ {
		temp[i] = data[i]
	}

	for i := 0; i < n-1; i++ {

		max := i

		for j := i + 1; j < n; j++ {

			if temp[j].suara > temp[max].suara {
				max = j
			}
		}

		temp[i], temp[max] = temp[max], temp[i]
	}

	fmt.Println("\nSuara dari yang Terbesar")

	tampilArray(temp)
}
func insertionSortSuaraKecil() {

	var temp [100]Kandidat

	for i := 0; i < n; i++ {
		temp[i] = data[i]
	}

	for i := 1; i < n; i++ {

		key := temp[i]
		j := i - 1

		for j >= 0 && temp[j].suara > key.suara {
			temp[j+1] = temp[j]
			j--
		}

		temp[j+1] = key
	}

	fmt.Println("\nSuara dari yang Terkecil")
	tampilArray(temp)
}

func statistik() {

	total := 0

	for i := 0; i < n; i++ {
		total += data[i].suara
	}

	if total == 0 {
		fmt.Println("Belum ada suara masuk")
		return
	}

	fmt.Println("\nSTATISTIK")

	for i := 0; i < n; i++ {

		persen := float64(data[i].suara) * 100 / float64(total)

		fmt.Printf("%s : %.2f%%\n",
			data[i].nama,
			persen)
	}

	fmt.Println("Total suara :", total)
}
func pemenang() {

	if n == 0 {
		fmt.Println("Belum ada kandidat")
		return
	}

	pemenang := 0

	for i := 1; i < n; i++ {

		if data[i].suara > data[pemenang].suara {
			pemenang = i
		}
	}

	fmt.Println("\n===== HASIL AKHIR PEMILIHAN =====")
	fmt.Println("Pemenang : ", data[pemenang].nama)
	fmt.Println("No Urut  : ", data[pemenang].nourut)
	fmt.Println("Suara    : ", data[pemenang].suara)
}

func tampilArray(arr [100]Kandidat) {

	for i := 0; i < n; i++ {
		fmt.Println("-------------------")
		fmt.Println("No Urut :", arr[i].nourut)
		fmt.Println("Nama :", arr[i].nama)
		fmt.Println("Visi :", arr[i].visimisi)
		fmt.Println("Suara :", arr[i].suara)
	}
}

func menuUrutanDanPencarian() {
	var pilih int

	for {
		fmt.Println("\n===== URUTAN DAN PENCARIAN =====")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Println("3. Urut Nomor Urut (Kecil -> Besar)")
		fmt.Println("4. Urut Nomor Urut (Besar -> Kecil)")
		fmt.Println("5. Urut Suara (Besar -> Kecil)")
		fmt.Println("6. Urut Suara (Kecil -> Besar)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilih)

		switch pilih {

		case 1:
			sequentialSearch()

		case 2:
			binarySearch()

		case 3:
			insertionSortNomorKecil()

		case 4:
			selectionSortNomorBesar()

		case 5:
			selectionSortSuaraBesar()

		case 6:
			insertionSortSuaraKecil()

		case 0:
			return

		default:
			fmt.Println("Menu tidak ada")
		}
	}
}

func main() {
	var menu int

	for {

		fmt.Println("\n===== E VOTING =====")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Ubah Kandidat")
		fmt.Println("3. Hapus salah satu Kandidat")
		fmt.Println("4. Tampilkan Kandidat")
		fmt.Println("5. Voting suara")
		fmt.Println("6. Urutan dan Pencarian")
		fmt.Println("7. Statistik presentase perolehan suara")
		fmt.Println("0. Program selesai Tampilkan pemenang")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&menu)

		switch menu {

		case 1:
			tambah()

		case 2:
			ubah()

		case 3:
			hapus()

		case 4:
			tampil()

		case 5:
			voting()

		case 6:
			menuUrutanDanPencarian()

		case 7:
			statistik()

		case 0:
			pemenang()
			fmt.Println("\nProgram selesai")
			return

		default:
			fmt.Println("Menu tidak ada (pilih menu 0-7)")
		}
	}
}