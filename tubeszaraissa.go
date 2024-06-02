package main

import "fmt"

const NMAX int = 100

type mataKuliah struct {
	matkul string
	sks    int
}

type mahasiswa struct {
	nama                  string
	nim                   string
	jurusan               string
	matkul                []mataKuliah
	uts, uas, quiz, total int
	grade                 string
}

type TabMahasiswa [NMAX]mahasiswa

var data TabMahasiswa
var jumlahData int

func main() {
	var x, y int
	x = menu()
	for x != 1 && x != 4 {
		fmt.Println("Silahkan Menginput Data Terlebih Dahulu")
		x = menu()
	}
	for x != 4 {
		if x == 1 {
			y = menuInput()
			if y == 1 {
				inputDataMHS(&data, &jumlahData)
			} else if y == 2 {
				inputDataMatkul(&data, jumlahData)
			} else if y == 3 {
				inputNilai(&data, jumlahData)
			}
		} else if x == 2 {
			y = menuEdit()
			if y == 1 {
				hapus(&data, &jumlahData)
			} else if y == 2 {
				ganti(&data, jumlahData)
			} else if y == 3 {
				nambah(&data, &jumlahData)
			} else if y == 4 {
				hapusMatkul(&data, jumlahData)
			} else if y == 5 {
				editMatkul(&data, jumlahData)
			} else if y == 6 {
				tambahMatkul(&data, jumlahData)
			}
		} else if x == 3 {
			y = menuMenampilkan()
			if y == 1 {
				MenampilkanTerurutNilai(&data, jumlahData)
			} else if y == 2 {
				fmt.Println("Fitur ini belum diimplementasikan.")
			} else if y == 3 {
				MenampilkanBerdasarMatkul(&data, jumlahData)
			} else if y == 4 {
				TranskipNilai(&data, jumlahData)
			}
		}
		x = menu()
	}
	if x == 4 {
		fmt.Println("Terimakasih Sudah Menggunakan Aplikasi Data Mahasiswa")
	}
}

func menu() int {
	var n int
	fmt.Println("*******************************")
	fmt.Println("Aplikasi Data Nilai Mahasiswa")
	fmt.Println("*******************************")
	fmt.Println("Pilih Menu :")
	fmt.Println("1. Input")
	fmt.Println("2. Edit")
	fmt.Println("3. Menampilkan")
	fmt.Println("4. Keluar")
	fmt.Println("---------------------------------")
	fmt.Print("Pilih(1/2/3/4): ")
	fmt.Scanln(&n)
	return n
}

func menuInput() int {
	var n int
	fmt.Println("***************")
	fmt.Println("     INPUT     ")
	fmt.Println("***************")
	fmt.Println("1. Input Data Mahasiswa")
	fmt.Println("2. Input Data Mata Kuliah")
	fmt.Println("3. Input Data Nilai")
	fmt.Println("4. Kembali")
	fmt.Println("---------------")
	fmt.Print("Pilih(1/2/3/4): ")
	fmt.Scan(&n)
	return n
}

func menuEdit() int {
	var n int
	fmt.Println("***************")
	fmt.Println("     EDIT      ")
	fmt.Println("***************")
	fmt.Println("1. Hapus Data")
	fmt.Println("2. Ganti Data")
	fmt.Println("3. Tambah Data")
	fmt.Println("4. Hapus Mata Kuliah")
	fmt.Println("5. Ganti Mata Kuliah")
	fmt.Println("6. Tambah Mata Kuliah")
	fmt.Println("7. Kembali")
	fmt.Println("---------------")
	fmt.Print("Pilih(1/2/3/4/5/6/7): ")
	fmt.Scan(&n)
	return n
}

