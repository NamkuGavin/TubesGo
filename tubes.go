package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/google/uuid"
)

const nmax int = 200

type seagames struct {
	negara, id                                   string
	gold, silver, bronze, rank, tMedali, rankByT int
}

type pesertaSeagames [nmax]seagames

type arrayKosong [nmax]seagames

type arraySearch [nmax]seagames

func main() {
	var data pesertaSeagames
	var nData int
	data[0].tMedali = 0
	mainMenu(data, nData)
}

func mainHeader(data pesertaSeagames, n int) {
	fmt.Println("==================================================")
	fmt.Println("          Aplikasi Seagames Manager               ")
	fmt.Println("==================================================")
	if n != 0 {
		fmt.Println("data seagames:")
		cetakData(data, n)
		fmt.Println()
	} else {
		fmt.Println("data seagames:")
		fmt.Println("Belum ada data peserta Seagames yang dimasukkan. Silahkan lakukan input data.")
		fmt.Println()
	}
}

func searchHeader(data arraySearch, n int) {
	fmt.Println("==================================================")
	fmt.Println("          Aplikasi Seagames Manager               ")
	fmt.Println("==================================================")
	if n != 0 {
		fmt.Println("search seagames:")
		cetakSearch(data, n)
		fmt.Println()
	} else {
		fmt.Println("search seagames:")
		fmt.Println("Belum ada data pencarian yang dimasukkan.")
		fmt.Println()
	}
}

func popUp(kalimat string) {
	fmt.Println(kalimat)
	fmt.Scanln()
	fmt.Scanln()
}

func searchData(data pesertaSeagames, n int, dataSearch *arraySearch, nCari *int) {
	clearScreen()
	var pilihMainMenu string
	fmt.Println()
	searchHeader(*dataSearch, *nCari)
	fmt.Println("============ MENU PENCARIAN ============")
	fmt.Println("|1. Alphabet                           |")
	fmt.Println("|2. Nama                               |")
	fmt.Println("|3. batas Rank                         |")
	fmt.Println("|4. Perolehan Medali                   |")
	fmt.Println("|5. <- Kembali                         |")
	fmt.Println("========================================")
	fmt.Print("Pilih opsi (1/2/3/4/5): ")
	fmt.Scan(&pilihMainMenu)

	if pilihMainMenu == "1" {
		searchAlphabet(data, dataSearch, n, nCari)
	} else if pilihMainMenu == "2" {
		searchName(data, dataSearch, n, nCari)
	} else if pilihMainMenu == "3" {
		searchRank(data, dataSearch, n, nCari)
	} else if pilihMainMenu == "4" {
		searchMedal(data, dataSearch, n, nCari)
	} else if pilihMainMenu == "5" {
		mainMenu(data, n)
	} else {
		popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
		searchData(data, n, dataSearch, nCari)
	}
}

