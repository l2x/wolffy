"use strict";

define(['app', '../service/user'], function (app) {
    return ['$scope', '$rootScope', '$route', '$window', '$mdDialog', 'User.Get', 'User.Save', 'User.Delete',
        function ($scope, $rootScope, $route, $window, $mdDialog, Get, Save, Delete) {
            $scope.args = {}
            $scope.ev = {}
            $scope.args.user = {}


            var $id = $route.current.params.id
            if ($id) {
                Get.query({id: $id}, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }
                    $scope.args.user = json.data
                    $scope.args.user.administrator = $scope.args.user.administrator ? true : false
                })
            }

            $scope.ev.save = function () {
                if (!$scope.userform.$valid) {
                    return false
                }

                Save.query($scope.args.user, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }

                    $window.history.back()
                })
            }

            $scope.ev.delete = function (ev) {
                $mdDialog.show($rootScope.confirmDialog.targetEvent(ev)).then(function () {
                    Delete.query({id: $scope.args.user.id}, function (json) {
                        if ($rootScope.checkErr(json)) {
                            return
                        }

                        $window.history.back()
                    })
                })
            }


        }];
});
