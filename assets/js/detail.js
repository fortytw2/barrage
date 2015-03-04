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
  return <img src={series.poster}></img>;
};

resButtons = function(series, episode) {
  return <div>
            <a href="">High (web) </a>
            <a href="">Low (web) </a>
            <a href="">Source (mkv) </a>
          </div>;
};

// ensure a string isn't too long to display properly
shorten = function(str) {
  if (str.length < 55) {
    return str;
  } else {
    return str.substring(0,55) + "...";
  }
};

//here's the view
module.exports.view = function(ctrl) {
  episodeList = [];
  ctrl.series().Episodes.map(function(episode, index) {
    episodeList.push(<tr> <td>{episode.id}</td> <td>{shorten(episode.title)}</td><td> {moment(episode.date).format("MMMM Do YYYY")}</td> <td>{resButtons(ctrl.series(), episode)}</td> </tr>);
  });

  return <div>
    {nav.view(ctrl.navbar)}
    <div class="container" style="margin-top: 15px">
      {poster(ctrl.series())}
      <h2>{ctrl.series().title}</h2>
      <p>{ctrl.series().description}</p>
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
