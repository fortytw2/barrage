/** @jsx m */
var m = require("./vendor/mithril.js");
var nav = require("./nav.js");

module.exports.controller = function() {

  this.navbar = new nav.controller();
};


module.exports.view = function(ctrl) {
  return <div>
    {nav.view(ctrl.navbar)}
    <div class="container" style="margin-top: 15px">
      <h1>Settings Page</h1>
    </div>
  </div>;
};
