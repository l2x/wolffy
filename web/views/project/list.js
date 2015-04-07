"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', '$rootScope', 'Project.Search', 'Project.GetAll', function ($scope, $rootScope, Search, GetAll) {
        $scope.args = {}
        $scope.ev = {}

        $scope.args.list = []

        GetAll.query({}, function (json) {
            if ($rootScope.checkErr(json)) {
                return
            }

            $scope.args.list = json.data
        })

        $scope.ev.search = function () {
            Search.query({keywords: $scope.args.keywords}, function (json) {
                console.log(json)
            })
        }

    }];
});
