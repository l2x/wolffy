define(['angularAMD'], function (angularAMD) {
    angularAMD.controller("sidebarCtrl", ['$scope', '$rootScope', '$timeout', '$mdSidenav', '$anchorScroll', '$location',
        function ($scope, $rootScope, $timeout, $mdSidenav, $anchorScroll, $location) {
            $rootScope.toggleLeft = function () {
                $mdSidenav('left').toggle()
            };

            $scope.close = function () {
                $mdSidenav('left').close()
            };
            var menus = [
                {name: "PROJECT_PUSH", type: 1, url: "/project_push"},
                {
                    name: "ADMIN", type: 0, children: [
                        {name: "PROJECT", url: "/view1"},
                        {name: "CLUSTER", url: "/"},
                        {name: "USER", url: "/"}
                    ]
                }
            ]

            $scope.menus = menus

            $scope.goto = function (url) {
                $location.path(url)
            }
        }])
});
