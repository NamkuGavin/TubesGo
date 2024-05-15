package main

import "fmt"

const nmax int = 200

type seagames struct {
	negara               string
	gold, silver, bronze int
}

type pesertaSeagames [nmax]seagames

type arrayKosong [nmax]seagames

func main() {
	var data pesertaSeagames
	var nData int
	mainMenu(data, nData)
}

func mainMenu(data pesertaSeagames, n int) {
	var pilihMainMenu string
	fmt.Println()
	fmt.Println("*************************")
	fmt.Println("Aplikasi Seagames Manager")
	fmt.Println("*************************")
	fmt.Println("Pilih menu:")
	fmt.Println("1. Kustomisasi data peserta Seagames")
	fmt.Println("2. Print Data")
	fmt.Println("3. keluar")
	fmt.Println("-------------------------")
	fmt.Print("Pilih (1/2/3): ")
	fmt.Scan(&pilihMainMenu)

	if pilihMainMenu == "1" {
		kustomisasiData(&data, &n)
	} else if pilihMainMenu == "2" {
		cetakData(data, n)
	} else if pilihMainMenu == "3" {
		fmt.Println("Keluar. Terima kasih telah mencoba aplikasi ini")
	} else {
		fmt.Println("Silahkan pilihan angka yang benar.")
		mainMenu(data, n)
	}
}

func kustomisasiData(data *pesertaSeagames, n *int) {
	var pilihCustom string
	fmt.Println()
	fmt.Println("************************")
	fmt.Println("Kustomisasi data peserta")
	fmt.Println("************************")
	fmt.Println("Pilih menu:")
	fmt.Println("1. Add")
	fmt.Println("2. Edit")
	fmt.Println("3. Delete")
	fmt.Println("4. <- kembali")
	fmt.Println("-------------------------")
	fmt.Print("Pilih (1/2/3/4): ")
	fmt.Scan(&pilihCustom)

	if pilihCustom == "1" {
		addData(data, n)
	} else if pilihCustom == "2" {
		editData(data, *n)
	} else if pilihCustom == "3" {
		fmt.Println("delete")
	} else if pilihCustom == "4" {
		mainMenu(*data, *n)
	} else {
		fmt.Println("Silahkan input pilihan yang benar.")
		kustomisasiData(data, n)
	}
}

func cetakData(data pesertaSeagames, n int) {
	var maxGold, maxSilver, maxBronze bool
	var maxG, maxS, maxB string
	if n != 0 {
		searchFindMax(data, &maxGold, &maxSilver, &maxBronze, n)
		if maxGold {
			sortingGoldDesc(&data, n)
		} else if maxSilver {
			sortingSilverDesc(&data, n)
		} else if maxBronze {
			sortingBronzeDesc(&data, n)
		}
		maxGSB(data, &maxG, &maxS, &maxB, n)
		fmt.Println()
		fmt.Println()
		fmt.Println("--------------------------------")
		fmt.Printf("%-10s %-4s %-6s %-6s\n", "Negara", "Gold", "Silver", "Bronze")
		fmt.Println("--------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-10s %-4d %-6d %-6d\n", data[i].negara, data[i].gold, data[i].silver, data[i].bronze)
		}
		fmt.Println("--------------------------------")
		fmt.Printf("Negara dengan perolehan Gold terbanyak: %s\n", maxG)

		fmt.Printf("Negara dengan perolehan Silver terbanyak: %s\n", maxS)
		fmt.Printf("Negara dengan perolehan Bronze terbanyak: %s\n", maxB)
		mainMenu(data, n)
	} else {
		fmt.Println("Belum ada data peserta Seagames yang dimasukkan. Silahkan lakukan input data.")
		mainMenu(data, n)
	}
}

