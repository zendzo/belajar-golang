package main

import "fmt"

func main() {
	var months = [12]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var sliceOne = months[4:7]
	fmt.Println(sliceOne)
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	// ketika utama array diubah maka slice refrence akan ikut berubah
	// months[5] = "mei lagi"
	// fmt.Println(sliceOne)

	// ketika slice refrence diubah maka array utama akan ikut berubah
	// sliceOne[2] = "diubah"
	// fmt.Println(months)
	// fmt.Println(sliceOne)

	var sliceTwo = months[10:]
	fmt.Println(sliceTwo)

	// jika append di array yang cap data masih cukup maka akan menggunakan array yang sama data-terkait
	// jika tidak maka akan dibuat array baru data-tidak-terkait
	var sliceThree = append(sliceTwo, "Eko")
	fmt.Println(sliceThree)
	sliceThree[1] = "Bukan desember"
	fmt.Println(sliceTwo)
	fmt.Println(months)

	// membuat slice
	newSlice := make([]string, 2, 5) // make(type, len,cap)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println(newSlice)

	// copy slice
	// source dan destination harus sama, jika kurang makan akan terpotong
	var copyslice = make([]string, len(newSlice), cap(newSlice))
	copy(copyslice, newSlice)
	fmt.Println(copyslice)

	// perbedaan dalam pembuatan array dan slice
	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
