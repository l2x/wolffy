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
        function ($routeProvider, $locationProvider, cfpLoadingBarProvider) {
            $routeProvider
                .when("/deploy/list/:id", angularAMD.route({
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
                .when("/node/list", angularAMD.route({
                    templateUrl: './views/node/list.html',
                    controllerUrl: '../views/node/list'
                }))
                .when("/node/add", angularAMD.route({
                    templateUrl: './views/node/add.html',
                }))
                .when("/node/add/:id", angularAMD.route({
                    templateUrl: './views/node/add.html',
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
                .when("/user/changepwd", angularAMD.route({
                    templateUrl: './views/user/changepwd.html',
                    controllerUrl: '../views/user/changepwd'
                }))
                .when("/login", angularAMD.route({
                    templateUrl: './views/user/login.html',
                    controllerUrl: '../views/user/login'
                }))
                .when("/logout", angularAMD.route({
                    templateUrl: './views/user/logout.html',
                    controllerUrl: '../views/user/logout'
                }))
                .otherwise({redirectTo: '/project/list'});

            //$locationProvider.html5Mode(true);
            cfpLoadingBarProvider.includeSpinner = false;
        }]);

    app.config(['$translateProvider',
        function ($translateProvider) {
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

    app.run(function ($rootScope, $location, $mdDialog, $mdToast, $animate) {
        $rootScope.checkErr = function (json) {
            if (json && json.errno == 401) {
                $location.path("/login")
            }
            if (json && json.errno == 1) {
                console.log(json);
                $mdToast.show(
                    $mdToast.simple().content(json.errmsg).action("close").hideDelay(0)
                );
                return true
            }
            return false
        }
        $rootScope.confirmDialog = $mdDialog.confirm()
            .title('Are you ABSOLUTELY sure?')
            .content('This action CANNOT be undone. This will permanently delete this item.')
            .ok('Delete')
            .cancel('Cancel')

        $rootScope.loadingDialog = {
            tpl: '<md-dialog aria-label="">' +
            '<md-content class="">' +
            '<div layout="row" layout="row" layout-align="center center">' +
            '<md-progress-circular md-mode="indeterminate"></md-progress-circular>' +
            '<p>{{ctrl.content}}</p>' +
            '</div> ' +
            '</md-content>' +
            '</md-dialog>',
            show: function (content) {
                var self = this;
                var dialog = {
                    controller: self.showCtrl,
                    template: self.tpl,
                    bindToController: true,
                    controllerAs: "ctrl",
                    locals: {
                        content: content
                    }
                }
                $mdDialog.show(dialog).then()
            },
            hide: function () {
                $mdDialog.hide()
            },
            showCtrl: function ($scope, $mdDialog) {
            }
        }
    })

    return angularAMD.bootstrap(app);
});
