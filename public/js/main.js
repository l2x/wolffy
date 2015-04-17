"use strict";

require.config({
    paths: {
        async: 'libs/requirejs-plugin/src/async',
        json: 'libs/requirejs-plugin/src/json',
        text: 'libs/requirejs-plugin/lib/text',
        hammer: 'libs/hammerjs/hammer.min',
        hammerProxy: 'libs/hammerjs/hammer-proxy',
        angular: 'libs/angular/angular.min',
        angularAMD: 'libs/angularAMD/angularAMD.min',
        ngload: 'lib/angularAMD/ngload.min',
        ngRoute: 'libs/angular-route/angular-route.min',
        ngResource: 'libs/angular-resource/angular-resource.min',
        ngMocks: 'libs/angular-mocks/angular-mocks',
        ngAria: 'libs/angular-aria/angular-aria.min',
        ngAnimate: 'libs/angular-animate/angular-animate.min',
        angularLoadingBar: 'libs/angular-loading-bar/build/loading-bar.min',
        ngMaterial: 'libs/angular-material/angular-material',
        ngTranslate: 'libs/angular-translate/angular-translate.min',
        ngTranslateLoader: 'libs/angular-translate-loader-static-files/angular-translate-loader-static-files.min',
        ngSanitize: 'libs/angular-sanitize/angular-sanitize.min'
    },
    shim: {
        angular: {
            exports: 'angular',
            deps: ['hammer']
        },
        ngRoute: ['angular'],
        ngResource: ['angular'],
        angularAMD: ['angular'],
        ngMocks: ['angular'],
        ngAnimate: ['angular'],
        ngAria: ['angular'],
        angularLoadingBar: ['angular', 'ngAnimate'],
        ngload: ['angularAMD'],
        hammerProxy: ['hammer'],
        ngMaterial: ['angular', 'hammerProxy', 'ngAnimate', 'ngAria'],
        ngTranslate: ['angular'],
        ngTranslateLoader: ['angular', 'ngTranslate'],
        ngSanitize: ['angular']
    },
    deps: ['app'],
    baseUrl: 'js/',
	waitSeconds: 20
});

window.name = "NG_DEFER_BOOTSTRAP!";
