package main

import (
      "fmt"
      "strings"
      "math/rand"
      "time"
      "math"
)

// line 4 - 8 shows how to include library. just like require on node.js
func main() {
//   kata := 10
//   if kata == 10 {
//     fmt.Println("lulus dengan nilai sempurna")
// } else if kata > 5 {
//     fmt.Println("lulus")
// } else if kata == 4 {
//     fmt.Println("hampir lulus")
// } else {
//     fmt.Printf("tidak lulus. nilai anda %d\n", kata)
//   }
  // for i := 0; i < 5; i++ {
  //     fmt.Println("Angka", i)
  // }
  //======================= for with condition ==============================
  // var i = 0
  //
  // for i < 5 {
  //     fmt.Println("Angka", i)
  //     i++
  // }
  // for with condition as break
  // var i = 0
  //
  // for {
  //     fmt.Println("Angka", i)
  //
  //     i++
  //     if i == 5 {
  //         break
  //     }
  // }
  // ==========================  for with break and continue =========================
  //   for i := 1; i <= 10; i++ {
  //     if i % 2 == 0 {
  //         continue
  //     }
  //
  //     if i > 8 {
  //         break
  //     }
  //
  //     fmt.Println("Angka", i)
  // }
  // ============================== for without argumen ==============================
  // var i = 0
  //
  // for {
  //     fmt.Println("Angka", i)
  //
  //     i++
  //     if i == 5 {
  //         break
  //     }
  // }
  // ========================== nested for ===========================================
  //   for i := 0; i < 5; i++ {
  //     for j := i; j < 5; j++ {
  //         fmt.Print(j, " ")
  //     }
  //
  //     fmt.Println()
  // }
  // =========================== for loop with outer label ==========================
  // outerLoop:
  // for i := 0; i < 5; i++ {
  //     for j := 0; j < 5; j++ {
  //         if i == 3 {
  //             break outerLoop
  //         }
  //         fmt.Print("matriks [", i, "][", j, "]", "\n")
  //     }
  // }
  // ================================ how to assign array in go ==============================
  // var names [4]string
  // names[0] = "trafalgar"
  // names[1] = "d"
  // names[2] = "water"
  // names[3] = "law"
  //
  // fmt.Println(names[0], names[1], names[2], names[3])
  // ========================= how to assign array when declare the variables =================
  // cara horizontal
  //   var fruits = [4]string{"apple", "grape", "banana", "melon"}
  // cara vertical
  // var fruits = [4]string{
  //   "apple",
  //   "grape",
  //   "banana",
  //   "melon"
  // }
  //  Note: \t = tab so when you used \t it gives you 1 line of tab
  //   fmt.Println("Jumlah element \t\t", len(fruits))
  //   fmt.Println("Isi semua element \t", fruits)
  // ===================================== declare array variable without fil length of array ===============
  // var numbers = [...]int{2, 3, 2, 4, 3}
  //  note: use (...) and it will auto-calculate length of array
  // fmt.Println("data array \t:", numbers)
  // fmt.Println("jumlah elemen \t:", len(numbers))
  // ================== declare array with slice ===================================================
  // var fruits = []string{"apple", "grape", "banana", "melon"}
  // var fruitsA = []string{"apple", "grape"}      // slice
  // var fruitsB = [2]string{"banana", "melon"}    // array
  // var fruitsC = [...]string{"papaya", "grape"}  // array
  // fmt.Println(fruits) // "apple"
  // fmt.Println(fruitsA) // "[apple grape]"
  // var fruits = []string{"apple", "grape", "banana", "melon"}
  // var newFruits = fruits[0:4]
  //
  // fmt.Println(newFruits) // ["apple", "grape"]
  // dst := make([]string, 3)
  // src := []string{"watermelon", "pinnaple", "apple", "orange"}
  // n := copy(dst, src)
  //
  // fmt.Println(dst) // watermelon pinnaple apple
  // fmt.Println(src) // watermelon pinnaple apple orange
  // fmt.Println(n)   // 3
  // var chicken map[string]int
  // chicken = map[string]int{}
  //
  // chicken["januari"] = 50
  // chicken["februari"] = 40
  //
  // fmt.Println("januari", chicken["januari"]) // januari 50
  // fmt.Println("mei",     chicken["mei"])     // mei 0
    var names = []string{"John", "Wick"}
    printMessage("halo", names)
    rand.Seed(time.Now().Unix()) // untuk memastikan angka yang keluar benar" bernilai randoom apa tidak
    var randomValue int

    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)
    randomValue = randomWithRange(2, 10)
    fmt.Println("random number:", randomValue)

    var diameter float64 = 15
    var area, circumference = calculate(diameter)
    fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}
func printMessage(message string, arr []string) {
  var nameString = strings.Join(arr, " ") // Join array = Join.(<nama Variabel>, <join with>) in this case they separate array with whitespace
  fmt.Println(message, nameString)
}
func randomWithRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}
func calculate(d float64) (float64, float64) {
    // hitung luas
    var area = math.Pi * math.Pow(d / 2, 2)
    // hitung keliling
    var circumference = math.Pi * d
    // kembalikan 2 nilai
    return area, circumference
}
