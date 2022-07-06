package main
 
import (
    "time"
    "fmt"
    "os"
    "net/url"
    "net/http"
    "log"
    "net/http/httputil"
    "io/ioutil"

)
 
func request(link,cook string){

               client := &http.Client{}
               req, err := http.NewRequest("GET",link, nil)
               req.Header.Add("Cookie", cook)
               start := time.Now()
               res, err := client.Do(req) 
               if err != nil{
                  log.Fatal(err)
               }
               t:=time.Since(start).Seconds()
              
               requestDump, err := httputil.DumpRequest(req, true)
               if err != nil {
                   fmt.Println(err)
               }
               fmt.Println(string(requestDump))
               fmt.Println(link,res.StatusCode,t)
               reqBody, err := ioutil.ReadAll(res.Body)
               if err != nil {
                            log.Fatal(err)
               }
               fmt.Printf("%s", reqBody)
               fmt.Println(link,res.StatusCode,t,len(reqBody))

}

func main() {
 
               str := os.Args[1] 
               str1 := os.Args[2] 
               cookie:=os.Args[3]
               u, _ := url.Parse(str)
               q, _ := url.ParseQuery(u.RawQuery)
               for key := range q {
                   q.Set(key,str1)
               }

               u.RawQuery = q.Encode()
               l:=u.String()
//               fmt.Println(l)
               if cookie == "1" {
                  request(l,"test")
               }else{
                  request(l,os.Args[3])
               }


}





