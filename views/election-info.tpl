<!DOCTYPE html>
<html>
    <head>

      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" />
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" type="text/javascript"></script>
      <style>
      body {
        box-sizing: border-box;
      }

      #contests {
        margin-left: 10px;
      }
      </style>
    </head>
    <body>
      <div id="contests">
        <h2>Contests</h2>
          {{range $contest := .Contests}}
          {{if $contest.Office}}
          <h3>{{$contest.Office}}</h3>
          <ul>
            {{range $candidate := $contest.Candidates}}
            <li>{{$candidate.Name}}
            {{if $candidate.Party}}<br />{{$candidate.Party}}{{end}} <br /></li>
            {{end}}
          </ul>
          {{else if $contest.ReferendumTitle}}
          <h3>{{$contest.ReferendumTitle}} </h3>
              <p>{{$contest.ReferendumText}}</p>
          {{end}}
        {{end}}
      </div>

      <div id="voting-sites">
        <h2>Your Voting Site(s)</h2>

          {{if .EarlyVoteSites}}
          <h3>Early Voting Sites</h3>
          <ul>
          {{range $site := .EarlyVoteSites}}
          <li>{{$site.Address.Line1}}<br />
            {{$site.Address.City}}, {{$site.Address.State}} {{$site.Address.Zip}}<br />
          Hours: {{$site.PollingHours}}</li>
          {{end}}
          </ul>
          {{end}}
          {{if .DropOffLocations}}
          <h3>Drop-Off Locations</h3>
          <ul>
          {{range $site := .DropOffLocations}}
          <li>{{$site.Address.Line1}}<br />
            {{$site.Address.City}}, {{$site.Address.State}} {{$site.Address.Zip}}<br />
          Hours: {{$site.PollingHours}}</li>
          {{end}}
          </ul>
          {{end}}
          {{if .PollingLocations}}
          <h3>Polling Locations</h3>
          <ul>
            {{range $site := .PollingLocations}}
            <li>{{$site.Address.Line1}}<br />
              {{$site.Address.City}}, {{$site.Address.State}} {{$site.Address.Zip}}<br />
            Hours: {{$site.PollingHours}}</li>
            {{end}}
          </ul>
          {{end}}
      </div>
    </body>
</html>
