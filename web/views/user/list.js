"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', 'Project.Search', function ($scope, Search) {
			$scope.args = {}
			$scope.ev = {}

			$scope.args.list = [
{id:1, username:"123", name:"test user", created:"2014-01-01 12:00:00", "last_login_ip":"127.0.0.1", modified:"2015-01-01 12:00:00"}
			]

			$scope.ev.search = function(){
				Search.query({keywords: $scope.args.keywords}, function(json){
					console.log(json)
				})
			}

        }];
});
