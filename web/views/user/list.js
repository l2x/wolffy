"use strict";

define(['app', '../service/user'], function (app) {
    return ['$scope', '$rootScope', 'User.GetAll', function ($scope, $rootScope, GetAll) {
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
