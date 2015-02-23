"use strict";

define(['app'], function(app) {

    app.factory('Project.Search', function($resource) {
		return $resource("/project/search/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.GetAll', function($resource) {
		return $resource("/project/search/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.Get', function($resource) {
		return $resource("/project/get/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.Save', function($resource) {
		return $resource("/project/save/", {}, {
            query: {isArray: false}
        });
    })

})
