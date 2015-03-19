"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope','$route', 'Project.Save', 'Project.Get', function ($scope, $route, Save, Get) {
			$scope.args = {
				"project":{},
				"clusters":[]
			}
			$scope.ev = {}

			var $id = $route.current.params.id
			if ($id) {
				Get.query({id: $id}, function(json) {
					$scope.args.project = json.data
				})
			}

$scope.args.clusters = [
{
	"id":1,
	"name": "test1"
},
{
	"id":2,
	"name": "test2"
}
]

			$scope.ev.addCluster = function() {
				if(!$scope.args.project.projectClusters)
					$scope.args.project.projectClusters = []
				$scope.args.project.projectClusters.push([])
			}

			$scope.ev.delCluster = function(cluster) {
				var idx = $scope.args.project.projectClusters.indexOf(cluster)
				$scope.args.project.projectClusters.splice(idx, 1)
			}

			$scope.ev.save = function() {
				if(!$scope.projectform.$valid) {
					return false
				}

				var $data = {
					name: $scope.args.project.name,
					path: $scope.args.project.path,
					pushPath: $scope.args.project.pushPath,
					note: $scope.args.project.note ? $scope.args.project.note : "",
					tags: $scope.args.project.tags ? $scope.args.project.tags : "",
					projectClusters: JSON.stringify(getClusters())
				}
				Save.query($data, function(json) {

				})
			}

			function getClusters() {
				var $data = []
				angular.forEach($scope.args.project.projectClusters,  function(v) {
					var cluster = {
						cid: v.cid,
						bshell: v.bshell ? v.bshell:"",
						eshell: v.eshell ? v.eshell:"",
						note: v.note ? v.note:""
					}
					$data.push(cluster)
				})
				return $data
			}

        }];
});
