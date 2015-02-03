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
    'angularLoadingBar'

], function (angularAMD, sidebar) {
    var app = angular.module("myApp", [
        'ngRoute',
        'ngAria',
        'ngAnimate',
        'ngResource',
        'angular-loading-bar',
        'ngMaterial'
    ]);

    app.config(['$routeProvider', '$locationProvider', 'cfpLoadingBarProvider',
        function($routeProvider, $locationProvider, cfpLoadingBarProvider) {
        $routeProvider
            .when("/", angularAMD.route({
                templateUrl: './views/latest/index.html',
                controllerUrl: '../views/latest/ctrl'
            }))
            .when("/view1", angularAMD.route({
                templateUrl: './views/view1/index.html',
                controllerUrl: '../views/view1/ctrl'
            }))
            .otherwise({redirectTo: '/'});

        //$locationProvider.html5Mode(true);
        cfpLoadingBarProvider.includeSpinner = false;
    }]);



    return angularAMD.bootstrap(app);
});