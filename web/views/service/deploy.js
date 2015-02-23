"use strict";

define(['app'], function(app) {

    app.factory('Deploy.Save', function($resource) {
		return $resource("/deploy/save/", {}, {
            query: {isArray: false}
        });
    })

})
