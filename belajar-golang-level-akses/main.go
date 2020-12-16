package main
import(
  . "belajar-golang-level-akses/library"
  f "fmt"
)

func main() {
    var s1 = Student{"ethan", 21}
    f.Println("name ", s1.Name)
    f.Println("grade", s1.Grade)
}
