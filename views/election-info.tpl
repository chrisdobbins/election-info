<!DOCTYPE html>
<html>
    <head>

      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" />
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" type="text/javascript"></script>
      <style>
        /*li {
          list-style-type: none;
        }*/
      </style>
    </head>
    <body>
      <div class="container">
        <div class="row">
          <h1 class="text-center">Election 2016</h1>
        </div>
        <div class="row">
      <div id="contests" class="col-md-6">
        <h2>Elections</h2>
          {{range $contest := .GeneralContests}}
          <h3>{{$contest.Office}}</h3>
          <ul class="list-group">
            {{range $candidate := $contest.Candidates}}
            <li class="list-group-item"><a href="{{$candidate.CandidateUrl}}"><h4>{{$candidate.Name}}</h4>
              {{if $candidate.Party}}
                {{$candidate.Party}} <br />
              {{else}}
                No party listed
              {{end}}
           </a></li>
            {{end}}
          </ul>
          {{end}}
          </div>
          {{if .Referendums}}
          <div id="referendums" class="col-md-6">
            <h2>Referendums</h2>
            <ul class="list-group">
              {{range $referendum := .Referendums}}
              <li class="list-group-item"><h3>{{$referendum.Title}}</h3>
                      <h4>{{$referendum.Subtitle}}</h4>
                      <p>{{$referendum.Text}}</p></li>
              {{end}}
           </ul>
        </div>
        {{end}}

      <div id="voting-sites" class="col-md-6">
        <h2>Your Voting Site(s)</h2>
          {{if .EarlyVoteSites}}
          <h3>Early Voting Sites</h3>
          <ul class="list-group">
          {{range $site := .EarlyVoteSites}}
          <li class="list-group-item">{{$site.Address.Line1}}<br />
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
      </div>
    </div>
    </body>
</html>
