"use strict";

define(['app', '../service/project'], function (app) {
    return ['$scope', 'Project.Search', function ($scope, Search) {
			$scope.args = {}
			$scope.ev = {}

			$scope.deployList = [
				{name:"后台", tags:"后台, 消息", version:"1.1.1", created:"2014-01-01 12:00:00", modified:"2014-01-02 12:00:00"},
				{name:"后台2", tags:"后台, 消息", version:"1.1.2", created:"2014-06-01 19:00:00", modified:"2014-06-02 17:00:00"}
			]

			$scope.ev.search = function(){
				Search.query({keywords: $scope.args.keywords}, function(json){
					console.log(json)
				})
			}

        }];
});
