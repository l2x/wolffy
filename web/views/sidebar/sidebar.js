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

            $scope.goto = function (menu) {
                if (!menu || !menu.url) {
                    return
                }
                $location.path(menu.url)
                resetMenuActive($scope.menus)
            }

            function resetMenuActive(menus) {
                var path = $location.path()
                angular.forEach(menus, function (menu, k) {
                    menus[k].active = path == menus[k].url ? true : false
                    angular.forEach(menu.children, function (child, k2) {
                        menus[k].children[k2].active = child.url == path ? true: false
                    })
                })
            }
        }])
});
