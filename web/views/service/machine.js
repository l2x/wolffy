"use strict";

define(['app'], function(app) {


    app.factory('Machine.GetAll', function($resource) {
		return $resource("/machine/getall", {}, {
            query: {isArray: false}
        });
    })

})
