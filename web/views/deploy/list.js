"use strict";

define(['app', '../service/project', '../service/deploy'], function (app) {
    return ['$scope','$rootScope', '$route', '$window', '$mdDialog', 'Project.Get','Project.GetTags', 'Deploy.History', 'Deploy.AddTag', 'Deploy.Push', 'Deploy.GetDiff', 'Deploy.HistoryDetail',
	function ($scope, $rootScope, $route, $window, $mdDialog, Project_Get, Project_GetTags, History, AddTag, Push, GetDiff, HistoryDetail) {
			$scope.args = {}
			$scope.ev = {}
			$scope.args.project = {}
			$scope.args.list = []
			$scope.args.tag  = ""

			var $id = $route.current.params.id
			if($id) {
				History.query({id: $id}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}

					$scope.args.list = json.data
				})

				Project_Get.query({id: $id}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}

					$scope.args.project = json.data
				})

				Project_GetTags.query({id: $id}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}

					$scope.args.tags = json.data
					$scope.args.tag = $scope.args.tags[0]
				})
			}

			$scope.ev.addTag = function() {
				var btag = getPrevTag($scope.args.tag)
				AddTag.query({id: $scope.args.project.id, tag: $scope.args.tag, btag: btag}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}

					$route.reload()
				})
			}

			function getPrevTag($tag) {
				if($scope.args.list.length > 0) {
					return $scope.args.list[0].commit
				}
				return ""
			}

			$scope.ev.showDiff = function(ev, id) {
				var dialog = {
				  controller: showDiffController,
				  template: document.getElementById('diffTpl').innerHTML,
				  targetEvent: ev,
				  bindToController:true,
				  controllerAs:"ctrl",
				  locals: {
					  id: id,
					  diff: "loading..."
				  }
				}
				$mdDialog.show(dialog).then()
			}

			$scope.ev.showStatus = function(ev, id) {
				var dialog = {
				  controller: showStatusController,
				  template: document.getElementById('statusTpl').innerHTML,
				  targetEvent: ev,
				  bindToController:true,
				  controllerAs:"ctrl",
				  locals: {
					  id: id,
					  list: []
				  }
				}
				$mdDialog.show(dialog).then()
			}

			$scope.ev.deploy = function(ev, id, $commit) {
				Push.query({pid: $scope.args.project.id, id: id, commit: $commit}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}
					$route.reload()
				})
			}

			function showDiffController(scope, $mdDialog, id) {
				GetDiff.query({id: id}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}
					scope.ctrl.diff = json.data.diff
				})

				scope.hide = function() {
					$mdDialog.hide()
				}
			}
			function showStatusController(scope, $mdDialog, id) {
				HistoryDetail.query({id: id}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}
					scope.ctrl.list = json.data
					console.log(scope.ctrl.list)
				})

				scope.hide = function() {
					$mdDialog.hide()
				}
			}


        }];
});
