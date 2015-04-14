"use strict";

define(['app'], function (app) {
    app.factory('Node.GetAll', function ($resource) {
        return $resource("/node/getall", {}, {
            query: {isArray: false}
        });
    })
    app.factory('Node.Delete', function ($resource) {
        return $resource("/node/delete", {}, {
            query: {isArray: false}
        });
    })
})
