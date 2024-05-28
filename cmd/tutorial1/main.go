//https://www.youtube.com/watch?v=8uiZC0l4Ajw
package main
import "fmt"

func main(){
  var intNum uint16 = 32767 
  intNum = intNum + 1
  fmt.Println(intNum)

  var floatNum float64 = 12345.9
  fmt.Println(floatNum)

  //casting
  var result float64 = floatNum + float64(intNum)
  fmt.Println(result)

  var mystring string = "Hello \n World"
  mystring = "hello" + "\n" + "world"
  mystring = `
  Hello
  World
  `
  fmt.Println(mystring) 
  //why go through all this
  var1, var2 := 1, 2
  fmt.Println(var1, var2)
}