func searchMedal(data pesertaSeagames, dataSearch *arraySearch, n int, nCari *int) {
	var dataKosong arrayKosong
	var g, s, b, index int
	var medaliApa string
	for i := 0; i < n; i++ {
		dataSearch[i] = dataKosong[i]
	}
	*nCari = 0
	fmt.Scanln()
	fmt.Println()
	fmt.Println("============ MENU PENCARIAN ============")
	fmt.Println("|1. Gold                               |")
	fmt.Println("|2. Silver                             |")
	fmt.Println("|3. Bronze                             |")
	fmt.Println("|4. Semua medali                       |")
	fmt.Println("========================================")
	fmt.Print("Pilih opsi (1/2/3/4): ")
	fmt.Scan(&medaliApa)

	if medaliApa == "4" {
		fmt.Println("Masukkan medali yang anda cari!")
		fmt.Print("dengan format <Gold> <Silver> <Bronze>: ")
		fmt.Scan(&g, &s, &b)
		for i := 0; i < n; i++ {
			if g == data[i].gold && s == data[i].silver && b == data[i].bronze {
				dataSearch[index] = data[i]
				index++
				*nCari++
			}
		}
		if *nCari != 0 {
			sortingMedaliSearch(dataSearch, n)
			popUp("Pencarian berhasil!. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		} else {
			popUp("Pencarian gagal!. Tidak ada data gold, silver, dan bronze yang anda masukkan. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		}
	} else if medaliApa == "3" {
		fmt.Println("Masukkan medali yang anda cari!")
		fmt.Print("dengan format <Bronze>: ")
		fmt.Scan(&b)
		for i := 0; i < n; i++ {
			if b == data[i].bronze {
				dataSearch[index] = data[i]
				index++
				*nCari++
			}
		}
		if *nCari != 0 {
			sortingMedaliSearch(dataSearch, n)
			popUp("Pencarian berhasil!. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		} else {
			popUp("Pencarian gagal!. Tidak ada data bronze yang anda masukkan. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		}
	} else if medaliApa == "2" {
		fmt.Println("Masukkan medali yang anda cari!")
		fmt.Print("dengan format <Silver>: ")
		fmt.Scan(&s)
		for i := 0; i < n; i++ {
			if s == data[i].silver {
				dataSearch[index] = data[i]
				index++
				*nCari++
			}
		}
		if *nCari != 0 {
			sortingMedaliSearch(dataSearch, n)
			popUp("Pencarian berhasil!. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		} else {
			popUp("Pencarian gagal!. Tidak ada data silver yang anda masukkan. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		}
	} else if medaliApa == "1" {
		fmt.Println("Masukkan medali yang anda cari!")
		fmt.Print("dengan format <Gold>: ")
		fmt.Scan(&g)
		for i := 0; i < n; i++ {
			if g == data[i].gold {
				dataSearch[index] = data[i]
				index++
				*nCari++
			}
		}
		if *nCari != 0 {
			sortingMedaliSearch(dataSearch, n)
			popUp("Pencarian berhasil!. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		} else {
			popUp("Pencarian gagal!. Tidak ada data gold yang anda masukkan. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		}
	} else {
		fmt.Println("Silahkan input pilihan yang benar.")
		searchMedal(data, dataSearch, n, nCari)
	}
}

func searchRank(data pesertaSeagames, dataSearch *arraySearch, n int, nCari *int) {
	var index, min, max int
	var dataKosong arrayKosong
	for i := 0; i < n; i++ {
		dataSearch[i] = dataKosong[i]
	}
	*nCari = 0
	fmt.Scanln()
	fmt.Println("Masukkan batas Rank minimum dan maksimum!")
	fmt.Print("Minimum: ")
	fmt.Scan(&min)
	fmt.Print("Maksimum: ")
	fmt.Scan(&max)
	for i := 0; i < n; i++ {
		if data[i].rank >= min && data[i].rank <= max {
			dataSearch[index] = data[i]
			index++
			*nCari++
		}
	}
	if *nCari != 0 {
		sortingMedaliSearch(dataSearch, n)
		popUp("Pembatasan berhasil!. Tekan 'enter' untuk melanjutkan")
		searchData(data, n, dataSearch, nCari)
	} else {
		popUp("Pembatasan gagal!. Nilai batas Rank yang anda masukkan tidak valid. Tekan 'enter' untuk melanjutkan")
		searchData(data, n, dataSearch, nCari)
	}
}

func searchAlphabet(data pesertaSeagames, dataSearch *arraySearch, n int, nCari *int) {
	var alphabet byte
	var index int
	var dataKosong arrayKosong
	for i := 0; i < n; i++ {
		dataSearch[i] = dataKosong[i]
	}
	*nCari = 0
	fmt.Scanln()
	fmt.Print("Masukkan alphabet (A/B/C/D/E/F/G/H/I/J/K/L/M/N/O/P/Q/R/S/T/U/V/W/X/Y/Z): ")
	fmt.Scanf("%c", &alphabet)
	if (alphabet >= 65 && alphabet <= 90) || (alphabet >= 97 && alphabet <= 122) {
		if alphabet >= 97 && alphabet <= 122 {
			alphabet -= 32
		}
		for i := 0; i < n; i++ {
			if alphabet == data[i].negara[0] {
				dataSearch[index] = data[i]
				index++
				*nCari++
			}
		}
		if *nCari != 0 {
			popUp("Pencarian berhasil!. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		} else {
			popUp("Pencarian gagal!. Negara dengan alphabet ini tidak ada. Tekan 'enter' untuk melanjutkan")
			searchData(data, n, dataSearch, nCari)
		}
	} else {
		fmt.Println("Silahkan input alphabet yang benar.")
		searchAlphabet(data, dataSearch, n, nCari)
	}

}

func contains(namaNegara, diCari string) bool {
	var contain bool
	for i := 0; i <= len(namaNegara)-len(diCari); i++ {
		if namaNegara[i:i+len(diCari)] == diCari {
			contain = true
		}
	}
	return contain
}

func searchName(data pesertaSeagames, dataSearch *arraySearch, n int, nCari *int) {
	var country string
	var index int
	var dataKosong arrayKosong
	for i := 0; i < n; i++ {
		dataSearch[i] = dataKosong[i]
	}
	*nCari = 0
	fmt.Scanln()
	fmt.Print("Masukkan nama negara: ")
	fmt.Scan(&country)
	for i := 0; i < n; i++ {
		if contains(data[i].negara, country) {
			dataSearch[index] = data[i]
			index++
			*nCari++
		}
	}
	if *nCari != 0 {
		popUp("Pencarian berhasil!. Tekan 'enter' untuk melanjutkan")
		searchData(data, n, dataSearch, nCari)
	} else {
		popUp("Pencarian gagal!. Tidak ada nama negara yang dicari. Tekan 'enter' untuk melanjutkan")
		searchData(data, n, dataSearch, nCari)
	}
}

func mainMenu(data pesertaSeagames, n int) {
	clearScreen()
	var pilihMainMenu, pilih string
	var dataSearch arraySearch
	var nSearch int
	fmt.Println()
	mainHeader(data, n)
	fmt.Println("================= MENU =================")
	fmt.Println("|1. Kustomisasi data peserta Seagames  |")
	fmt.Println("|2. Ubah sorting                       |")
	fmt.Println("|3. Cari peserta                       |")
	fmt.Println("|4. Keluar                             |")
	fmt.Println("========================================")
	fmt.Print("Pilih opsi (1/2/3/4): ")
	fmt.Scan(&pilihMainMenu)

	if pilihMainMenu == "1" {
		kustomisasiData(&data, &n)
	} else if pilihMainMenu == "2" {
		if n != 0 {
			opsiSorting(data, n)
		} else {
			popUp("Belum ada data peserta Seagames yang dimasukkan. Tekan 'enter' untuk melanjutkan")
			mainMenu(data, n)
		}
	} else if pilihMainMenu == "3" {
		if n != 0 {
			searchData(data, n, &dataSearch, &nSearch)
		} else {
			popUp("Belum ada data peserta Seagames yang dimasukkan. Tekan 'enter' untuk melanjutkan")
			mainMenu(data, n)
		}
	} else if pilihMainMenu == "4" {
		fmt.Print("Apakah anda ingin keluar (Y/N)? ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			fmt.Println("Keluar. Terima kasih telah mencoba aplikasi ini.")
		} else if pilih == "N" || pilih == "n" {
			mainMenu(data, n)
		} else {
			popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
			mainMenu(data, n)
		}
	} else {
		popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
		mainMenu(data, n)
	}
}

func kustomisasiData(data *pesertaSeagames, n *int) {
	clearScreen()
	var pilihCustom string
	fmt.Println()
	mainHeader(*data, *n)
	fmt.Println("=========== MENU KUSTOMASI DATA ===========")
	fmt.Println("|1. Tambahkan data peserta                |")
	fmt.Println("|2. Edit data peserta                     |")
	fmt.Println("|3. Delete data peserta                   |")
	fmt.Println("|4. <- kembali                            |")
	fmt.Println("===========================================")
	fmt.Print("Pilih opsi (1/2/3/4): ")
	fmt.Scan(&pilihCustom)

	if pilihCustom == "1" {
		addData(data, n)
	} else if pilihCustom == "2" {
		editData(data, *n)
	} else if pilihCustom == "3" {
		deleteData(data, *n)
	} else if pilihCustom == "4" {
		mainMenu(*data, *n)
	} else {
		popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
		kustomisasiData(data, n)
	}
}

func editData(data *pesertaSeagames, n int) {
	if n != 0 {
		clearScreen()
		var pilihEdit string
		fmt.Println()
		mainHeader(*data, n)
		fmt.Println("============ MENU KUSTOMASI DATA ===========")
		fmt.Println("|1. Edit nama peserta Negara               |")
		fmt.Println("|2. Edit data Medali peserta               |")
		fmt.Println("|3. <- kembali                             |")
		fmt.Println("============================================")
		fmt.Print("Pilih opsi (1/2/3): ")
		fmt.Scan(&pilihEdit)

		if pilihEdit == "1" {
			editNegara(data, n)
		} else if pilihEdit == "2" {
			editMedali(data, n)
		} else if pilihEdit == "3" {
			kustomisasiData(data, &n)
		} else {
			popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
			editData(data, n)
		}
	} else {
		popUp("Data peserta Seagames kosong. Pengeditan tidak dapat dilakukan. Tekan 'enter' untuk melanjutkan")
		kustomisasiData(data, &n)
	}
}

func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

func cetakData(data pesertaSeagames, n int) {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Printf("|%-4v|%-5v|%-30v|%-5v|%-7v|%-7v|%-6v|%-14v|\n", "Id", "Rank", "Team/NOC", "Gold", "Silver", "Bronze", "Total", "Rank by Total")
	fmt.Println("---------------------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("|%v|%v|%-30v|%v|%v|%v|%v|%v|\n", center(data[i].id, 4), center(strconv.Itoa(data[i].rank), 5), data[i].negara, center(strconv.Itoa(data[i].gold), 5), center(strconv.Itoa(data[i].silver), 7), center(strconv.Itoa(data[i].bronze), 7), center(strconv.Itoa(data[i].tMedali), 6), center(strconv.Itoa(data[i].rankByT), 14))
		fmt.Println("---------------------------------------------------------------------------------------")
	}
}

func cetakSearch(data arraySearch, n int) {
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Printf("|%-4v|%-5v|%-30v|%-5v|%-7v|%-7v|%-6v|%-14v|\n", "Id", "Rank", "Team/NOC", "Gold", "Silver", "Bronze", "Total", "Rank by Total")
	fmt.Println("---------------------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("|%v|%v|%-30v|%v|%v|%v|%v|%v|\n", center(data[i].id, 4), center(strconv.Itoa(data[i].rank), 5), data[i].negara, center(strconv.Itoa(data[i].gold), 5), center(strconv.Itoa(data[i].silver), 7), center(strconv.Itoa(data[i].bronze), 7), center(strconv.Itoa(data[i].tMedali), 6), center(strconv.Itoa(data[i].rankByT), 14))
		fmt.Println("---------------------------------------------------------------------------------------")
	}
}

func opsiSorting(data pesertaSeagames, n int) {
	clearScreen()
	var pilihSort string
	mainHeader(data, n)
	fmt.Println("================== MENU ==================")
	fmt.Println("|1. Sorting berdasarkan 'Medali'         |")
	fmt.Println("|   a. ASC (Ascending)                   |")
	fmt.Println("|   b. DESC (Descending)                 |")
	fmt.Println("|2. Sorting berdasarkan 'Rank by Total'  |")
	fmt.Println("|   a. ASC (Ascending)                   |")
	fmt.Println("|   b. DESC (Descending)                 |")
	fmt.Println("|3. Sorting berdasarkan 'Team/NOC'       |")
	fmt.Println("|   a. ASC (Ascending)                   |")
	fmt.Println("|   b. DESC (Descending)                 |")
	fmt.Println("|4. <- kembali                           |")
	fmt.Println("==========================================")
	fmt.Print("Pilih opsi (1a/1b/2a/2b/3a/3b/4): ")
	fmt.Scan(&pilihSort)

	if pilihSort == "1a" || pilihSort == "1b" {
		sortingOptMedali(&data, n, pilihSort)
		opsiSorting(data, n)
	} else if pilihSort == "2a" || pilihSort == "2b" {
		sortingRankbyT(&data, n, pilihSort)
		opsiSorting(data, n)
	} else if pilihSort == "3a" || pilihSort == "3b" {
		sortingNegara(&data, n, pilihSort)
		opsiSorting(data, n)
	} else if pilihSort == "4" {
		mainMenu(data, n)
	} else {
		popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
		opsiSorting(data, n)
	}
}

func addData(data *pesertaSeagames, n *int) {
	var dataKosong arrayKosong
	var adaKah, dataDuplicate bool
	var tambah int
	var negara, pilih string
	var gold, silver, bronze int
	adaKah = checkingData(*n)
	if adaKah {
		fmt.Print("Terdapat data yang sudah tersimpan. Apakah anda ingin menambahkan data (Y/N)? ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			fmt.Println("Silahkan masukkan data peserta tambahan!")
			fmt.Print("Ingin menambahkan berapa data: ")
			fmt.Scan(&tambah)
			if *n+tambah > nmax {
				fmt.Println("Mohon maaf. Data tidak bisa lebih dari 200. Silahkan input lagi")
				addData(data, n)
			} else {
				fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
				fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
				for i := *n; i < *n+tambah; i++ {
					fmt.Printf("%d. ", i+1)
					fmt.Scan(&negara, &gold, &silver, &bronze)
					id := uuid.New()
					data[i].id = id.String()[:3]
					data[i].negara = negara
					if data[i].negara[0] >= 97 && data[i].negara[0] <= 122 {
						x := data[i].negara[0]
						x -= 32
						data[i].negara = string(x) + data[i].negara[1:]
					}
					data[i].gold = gold
					data[i].silver = silver
					data[i].bronze = bronze
					dataDuplicate = searchDuplicate(*data, *data, i+1)
					for dataDuplicate {
						data[i] = dataKosong[i]
						fmt.Printf("Negara %s sudah ada di dalam data. Data peserta negara tidak boleh sama. Silahkan coba lagi.\n", negara)
						fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
						fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
						fmt.Printf("%d. ", i+1)
						fmt.Scan(&negara, &gold, &silver, &bronze)
						id := uuid.New()
						data[i].id = id.String()[:3]
						data[i].negara = negara
						if data[i].negara[0] >= 97 && data[i].negara[0] <= 122 {
							x := data[i].negara[0]
							x -= 32
							data[i].negara = string(x) + data[i].negara[1:]
						}
						data[i].gold = gold
						data[i].silver = silver
						data[i].bronze = bronze
						dataDuplicate = searchDuplicate(*data, *data, i+1)
					}
				}
				*n += tambah
				if data[*n-1].tMedali == 0 {
					sortingTotal(data, *n)
				}
				sortingMedali(data, *n)
				popUp("Pemasukan data selesai!. Tekan 'enter' untuk melanjutkan")
				mainMenu(*data, *n)
			}
		} else if pilih == "N" || pilih == "n" {
			kustomisasiData(data, n)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			addData(data, n)
		}
	} else {
		fmt.Print("Belum ada data yang tersimpan. Apakah anda ingin menambahkan data (Y/N)? ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			fmt.Print("Ingin memasukkan berapa data: ")
			fmt.Scan(&*n)
			if *n > nmax {
				fmt.Println("Mohon maaf. Data tidak bisa menampung lebih dari 200. Silahkan input lagi")
				*n = 0
				addData(data, n)
			} else {
				fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
				fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
				for i := 0; i < *n; i++ {
					fmt.Printf("%d. ", i+1)
					fmt.Scan(&negara, &gold, &silver, &bronze)
					id := uuid.New()
					data[i].id = id.String()[:3]
					data[i].negara = negara
					if data[i].negara[0] >= 97 && data[i].negara[0] <= 122 {
						x := data[i].negara[0]
						x -= 32
						data[i].negara = string(x) + data[i].negara[1:]
					}
					data[i].gold = gold
					data[i].silver = silver
					data[i].bronze = bronze
					dataDuplicate = searchDuplicate(*data, *data, i+1)
					for dataDuplicate {
						data[i] = dataKosong[i]
						fmt.Printf("Data Negara %s sudah ada di dalam data. Data peserta negara tidak boleh sama. Silahkan coba lagi.\n", negara)
						fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
						fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
						fmt.Printf("%d. ", i+1)
						fmt.Scan(&negara, &gold, &silver, &bronze)
						id := uuid.New()
						data[i].id = id.String()[:3]
						data[i].negara = negara
						if data[i].negara[0] >= 97 && data[i].negara[0] <= 122 {
							x := data[i].negara[0]
							x -= 32
							data[i].negara = string(x) + data[i].negara[1:]
						}
						data[i].gold = gold
						data[i].silver = silver
						data[i].bronze = bronze
						dataDuplicate = searchDuplicate(*data, *data, i+1)
					}
				}
				if data[*n-1].tMedali == 0 {
					sortingTotal(data, *n)
				}
				sortingMedali(data, *n)
				popUp("Pemasukan data selesai!. Tekan 'enter' untuk melanjutkan")
				mainMenu(*data, *n)
			}
		} else if pilih == "N" || pilih == "n" {
			kustomisasiData(data, n)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			addData(data, n)
		}
	}
}

func checkingData(n int) bool {
	var ada bool
	if n != 0 {
		ada = true
	}
	return ada
}

func sortingTotal(data *pesertaSeagames, n int) {
	var idx, temp int
	var tempData seagames
	for i := 0; i < n; i++ {
		data[i].tMedali = data[i].gold + data[i].silver + data[i].bronze
	}
	for pass := 1; pass < n; pass++ {
		idx = pass
		temp = data[pass].tMedali
		tempData = data[pass]
		for idx > 0 && data[idx-1].tMedali < temp {
			data[idx] = data[idx-1]
			idx -= 1
		}
		data[idx] = tempData
	}
	for i := 0; i < n; i++ {
		data[i].rankByT = i + 1
	}
}

func sortingNegara(data *pesertaSeagames, n int, pilihanSort string) {
	var idx int
	var temp string
	var tempData seagames
	if pilihanSort == "3a" {
		for pass := 1; pass < n; pass++ {
			idx = pass
			temp = data[pass].negara
			tempData = data[pass]
			for idx > 0 && data[idx-1].negara > temp {
				data[idx] = data[idx-1]
				idx -= 1
			}
			data[idx] = tempData
		}
	} else if pilihanSort == "3b" {
		for pass := 1; pass < n; pass++ {
			idx = pass
			temp = data[pass].negara
			tempData = data[pass]
			for idx > 0 && data[idx-1].negara < temp {
				data[idx] = data[idx-1]
				idx -= 1
			}
			data[idx] = tempData
		}
	}
}

func sortingRankbyT(data *pesertaSeagames, n int, pilihanSort string) {
	var idx int
	var temp seagames
	if pilihanSort == "2a" {
		for pass := 1; pass < n; pass++ {
			idx = pass - 1
			for j := pass; j < n; j++ {
				if data[idx].rankByT > data[j].rankByT {
					idx = j
				}
			}
			temp = data[idx]
			data[idx] = data[pass-1]
			data[pass-1] = temp
		}
	} else if pilihanSort == "2b" {
		for pass := 1; pass < n; pass++ {
			idx = pass - 1
			for j := pass; j < n; j++ {
				if data[idx].rankByT < data[j].rankByT {
					idx = j
				}
			}
			temp = data[idx]
			data[idx] = data[pass-1]
			data[pass-1] = temp
		}
	}
}

func sortingOptMedali(data *pesertaSeagames, n int, pilihanSort string) {
	var idx int
	var tempData, temp seagames
	if pilihanSort == "1a" {
		for pass := 1; pass < n; pass++ {
			idx = pass
			temp = data[pass]
			tempData = data[pass]
			for idx > 0 && (data[idx-1].gold < temp.gold || (data[idx-1].gold == temp.gold && data[idx-1].silver < temp.silver) || (data[idx-1].gold == temp.gold && data[idx-1].silver == temp.silver && data[idx-1].bronze < temp.bronze) || (data[idx-1].gold == temp.gold && data[idx-1].silver == temp.silver && data[idx-1].bronze == temp.bronze && data[idx-1].rank > temp.rank)) {
				data[idx] = data[idx-1]
				idx -= 1
			}
			data[idx] = tempData
		}
	} else if pilihanSort == "1b" {
		for pass := 1; pass < n; pass++ {
			idx = pass - 1
			for j := pass; j < n; j++ {
				if data[idx].gold > data[j].gold || (data[idx].gold == data[j].gold && data[idx].silver > data[j].silver) || (data[idx].gold == data[j].gold && data[idx].silver == data[j].silver && data[idx].bronze > data[j].bronze) || (data[idx].gold == data[j].gold && data[idx].silver == data[j].silver && data[idx].bronze == data[j].bronze && data[idx].rank < data[j].rank) {
					idx = j
				}
			}
			temp = data[idx]
			data[idx] = data[pass-1]
			data[pass-1] = temp
		}
	}
}

func sortingMedaliSearch(data *arraySearch, n int) {
	var idx int
	var tempData, temp seagames
	for pass := 1; pass < n; pass++ {
		idx = pass
		temp = data[pass]
		tempData = data[pass]
		for idx > 0 && (data[idx-1].gold < temp.gold || (data[idx-1].gold == temp.gold && data[idx-1].silver < temp.silver) || (data[idx-1].gold == temp.gold && data[idx-1].silver == temp.silver && data[idx-1].bronze < temp.bronze) || (data[idx-1].gold == temp.gold && data[idx-1].silver == temp.silver && data[idx-1].bronze == temp.bronze && data[idx-1].rank > temp.rank)) {
			data[idx] = data[idx-1]
			idx -= 1
		}
		data[idx] = tempData
	}
}

func sortingMedali(data *pesertaSeagames, n int) {
	var idx int
	var tempData, temp seagames
	for pass := 1; pass < n; pass++ {
		idx = pass
		temp = data[pass]
		tempData = data[pass]
		for idx > 0 && (data[idx-1].gold < temp.gold || (data[idx-1].gold == temp.gold && data[idx-1].silver < temp.silver) || (data[idx-1].gold == temp.gold && data[idx-1].silver == temp.silver && data[idx-1].bronze < temp.bronze)) {
			data[idx] = data[idx-1]
			idx -= 1
		}
		data[idx] = tempData
	}
	for i := 0; i < n; i++ {
		data[i].rank = i + 1
	}
}

func searchDuplicate(data, tampungan pesertaSeagames, n int) bool {
	var cek int
	var checking string
	for k := 0; k < n; k++ {
		if cek < 2 {
			cek = 0
			checking = data[k].negara
			for m := 0; m < n; m++ {
				if checking == tampungan[m].negara {
					cek += 1
				}
			}
		}
	}
	return cek >= 2
}

func editNegara(data *pesertaSeagames, n int) {
	var negaraBaru, pilih, cariApa string
	var ketemu, dataDuplicate bool
	var indexDi int
	var titip pesertaSeagames
	for i := 0; i < n; i++ {
		titip[i] = data[i]
	}
	fmt.Print("Ingin mengedit Id Negara: ")
	fmt.Scan(&cariApa)
	searchingNegara(*data, n, cariApa, &ketemu, &indexDi)
	if ketemu {
		fmt.Print("Id Negara ditemukan. Ubah menjadi: ")
		fmt.Scan(&negaraBaru)
		fmt.Print("Apakah anda ingin mengedit ini (Y/N)? ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			data[indexDi].negara = negaraBaru
			dataDuplicate = searchDuplicate(*data, *data, n)
			if dataDuplicate {
				for i := 0; i < n; i++ {
					data[i].negara = titip[i].negara
				}
				fmt.Printf("Data Negara %s sudah ada di dalam data. Data peserta negara tidak boleh sama. Tekan 'enter' untuk melanjutkan\n", negaraBaru)
				fmt.Scanln()
				fmt.Scanln()
				editNegara(data, n)
			} else {
				popUp("Pengeditan selesai. Tekan 'enter' untuk melanjutkan")
				mainMenu(*data, n)
			}
		} else if pilih == "N" || pilih == "n" {
			editData(data, n)
		} else {
			popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
			editData(data, n)
		}
	} else {
		popUp("Id Negara tidak ditemukan. Silahkan cek kembali. Tekan 'enter' untuk melanjutkan")
		editData(data, n)
	}
}

func searchingNegara(data pesertaSeagames, n int, cari string, ketemu *bool, indexDi *int) {
	for i := 0; i < n; i++ {
		if cari == data[i].id {
			*ketemu = true
			*indexDi = i
		}
	}
}

func editMedali(data *pesertaSeagames, n int) {
	var pilih, cariApa string
	var ketemu bool
	var indexDi, gold, silver, bronze int
	fmt.Print("Ingin mengedit medali dari Id Negara: ")
	fmt.Scan(&cariApa)
	searchingNegara(*data, n, cariApa, &ketemu, &indexDi)
	if ketemu {
		fmt.Println("Silahkan masukkan data Medali nya. Masukkan -1 jika tidak ingin mengedit medali tertentu")
		fmt.Println("dengan format: <Gold> <Silver> <Bronze>")
		fmt.Scan(&gold, &silver, &bronze)
		fmt.Print("Apakah anda ingin mengedit ini (Y/N)? ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			if gold != -1 {
				data[indexDi].gold = gold
			}
			if silver != -1 {
				data[indexDi].silver = silver
			}
			if bronze != -1 {
				data[indexDi].bronze = bronze
			}
			popUp("Pengeditan selesai. Tekan 'enter' untuk melanjutkan")
			sortingTotal(data, n)
			sortingMedali(data, n)
			mainMenu(*data, n)
		} else if pilih == "N" || pilih == "n" {
			editData(data, n)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editMedali(data, n)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		editMedali(data, n)
	}
}

func deleteData(data *pesertaSeagames, n int) {
	if n != 0 {
		clearScreen()
		var pilihEdit string
		fmt.Println()
		mainHeader(*data, n)
		fmt.Println("=========== MENU KUSTOMASI DATA ===========")
		fmt.Println("1. Delete nama peserta Negara")
		fmt.Println("2. Delete data Medali peserta")
		fmt.Println("3. <- kembali")
		fmt.Println("===========================================")
		fmt.Print("Pilih opsi (1/2/3): ")
		fmt.Scan(&pilihEdit)

		if pilihEdit == "1" {
			deleteNegara(data, n)
		} else if pilihEdit == "2" {
			deleteMedali(data, n)
		} else if pilihEdit == "3" {
			kustomisasiData(data, &n)
		} else {
			popUp("Silahkan input pilihan yang benar. tekan 'enter' untuk melanjutkan")
			editData(data, n)
		}
	} else {
		popUp("Data peserta Seagames kosong. Pengeditan tidak dapat dilakukan. Tekan 'enter' untuk melanjutkan")
		kustomisasiData(data, &n)
	}
}

func deleteMedali(data *pesertaSeagames, n int) {
	var medaliApa, pilih, cariApa string
	var ketemu bool
	var indexDi int
	fmt.Print("Ingin menghapus medali dari Id Negara: ")
	fmt.Scan(&cariApa)
	searchingNegara(*data, n, cariApa, &ketemu, &indexDi)
	if ketemu {
		fmt.Println("=========== MENU DELETE ===========")
		fmt.Println("|1. Gold                          |")
		fmt.Println("|2. Silver                        |")
		fmt.Println("|3. Bronze                        |")
		fmt.Println("|4. Semua medali                  |")
		fmt.Println("===================================")
		fmt.Print("Pilih opsi (1/2/3/4): ")
		fmt.Scan(&medaliApa)

		if medaliApa == "1" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)? ")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].gold = 0
				popUp("Penghapusan berhasil. Tekan 'enter' untuk melanjutkan")
				sortingTotal(data, n)
				sortingMedali(data, n)
				mainMenu(*data, n)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n)
			}
		} else if medaliApa == "2" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)?")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].silver = 0
				popUp("Penghapusan berhasil. Tekan 'enter' untuk melanjutkan")
				sortingTotal(data, n)
				sortingMedali(data, n)
				mainMenu(*data, n)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n)
			}
		} else if medaliApa == "3" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)?")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].bronze = 0
				popUp("Penghapusan berhasil. Tekan 'enter' untuk melanjutkan")
				sortingTotal(data, n)
				sortingMedali(data, n)
				mainMenu(*data, n)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n)
			}
		} else if medaliApa == "4" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)?")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].gold = 0
				data[indexDi].silver = 0
				data[indexDi].bronze = 0
				popUp("Penghapusan berhasil. Tekan 'enter' untuk melanjutkan")
				sortingTotal(data, n)
				sortingMedali(data, n)
				mainMenu(*data, n)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n)
			}
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			deleteMedali(data, n)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		deleteMedali(data, n)
	}
}

func deleteNegara(data *pesertaSeagames, n int) {
	var pilih, cariApa string
	var ketemu bool
	var indexDi int
	fmt.Print("Ingin menghapus Id Negara: ")
	fmt.Scan(&cariApa)
	searchingNegara(*data, n, cariApa, &ketemu, &indexDi)
	if ketemu {
		fmt.Print("Apakah anda ingin menghapus data ini (Y/N)? ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			for i := indexDi; i < n; i++ {
				data[i] = data[i+1]
			}
			n -= 1
			popUp("Penghapusan berhasil. Tekan 'enter' untuk melanjutkan")
			sortingTotal(data, n)
			sortingMedali(data, n)
			mainMenu(*data, n)
		} else if pilih == "N" || pilih == "n" {
			deleteData(data, n)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			deleteNegara(data, n)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		deleteNegara(data, n)
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
