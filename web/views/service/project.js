"use strict";

define(['app'], function (app) {

    app.factory('Project.Search', function ($resource) {
        return $resource("/project/search", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.Get', function ($resource) {
        return $resource("/project/get", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.GetTags', function ($resource) {
        return $resource("/project/gettags", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.GetAll', function ($resource) {
        return $resource("/project/getall", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.Save', function ($resource) {
        return $resource("/project/add", {}, {
            query: {isArray: false}
        });
    })

    app.factory('Project.Delete', function ($resource) {
        return $resource("/project/delete", {}, {
            query: {isArray: false}
        });
    })

})
