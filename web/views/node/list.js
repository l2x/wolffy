"use strict";

define(['app', '../service/node'], function (app) {
    return ['$scope', '$rootScope', '$mdDialog', 'Node.GetAll', 'Node.Delete',
		function ($scope, $rootScope, $mdDialog, GetAll, Delete) {
			$scope.args = {}
			$scope.args.list = []
			$scope.ev = {}

			GetAll.query({}, function (json) {
				if ($rootScope.checkErr(json)) {
					return
				}
				$scope.args.list = json.data
			})

			$scope.ev.delete = function (ev, $idx) {
				var dialog = $rootScope.confirmDialog.targetEvent(ev)
				$mdDialog.show(dialog).then(function () {
					var item = $scope.args.list[$idx]
					Delete.query({id: item.id}, function (json) {
						if ($rootScope.checkErr(json)) {
							return
						}
						$scope.args.list.splice($idx, 1)
					})
				})
			}
    }];
});
