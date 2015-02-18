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
            .when("/", angularAMD.route({
                templateUrl: './views/project_deploy/index.html',
                controllerUrl: '../views/project_deploy/ctrl'
            }))
            .when("/project_push", angularAMD.route({
                templateUrl: './views/project_deploy/index.html',
                controllerUrl: '../views/project_deploy/ctrl'
            }))
            .when("/view1", angularAMD.route({
                templateUrl: './views/view1/index.html',
                controllerUrl: '../views/view1/ctrl'
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
