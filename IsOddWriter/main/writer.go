package main
/*
  A fun little program that writes a very optimal IsOdd() function
*/
import (
  "fmt"
  "os"
  "log"
  "strconv"
  "time"
)

const fileName = "IsOdd"

func WriteToFile(fileName string, contents string) {
  file, err := os.Create(fileName + ".go")

  if err != nil {
    log.Fatal("Something bad happened during txt file creation", err)
  }

  _, err2 := file.WriteString(contents)

  if err2 != nil {
    log.Fatal("Something bad happened during writing proccess", err2)
  }

  file.Close()

  fmt.Println("Successfully created " + fileName + ".go")
}

func main() {
  startTime := time.Now()
  contents := "func IsOdd(number int) {"

  // you can use Repeat in strings library to do this part
  // For purposes of education I'll use for loops
  for i := 1; i < 400000; i += 2 {
    contents += "\n\tif number == " + strconv.Itoa(i) + " {\n\t\treturn true\n\t}"
  }

  contents += "\n\treturn false"
  contents += "\n}"

  WriteToFile(fileName, contents)
  elapsedTime := time.Since(startTime)
  fmt.Printf("Writing file took %s", elapsedTime)
}
