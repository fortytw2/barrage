// barrage routes
var m = require("./vendor/mithril.js");
var dashboard = require("./dashboard.js");
var detail = require("./detail.js");
var settings = require("./settings.js");

// keep default routing mode
m.route.mode = "search";

m.route(document.body, "/", {
  "/": dashboard,
  "/settings": settings,
  "/detail/:seriesId": detail,
});
