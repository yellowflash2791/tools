package main

import (
    "fmt"
    "os"
    "net/url"
    "net/http"
    "log"
    "net/http/httputil"
    "io/ioutil"
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

func sso(url1,cookie1 string) int{

//         webhookUrl := "https://hooks.slack.com/services/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"        
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

         reqBody, err := ioutil.ReadAll(res.Body)
         if err != nil {
                   log.Fatal(err)
         }
         fmt.Println(res.StatusCode,len(reqBody))

         return(len(reqBody)) 



}

func main() {

               webhookUrl := "https://hooks.slack.com/services/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"
               var s []int
               str := os.Args[1]
               cookie:=os.Args[2]
               u, _ := url.Parse(str)
               q, _ := url.ParseQuery(u.RawQuery)
               u.RawQuery = q.Encode()
               l:=u.String()
               x:=sso(l,"test")
               s=append(s,x)
               y:=sso(l,cookie)
               s=append(s,y)
               fmt.Println(s)
               if s[0]!=s[1] && (s[1]-s[0] >= 10 || s[0]-s[1] >=10){
                    text:=fmt.Sprintf("probable SSO login enabled on %s and content length is %d\n", str,s)
                    SendSlackNotification(webhookUrl,text) 
               }

}
