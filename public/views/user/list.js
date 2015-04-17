"use strict";

define(['app', '../service/user'], function (app) {
    return ['$scope', '$rootScope', '$route', 'User.GetAll', 'User.GetUserInfo',
        function ($scope, $rootScope, $route, GetAll, GetUserInfo) {
            $scope.args = {}
            $scope.ev = {}
            $scope.args.list = []
            if (!$rootScope.user) {
                GetUserInfo.query({}, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }
                    $rootScope.user = json.data
                    $route.reload()
                })
                return
            }

            if (!$rootScope.user.administrator) {
                $location.path("/user/changepwd")
                return
            }

            GetAll.query({}, function (json) {
                if ($rootScope.checkErr(json)) {
                    return
                }

                $scope.args.list = json.data
            })

        }];
});
