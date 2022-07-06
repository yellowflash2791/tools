package main

import (

    "os/exec"
    "io/ioutil"
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
    "strings"

)

type SlackRequestBody struct {
    Text string `json:"text"`
}



func execute(url1 string) {

    webhookUrl := "https://hooks.slack.com/services/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"
    out, err := exec.Command("touch","123.txt").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Command Successfully Executed")
    output := string(out[:])
    fmt.Println(output)


    test, err := exec.Command("aws","s3","cp","123.txt","s3://"+url1).Output()

    // if there is an error with our execution
    // handle it here
    if err != nil {
        fmt.Printf("%s", err)
    }
    output1 := string(test[:])
    if (strings.Contains(output1,"upload:")){
       text:=fmt.Sprintf("check s3-bucket with s3 url writeable %v", url1)
       SendSlackNotification(webhookUrl,text)
    }
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
     

func s3_check(url1 string) {

         //webhookUrl := "https://hooks.slack.com/services/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"
         u, _ := url.Parse(url1)
         //path := u.Path
         url2 := u.Scheme+"://"+u.Host+"/%c0"
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
         fmt.Println(string(requestDump))
         fmt.Println(res.StatusCode)
         reqBody, err := ioutil.ReadAll(res.Body)
         if err != nil {
                          log.Fatal(err)
         }
         if (bytes.Contains(reqBody, []byte("Couldn't parse the specified URI")))&&(bytes.Contains(reqBody, []byte("InvalidURI"))){
            fmt.Println("probable s3-bucket found in %v", url1)           
//            text:=fmt.Sprintf("probable s3-bucket found in %v", url1)
//            SendSlackNotification(webhookUrl,text)

         }



}


func s3_open_check(url1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"
         u, _ := url.Parse(url1)
         //path := u.Path
         url2 := u.Scheme+"://"+u.Host
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
         fmt.Println(string(requestDump))
         fmt.Println(res.StatusCode)
         reqBody, err := ioutil.ReadAll(res.Body)
         if err != nil {
                          log.Fatal(err)
         }
         if (bytes.Contains(reqBody, []byte("ListBucket"))){

            text:=fmt.Sprintf("open s3-bucket found %v", url1)
            SendSlackNotification(webhookUrl,text)

         }



}


func s3_aws_url(url1 string) {

         webhookUrl := "https://hooks.slack.com/services/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"
         u, _ := url.Parse(url1)
         //path := u.Path
         url2 := u.Scheme+"://"+u.Host+".s3.amazonaws.com/"
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
         fmt.Println(string(requestDump))
         fmt.Println(res.StatusCode)
         reqBody, err := ioutil.ReadAll(res.Body)
         if err != nil {
                          log.Fatal(err)
         }
         if (bytes.Contains(reqBody, []byte("ListBucket"))){
            text:=fmt.Sprintf("check s3-bucket with s3 url listing enabled checking if writeable %v", url2)
            SendSlackNotification(webhookUrl,text)
            execute(u.Host)
         }

         if (bytes.Contains(reqBody, []byte("<Code>AccessDenied</Code>"))){
             execute(u.Host)            
         }

}


func main() {
         url_test := os.Args[1]
//         fmt.Println(url_test)
         s3_check(url_test)
         s3_open_check(url_test)
         s3_aws_url(url_test)

}

