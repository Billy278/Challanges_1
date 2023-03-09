package main

import "fmt"

func main() {

	// 	1. menampilkan nilai i : 21 fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
	i := 21
	var j bool = true
	fmt.Printf("%v \n", i)
	// 2. menampilkan tipe data dari variabel i
	fmt.Printf("%T \n", i)
	// 3. menampilkan tanda %
	fmt.Println("%")
	// 4. menampilkan nilai boolean j : true
	fmt.Printf("%v \n", j)
	// 5. menampilkan nilai boolean j : true
	fmt.Printf("%v \n", j)
	// 6. menampilkan unicode russia : Я (ya)
	fmt.Printf("%c \n", '\u042F')
	// 7. menampilkan nilai base 2 : 21
	fmt.Printf("%b \n", 21)
	// 8. menampilkan nilai base 10 : 21
	fmt.Printf("%d \n", 21)
	//9.menampilkan nilai base 8 :25
	// kak di output soal ada nilai 25 itu dari mana ya?
	//kalo nilai base 8 : 25 bukannya 31
	//untuk jaga2 ku print juga desimal dari 25 ya kak
	fmt.Printf("%d \n", 25)
	fmt.Printf("%o \n", 25)
	// 10. menampilkan nilai base 16 : f
	fmt.Printf("%x \n", 15)
	// 11. menampilkan nilai base 16 : F13 desimalnya = 3859
	fmt.Printf("%x \n", 3859)
	// 12. menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
	fmt.Printf("%U \n", 'Я')
	var k float64 = 123.456
	// 13. menampilkan float : 123.456000
	fmt.Printf("%f \n", k)
	// 14. menampilkan float scientific : 1.234560E+02
	fmt.Printf("%E \n", k)

}
