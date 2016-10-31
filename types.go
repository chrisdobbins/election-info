package main
// types

type PollingInfo struct {
    Election election `json:"election"`
    EarlyVoteSites []earlyVoteSite `json:"earlyVoteSites"`
    Contests []contest `json:"contests"`
    State []stateInfo`json:"state"`
}

type election struct {
    Name string `json:"name"`
    ElectionDay string `json:"electionDay"`
}

type earlyVoteSite struct {
    Address address `json:"address"`
    PollingHours string `json:"pollingHours"`
    StartDate string `json:"startDate"`
    EndDate string `json:"endDate"`
}

type contest struct {
    ContestType string `json:"type"`
    Office string `json:"office"`
    Candidates []candidate `json:"candidates"`
}

type candidate struct {
    Name string `json:"name"`
    Party string `json:"party"`
    CandidateUrl string `json:"candidateUrl"`
}
type address struct {
  Line1 string `json:"line1"`
  City string `json:"city"`
  State string `json:"state"`
  Zip string `json:"zip"`
}
type stateInfo struct {
    Name string `json:"name"`
}
