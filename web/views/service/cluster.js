"use strict";

define(['app'], function(app) {

    app.factory('Cluster.Search', function($resource) {
		return $resource("/cluster/search/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Cluster.GetAll', function($resource) {
		return $resource("/cluster/search/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Cluster.Get', function($resource) {
		return $resource("/cluster/get/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Cluster.Save', function($resource) {
		return $resource("/cluster/save/", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.Delete', function($resource) {
		return $resource("/project/delete/", {}, {
            query: {isArray: false}
        });
    })

})
