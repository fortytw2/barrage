var home = {};

home.controller = function() {
  this.series = m.request({
    method: "GET",
    url: "/api/series"
  });
  this.movies = m.request({
    method: "GET",
    url: "/api/movies"
  });
};

// poster helper for the view
home.poster = function(series) {
  posterUri = "/video" + series.RootURI + "poster.png"
  return m("img", {src: posterUri, class: "img-responsive"})
};

//here's the view
home.view = function(controller) {
  return m("div", {class: "container"}, [
  controller.series().map(function(series, index) {
    return m("div",{class: "col-md-3"}, m("div", {class: "panel panel-default"}, [
    m("div", {class: "panel-heading"},series.Title),
    home.poster(series),
    m("div", {class: "panel-body"},
      m("a[href='/detail/" + series.Id + "']", {config: m.route},
      m("button", {class: "btn btn-default"}, "Watch"))),]));
  }),
  ]);
};
