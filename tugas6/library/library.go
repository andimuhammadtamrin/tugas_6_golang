package library

import "fmt"

func private(){
  var m1 mahasiswa
  m1.nama = "Zilong"
  m1.umur = 20
  m1.data()

}
func Public(){
  var m2 = mahasiswa{"Karina",17}
  m2.data()
  fmt.Println()
  private()
}


type mahasiswa struct{
  nama string
  umur int
}

func (s mahasiswa) data(){
    fmt.Println("Nama saya adalah",s.nama)
    fmt.Println("Umur saya adalah",s.umur)
  }
