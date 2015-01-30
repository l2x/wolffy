"use strict";

define(['app', './service'], function (app) {
    return ['$scope', 'ServiceView1',
        function ($scope, ServiceView1) {
            $scope.name = "haha"

            ServiceView1.query()
        }];
});
