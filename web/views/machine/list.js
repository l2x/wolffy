"use strict";

define(['app', '../service/cluster'], function (app) {
    return ['$scope', 'Cluster.Search', function ($scope, Search) {
			$scope.args = {}
			$scope.ev = {}

			$scope.args.list = [
				{"ip":"127.0.0.1", status:1, created:"2014-01-01 12:00:00", modified:"2014-01-02 12:00:00"},
				{"ip":"127.0.0.2",status:0, created:"2014-06-01 19:00:00", modified:"2014-06-02 17:00:00"}
			]

			$scope.ev.search = function(){
				Search.query({keywords: $scope.args.keywords}, function(json){
					console.log(json)
				})
			}

        }];
});
