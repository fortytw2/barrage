// barrage routes

// keep default routing mode
m.route.mode = "search";

m.route(document.getElementById('barrage'), "/", {
  "/": home,
  "/detail/:seriesId": detail,
});
