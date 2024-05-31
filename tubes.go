package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/google/uuid"
)

const nmax int = 200

type seagames struct {
	negara, id                                   string
	gold, silver, bronze, rank, tMedali, rankByT int
}

type pesertaSeagames [nmax]seagames

type arrayKosong [nmax]seagames

type recycleID [nmax]string

func main() {
	var data pesertaSeagames
	var nData int
	var recycleId recycleID
	for i := 0; i < nmax; i++ {
		recycleId[i] = "-"
	}
	data[0].tMedali = 0
	mainMenu(data, nData, recycleId)
}

func mainMenu(data pesertaSeagames, n int, recycle recycleID) {
	var pilihMainMenu, pilih string
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
		kustomisasiData(&data, &n, &recycle)
	} else if pilihMainMenu == "2" {
		cetakData(data, n, recycle)
	} else if pilihMainMenu == "3" {
		fmt.Print("Apakah anda ingin keluar (Y/N)? ")
		fmt.Scan(&pilih)
		if pilih == "Y" || pilih == "y" {
			fmt.Println("Keluar. Terima kasih telah mencoba aplikasi ini.")
		} else if pilih == "N" || pilih == "n" {
			mainMenu(data, n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			mainMenu(data, n, recycle)
		}
	} else {
		fmt.Println("Silahkan input pilihan yang benar.")
		mainMenu(data, n, recycle)
	}
}

func kustomisasiData(data *pesertaSeagames, n *int, recycle *recycleID) {
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
		addData(data, n, recycle)
	} else if pilihCustom == "2" {
		editData(data, *n, *recycle)
	} else if pilihCustom == "3" {
		deleteData(data, *n, recycle)
	} else if pilihCustom == "4" {
		mainMenu(*data, *n, *recycle)
	} else {
		fmt.Println("Silahkan input pilihan yang benar.")
		kustomisasiData(data, n, recycle)
	}
}

func editData(data *pesertaSeagames, n int, recycle recycleID) {
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
			editNegara(data, n, recycle)
		} else if pilihEdit == "2" {
			editMedali(data, n, recycle)
		} else if pilihEdit == "3" {
			kustomisasiData(data, &n, &recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editData(data, n, recycle)
		}
	} else {
		fmt.Println("Data peserta Seagames kosong. Pengeditan tidak dapat dilakukan.")
		kustomisasiData(data, &n, &recycle)
	}
}

