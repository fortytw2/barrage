/** @jsx m */
var m = require("./vendor/mithril.js");
var moment = require("./vendor/moment.js");
var nav = require("./nav.js");

module.exports.controller = function() {
  this.seriesId = m.route.param("seriesId");

  this.series = m.request({
    method: "GET",
    url: "/api/series/" + this.seriesId
  });

  this.navbar = new nav.controller();
};

// poster helper for the view
poster = function(series) {
  posterUri = "/video/" + series.RootURI + "poster.png";
  return <img src={posterUri}></img>;
};

resButtons = function(series, episode) {
  sourceUri = "/video/" + series.RootURI + episode.File;
  baseUri = "/video/" + series.RootURI + episode.File.substring(0, episode.File.indexOf('.'));
  return <div>
            <a href={baseUri + " - High.mp4" }>High (web) </a>
            <a href={baseUri + " - Low.mp4"}>Low (web) </a>
            <a href={baseUri + ".mkv"}>Source (mkv) </a>
          </div>;
};

// ensure a string isn't too long to display properly
shorten = function(string) {
  if (string.length < 55) {
    return string
  } else {
    return string.substring(0,55) + "...";
  }
};

//here's the view
module.exports.view = function(ctrl) {
  episodeList = [];
  ctrl.series().Episodes.map(function(episode, index) {
    episodeList.push(<tr> <td>{index + 1}</td> <td>{shorten(episode.Title)}</td><td> {moment(episode.ReleaseDate).format("MMMM Do YYYY")}</td> <td>{resButtons(ctrl.series(), episode)}</td> </tr>);
  });

  return <div>
    {nav.view(ctrl.navbar)}
    <div class="container" style="margin-top: 15px">
      {poster(ctrl.series())}
      <h2>{ctrl.series().Title}</h2>
      <table class="table">
        <thead>
          <tr>
            <td>Episode</td>
            <td>Title</td>
            <td>Release Date</td>
            <td></td>
          </tr>
        </thead>
        <tbody>
          {episodeList}
        </tbody>
      </table>
    </div>
  </div>;
};
