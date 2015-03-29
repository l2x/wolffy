"use strict";

define(['app'], function(app) {

    app.factory('Deploy.Get', function($resource) {
		return $resource("/deploy/get/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Deploy.History', function($resource) {
		return $resource("/deploy/history/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Deploy.AddTag', function($resource) {
		return $resource("/deploy/addtag/", {}, {
            query: {isArray: false}
        });
    })

})