func addData(data *pesertaSeagames, n *int) {
	var dataKosong arrayKosong
	var adaKah, dataDuplicate bool
	var tambah int
	var negara string
	var gold, silver, bronze int
	adaKah = checkingData(*n)
	if adaKah {
		fmt.Println("Terdapat data yang sudah tersimpan. Silahkan masukkan data peserta tambahan!")
		fmt.Print("Ingin menambahkan berapa data: ")
		fmt.Scan(&tambah)
		if *n+tambah > nmax {
			fmt.Println("Mohon maaf. Data tidak bisa lebih dari 200. Silahkan input lagi.")
			addData(data, n)
		} else {
			fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
			fmt.Printf("%-10s %-4s %-6s %-6s\n", "Negara", "Gold", "Silver", "Bronze")
			for i := *n; i < *n+tambah; i++ {
				fmt.Scan(&negara, &gold, &silver, &bronze)
				data[i].negara = negara
				data[i].gold = gold
				data[i].silver = silver
				data[i].bronze = bronze
			}
			dataDuplicate = searchDuplicate(*data, *data, *n+tambah)
			if dataDuplicate {
				for i := *n; i < *n+tambah; i++ {
					data[i] = dataKosong[i]
				}
				fmt.Println("Error. Data peserta negara tidak boleh sama. Silahkan coba lagi.")
				kustomisasiData(data, n)
			} else {
				*n += tambah
				fmt.Println("Pemasukan data selesai!")
				kustomisasiData(data, n)
			}
		}
	} else {
		fmt.Print("Ingin memasukkan berapa data: ")
		fmt.Scan(&*n)
		if *n > nmax {
			fmt.Println("Mohon maaf. Data tidak bisa lebih dari 200. Silahkan input lagi.")
			*n = 0
			addData(data, n)
		} else {
			fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
			fmt.Printf("%-10s %-4s %-6s %-6s\n", "Negara", "Emas", "Silver", "Bronze")
			for i := 0; i < *n; i++ {
				fmt.Scan(&negara, &gold, &silver, &bronze)
				data[i].negara = negara
				data[i].gold = gold
				data[i].silver = silver
				data[i].bronze = bronze
			}
			dataDuplicate = searchDuplicate(*data, *data, *n)
			if dataDuplicate {
				for i := 0; i < *n; i++ {
					data[i] = dataKosong[i]
				}
				*n = 0
				fmt.Println("Error. Data peserta negara tidak boleh sama. Silahkan coba lagi.")
				kustomisasiData(data, n)
			} else {
				fmt.Println("Pemasukan data selesai!")
				kustomisasiData(data, n)
			}
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

func sortingGoldDesc(data *pesertaSeagames, n int) {
	var idx, temp int
	var tempData seagames
	for pass := 1; pass < n; pass++ {
		idx = pass
		temp = data[pass].gold
		tempData = data[pass]
		for idx > 0 && data[idx-1].gold < temp {
			data[idx] = data[idx-1]
			idx -= 1
		}
		data[idx] = tempData
	}
}

func sortingSilverDesc(data *pesertaSeagames, n int) {
	var idx, temp int
	var tempData seagames
	for pass := 1; pass < n; pass++ {
		idx = pass
		temp = data[pass].silver
		tempData = data[pass]
		for idx > 0 && data[idx-1].silver < temp {
			data[idx] = data[idx-1]
			idx -= 1
		}
		data[idx] = tempData
	}
}

func sortingBronzeDesc(data *pesertaSeagames, n int) {
	var idx, temp int
	var tempData seagames
	for pass := 1; pass < n; pass++ {
		idx = pass
		temp = data[pass].bronze
		tempData = data[pass]
		for idx > 0 && data[idx-1].bronze < temp {
			data[idx] = data[idx-1]
			idx -= 1
		}
		data[idx] = tempData
	}
}

func searchFindMax(data pesertaSeagames, maxGold, maxSilver, maxBronze *bool, n int) {
	var max int
	max = data[0].gold
	for i := 0; i < n; i++ {
		if data[i].gold >= max {
			max = data[i].gold
			*maxGold = true
			*maxSilver = false
			*maxBronze = false
		}
		if data[i].silver >= max {
			max = data[i].silver
			*maxGold = false
			*maxSilver = true
			*maxBronze = false
		}
		if data[i].bronze >= max {
			max = data[i].bronze
			*maxGold = false
			*maxSilver = false
			*maxBronze = true
		}
		if max == data[i].bronze || max == data[i].silver {
			*maxGold = true
			*maxSilver = false
			*maxBronze = false
		}
	}
}

func maxGSB(data pesertaSeagames, g, s, b *string, n int) {
	var maxG, maxS, maxB int
	maxG = data[0].gold
	maxS = data[0].silver
	maxB = data[0].bronze
	for i := 0; i < n; i++ {
		if data[i].gold >= maxG {
			maxG = data[i].gold
			*g = data[i].negara
		}
		if data[i].silver >= maxS {
			maxS = data[i].silver
			*s = data[i].negara
		}
		if data[i].bronze >= maxB {
			maxB = data[i].bronze
			*b = data[i].negara
		}
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

func editData(data *pesertaSeagames, n int) {
	if n != 0 {
		var pilihEdit string
		fmt.Println("-------------------------")
		fmt.Println("Pilih:")
		fmt.Println("1. Edit nama peserta Negara")
		fmt.Println("2. Edit data Medali peserta")
		fmt.Println("3. <- kembali")
		fmt.Println("-------------------------")
		fmt.Print("Pilih (1/2/3): ")
		fmt.Scan(&pilihEdit)

		if pilihEdit == "1" {
			editNegara(data, n)
		} else if pilihEdit == "2" {
			fmt.Print()
		} else if pilihEdit == "3" {
			kustomisasiData(data, &n)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editData(data, n)
		}
	} else {
		fmt.Println("Data peserta Seagames kosong. Pengeditan tidak dapat dilakukan.")
		kustomisasiData(data, &n)
	}
}

func editNegara(data *pesertaSeagames, n int) {
	var cariApa, negaraBaru string
	var ketemu, dataDuplicate bool
	var indexDi int
	var titip pesertaSeagames
	for i := 0; i < n; i++ {
		titip[i] = data[i]
	}
	fmt.Print("Ingin mengedit nama Negara apa: ")
	fmt.Scan(&cariApa)
	searchingNegara(*data, n, cariApa, &ketemu, &indexDi)
	if ketemu {
		fmt.Print("Nama Negara ditemukan. Ubah menjadi: ")
		fmt.Scan(&negaraBaru)
		data[indexDi].negara = negaraBaru
		dataDuplicate = searchDuplicate(*data, *data, n)
		if dataDuplicate {
			for i := 0; i < n; i++ {
				data[i].negara = titip[i].negara
			}
			fmt.Println("Error. Data peserta negara tidak boleh sama. Silahkan coba lagi.")
			editNegara(data, n)
		} else {
			fmt.Println("Pengeditan selesai.")
			editData(data, n)
		}
	} else {
		fmt.Println("Nama Negara tidak ditemukan. Silahkan cek kembali.")
		editData(data, n)
	}
}

func searchingNegara(data pesertaSeagames, n int, cari string, ketemu *bool, indexDi *int) {
	for i := 0; i < n; i++ {
		if cari == data[i].negara {
			*ketemu = true
			*indexDi = i
		}
	}
}
