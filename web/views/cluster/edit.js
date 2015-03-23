"use strict";

define(['app', '../service/cluster', '../filter/filter'], function (app) {
	return ['$scope', '$rootScope', '$route','$window', 'Cluster.Get', 'Cluster.Save',
		function ($scope, $rootScope, $route, $window, Get, Save) {
			$scope.args = {}
			$scope.args.cluster = {}
			$scope.ev = {}

			var $id = $route.current.params.id
			if ($id) {
				Get.query({id: $id}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}
					$scope.args.cluster = json.data
					var ip  = []
				angular.forEach($scope.args.cluster.machines.split(","), function(v) {
						ip.push({ip: v})
					})
					$scope.args.cluster.machines = ip

					console.log($scope.args.cluster)
				})
			}

			$scope.args.machines = [
			{ip:"127.0.1"},
			{ip:"127.0.2"},
			{ip:"127.0.3"}
			]

			$scope.ev.addMachine = function($idx) {
				if (!$scope.args.cluster.machines) {
					$scope.args.cluster.machines = []
				}

				for(var i=0; i<$scope.args.cluster.machines.length; i++) {
					if($scope.args.cluster.machines[i].ip == $scope.args.machines[$idx].ip) {
						return
					}
				}

				$scope.args.cluster.machines.push($scope.args.machines[$idx])
			}

			$scope.ev.removeMachine = function($idx) {
				$scope.args.cluster.machines.splice($idx, 1)
			}

			$scope.ev.save = function() {
				if(!$scope.clusterform.$valid) {
					return false
				}
				var data = {}
				angular.copy($scope.args.cluster, data)
				var ip = []
				angular.forEach(data.machines, function(v) {
					ip.push(v.ip)
				})
				data.machines = ip.join(",")

				Save.query(data, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}
					$window.history.back()
				})
			}

        }];
});
