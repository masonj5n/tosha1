package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
  "os"
  "flag"
  "bufio"
  "log"
)


func tosha1(pw string) string {

	hash := sha1.New()
	hash.Write([]byte((string(pw))))
	str := strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
  return str


}

func main() {

  var path = flag.String("file", "none", "name of file to be read")
  flag.Parse();

  file, err := os.Open(*path)
  if err != nil {
    fmt.Println("Error opening file")
    return
  }
  defer  file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(tosha1(scanner.Text()))
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }


}
