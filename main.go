package main

import "fmt"

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

}
