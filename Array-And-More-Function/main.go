package main

import (
    "fmt"
    // "strings"
    // "math/rand"
    // "math"
  )

  type student struct {
      name string
      grade int
  }

func main() {
    var number = [16]int{2,4,6,2,4,7,8,6,7,8,8,9,9,9,9,9}
    var objNumber map[int]int
    objNumber = map[int]int{}
    for _, each := range number {
          if val, ok := objNumber[each]; ok {
            fmt.Println(val)
            objNumber[each] = objNumber[each] + 1
          } else {
            objNumber[each] = 1
          }
    }
    fmt.Println(objNumber)
    // var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
    // var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
    // fmt.Println(msg)
    // var getMinMax = func(n []int) (int, int) {
    //     var min, max int
    //     for i, e := range n {
    //         switch {
    //         case i == 0:
    //             max, min = e, e
    //         case e > max:
    //             max = e
    //         case e < min:
    //             min = e
    //         }
    //     }
    //     return min, max
    //   }
    // contoh dari closure function, dimana sebuah fungsi diwakilkan/dimasukkan kedalam variabel.
    // pemanggilannya mirip seperti pemanggilan fungsi biasa, tinggal mencantumkan nama function dan parameter yang dibutuhkan
    //   var numbers = []int{2, 3, 4, 3, 4, 2, 3,1,0,8,9,10,15,20,100}
    //   var min, max = getMinMax(numbers)
    //   fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
    // var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    //
    // var newNumbers = func(min int) []int {
    //     var r []int
    //     for _, e := range numbers {
    //       fmt.Println("nilai minnya adalah %v", min)
    //       fmt.Println("nilai e nya adalah %v", e)
    //         if e < min {
    //             continue
    //         }
    //         r = append(r, e)
    //     }
    //     return r
    // }(3) // contoh penulisan parameter pada closure function jenis Immediately-Invoked Function Expression (IIFE)
    //
    // fungsi closure yang langsung dieksekusi ketika di deklarasikan.
    // bila memerluka parameter, maka param hanya perlu dituliskan setelah curly bracket
    //
    // fmt.Println("original number :", numbers)
    // fmt.Println("filtered number :", newNumbers)
    // var max = 3
    // var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    // var howMany, getNumbers = findMax(numbers, max)
    // var theNumbers = getNumbers()
    //
    // fmt.Println("numbers\t:", numbers)
    // fmt.Printf("find \t: %d\n\n", max)
    //
    // fmt.Println("found \t:", howMany)    // 9
    // fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
    // var data = []string{"wick", "jason", "ethan"}
    // var dataContainsO = filter(data, func(each string) bool {
    //     return strings.Contains(each, "o")
    // })
    // var dataLenght5 = filter(data, func(each string) bool {
    //     return len(each) == 5
    // })
    //
    // fmt.Println("data asli \t\t:", data)
    // // data asli : [wick jason ethan]
    //
    // fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
    // // filter ada huruf "o" : [jason]
    //
    // fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
    // var number = 4
    // fmt.Println("before :", number) // 4
    //
    // change(&number, 10)
    // fmt.Println("after  :", number) // 10
    // var numberA int = 250
    // var numberB *int = &numberA
    // fmt.Println("numberA (value)   :", numberA)  // 4
    // fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

    // fmt.Println("numberB (value)   :", *numberB) // 4
    // fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
    var data map[string]int
    data = map[string]int{}
    data["one"] = 1
    fmt.Println(data["one"])

    var s1 student
    s1.name = "john wick"
    s1.grade = 2

    fmt.Println("name  :", s1.name)
    fmt.Println("grade :", s1.grade)
}
// func change(original *int, value int) {
//     *original = value
// }
// func filter(data []string, callback func(string) bool) []string {
//     var result []string
//     for _, each := range data {
//         if filtered := callback(each); filtered {
//             result = append(result, each)
//         }
//     }
//     return result
// }
// func findMax(numbers []int, max int) (int, func() []int) {
//     var res []int
//     for _, e := range numbers {
//         if e <= max {
//             res = append(res, e)
//         }
//     }
//     return len(res), func() []int {
//         return res
//     }
// }
// func calculate(numbers ...int) float64 {
//     var total int = 0
//     var count int = 0
//     for _, number := range numbers {
//         fmt.Println("Iterasi ke-",count)
//         total += number
//         count++
//     }
//
//     var avg = float64(total) / float64(len(numbers))
//     return avg
// }
