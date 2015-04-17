"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', '$rootScope', 'Project.GetAll', function ($scope, $rootScope, GetAll) {
        $scope.args = {}
        $scope.ev = {}

        $scope.args.list = []

        GetAll.query({}, function (json) {
            if ($rootScope.checkErr(json)) {
                return
            }

            $scope.args.list = json.data
        })
    }];
});
