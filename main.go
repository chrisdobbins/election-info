package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "log"
  "bytes"
  "strings"
  // "os"
  _ "html/template"
  text "text/template"
    "net/url"
    "encoding/json"
)
type FormVars struct {
  Address string
}

func main() {
  http.HandleFunc("/", serveIndex)
  http.HandleFunc("/search", search)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func search(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    address := r.Form["address"][0]
    fmt.Println(address)
    city := r.Form["city"][0]
    state := r.Form["state"][0]
    postal := r.Form["postal-code"][0]
    addressQuery := escapeQuery(address, city, state, postal)
    electionObj := getPollingInfo(addressQuery)
    tmpl, err := text.New("election-info").Parse(getPage("election-info"))
    if err != nil {
      fmt.Println(err)
    }
    // fmt.Println(electionObj.EarlyVoteSites[0].Address)
    err = tmpl.Execute(w, &electionObj)

}

func getPollingInfo(addressParam string) PollingInfo {
  resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/civicinfo/v2/voterinfo?address=%s&key=%s", addressParam, key))
  if err != nil {
    fmt.Println(err)
  }
  buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  var info PollingInfo
  if err := json.Unmarshal([]byte(buf.String()), &info); err != nil {
    fmt.Println(err)
  }
  defer resp.Body.Close()
  return info
}

func escapeQuery(args ...string) string {
  joinedQuery := strings.Join(args, " ")
  return url.QueryEscape(joinedQuery)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, getPage("index"))
}

func getPage(name string) string {
  contents, _ := ioutil.ReadFile(fmt.Sprintf("views/%s.tpl", name))
  return string(contents[:])
}
