var detail = {};

detail.controller = function() {
  this.seriesId = m.route.param("seriesId");

  this.series = m.request({
    method: "GET",
    url: "/api/series/" + this.seriesId
  });
};

// poster helper for the view
detail.poster = function(series) {
  posterUri = "/video" + series.RootURI + "poster.png"
  return m("div", {class: "col-md-4"},m("img", {src: posterUri, style: "float:left; height:50%, width:50%;",class: "img-responsive img-thumbnail"}))
};

detail.resButtons = function(series, episode) {
  sourceUri = "/video" + series.RootURI + episode.File
  return [
    m("a[href=" + sourceUri + "]", m("span", {class: "label label-primary pull-right"}, "Source")),
    m("a[href=" + sourceUri + "]", m("span", {class: "label label-success pull-right"}, "High")),
    m("a[href=" + sourceUri + "]", m("span", {class: "label label-warning pull-right"}, "Med")),
    m("a[href=" + sourceUri + "]", m("span", {class: "label label-danger pull-right"}, "Low")),
    ]
}

//here's the view
detail.view = function(controller) {
  return m("div", {class: "container"},
    m("div", {class: "panel panel-primary"},[
      m("div", {class: "panel-heading"}, controller.series().Title),
      m("div", {class: "panel-body"},
        detail.poster(controller.series()),
        m("div", {class:"col-md-8"},controller.series().Description)),
        m("ul", {class: "list-group"},[
        m("li", {class: "list-group-item active"}, "Episodes"),
      controller.series().Episodes.map(function(episode, index){
        return m("li",{class: "list-group-item"}, episode.Title, detail.resButtons(controller.series(), episode))
      })]),
      ])
)};
