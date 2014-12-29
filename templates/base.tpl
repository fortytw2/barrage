<!DOCTYPE html>
<!--[if lt IE 7]>      <html class="no-js lt-ie9 lt-ie8 lt-ie7"> <![endif]-->
<!--[if IE 7]>         <html class="no-js lt-ie9 lt-ie8"> <![endif]-->
<!--[if IE 8]>         <html class="no-js lt-ie9"> <![endif]-->
<!--[if gt IE 8]><!-->
<html class="no-js"> <!--<![endif]-->

  <head>
      <meta charset="utf-8">
      <meta http-equiv="X-UA-Compatible" content="IE=edge">
      <title>{{ template "title" }}</title>
      <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">

      <!-- minified css -->
      <link rel="stylesheet" href="static/css/main.min.css">

      <link href="//vjs.zencdn.net/4.11/video-js.css" rel="stylesheet">
      <script src="//vjs.zencdn.net/4.11/video.js"></script>
      <!-- favicon -->
      <link rel="icon" type="image/png" href="static/favicon.png">
  </head>

  <body>
    <nav class="navbar navbar-default navbar-fixed-top" role="navigation">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="/">Barrage</a>
        </div>
        <div id="navbar" class="collapse navbar-collapse">
          <form class="navbar-form navbar-left" role="search">
            <div class="form-group">
              <input type="text" class="form-control" placeholder="Search Videos">
            </div>
            <button type="submit" class="btn btn-primary">Search</button>
          </form>
          <ul class="nav navbar-nav pull-right">
            <li></li>
            <li><a href="/about">About</a></li>
            <li><a href="https://github.com/fortytw2/barrage"><span class="octicon octicon-mark-github"></span></a></li>
          </ul>
        </div><!--/.nav-collapse -->
      </div>
    </nav>

    <div class="container" style="margin-top:65px;">
      {{ template "body" }}
    </div>

    <script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
    <script type="text/javascript" src="static/js/bootstrap.min.js"></script>
  </body>

</html>
