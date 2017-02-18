package main

import(
    "fmt"
    "bytes"
    //"strings"
)
func main()  {
    var buf bytes.Buffer
    var teste = "Isso é uma string de caracteres"
    buf.WriteString(teste)
    fmt.Println(buf)
}