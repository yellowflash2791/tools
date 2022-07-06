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
//    "bufio"
    "bytes"
    "crypto/tls"
    "encoding/json"
    "errors"
    "time"
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

func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}



func jsonp(url1,cookie1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"
         http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
         client := &http.Client{}
         req, err := http.NewRequest("GET",url1, nil)
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
         reqBody, err := ioutil.ReadAll(res.Body)
         if err != nil {
                          log.Fatal(err)
         }
         //fmt.Printf("%s", reqBody)
         if (bytes.Contains(reqBody, []byte("qwerty1357("))){
            text:=fmt.Sprintf("probable jsonp in %v", url1)
            SendSlackNotification(webhookUrl,text)

         }


}

func main() {
                 
                 var s []string
                 str:=os.Args[1]
                 cookie:=os.Args[2]
                 u, err := url.Parse(str)
                 if err != nil {
                      panic(err)
                 }
                 q, _ := url.ParseQuery(u.RawQuery)
                 for key := range q {
                         //q.Set(key,payload)
                     s = append(s,key) 

                 }
                 _, found := Find(s, "callback")

                 if found{
                        q.Set("callback", "qwerty1357")
                        u.RawQuery = q.Encode()                       
                 }else{
                        q.Add("callback", "qwerty1357")
                        u.RawQuery = q.Encode()
                 } 


                 if cookie == "1" {
                       jsonp(u.String(),"test")
                 }else{
                       jsonp(u.String(),cookie)
                 }




}
