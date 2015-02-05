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
                if(!menu || !menu.url) {
                    return
                }

                resetMenuActive($scope.menus, menu)
                $location.path(menu.url)
            }

            function resetMenuActive(menus, menu) {
                angular.forEach(menus, function(menu, k) {
                    menus[k].active = false
                    angular.forEach(menu.children, function(child, k2){
                        menus[k].children[k2].active = false
                    })
                })
                menu.active = true
            }
        }])
});