func cetakData(data pesertaSeagames, n int, recycle recycleID) {
	var maxG, maxS, maxB, pilihSort string
	if n != 0 {
		maxGSB(data, &maxG, &maxS, &maxB, n)
		fmt.Println()
		fmt.Println()
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("%-3s %-4s %-11s %-4s %-6s %-6s %-5s %-13s\n", "Id", "Rank", "Team/NOC", "Gold", "Silver", "Bronze", "Total", "Rank by Total")
		fmt.Println("----------------------------------------------------------")
		for i := 0; i < n; i++ {
			fmt.Printf("%-3s %-4d %-11s %-4d %-6d %-6d %-5d %-13d\n", data[i].id, data[i].rank, data[i].negara, data[i].gold, data[i].silver, data[i].bronze, data[i].tMedali, data[i].rankByT)
		}
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("Negara dengan perolehan Gold terbanyak: %s\n", maxG)
		fmt.Printf("Negara dengan perolehan Silver terbanyak: %s\n", maxS)
		fmt.Printf("Negara dengan perolehan Bronze terbanyak: %s\n", maxB)
		fmt.Println()
		fmt.Println()
		fmt.Println("Pilih:")
		fmt.Println("1. Sorting berdasarkan 'Medali'")
		fmt.Println("2. Sorting berdasarkan 'Rank by Total'")
		fmt.Println("3. Sorting berdasarkan 'Team/NOC'")
		fmt.Println("4. <- kembali")
		fmt.Println("-------------------------")
		fmt.Print("Pilih (1/2/3/4): ")
		fmt.Scan(&pilihSort)

		if pilihSort == "1" {
			sortingMedali(&data, n)
			cetakData(data, n, recycle)
		} else if pilihSort == "2" {
			sortingRankbyT(&data, n)
			cetakData(data, n, recycle)
		} else if pilihSort == "3" {
			sortingNegara(&data, n)
			cetakData(data, n, recycle)
		} else if pilihSort == "4" {
			mainMenu(data, n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			cetakData(data, n, recycle)
		}
	} else {
		fmt.Println("Belum ada data peserta Seagames yang dimasukkan. Silahkan lakukan input data.")
		mainMenu(data, n, recycle)
	}
}

func addData(data *pesertaSeagames, n *int, recycle *recycleID) {
	var dataKosong arrayKosong
	var adaKah, dataDuplicate, nemuId bool
	var tambah int
	var negara, pilih string
	var gold, silver, bronze, index int
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
				addData(data, n, recycle)
			} else {
				fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
				fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
				for i := *n; i < *n+tambah; i++ {
					fmt.Printf("%d. ", i+1)
					fmt.Scan(&negara, &gold, &silver, &bronze)
					id := uuid.New()
					for !nemuId && index < nmax {
						if recycle[index] != "-" {
							nemuId = true
							data[i].id = recycle[index]
							recycle[index] = ""
						}
						index++
					}
					if !nemuId {
						data[i].id = id.String()[:3]
					}
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
						id := uuid.New()
						for !nemuId && index < nmax {
							if recycle[index] != "-" {
								nemuId = true
								data[i].id = recycle[index]
								recycle[index] = ""
							}
							index++
						}
						if !nemuId {
							data[i].id = id.String()[:3]
						}
						data[i].negara = negara
						data[i].gold = gold
						data[i].silver = silver
						data[i].bronze = bronze
						dataDuplicate = searchDuplicate(*data, *data, i+1)
					}
				}
				*n += tambah
				fmt.Println("Pemasukan data selesai!")
				if data[*n-1].tMedali == 0 {
					sortingTotal(data, *n)
				}
				sortingMedali(data, *n)
				kustomisasiData(data, n, recycle)
			}
		} else if pilih == "N" || pilih == "n" {
			kustomisasiData(data, n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			addData(data, n, recycle)
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
				addData(data, n, recycle)
			} else {
				fmt.Println("Silahkan masukkan data peserta Negara dan Medali nya!")
				fmt.Println("dengan format: <Negara> <Gold> <Silver> <Bronze>")
				for i := 0; i < *n; i++ {
					fmt.Printf("%d. ", i+1)
					fmt.Scan(&negara, &gold, &silver, &bronze)
					id := uuid.New()
					data[i].id = id.String()[:3]
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
						id := uuid.New()
						data[i].id = id.String()[:3]
						data[i].negara = negara
						data[i].gold = gold
						data[i].silver = silver
						data[i].bronze = bronze
						dataDuplicate = searchDuplicate(*data, *data, i+1)
					}
				}
				fmt.Println("Pemasukan data selesai!")
				if data[*n-1].tMedali == 0 {
					sortingTotal(data, *n)
				}
				sortingMedali(data, *n)
				kustomisasiData(data, n, recycle)
			}
		} else if pilih == "N" || pilih == "n" {
			kustomisasiData(data, n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			addData(data, n, recycle)
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

func sortingNegara(data *pesertaSeagames, n int) {
	var idx int
	var temp string
	var tempData seagames
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
}

func sortingRankbyT(data *pesertaSeagames, n int) {
	var idx, temp int
	var tempData seagames
	for pass := 1; pass < n; pass++ {
		idx = pass
		temp = data[pass].rankByT
		tempData = data[pass]
		for idx > 0 && data[idx-1].rankByT > temp {
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

func editNegara(data *pesertaSeagames, n int, recycle recycleID) {
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
				fmt.Printf("Data Negara %s sudah ada di dalam data. Data peserta negara tidak boleh sama. Silahkan coba lagi.\n", negaraBaru)
				editNegara(data, n, recycle)
			} else {
				fmt.Println("Pengeditan selesai.")
				editData(data, n, recycle)
			}
		} else if pilih == "N" || pilih == "n" {
			editData(data, n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editData(data, n, recycle)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		editData(data, n, recycle)
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

func editMedali(data *pesertaSeagames, n int, recycle recycleID) {
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
			fmt.Println("Pengeditan Selesai")
			editData(data, n, recycle)
		} else if pilih == "N" || pilih == "n" {
			editData(data, n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editMedali(data, n, recycle)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		editMedali(data, n, recycle)
	}
}

func deleteData(data *pesertaSeagames, n int, recycle *recycleID) {
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
			deleteNegara(data, n, recycle)
		} else if pilihEdit == "2" {
			deleteMedali(data, n, *recycle)
		} else if pilihEdit == "3" {
			kustomisasiData(data, &n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			editData(data, n, *recycle)
		}
	} else {
		fmt.Println("Data peserta Seagames kosong. Penghapusan tidak dapat dilakukan.")
		kustomisasiData(data, &n, recycle)
	}
}

func deleteMedali(data *pesertaSeagames, n int, recycle recycleID) {
	var medaliApa, pilih, cariApa string
	var ketemu bool
	var indexDi int
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
				deleteData(data, n, &recycle)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n, &recycle)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n, recycle)
			}
		} else if medaliApa == "2" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)?")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].silver = 0
				fmt.Println("Penghapusan berhasil.")
				deleteData(data, n, &recycle)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n, &recycle)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n, recycle)
			}
		} else if medaliApa == "3" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)?")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].bronze = 0
				fmt.Println("Penghapusan berhasil.")
				deleteData(data, n, &recycle)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n, &recycle)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n, recycle)
			}
		} else if medaliApa == "4" {
			fmt.Print("Apakah anda ingin menghapus data ini (Y/N)?")
			fmt.Scan(&pilih)
			if pilih == "Y" || pilih == "y" {
				data[indexDi].gold = 0
				data[indexDi].silver = 0
				data[indexDi].bronze = 0
				fmt.Println("Penghapusan berhasil.")
				deleteData(data, n, &recycle)
			} else if pilih == "N" || pilih == "n" {
				deleteData(data, n, &recycle)
			} else {
				fmt.Println("Silahkan input pilihan yang benar.")
				deleteMedali(data, n, recycle)
			}
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			deleteMedali(data, n, recycle)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		deleteMedali(data, n, recycle)
	}
}

func deleteNegara(data *pesertaSeagames, n int, recycle *recycleID) {
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
			recycle[indexDi] = data[indexDi].id
			for i := indexDi; i < n; i++ {
				data[i] = data[i+1]
			}
			n -= 1
			fmt.Println("Penghapusan berhasil.")
			deleteData(data, n, recycle)
		} else if pilih == "N" || pilih == "n" {
			deleteData(data, n, recycle)
		} else {
			fmt.Println("Silahkan input pilihan yang benar.")
			deleteNegara(data, n, recycle)
		}
	} else {
		fmt.Println("Id Negara tidak ditemukan. Silahkan cek kembali.")
		deleteNegara(data, n, recycle)
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
