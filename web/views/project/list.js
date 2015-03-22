"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', 'Project.Search', 'Project.GetAll', function ($scope, Search, GetAll) {
			$scope.args = {}
			$scope.ev = {}

			$scope.args.list = []

			GetAll.query({}, function(json) {
				if (!json || json.errno != 0) {
					return
				}

				$scope.args.list = json.data
			})

			$scope.ev.search = function(){
				Search.query({keywords: $scope.args.keywords}, function(json){
					console.log(json)
				})
			}

        }];
});
