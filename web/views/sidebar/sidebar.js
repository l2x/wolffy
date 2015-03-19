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
				{name: "PROJECT", type: 1, url: "/project/list"},
                {
                    name: "ADMIN", type: 0, children: [
                        {name: "CLUSTER", url: "/cluster/list"},
                        {name: "MACHINE", url: "/machine/list"},
                        {name: "USER", url: "/user/list"}
                    ]
                }
            ]
            $scope.menus = menus

            $scope.goto = function (menu) {
                if (!menu || !menu.url) {
                    return
                }
                $location.path(menu.url)
            }

			$scope.$on('$routeChangeSuccess', function(next, current) {
                resetMenuActive($scope.menus, current.originalPath)
			});

            function resetMenuActive(menus, current) {
				var flag = false
                var path = current

                angular.forEach(menus, function (menu, k) {
					if (!path) {
						return
					}

					if (path == menus[k].url) {
						flag = true
						return
					}

                    angular.forEach(menu.children, function (child, k2) {
						if (path == child.url) {
							flag = true
							return
						}
                    })
                })

				if (flag == false) {
					return
				}

                angular.forEach(menus, function (menu, k) {
                    menus[k].active = path == menus[k].url ? true : false
                    angular.forEach(menu.children, function (child, k2) {
                        menus[k].children[k2].active = child.url == path ? true: false
                    })
                })
            }
        }])
});
