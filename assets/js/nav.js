/** @jsx m */
var m = require("./vendor/mithril.js");

module.exports.controller = function() {
};
//here's the view
module.exports.view = function(ctrl) {
  return <nav class="nav" tabindex="-1">
    <div class="container">
      <a href="/?/" class="pagename current"><img src="/img/banner.png" style="display:inline;"></img></a>
      <div class="right">
        <a href="/?/settings">Settings</a>
        <a href="https://github.com/fortytw2/barrage"><span class="octicon octicon-mark-github"></span></a>
      </div>
    </div>
  </nav>;
};
