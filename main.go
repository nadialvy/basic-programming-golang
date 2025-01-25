package main

import "fmt"

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

}
