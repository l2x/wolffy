"use strict";

define(['app', '../service/cluster', '../filter/filter'], function (app) {
    return ['$scope', function ($scope) {
			$scope.args = {
				cluster:{}
			}
			$scope.ev = {}

			$scope.args.cluster = {
				"name" : "test",
				"tags":"后台",
				"machines": [
			{"ip":"127.0.0.1", "status":1},
			{"ip":"127.0.0.2", "status":1}
	]
			}

			$scope.args.machines = [
			{"id":1, "ip":"127.0.0.1", "status":1},
			{"id":2, "ip":"127.0.0.2", "status":1},
			{"id":4, "ip":"127.0.0.3", "status":1}
			]

			$scope.ev.addMachine = function($idx) {
				for(var i=0; i<$scope.args.cluster.machines.length; i++) {
					if($scope.args.cluster.machines[i].ip == $scope.args.machines[$idx].ip) {
						return
					}
				}

				$scope.args.cluster.machines.push($scope.args.machines[$idx])
			}

			$scope.ev.save = function() {
				Save.query()
			}

        }];
});
