<!DOCTYPE html>
<html>
    <head>
    </head>
    <body>
      <div id="voting-sites">
        <h2>Your Voting Site(s)</h2>
        <ul>
          {{range $site := .EarlyVoteSites}}
          <li>{{$site.Address.Line1}}<br />
            {{$site.Address.City}}, {{$site.Address.State}} {{$site.Address.Zip}}<br />
          Hours: {{$site.PollingHours}}</li>
          {{end}}
        </ul>
      </div>
      <div id="contests">
        <h2>Contests</h2>
      </div>
    </body>
</html>
