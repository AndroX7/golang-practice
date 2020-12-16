package main

import "fmt"
import "reflect"

func main() {
    var number = 23
    var reflectValue = reflect.ValueOf(number)

    fmt.Println("tipe  variabel :", reflectValue.Type())

    if reflectValue.Kind() == reflect.Int {
      // .Kind used to compare types of data-types.
      // cannot use .Type to compare data-types
        fmt.Println(reflectValue.Type())
        // for checking data types in go
        // untuk ngecek tipe data di go
        fmt.Println("nilai variabel :", reflectValue.Int())
    }
}
