package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "log"
  "bytes"
  "strings"
_  "time"
  text "text/template"
    "net/url"
    "encoding/json"
_    "strconv"
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
    electionObj := sortContests(getPollingInfo(addressQuery))
    // electionObj.FormattedDate = formatDate(electionObj.Election.ElectionDay)
    // day, err := strconv.Atoi(strings.Split(electionObj.Election.ElectionDay, "-")[1])
    // if err != nil {
    //   fmt.Println("error: ", err)
    // }
    // dateSlice := strings.Split(electionObj.Election.ElectionDay, "-")
    // month, err := strconv.Atoi(dateSlice[1])
    // year, err := strconv.Atoi(dateSlice[0])
    // day, err := strconv.Atoi(dateSlice[2])

    // convMonth := new(time.Month)
    // fmt.Println(convMonth)
    // parsedElectionDay := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
    // fmt.Println(parsedElectionDay)
    // // strings.Replace(electionObj.Election.ElectionDay, "-", " ", -1)

    // fmt.Println(formatDate(electionObj.Election.ElectionDay))
    tmpl, err := text.New("election-info").Parse(getPage("election-info"))
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(electionObj)
    // fmt.Println(electionObj.EarlyVoteSites[0].Address)
    err = tmpl.Execute(w, &electionObj)
}

// func formatDate(date string) string {
//   dateSlice := strings.Split(date, "-")
//   var year, month string
//   var monthInt int
//   var err error
//   year = dateSlice[0]
//   monthInt, err = strconv.Atoi(dateSlice[1])
//   if err != nil {
//     fmt.Println(err)
//   }
//   month = time.Month(monthInt).String()
//
//   return fmt.Sprintf("%s %s, %s", month, dateSlice[2], year)
// }

// func calcDaysUntilElection(electionDay time.Time) {
//   today := time.Now().Day()
//   electionDate := electionDay.Day()
//   fmt.Println(electionDate)
//   fmt.Println(today)
//   fmt.Println(electionDay)
// }

type organizedPollingInfo struct {
    Election election
    DaysUntilElection int
    EarlyVoteSites []earlyVoteSite
    DropOffLocations []dropOffLocation
    PollingLocations []pollingLocation
    State []stateInfo
    GeneralContests []generalContest
    Referendums []referendumContest
}

func sortContests (p PollingInfo) (organizedPollingInfo) {
  generalContests := []generalContest{}
  referendums := []referendumContest{}
  fmt.Println(p)
  for _, contest := range p.Contests {
    if contest.ContestType == "General" {
      general := generalContest{
        Office: contest.Office,
        Candidates: contest.Candidates,
      }
      generalContests = append(generalContests, general)
    }
      if contest.ContestType == "Referendum" {
        referendum := referendumContest{
          Title: contest.ReferendumTitle,
          Subtitle: contest.ReferendumSubtitle,
          Text: contest.ReferendumText,
        }
        referendums = append(referendums, referendum)
    }
  }
  op := organizedPollingInfo{
    Election: p.Election,
    EarlyVoteSites: p.EarlyVoteSites,
    DropOffLocations: p.DropOffLocations,
    State: p.State,
    GeneralContests: generalContests,
    Referendums: referendums,
  }
  return op
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
  fmt.Println(fmt.Sprintf("https://www.googleapis.com/civicinfo/v2/voterinfo?address=%s&key=%s", addressParam, key))
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