func menuMenampilkan() int {
	var n int
	fmt.Println("*******************************")
	fmt.Println("          Menampilkan          ")
	fmt.Println("*******************************")
	fmt.Println("1. Menampilkan Terurut Nilai")
	fmt.Println("2. Menampilkan Terurut SKS")
	fmt.Println("3. Menampilkan Berdasar Matkul")
	fmt.Println("4. Transkip Nilai (Grade)")
	fmt.Println("5. Kembali")
	fmt.Println("---------------")
	fmt.Print("Pilih(1/2/3/4/5): ")
	fmt.Scan(&n)
	return n
}

func inputDataMHS(data *TabMahasiswa, jumlahData *int) {
	var done bool

	for !done && *jumlahData < NMAX {
		fmt.Printf("Masukkan data mahasiswa ke-%d (ketik - pada NIM jika sudah selesai):\n", *jumlahData+1)
		fmt.Println("Untuk spasi gunakan underscore")
		fmt.Print("NIM: ")
		var nim string
		fmt.Scan(&nim)

		if nim == "-" {
			done = true
		} else {
			fmt.Print("Nama: ")
			var nama string
			fmt.Scan(&nama)

			fmt.Print("Jurusan: ")
			var jurusan string
			fmt.Scan(&jurusan)

			data[*jumlahData].nim = nim
			data[*jumlahData].nama = nama
			data[*jumlahData].jurusan = jurusan
			*jumlahData++
		}
	}
}

