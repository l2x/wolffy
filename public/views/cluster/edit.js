"use strict";

define(['app', '../service/cluster', '../service/node', '../filter/filter'], function (app) {
    return ['$scope', '$rootScope', '$route', '$window', '$mdDialog', 'Cluster.Get', 'Cluster.Save', 'Cluster.Delete', 'Node.GetAll',
        function ($scope, $rootScope, $route, $window, $mdDialog, Get, Save, Delete, Node_GetAll) {
            $scope.args = {}
            $scope.args.cluster = {}
            $scope.args.nodes = []
            $scope.ev = {}

            var $id = $route.current.params.id
            if ($id) {
                Get.query({id: $id}, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }
                    $scope.args.cluster = json.data
                })
            }

            Node_GetAll.query({}, function (json) {
                if ($rootScope.checkErr(json)) {
                    return
                }
                $scope.args.nodes = json.data
            })

            $scope.ev.addNode = function ($idx) {
                if (!$scope.args.cluster.nodes) {
                    $scope.args.cluster.nodes = []
                }

                for (var i = 0; i < $scope.args.cluster.nodes.length; i++) {
                    if ($scope.args.cluster.nodes[i].ip == $scope.args.nodes[$idx].ip) {
                        return
                    }
                }

                $scope.args.cluster.nodes.push($scope.args.nodes[$idx])
            }

            $scope.ev.removeNode = function ($idx) {
                $scope.args.cluster.nodes.splice($idx, 1)
            }

            $scope.ev.delCluster = function (ev) {
                var dialog = $rootScope.confirmDialog.targetEvent(ev)
                $mdDialog.show(dialog).then(function () {
                    Delete.query({id: $scope.args.cluster.id}, function (json) {
                        if ($rootScope.checkErr(json)) {
                            return
                        }
                        $window.history.back()
                    })
                })
            }

            $scope.ev.save = function () {
                if (!$scope.clusterform.$valid) {
                    return false
                }
                var data = {}
                angular.copy($scope.args.cluster, data)
                var ip = []
                angular.forEach(data.nodes, function (v) {
                    ip.push(v.id)
                })
                data.nodes = ip.join(",")

                Save.query(data, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }
                    $window.history.back()
                })
            }

        }];
});
