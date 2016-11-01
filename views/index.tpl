<!DOCTYPE html>
<html>
    <head>
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" />
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" type="text/javascript"></script>
    </head>
    <body>
      <h1>Find Your Nearest Polling Place</h1>
      <div id="main-container">
        <form action="/search" method="post">
          <input type="text" name="address" placeholder="1313 Mockingbird Ln" />
          <input type="text" name="city" placeholder="Cordele" />
          <input type="text" name="state" placeholder="GA" />
          <input type="text" name="postal-code" placeholder="31010" />
          <input type="submit" />
        </form>
      </div>
    </body>
</html>
