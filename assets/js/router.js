// barrage routes
var m = require("./vendor/mithril.js");
var dashboard = require("./dashboard.js");
var detail = require("./detail.js");

// keep default routing mode
m.route.mode = "search";

m.route(document.body, "/", {
  "/": dashboard,
  "/detail/:seriesId": detail,
});
