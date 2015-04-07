"use strict";

define(['app', '../service/user'], function (app) {
    return ['$scope', '$rootScope', '$location', 'User.Login',
        function ($scope, $rootScope, $location, Login) {
            $scope.args = {}
            $scope.ev = {}
            $scope.args.user = {}

            $scope.ev.login = function () {
                if (!$scope.loginform.$valid) {
                    return false
                }

                Login.query($scope.args.user, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }

                    if (json.errno == 2001) {
                        $location.path('/user/changepwd')
                        return
                    }

                    $location.path("/")
                })

            }


        }];
});
