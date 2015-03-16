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

			$scope.ev.addLatest = function($idx) {
			}

			$scope.ev.showDiff = function(ev, $idx) {
				var dialog = {
				  controller: DialogController,
				  template: document.getElementById('diffTpl').innerHTML,
				  targetEvent: ev,
				  bindToController:true,
				  controllerAs:"ctrl",
				  locals: {
					  diff: $scope.args.list[$idx].diff
				  }
				}
				$mdDialog.show(dialog).then()
			}

			$scope.ev.showStatus = function(ev, $idx) {
			}

			$scope.ev.deploy = function(ev, $id, $commit) {

			}

			function DialogController($scope, $mdDialog) {
			  $scope.hide = function() {
				$mdDialog.hide()
			  }
			}


        }];
});
