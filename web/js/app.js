"use strict";

define([
    'angularAMD',
    '../views/sidebar/sidebar',
    'hammer',
    'ngRoute',
    'ngMocks',
    'ngResource',
    'ngAnimate',
    'ngAria',
    'ngMaterial',
    'angularLoadingBar',
    'ngTranslate',
    'ngTranslateLoader'
], function (angularAMD, sidebar) {
    var app = angular.module("myApp", [
        'ngRoute',
        'ngAria',
        'ngAnimate',
        'ngResource',
        'angular-loading-bar',
        'ngMaterial',
        'pascalprecht.translate'
    ]);

    app.config(['$routeProvider', '$locationProvider', 'cfpLoadingBarProvider',
        function($routeProvider, $locationProvider, cfpLoadingBarProvider) {
        $routeProvider
			.when("/deploy/list", angularAMD.route({
                templateUrl: './views/deploy/list.html',
                controllerUrl: '../views/deploy/list'
            }))
            .when("/project/list", angularAMD.route({
				templateUrl: './views/project/list.html',
				controllerUrl: '../views/project/list'
            }))
            .when("/project/edit", angularAMD.route({
				templateUrl: './views/project/edit.html',
				controllerUrl: '../views/project/edit'
            }))
            .when("/project/edit/:id", angularAMD.route({
				templateUrl: './views/project/edit.html',
				controllerUrl: '../views/project/edit'
            }))
			.when("/cluster/list", angularAMD.route({
				templateUrl: './views/cluster/list.html',
				controllerUrl: '../views/cluster/list'
            }))
			.when("/cluster/edit", angularAMD.route({
				templateUrl: './views/cluster/edit.html',
				controllerUrl: '../views/cluster/edit'
            }))
			.when("/cluster/edit/:id", angularAMD.route({
				templateUrl: './views/cluster/edit.html',
				controllerUrl: '../views/cluster/edit'
            }))
			.when("/machine/list", angularAMD.route({
				templateUrl: './views/machine/list.html',
				controllerUrl: '../views/machine/list'
            }))
			.when("/machine/add", angularAMD.route({
				templateUrl: './views/machine/add.html',
            }))
			.when("/machine/add/:id", angularAMD.route({
				templateUrl: './views/machine/add.html',
            }))
			.when("/user/list", angularAMD.route({
				templateUrl: './views/user/list.html',
				controllerUrl: '../views/user/list'
            }))
			.when("/user/edit/:id", angularAMD.route({
				templateUrl: './views/user/edit.html',
				controllerUrl: '../views/user/edit'
            }))
			.when("/user/edit", angularAMD.route({
				templateUrl: './views/user/edit.html',
				controllerUrl: '../views/user/edit'
            }))
			.when("/login", angularAMD.route({
				templateUrl: './views/index/login.html',
				controllerUrl: '../views/index/login'
            }))
			.when("/logout", angularAMD.route({
				templateUrl: './views/index/logout.html',
				controllerUrl: '../views/index/logout'
            }))
            .otherwise({redirectTo: '/'});

        //$locationProvider.html5Mode(true);
        cfpLoadingBarProvider.includeSpinner = false;

    }]);

    app.config(['$translateProvider',
        function($translateProvider) {
            $translateProvider.useStaticFilesLoader({
                prefix: './languages/',
                suffix: '.json'
            });

            $translateProvider.determinePreferredLanguage(function () {
                var supportLanguage = [
                    'zh-cn'
                ];

                return 'zh-cn'
            });
        }]);

    return angularAMD.bootstrap(app);
});
