"use strict";

define(['app', '../service/cluster'], function (app) {
    return ['$scope', 'Cluster.Search', function ($scope, Search) {
			$scope.args = {}
			$scope.ev = {}

			$scope.args.list = [
				{name:"集群1", tags:"后台", created:"2014-01-01 12:00:00", modified:"2014-01-02 12:00:00"},
				{name:"后台2", tags:"后台, 消息", created:"2014-06-01 19:00:00", modified:"2014-06-02 17:00:00"}
			]

			$scope.ev.search = function(){
				Search.query({keywords: $scope.args.keywords}, function(json){
					console.log(json)
				})
			}

        }];
});
