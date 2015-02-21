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

})
