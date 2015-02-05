"use strict";

define(['app'], function (app) {
    return ['$scope',
        function ($scope) {
            $scope.name = "haha"

            $scope.country = {};
            $scope.countries = [
                {name: 'Venezuela', code: 'VE'},
                {name: 'Vietnam', code: 'VN'},
                {name: 'Virgin Islands, British', code: 'VG'},
                {name: 'Virgin Islands, U.S.', code: 'VI'},
                {name: 'Wallis and Futuna', code: 'WF'},
                {name: 'Western Sahara', code: 'EH'},
                {name: 'Yemen', code: 'YE'},
                {name: 'Zambia', code: 'ZM'},
                {name: 'Zimbabwe', code: 'ZW'}
            ];
        }];
});
