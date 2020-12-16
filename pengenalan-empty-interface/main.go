package main

import "fmt"

func main() {
    var secret interface{}

    // interface data-type doesn't support for indexing.
    secret = "ethan hunt"
    fmt.Println(secret)

    secret = []string{"apple", "manggo", "banana"}
    fmt.Println(secret)
    //to prove it try to change secret=>secret[1]
    // it will return error because interface not support indexing
    secret = 12.4
    fmt.Println(secret)
    secret = map[string]interface{}{
      "chicken": 1,
      "monday": 2,
      }
    fmt.Println(secret)
    var data map[string]interface{}
    data = map[string]interface{}{
        "name":      "ethan hunt",
        "grade":     2,
        "breakfast": []string{"apple", "manggo", "banana"},
        "ant" :[]string{"semut merah", "semut hitam"},
    }
    fmt.Println(data["grade"])
    // if u wanna use interface for indexing then you have to define it into map data-type
    // then use inteface to define each element on it
    // jadi nanti defenisikan dulu dia sebagai map. trs untuk tiap" element datanya di defenisikan sebagai interface
}
