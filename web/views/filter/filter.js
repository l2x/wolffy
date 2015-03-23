"use strict";

define(['app'], function(app) {

	app.filter('ipFilter', function() {
		return function(ip, machines) {
		if (!machines || !ip){
			return true
		}
			for(var $i=0; $i<machines.length; $i++) {
				if (machines[$i].ip == ip) {
					return false
				}
			}
			return true
		}
	})
})
