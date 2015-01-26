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
    'ngStrap',
    'ngStrapTpl'

], function (angularAMD, sidebar) {
    var app = angular.module("myApp", [
        'ngRoute',
        'ngAria',
        'ngAnimate',
        'ngResource',
        'angular-loading-bar',
        'ngMaterial',
        'mgcrea.ngStrap'
    ]);

    app.config(['$routeProvider', '$locationProvider', 'cfpLoadingBarProvider', function($routeProvider, $locationProvider, cfpLoadingBarProvider) {
        $routeProvider
            .when("/view1", angularAMD.route({
                templateUrl: './views/view1/index.html',
                controllerUrl: '../views/view1/ctrl'
            }))
            .otherwise({redirectTo: '/view1'});

        //$locationProvider.html5Mode(true);
        cfpLoadingBarProvider.includeSpinner = false;
    }]);


    return angularAMD.bootstrap(app);
});