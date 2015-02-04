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
                {name: "项目发布", type: 1, url: "/project_push"},
                {
                    name: "管理", type: 0, children: [
                        {name: "项目", url: "/view1"},
                        {name: "集群", url: "/"},
                        {name: "用户", url: "/"}
                    ]
                }
            ]

            $scope.menus = menus

            $scope.goto = function (url) {
                $location.path(url)
            }
        }])
});
