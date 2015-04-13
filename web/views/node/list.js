"use strict";

define(['app', '../service/node'], function (app) {
    return ['$scope', '$rootScope', 'Node.GetAll', function ($scope, $rootScope, GetAll) {
        $scope.args = {}
        $scope.args.list = []
        $scope.ev = {}

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
