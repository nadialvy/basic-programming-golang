package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// ============== A.9 VARIABLE ==============
	// manifest typing"
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// type inference = tipe data otomatis
	middleName := "bin"
	fmt.Printf("first name = %s \nmiddle name = %s\nlast name = %s\n", firstName, middleName, lastName)

	// multi variable declaration
	var first, sec, third string = "satu", "dua", "tiga"
	fmt.Printf(first, sec, third)

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hi"
	fmt.Printf(fmt.Sprint(one), isFriday, twoPointTwo, say)
	fmt.Println()

	// underscore var = tempat sampah, kalo declare var dan gadipake. yaa mirip ts lah
	_ = "belajar"

	// declare var with new keyword
	// new is used to create a pointer variable with exact data type
	name := new(string)                          // default data type of name is pointer string
	fmt.Println("reference of name = ", name)    // reference, memory address
	fmt.Println("dereference of name = ", *name) //dereference
	// ============== A.9 VARIABLE ==============

	// ============== A.10 DATA TYPE ==============
	fmt.Println("DATA")
	// numeric non decimal = uint (only positive number) and int(negative and positive)
	var positiveNumber uint8 = 89
	var negativeNumber = -123456789

	fmt.Printf("\n\nbilangan positif %d\n", positiveNumber)
	fmt.Printf("bilangan negatif %d\n", negativeNumber)

	// numeric decimal = float32 and float64
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal %f\n", decimalNumber)
	fmt.Printf("bilangan desimal %.3f\n", decimalNumber)
	fmt.Printf("bilangan desimal %.10f\n", decimalNumber)

	// nil and zero value
	// nil not a data type, tapi dia sebuah nilai. nil = nilai kosong
	// tipe data yang diset nilainya dengan nil adalah pointer, tipe data fungsi, slice, map, channel, interface kosong atau any atau interface{}
	// ============== A.10 DATA TYPE ==============

	// ============== A.11 CONST ==============
	const pi = 3.14

	// declare multi const
	const (
		square         = "kotak" //type inference
		isToday  bool  = true    //manifest typing
		numeric  uint8 = 1       //manifest typing
		floatnum       = 2.2     //type inference
	)

	const (
		today    string = "senin"
		sekarang        //sekarang akan ngikut tipe dan nilai diatasnya, which is akan bernilai senin dengan tipe data string
		isToday2 = true
	)
	// ============== A.11 CONST ==============

	// ============== A.12 OPERATOR ==============
	// three type of operator = aritmatika, perbandingan dan logika
	// aritmatika = + - * / %
	// perbandingan == != < <= > >=
	// logika && || !
	// ============== A.12 OPERATOR ==============

	// ============== A.13 SELEKSI KONDISI ==============
	// cuma bisa if else dan switch dan gada ternary
	// fallthrough keyword dalam switch
	// di go gaperlu break di seatiap pernyataan switch, jadi kalo udah match sama suatu case maka case case selanjutnya gabakal dicompare
	// nah di fallthrough memaksa tetep ngecek kondisi setelahnya meskipun udah ada case yang match
	// efeknya case dipengecekan fllathrough akan dianggap true (meskipun hasil aslinya false)
	var point = 6
	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("not bad err")
		}
	}
	// ============== A.13 SELEKSI KONDISI ==============

	// ============== A.14 ARRAY ==============
	var names [4]string
	names[0] = "ab"
	names[1] = "bc"
	names[2] = "cd"
	names[3] = "de"

	// inidialisasi nilai awal
	var _ = [2]string{"banana", "apple"}

	// inisialisasi nilai awal tanpa nentuin jumlah elemen
	var _ = [...]int{1, 2, 3, 4, 5, 6, 7}

	// array multi dimensi
	var _ = [2][3]int{{2, 3, 3}, {1, 2, 3}}

	for i := 0; i < len(names); i++ {
		fmt.Printf("%s\n", names[i])
	}

	for i, name := range names {
		fmt.Printf("elemen %d : %s\n", i, name)
	}

	for _, name := range names {
		fmt.Printf("%s\n", name)
	}

	// alokasi elemen array dengan keyword make
	// make digunakan untuk membuat variable dengan tipe data slice, map dan channel
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"
	fmt.Println(fruits)
	// ============== A.14 ARRAY ==============

	// ============== A.15 SLICE ==============
	// slice adalah reference elemen array.
	// inisialisasi slice
	var _ = []string{"kimi no nama", "naruto", "jujutsu no kaisen"}

	// inisialisasi array
	// var _ = [2]string{"banana", "apple"}

	// perbedaan inisialisasinya adalah ga perlu dikasi tau jumlah elemen arraynya berapa
	var _ = []string{"apple", "banana"}        //slice
	var _ = [2]string{"tomato", "avocado"}     //array
	var _ = [...]string{"papaya", "pineapple"} //array

	// ARRAY VS SLICE
	// gabisa dibedanya, array = kumpulan nilai, slice = referensi tiap elemen
	// slice bisa dibentuk dari array yang sudah didefinisikan
	var forFruits = []string{"apple", "grape", "banana", "melon"}
	var copyForFruits = forFruits[0:2] //0 sampai sebelum 2

	// len vs cap
	// len dari copyTwoFruits adalah 2
	// cap dari copyTwoFruits adalah 4 => len ori dari twoFruits
	fmt.Println(len(copyForFruits))
	fmt.Println(cap(forFruits))

	// append => menambahkan elemen ke akhir slice
	var _ = append(forFruits, "papaya")

	// copy
	destination := make([]string, 3)
	source := []string{"watermelon", "pinneaple", "apple", "grape", "orange"}
	n := copy(destination, source) //masukin 3 elemn ke destination dari source

	fmt.Println(destination)
	fmt.Println(source)
	fmt.Println(n)

	// akses elemen dengan 3 index
	// fruits[0:1:1] => akses elemen 0 sampai sebelum 1, dan capacity nya adalah 1
	// angka capacity gaboleh melebihi capacity slice yang akan di slicing.

	// ============== A.15 SLICE ==============

	// ============== A.16 MAP ==============
	var chicken map[string]int //key nya string, valuenya int
	chicken = map[string]int{} //penambahan {} digunakan untuk menginisiasi nilai awal biar ga nil, kalau dibiarin nil bakalan error pas nampung data

	chicken["jan"] = 50
	chicken["feb"] = 40

	fmt.Println("januari = ", chicken["jan"])

	// atau kalau mau bikin var map dan langsung di assign nilai bisa kaya gini
	var chicken2 = map[string]int{
		"jan":  40,
		"feb":  30,
		"mar":  20,
		"apr":  90,
		"may":  50,
		"june": 20,
	}

	// variable map bisa di inisialiassi dengan tanpa nilai awal, yaitu dengan pake tanda kurung kurawal = map[string]int{}
	// atau bisa juga dengan menggunakan keyword make dan new, contoh
	var _ = map[string]int{}
	var _ = make(map[string]int)
	var _ = *new(map[string]int)
	// intinya ketiga cara diatas sama semua

	// iterasi map pake for range
	for key, val := range chicken2 {
		fmt.Println(key, " \t:", val)
	}

	delete(chicken2, "may")
	var value, isExist = chicken2["may"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	// kombinasi slice dan map
	// misalnya dipake buat nyimpen arrat yang isinya informasi siswa
	// ex: []map[string]int = slice yang tipe tiap elemennya map[string]int

	var absenChicken = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken green", "gender": "female"},
	}

	for _, chick := range absenChicken {
		fmt.Println(chick["gender"], " = ", chick["name"])
	}
	// ============== A.16 MAP ==============

	printMessage("nama nama buah", fruits)
	randomValue := randomWithRange(2, 10)
	fmt.Println("random number", randomValue)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Println("luas = ", area)
	fmt.Println("keliling = ", circumference)

	var avg = countAverage(2, 1, 3, 4, 2, 2, 3, 4, 7, 9, 8, 6, 5, 4, 5, 6, 5, 4)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

// ============== A.17 FUNCTION ==============
// void function
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// fungsi yang punya return value
var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

// ============== A.17 FUNCTION ==============

// ============== A.18 MULTIPLE RETURN FUNCTION ==============
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	// hitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// predefined return value
// artinya variable yang digunakan sebagai return bisa didefinisikan diawal
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d
	return
}

// ============== A.18 MULTIPLE RETURN FUNCTION ==============

// ============== A.19 FUNGSI VARIADIC ==============
// variadic function adalah fungsi dengan parametr bisa menampung nilai sejenis yang jumlahnya infinity
// cara akses nya dengan index

func countAverage(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// ============== A.19 FUNGSI VARIADIC ==============
