"use strict";

define(['app'], function (app) {
    return ['$scope', function ($scope) {
			$scope.args = {}
			$scope.ev = {}

			$scope.args.project = {
				name:"test project"
			}

			$scope.args.list = [
			{
				commit:"123456",
				diff:"test",
				status:0,
				created:"2015-01-01 12:00:00",
			}
			]

        }];
});
