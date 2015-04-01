"use strict";

define(['app'], function(app) {

    app.factory('Deploy.Get', function($resource) {
		return $resource("/deploy/get", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Deploy.History', function($resource) {
		return $resource("/deploy/history", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Deploy.GetDiff', function($resource) {
		return $resource("/deploy/getdiff", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Deploy.HistoryDetail', function($resource) {
		return $resource("/deploy/historyDetail", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Deploy.AddTag', function($resource) {
		return $resource("/deploy/addtag", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Deploy.Push', function($resource) {
		return $resource("/deploy/push", {}, {
            query: {isArray: false}
        });
    })

})
