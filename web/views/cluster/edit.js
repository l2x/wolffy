"use strict";

define(['app', '../service/cluster'], function (app) {
    return ['$scope', function ($scope) {
			$scope.args = {
				cluster:{}
			}
			$scope.ev = {}

			$scope.args.cluster = {
				"name" : "test",
				"tags":"后台",
				"machine": [
			{"ip":"127.0.0.1", "status":1},
			{"ip":"127.0.0.2", "status":1},
			{"ip":"127.0.0.2", "status":1}
	]
			}

			$scope.args.machines = [
			{"id":1, "ip":"127.0.0.1", "status":1},
			{"id":2, "ip":"127.0.0.2", "status":1},
			{"id":3, "ip":"127.0.0.2", "status":1}
			]

			$scope.ev.save = function() {
				Save.query()
			}

        }];
});
