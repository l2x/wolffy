"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', 'Project.Save', 'Project.Get', function ($scope, Save, Get) {
			$scope.args = {}
			$scope.ev = {}

			$scope.ev.save = function() {
				Save.query()
			}

			Get.query({id: 1}, function(json) {

			})

        }];
});
