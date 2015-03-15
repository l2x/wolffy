"use strict";

define(['app'], function (app) {
    return ['$scope', '$mdDialog', function ($scope, $mdDialog) {
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

			$scope.showDiff = function(ev, $idx) {
				var dialog = {
				  controller: DialogController,
				  template: document.getElementById('diffTpl').innerHTML,
				  targetEvent: ev,
				}
				.then()
			}

			function DialogController($scope, $mdDialog) {
			  $scope.hide = function() {
				$mdDialog.hide()
			  }
			}


        }];
});
