<!DOCTYPE html>
<html>
    <head>
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" />
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" type="text/javascript"></script>
      <style>
      body {
        box-sizing: border-box;
      }
      </style>
    </head>
    <body>
      <h1 class="text-center">Find Your Election Info</h1>
      <div id="main-container">
        <form action="/search" method="post" class="form-vertical text-center">
          <label class="control-label col-md-12"><input type="text" name="address" placeholder="1313 Mockingbird Ln" />    Address</label>
          <label class="control-label col-md-12"><input type="text" name="city" placeholder="Cordele" /> City</label>
          <label class="control-label col-md-12"> <input type="text" name="state" placeholder="GA" /> State</label>
          <label class="control-label col-md-12"><input type="text" name="postal-code" placeholder="31010" />   ZIP</label>
          <input type="submit" />
        </form>
      </div>
    </body>
</html>
