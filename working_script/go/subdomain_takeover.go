package main

import (
    "fmt"
    "os"
//    "net/url"
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

func subdomain(url1,cookie1 string) {

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
         if (bytes.Contains(reqBody, []byte("Sorry, this page is no longer available")))|| (bytes.Contains(reqBody, []byte("If this is your website and you've just created it, try refreshing in a minute")))||(bytes.Contains(reqBody, []byte("The specified bucket does not exist")))||(bytes.Contains(reqBody, []byte("Repository not found")))||(bytes.Contains(reqBody, []byte("Trying to access your account?")))||(bytes.Contains(reqBody, []byte("Fastly error: unknown domain")))||(bytes.Contains(reqBody, []byte("404: This page could not be found.")))||(bytes.Contains(reqBody, []byte("The thing you were looking for is no longer here, or never was")))||(bytes.Contains(reqBody, []byte("There isn't a Github Pages site here.")))||(bytes.Contains(reqBody, []byte("404 Blog is not found")))||(bytes.Contains(reqBody, []byte("We could not find what you're looking for.")))||(bytes.Contains(reqBody, []byte("No settings were found for this company:")))||(bytes.Contains(reqBody, []byte("No such app")))||(bytes.Contains(reqBody, []byte("Uh oh. That page doesn't exist.")))||(bytes.Contains(reqBody, []byte("is not a registered InCloud YouTrack")))||(bytes.Contains(reqBody, []byte("No Site For Domain")))||(bytes.Contains(reqBody, []byte("It looks like you may have taken a wrong turn somewhere. Don't worry...it happens to all of us.")))||(bytes.Contains(reqBody, []byte("Unrecognized domain")))||(bytes.Contains(reqBody, []byte("Tunnel *.ngrok.io not found")))||(bytes.Contains(reqBody, []byte("404 error unknown site!")))||(bytes.Contains(reqBody, []byte("This public report page has not been activated by the user")))||(bytes.Contains(reqBody, []byte("Project doesnt exist... yet!")))||(bytes.Contains(reqBody, []byte("Sorry, this shop is currently unavailable.")))||(bytes.Contains(reqBody, []byte("This job board website is either expired or its domain name is invalid.")))||(bytes.Contains(reqBody, []byte("project not found")))||(bytes.Contains(reqBody, []byte("Whatever you were looking for doesn't currently exist at this address")))||(bytes.Contains(reqBody, []byte("Please renew your subscription")))||(bytes.Contains(reqBody, []byte("Non-hub domain, The URL you've accessed does not provide a hub.")))||(bytes.Contains(reqBody, []byte("This UserVoice subdomain is currently available!")))||(bytes.Contains(reqBody, []byte("The page you are looking for doesn't exist or has been moved.")))||(bytes.Contains(reqBody, []byte("Do you want to register *.wordpress.com?")))||(bytes.Contains(reqBody, []byte("Hello! Sorry, but the website you&rsquo;re looking for doesn&rsquo;t exist."))){
            text:=fmt.Sprintf("probable Subdomain Takeover in %v", url1)
            SendSlackNotification(webhookUrl,text)           

         }


}

func main() {
         
                 str:=os.Args[1]
                 cookie:=os.Args[2]
                 if cookie == "1" {
                    subdomain(str,"test")
                 }else{
                    subdomain(str,cookie)
                 }

}
