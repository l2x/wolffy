"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', '$rootScope','$route', '$window', '$mdDialog', 'Project.Save', 'Project.Get', 'Project.Delete', function ($scope, $rootScope, $route, $window, $mdDialog, Save, Get, Delete) {
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
					id: $scope.args.project.id ? $scope.args.project.id : 0,
					name: $scope.args.project.name,
					path: $scope.args.project.path,
					pushPath: $scope.args.project.pushPath,
					note: $scope.args.project.note ? $scope.args.project.note : "",
					tags: $scope.args.project.tags ? $scope.args.project.tags : "",
					projectClusters: JSON.stringify(getClusters())
				}
				Save.query($data, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}
					$window.history.back()
				})
			}

			$scope.ev.delete = function(ev) {
				var $dialog = $mdDialog.alert()
				.title('Are you ABSOLUTELY sure?')
				.content('This action CANNOT be undone. This will permanently delete this project.')
				.ok('Delete this project')
				.targetEvent(ev)

				$mdDialog.show($dialog).then(function() {
					Delete.query({id: $scope.args.project.id}, function(json){
					if($rootScope.checkErr(json)) {
						return
					}
					$window.history.back()
					})
				})
			}

			function getClusters() {
				var $data = []
				angular.forEach($scope.args.project.projectClusters,  function(v) {
					var cluster = {
						cid: v.cid - 0,
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
