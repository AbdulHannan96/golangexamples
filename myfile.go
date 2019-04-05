package main
import ("fmt")
import ("github.com/Greetings")


func ConcatSlice(sliceToConcat [] byte) string {
  mystring:=""
  for i:=0;i<len(sliceToConcat);i++{
    mystring+= string(sliceToConcat[i])
    if(i!=len(sliceToConcat)-1){
      mystring+="-"
    }
  }
  return mystring
}
func Encrypt(sliceToEncrypt [] byte,ceaserCount int) string {
  mystring:=""
  for i:=0;i<len(sliceToEncrypt);i++{
    mystring+= string(int(sliceToEncrypt[i])+ceaserCount)
  }
  return mystring
}
func EZGreetings(name string) string{
  return greetings.PrintGreetings(name)
}

func main() {
  str:=""
  my_arr:= [] byte("Hannan")
  str = ConcatSlice(my_arr)
  fmt.Printf("%s \n",str)
  str = Encrypt(my_arr,3)
  fmt.Printf("%s \n",str)
  str=EZGreetings("Hannan")
  fmt.Printf("%s",str)

}
