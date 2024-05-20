package main

import "fmt"

const nmax int = 200

type seagames struct {
	negara                                           string
	gold, silver, bronze, id, rank, tMedali, rankByT int
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
		fmt.Println("Keluar. Terima kasih telah mencoba aplikasi ini.")
	} else {
		fmt.Println("Silahkan input pilihan yang benar.")
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
		deleteData(data, *n)
	} else if pilihCustom == "4" {
		mainMenu(*data, *n)
	} else {
		fmt.Println("Silahkan input pilihan yang benar.")
		kustomisasiData(data, n)
	}
}

func editData(data *pesertaSeagames, n int) {
	if n != 0 {
		var pilihEdit string
		fmt.Println()
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
			editMedali(data, n)
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

func cetakData(data pesertaSeagames, n int) {
	var maxG, maxS, maxB string
	if n != 0 {
		sortingTotal(&data, n)
		sortingData(&data, n)
		maxGSB(data, &maxG, &maxS, &maxB, n)
		fmt.Println()
		fmt.Println()
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("%-3s %-4s %-11s %-4s %-6s %-6s %-5s %-13s\n", "Id", "Rank", "Team/NOC", "Gold", "Silver", "Bronze", "Total", "Rank by Total")
		fmt.Println("----------------------------------------------------------")
		for i := 0; i < n; i++ {
			data[i].rank = i + 1
			fmt.Printf("%03d %-4d %-11s %-4d %-6d %-6d %-5d %-13d\n", data[i].id, data[i].rank, data[i].negara, data[i].gold, data[i].silver, data[i].bronze, data[i].tMedali, data[i].rankByT)
		}
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("Negara dengan perolehan Gold terbanyak: %s\n", maxG)

		fmt.Printf("Negara dengan perolehan Silver terbanyak: %s\n", maxS)
		fmt.Printf("Negara dengan perolehan Bronze terbanyak: %s\n", maxB)
		fmt.Println()
		fmt.Println()
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
				fmt.Println("Mohon maaf. Data tidak bisa lebih dari 200. Silahkan input lagi.")
				addData(data, n)
			} else {
				fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
				fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
				for i := *n; i < *n+tambah; i++ {
					fmt.Printf("%d. ", i+1)
					fmt.Scan(&negara, &gold, &silver, &bronze)
					data[i].id = i + 1
					data[i].negara = negara
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
						data[i].id = i + 1
						data[i].negara = negara
						data[i].gold = gold
						data[i].silver = silver
						data[i].bronze = bronze
						dataDuplicate = searchDuplicate(*data, *data, i+1)
					}
				}
				*n += tambah
				fmt.Println("Pemasukan data selesai!")
				kustomisasiData(data, n)
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
				fmt.Println("Mohon maaf. Data tidak bisa menampung lebih dari 200. Silahkan input lagi.")
				*n = 0
				addData(data, n)
			} else {
				fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
				fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
				for i := 0; i < *n; i++ {
					fmt.Printf("%d. ", i+1)
					fmt.Scan(&negara, &gold, &silver, &bronze)
					data[i].id = i + 1
					data[i].negara = negara
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
						data[i].id = i + 1
						data[i].negara = negara
						data[i].gold = gold
						data[i].silver = silver
						data[i].bronze = bronze
						dataDuplicate = searchDuplicate(*data, *data, i+1)
					}
				}
				fmt.Println("Pemasukan data selesai!")
				kustomisasiData(data, n)
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

func sortingData(data *pesertaSeagames, n int) {
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
}

func maxGSB(data pesertaSeagames, g, s, b *string, n int) {
	var maxG, maxS, maxB int
	for i := 0; i < n; i++ {
		if data[i].gold > maxG {
			maxG = data[i].gold
			*g = data[i].negara
		}
		if data[i].silver > maxS {
			maxS = data[i].silver
			*s = data[i].negara
		}
		if data[i].bronze > maxB {
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

func editNegara(data *pesertaSeagames, n int) {
	var negaraBaru, pilih string
	var ketemu, dataDuplicate bool
	var indexDi, cariApa int
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
				fmt.Printf("Data Negara %s sudah ada di dalam data. Data peserta negara tidak boleh sama. Silahkan coba lagi.\n", negaraBaru)
				editNegara(data, n)
			} else {
				fmt.Println("Pengeditan selesai.")
				editData(data, n)
			}
		} else if pilih == "N" || pilih == "n" {
			editData(data, n)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editData(data, n)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		editData(data, n)
	}
}

func searchingNegara(data pesertaSeagames, n, cari int, ketemu *bool, indexDi *int) {
	for i := 0; i < n; i++ {
		if cari == data[i].id {
			*ketemu = true
			*indexDi = i
		}
	}
}

func editMedali(data *pesertaSeagames, n int) {
	var pilih string
	var ketemu bool
	var indexDi, gold, silver, bronze, cariApa int
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
			fmt.Println("Pengeditan Selesai")
			editData(data, n)
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
		var pilihEdit string
		fmt.Println("-------------------------")
		fmt.Println("Pilih:")
		fmt.Println("1. Delete nama peserta Negara")
		fmt.Println("2. Delete data Medali peserta")
		fmt.Println("3. <- kembali")
		fmt.Println("-------------------------")
		fmt.Print("Pilih (1/2/3): ")
		fmt.Scan(&pilihEdit)

		if pilihEdit == "1" {
			deleteNegara(data, n)
		} else if pilihEdit == "2" {
			deleteMedali(data, n)
		} else if pilihEdit == "3" {
			kustomisasiData(data, &n)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editData(data, n)
		}
	} else {
		fmt.Println("Data peserta Seagames kosong. Penghapusan tidak dapat dilakukan.")
		kustomisasiData(data, &n)
	}
}

func deleteMedali(data *pesertaSeagames, n int) {
	var medaliApa, pilih string
	var ketemu bool
	var indexDi, cariApa int
	fmt.Print("Ingin menghapus medali dari Id Negara: ")
	fmt.Scan(&cariApa)
	searchingNegara(*data, n, cariApa, &ketemu, &indexDi)
	if ketemu {
		fmt.Println("-------------------------")
		fmt.Println("Pilih menghapus medali:")
		fmt.Println("1. Gold")
		fmt.Println("2. Silver")
		fmt.Println("3. Bronze")
		fmt.Println("4. Semua medali")
		fmt.Println("-------------------------")
		fmt.Print("Pilih (1/2/3/4): ")
		fmt.Scan(&medaliApa)

		if medaliApa == "1" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)? ")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].gold = 0
				fmt.Println("Penghapusan berhasil.")
				deleteData(data, n)
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
				fmt.Println("Penghapusan berhasil.")
				deleteData(data, n)
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
				fmt.Println("Penghapusan berhasil.")
				deleteData(data, n)
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
				fmt.Println("Penghapusan berhasil.")
				deleteData(data, n)
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
	var pilih string
	var ketemu bool
	var indexDi, cariApa int
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
			fmt.Println("Penghapusan berhasil.")
			deleteData(data, n)
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
