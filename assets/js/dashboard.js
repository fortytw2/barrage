/** @jsx m */
var m = require("./vendor/mithril.js");
var nav = require("./nav.js");

module.exports.controller = function() {
  this.series = m.request({
    method: "GET",
    url: "/api/series"
  });
  this.navbar = new nav.controller();
};

// poster helper for the view
gridView = function(series) {
  return <div class="panel panel-default col c4">
    <div class="panel-heading">
      {series.title}
    </div>
    <div class="panel-body">
    <a href={"/?/detail/" + series.id}>
      <img src={series.poster} ></img>
    </a>
    </div>
  </div>;
};

// add div .row to every 3 things
rows = function(gridView) {
  gridded = [];
  row = [];
  gridView.map(function(seriesGrid, index) {
    if ((index + 1)% 3 === 0) {
      row.push(seriesGrid);
      gridded.push(<div class="row">{row}</div>);
      row = [];
    } else {
      row.push(seriesGrid);
    }

    if (index + 1 == gridView.length) {
      gridded.push(<div class="row">{row}</div>);
    }
  });
  return gridded;
};

//here's the view
module.exports.view = function(ctrl) {
  seriesGrid = [];
  ctrl.series().map(function(series, index) {
    seriesGrid.push(gridView(series));
  });

  return <div>
    {nav.view(ctrl.navbar)}
    <div class="container" style="margin-top: 15px">
      {rows(seriesGrid)}
    </div>
  </div>;
};