func inputDataMatkul(data *TabMahasiswa, jumlahData int) {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data mahasiswa. Silahkan input data mahasiswa terlebih dahulu.")
		return
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin diinputkan mata kuliah: ")
	var nim string
	fmt.Scan(&nim)

	index := -1
	for i := 0; i < jumlahData; i++ {
		if data[i].nim == nim {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
		return
	}

	fmt.Printf("Masukkan data mata kuliah untuk mahasiswa %s (%s):\n", data[index].nama, data[index].nim)
	var done bool
	var matkul string
	for !done {
		fmt.Print("Mata Kuliah: ")
		fmt.Scan(&matkul)

		if matkul == "-" {
			done = true
		} else {
			fmt.Print("SKS: ")
			var sks int
			fmt.Scan(&sks)

			data[index].matkul = append(data[index].matkul, mataKuliah{matkul, sks})
		}
	}
}

func inputNilai(data *TabMahasiswa, jumlahData int) {
	if jumlahData == 0 {
		fmt.Println("Tidak ada data mahasiswa. Silahkan input data mahasiswa terlebih dahulu.")
		return
	}

	for i := 0; i < jumlahData; i++ {
		if len(data[i].matkul) == 0 {
			fmt.Printf("Mahasiswa %s belum memiliki mata kuliah yang diinputkan.\n", data[i].nama)
			continue
		}

		fmt.Printf("Masukkan nilai untuk mahasiswa %s (%s):\n", data[i].nama, data[i].nim)
		for j := 0; j < len(data[i].matkul); j++ {
			fmt.Printf("Mata Kuliah %s (%s) SKS:\n", data[i].matkul[j].matkul, data[i].matkul[j].sks)
			fmt.Print("Nilai Quiz: ")
			fmt.Scan(&data[i].quiz)
			fmt.Print("Nilai UTS: ")
			fmt.Scan(&data[i].uts)
			fmt.Print("Nilai UAS: ")
			fmt.Scan(&data[i].uas)

			total := data[i].quiz + data[i].uts + data[i].uas
			data[i].total = total
			data[i].grade = calculateGrade(total)

			fmt.Printf("Nilai total: %d, Grade: %s\n", total, data[i].grade)
		}
	}
}

func hapus(data *TabMahasiswa, jumlahData *int) {
	fmt.Println("Hapus Data Mahasiswa")

	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang akan dihapus: ")
	fmt.Scan(&nim)

	found := false
	for i := 0; i < *jumlahData; i++ {
		if data[i].nim == nim {
			found = true
			// Menghapus data mahasiswa dengan menggeser data ke depan
			for j := i; j < *jumlahData-1; j++ {
				data[j] = data[j+1]
			}
			*jumlahData--
			fmt.Println("Data mahasiswa berhasil dihapus.")
			break
		}
	}

	if !found {
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}
}

func ganti(data *TabMahasiswa, jumlahData int) {
	fmt.Println("Ganti Data Mahasiswa")

	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang akan diganti: ")
	fmt.Scan(&nim)

	found := false
	for i := 0; i < jumlahData; i++ {
		if data[i].nim == nim {
			found = true
			fmt.Printf("Data mahasiswa dengan NIM %s ditemukan.\n", nim)
			fmt.Println("Data saat ini:")
			fmt.Printf("Nama: %s, Jurusan: %s\n", data[i].nama, data[i].jurusan)
			for _, mk := range data[i].matkul {
				fmt.Printf("Mata Kuliah: %s, SKS: %d\n", mk.matkul, mk.sks)
			}

			fmt.Println("Silakan edit data berikut:")
			fmt.Print("Nama: ")
			var nama string
			fmt.Scan(&nama)
			data[i].nama = nama

			fmt.Print("Jurusan: ")
			var jurusan string
			fmt.Scan(&jurusan)
			data[i].jurusan = jurusan

			// Edit mata kuliah
			data[i].matkul = nil // Reset mata kuliah
			fmt.Println("Masukkan data mata kuliah:")
			var done bool
			for !done {
				fmt.Print("Mata Kuliah (kosongkan untuk selesai): ")
				var matkul string
				fmt.Scan(&matkul)
				if matkul == "" {
					done = true
				} else {
					fmt.Print("SKS: ")
					var sks int
					fmt.Scan(&sks)
					data[i].matkul = append(data[i].matkul, mataKuliah{matkul, sks})
				}
			}

			// Edit nilai
			fmt.Println("Masukkan nilai:")
			for j := 0; j < len(data[i].matkul); j++ {
				fmt.Printf("Mata Kuliah %s (%d SKS):\n", data[i].matkul[j].matkul, data[i].matkul[j].sks)
				fmt.Print("Nilai Quiz: ")
				fmt.Scan(&data[i].quiz)
				fmt.Print("Nilai UTS: ")
				fmt.Scan(&data[i].uts)
				fmt.Print("Nilai UAS: ")
				fmt.Scan(&data[i].uas)

				total := data[i].quiz + data[i].uts + data[i].uas
				data[i].total = total
				data[i].grade = calculateGrade(total)
			}

			fmt.Println("Data mahasiswa berhasil diganti.")
			break
		}
	}

	if !found {
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}
}

func nambah(data *TabMahasiswa, jumlahData *int) {
	if *jumlahData < NMAX {
		fmt.Println("Tambah Data Mahasiswa")

		fmt.Printf("Masukkan data mahasiswa ke-%d:\n", *jumlahData+1)
		fmt.Print("NIM: ")
		var nim string
		fmt.Scan(&nim)

		fmt.Print("Nama: ")
		var nama string
		fmt.Scan(&nama)

		fmt.Print("Jurusan: ")
		var jurusan string
		fmt.Scan(&jurusan)

		data[*jumlahData].nim = nim
		data[*jumlahData].nama = nama
		data[*jumlahData].jurusan = jurusan

		// Input mata kuliah
		fmt.Println("Masukkan data mata kuliah:")
		var done bool
		for !done {
			fmt.Print("Mata Kuliah (kosongkan untuk selesai): ")
			var matkul string
			fmt.Scan(&matkul)
			if matkul == "" {
				done = true
			} else {
				fmt.Print("SKS: ")
				var sks int
				fmt.Scan(&sks)
				data[*jumlahData].matkul = append(data[*jumlahData].matkul, mataKuliah{matkul, sks})
			}
		}

		// Input nilai
		fmt.Println("Masukkan nilai:")
		for j := 0; j < len(data[*jumlahData].matkul); j++ {
			fmt.Printf("Mata Kuliah %s (%d SKS):\n", data[*jumlahData].matkul[j].matkul, data[*jumlahData].matkul[j].sks)
			fmt.Print("Nilai Quiz: ")
			fmt.Scan(&data[*jumlahData].quiz)
			fmt.Print("Nilai UTS: ")
			fmt.Scan(&data[*jumlahData].uts)
			fmt.Print("Nilai UAS: ")
			fmt.Scan(&data[*jumlahData].uas)

			total := data[*jumlahData].quiz + data[*jumlahData].uts + data[*jumlahData].uas
			data[*jumlahData].total = total
			data[*jumlahData].grade = calculateGrade(total)
		}

		*jumlahData++
		fmt.Println("Data mahasiswa berhasil ditambahkan.")
	} else {
		fmt.Println("Jumlah maksimum data mahasiswa telah tercapai.")
	}
}

func editMatkul(data *TabMahasiswa, jumlahData int) {
	fmt.Println("Edit Data Mata Kuliah")

	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin diedit mata kuliahnya: ")
	fmt.Scan(&nim)

	found := false
	for i := 0; i < jumlahData; i++ {
		if data[i].nim == nim {
			found = true
			fmt.Printf("Data mahasiswa dengan NIM %s ditemukan.\n", nim)
			fmt.Println("Mata Kuliah saat ini:")
			for j, mk := range data[i].matkul {
				fmt.Printf("%d. Mata Kuliah: %s, SKS: %d\n", j+1, mk.matkul, mk.sks)
			}

			fmt.Println("Silakan pilih mata kuliah yang ingin diedit (masukkan nomor mata kuliah):")
			var mkIndex int
			fmt.Scan(&mkIndex)
			if mkIndex <= 0 || mkIndex > len(data[i].matkul) {
				fmt.Println("Nomor mata kuliah tidak valid.")
				return
			}

			// Edit mata kuliah
			fmt.Println("Masukkan data baru untuk mata kuliah yang dipilih:")
			fmt.Print("Mata Kuliah: ")
			var matkul string
			fmt.Scan(&matkul)

			fmt.Print("SKS: ")
			var sks int
			fmt.Scan(&sks)

			// Mengganti mata kuliah yang sudah ada dengan yang baru
			data[i].matkul[mkIndex-1].matkul = matkul
			data[i].matkul[mkIndex-1].sks = sks

			fmt.Println("Data mata kuliah berhasil diubah.")
			break
		}
	}

	if !found {
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}
}

func hapusMatkul(data *TabMahasiswa, jumlahData int) {
	fmt.Println("Hapus Data Mata Kuliah")

	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus mata kuliahnya: ")
	fmt.Scan(&nim)

	found := false
	for i := 0; i < jumlahData; i++ {
		if data[i].nim == nim {
			found = true
			fmt.Printf("Data mahasiswa dengan NIM %s ditemukan.\n", nim)
			fmt.Println("Mata Kuliah saat ini:")
			for j, mk := range data[i].matkul {
				fmt.Printf("%d. Mata Kuliah: %s, SKS: %d\n", j+1, mk.matkul, mk.sks)
			}

			fmt.Println("Silakan pilih mata kuliah yang ingin dihapus (masukkan nomor mata kuliah):")
			var mkIndex int
			fmt.Scan(&mkIndex)
			if mkIndex <= 0 || mkIndex > len(data[i].matkul) {
				fmt.Println("Nomor mata kuliah tidak valid.")
				return
			}

			// Menghapus mata kuliah yang dipilih dengan menggeser data ke depan
			for k := mkIndex - 1; k < len(data[i].matkul)-1; k++ {
				data[i].matkul[k] = data[i].matkul[k+1]
			}
			// Mengurangi panjang slice untuk menghapus data terakhir yang double
			data[i].matkul = data[i].matkul[:len(data[i].matkul)-1]

			fmt.Println("Data mata kuliah berhasil dihapus.")
			break
		}
	}

	if !found {
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}
}

func tambahMatkul(data *TabMahasiswa, jumlahData int) {
	fmt.Println("Tambah Data Mata Kuliah")

	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin ditambahkan mata kuliahnya: ")
	fmt.Scan(&nim)

	found := false
	for i := 0; i < jumlahData; i++ {
		if data[i].nim == nim {
			found = true
			fmt.Printf("Data mahasiswa dengan NIM %s ditemukan.\n", nim)
			fmt.Println("Data saat ini:")
			fmt.Printf("Nama: %s, Jurusan: %s\n", data[i].nama, data[i].jurusan)
			fmt.Println("Mata Kuliah saat ini:")
			for j, mk := range data[i].matkul {
				fmt.Printf("%d. Mata Kuliah: %s, SKS: %d\n", j+1, mk.matkul, mk.sks)
			}

			fmt.Println("Masukkan data mata kuliah yang ingin ditambahkan:")
			fmt.Print("Mata Kuliah: ")
			var matkul string
			fmt.Scan(&matkul)

			fmt.Print("SKS: ")
			var sks int
			fmt.Scan(&sks)

			// Menambahkan mata kuliah baru pada akhir slice
			newMatkul := mataKuliah{matkul, sks}
			data[i].matkul = append(data[i].matkul, newMatkul)

			fmt.Println("Data mata kuliah berhasil ditambahkan.")
			break
		}
	}

	if !found {
		fmt.Printf("Data mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}
}

func MenampilkanTerurutNilai(data *TabMahasiswa, jumlahData int) {
	// Bubble sort untuk mengurutkan data berdasarkan nilai total secara menurun
	for i := 0; i < jumlahData-1; i++ {
		for j := 0; j < jumlahData-i-1; j++ {
			if data[j].total < data[j+1].total {
				// Tukar posisi data jika nilai total lebih kecil
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	// Menampilkan data yang sudah diurutkan
	fmt.Printf("%-10s %-20s %-15s %-6s %-6s %-6s %-6s\n", "NIM", "Nama", "Jurusan", "Quiz", "UTS", "UAS", "Total")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%-10s %-20s %-15s %-6d %-6d %-6d %-6d\n",
			data[i].nim, data[i].nama, data[i].jurusan,
			data[i].quiz, data[i].uts, data[i].uas, data[i].total)
	}
}

func MenampilkanBerdasarMatkul(data *TabMahasiswa, jumlahData int) {
	fmt.Print("Masukkan nama mata kuliah yang ingin dicari: ")
	var matkul string
	fmt.Scan(&matkul)

	fmt.Printf("%-10s %-20s %-15s %-15s %-6s\n", "NIM", "Nama", "Jurusan", "Mata Kuliah", "SKS")
	for i := 0; i < jumlahData; i++ {
		for _, mk := range data[i].matkul {
			if mk.matkul == matkul {
				fmt.Printf("%-10s %-20s %-15s %-15s %-6d\n", data[i].nim, data[i].nama, data[i].jurusan, mk.matkul, mk.sks)
			}
		}
	}
}

func TranskipNilai(data *TabMahasiswa, jumlahData int) {
	fmt.Printf("%-10s %-20s %-15s %-6s %-6s %-6s %-6s %-6s\n", "NIM", "Nama", "Jurusan", "Quiz", "UTS", "UAS", "Total", "Grade")
	for i := 0; i < jumlahData; i++ {
		total := data[i].quiz + data[i].uts + data[i].uas
		grade := calculateGrade(total)
		data[i].total = total
		data[i].grade = grade
		fmt.Printf("%-10s %-20s %-15s %-6d %-6d %-6d %-6d %-6s\n", data[i].nim, data[i].nama, data[i].jurusan, data[i].quiz, data[i].uts, data[i].uas, total, grade)
	}
}

func calculateGrade(total int) string {
	if total >= 240 {
		return "A"
	} else if total >= 210 {
		return "B"
	} else if total >= 150 {
		return "C"
	} else if total >= 120 {
		return "D"
	} else {
		return "E"
	}
}
