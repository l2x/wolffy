"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', 'Project.Save', function ($scope, Save) {
			$scope.args = {}
			$scope.ev = {}

			$scope.ev.save = function() {
				Save.query()
			}

        }];
});
