"use strict";

define(['app', '../service/project', '../service/deploy'], function (app) {
    return ['$scope','$rootScope', '$route', '$mdDialog', 'Project.Get','Project.GetTags', 'Deploy.History', 'Deploy.AddTag',
	function ($scope, $rootScope, $route, $mdDialog, Project_Get, Project_GetTags, History, AddTag) {
			$scope.args = {}
			$scope.ev = {}
			$scope.args.project = {}
			$scope.args.list = []
			$scope.args.tag  = ""

			var $id = $route.current.params.id
			if($id) {
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

				History.query({id: $id}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}

					$scope.args.list = json.data
				})
			}

			$scope.ev.addTag = function() {
				var btag = getPrevTag($scope.args.tag)
				AddTag.query({id: $scope.args.project.id, tag: $scope.args.tag, btag: btag}, function(json) {
					if($rootScope.checkErr(json)) {
						return
					}

					History.query({id: $id}, function(json) {
						if($rootScope.checkErr(json)) {
							return
						}

						$scope.args.list = json.data
					})

				})
			}

			function getPrevTag($tag) {
				if($scope.args.list.length > 0) {
					return $scope.args.list[0].commit
				}
				return ""
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
