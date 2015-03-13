"use strict";

define(['app'], function(app) {

	app.filter('ipFilter', function() {
		return function(ip, machines) {
			for(var $i=0; $i<machines.length; $i++) {
				if (machines[$i].ip == ip) {
					return false
				}
			}
			return true
		}
	})
})
