package main

import "fmt"

func main() {

	// Dizi
	var a [2]string
	a[0] = "selam"
	a[1] = "Mehmet"

	// Slice
	sayilar := []int{1, 2, 3, 4}
	fmt.Println(a)

	fmt.Println(sayilar)

	// Referans tip

	names := []string{"deger"}
	systems := []string{"linux", "macos"}

	names = systems // names slice'ı systems slice'ının bellekteki adresini tutar
	// bu sebeple systems de yapılan bir değişiklik names'e de yansır

	fmt.Println("names : ", names, "\nsystems: ", systems)
	// OUTPUT
	// names :  [linux macos]
	// systems:  [linux macos]

	systems[0] = "Ubuntu"

	fmt.Println("names : ", names, "\nsystems: ", systems)
	// OUTPUT
	// names : [Ubuntu macos]
	// systems: [Ubuntu macos]

	// Dizi boyutunu artırma

	//systems[2] = "new index" // Hatalı kullanım

	systems = append(systems, "new index") // Olması gereken
	// dizi boyutu arttı ancak names isimli dizide herhangi bir değişiklik olmadı

	fmt.Println(systems)

	// alternatif slice tanımlama
	newArray := make([]int, 5, 10)
	// 5  sayısı dizinin uzunluğunu temsil eder , yani dizinini ilk 5 elemanı vardır (0)

	newArray = append(newArray, 1, 2)
	fmt.Println(newArray)
}
