package main
 
import (
//    "time"
    "fmt"
    "os"
    "bufio"
    "net/url"
//    "net/http"
    "log"
//      "io/ioutil"
)
 

func main() {
  
              file, err := os.Open(os.Args[1])
              if err != nil {
                 log.Fatal(err)
              }
              defer file.Close()
              scanner := bufio.NewScanner(file)
              for scanner.Scan() {
                 str := scanner.Text()
                 u, _ := url.Parse(str)
                 q, _ := url.ParseQuery(u.RawQuery)
                 for key := range q {
                    q.Set(key,"")
                 }

                 u.RawQuery = q.Encode()
                 fmt.Println(u)
              }
              if err := scanner.Err(); err != nil {
                 log.Fatal(err)
              }


//               str := string(data) 
//               u, _ := url.Parse(str)
//               q, _ := url.ParseQuery(u.RawQuery)
//               for key := range q {
//                   q.Set(key,"")
//               }

//               u.RawQuery = q.Encode()
//               fmt.Println(u)

}





