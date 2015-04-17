"use strict";

define(['app', '../service/user'], function (app) {
    return ['$scope', '$rootScope', '$location', 'User.Changepwd',
        function ($scope, $rootScope, $location, Changepwd) {
            $scope.args = {}
            $scope.ev = {}
            $scope.args.user = {}

            $scope.ev.save = function () {
                if (!$scope.changepwdform.$valid) {
                    return false
                }

                Changepwd.query($scope.args.user, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }

                    $location.path("/login")
                })

            }


        }];
});
