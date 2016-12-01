/* vim: set ft=javascript sw=2 ts=2 : */

var Gatsby = {};

Gatsby.storage = localStorage;

$(document).ready(function() {
  var ractive = Ractive({
    el: '#gatsby',
    template: '#gatsby-tmpl',
    data: function() {
      return {
        code: 0,
        status: "Gatsby's loading, ol'sport...",
      };
    },
    oncomplete: function() {
      $.ajax({
        url: '/code',
        method: 'GET',
        dataType: 'text'
      }).done(function(data) {
        this.set('code', parseInt(data));
        this.set('status', "Gatsby's just fine, ol'sport.");
      }.bind(this)).fail(function(jqxhr, status) {
        console.log('Failed to get code: ' + status);
        this.set('status', "Gatsby's dead, ol'sport!");
      }.bind(this)).always(function() {
        this.observe('code', function(new_value, old_value, keypath) {
          this.set('status', "Gatsby is missing, again.");
          $.ajax({
            url: '/code',
            method: 'PUT',
            dataType: 'json',
            data: JSON.stringify({
              'Code': parseInt(this.get('code'))
            })
          }).done(function(data) {
            this.set('status', "Gatsby is up-to-date, ol'sport.");
          }.bind(this)).fail(function(jqxhr, status) {
            console.log('Failed to set code: ' + status);
            this.set('status', "Gatsby's dead, ol'sport!");
          }.bind(this));
        });
      }.bind(this));

      this.on({
        stop: function() {
          this.set('code', 0);
        },
        pulse: function() {
          this.set('code', 1);
        },
        spaz: function() {
          this.set('code', 2);
        },
      });
    }
  });
});

