"use strict";

define(['app'], function (app) {
    app.factory('ServiceView1', function ($resource) {
        return $resource("/test.php", {}, {
            query: {isArray: false}
        });
    })
})
