"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', 'Project.Save', 'Project.Get', function ($scope, Save, Get) {
			$scope.args = {
				"project":{},
				"clusters":[]
			}
			$scope.ev = {}

			$scope.args.project = {
  "id": 1,
  "name": "项目1",
  "path": "git@xxxx",
  "pushPath": "/var/www/project1",
  "tags": "后台,测试",
  "note": "this is a test.",
  "created": "2015-01-01 12:00:00",
  "modified": "2015-01-01 12:00:00",
  "projectClusters": [
    {
      "id": 12,
      "pid": 1,
      "cid": 1,
      "customMachine": "",
      "bshell": "mkdir logs",
      "eshell": "ln -s /var/www/project1 /usr/www/project1",
      "note": "",
      "created": "2015-01-01 12:00:00",
      "modified": "2015-01-01 13:00:00"
    }
  ]
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

			$scope.ev.saveCluster = function() {
			}

			$scope.ev.save = function() {
				Save.query()
			}

			Get.query({id: 1}, function(json) {

			})

        }];
});
