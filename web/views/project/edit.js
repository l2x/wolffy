"use strict";

define(['app', '../service/project', '../service/cluster', '../filter/filter'], function (app) {
    return ['$scope', '$rootScope', '$route', '$window', '$location', '$mdDialog', 'Project.Save', 'Project.Get', 'Project.Delete', 'Cluster.GetAll',
	function ($scope, $rootScope, $route, $window, $location, $mdDialog, Save, Get, Delete, Cluster_GetAll) {
        $scope.args = {}
        $scope.args.project = []
        $scope.args.clusters = []
        $scope.ev = {}

        var $id = $route.current.params.id
        if ($id) {
            Get.query({id: $id}, function (json) {
                if ($rootScope.checkErr(json)) {
                    return
                }
                $scope.args.project = json.data
            })
        }

        Cluster_GetAll.query({}, function (json) {
            if ($rootScope.checkErr(json)) {
                return
            }
            $scope.args.clusters = json.data
        })


        $scope.ev.addCluster = function () {
            if (!$scope.args.project.projectClusters)
                $scope.args.project.projectClusters = []
            $scope.args.project.projectClusters.push([])
        }

        $scope.ev.delCluster = function (cluster) {
            var idx = $scope.args.project.projectClusters.indexOf(cluster)
            $scope.args.project.projectClusters.splice(idx, 1)
        }

        $scope.ev.save = function () {
            if (!$scope.projectform.$valid) {
                return false
            }

            var $data = {
                id: $scope.args.project.id ? $scope.args.project.id : 0,
                name: $scope.args.project.name,
                path: $scope.args.project.path,
                pushPath: $scope.args.project.pushPath,
                note: $scope.args.project.note ? $scope.args.project.note : "",
                tags: $scope.args.project.tags ? $scope.args.project.tags : "",
                projectClusters: JSON.stringify(getClusters())
            }

			$rootScope.loadingDialog.show('Cloning')
            Save.query($data, function (json) {
				$rootScope.loadingDialog.hide()

                if ($rootScope.checkErr(json)) {
                    return
                }
                $location.path("/project/list")
            })
        }

        $scope.ev.delete = function (ev) {
            $mdDialog.show($rootScope.confirmDialog.targetEvent(ev)).then(function () {
                Delete.query({id: $scope.args.project.id}, function (json) {
                    if ($rootScope.checkErr(json)) {
                        return
                    }
					$location.path("/project/list")
                })
            })
        }

        function getClusters() {
            var $data = []
            angular.forEach($scope.args.project.projectClusters, function (v) {
                var cluster = {
                    cid: v.cid - 0,
                    bshell: v.bshell ? v.bshell : "",
                    eshell: v.eshell ? v.eshell : "",
                    note: v.note ? v.note : ""
                }
                $data.push(cluster)
            })
            return $data
        }

    }];
});
