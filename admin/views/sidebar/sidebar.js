define(['angularAMD'], function (angularAMD) {
    angularAMD.controller("sidebarCtrl", ['$scope', '$rootScope', '$timeout', '$mdSidenav', '$anchorScroll', '$location',
        function($scope, $rootScope, $timeout, $mdSidenav, $anchorScroll, $location){
        $rootScope.toggleLeft = function() {
            $mdSidenav('left').toggle()
        };

        $scope.close = function() {
            $mdSidenav('left').close()
        };
        var menu = []
        for(var i=0;i<20;i++) {
            menu.push({id:i, name:"item" + i})
        }

        $scope.menu = menu

        $scope.gotoAnchor = function (x) {


        }
    }])



});