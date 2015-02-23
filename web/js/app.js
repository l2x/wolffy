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
            .when("/deploy", angularAMD.route({
				templateUrl: './views/deploy/index.html',
                controllerUrl: '../views/deploy/index'
            }))
            .when("/project/list", angularAMD.route({
				templateUrl: './views/project/list.html',
				controllerUrl: '../views/project/list'
            }))
            .when("/project/edit", angularAMD.route({
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
            .when("/view1", angularAMD.route({
                templateUrl: './views/view1/index.html',
                controllerUrl: '../views/view1/ctrl'
            }))
            .otherwise({redirectTo: '/deploy/list'});

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
