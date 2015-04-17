"use strict";

define(['app', '../service/node'], function (app) {
    return ['$scope', '$rootScope', '$mdDialog', 'Node.GetPrivateKey',
		function ($scope, $rootScope, $mdDialog, GetPrivateKey) {
            $scope.args = {}

            GetPrivateKey.query({}, function (json) {
                if ($rootScope.checkErr(json)) {
                    return
                }
                $scope.args.privateKey = json.data
            })
        }];
});
