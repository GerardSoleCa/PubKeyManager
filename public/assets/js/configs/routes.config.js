(function () {
    'use strict';
    var pubKeyManager = angular.module('PubKeyManager');

    pubKeyManager.config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
        $routeProvider
            .when('/login', {
                templateUrl: '/login.tpl',
                controller: 'LoginCtrl',
            })
            .when('/register', {
                templateUrl: '/register.tpl',
                controller: 'RegisterCtrl',
            })
            .when('/keys', {
                templateUrl: '/keylist.tpl',
                controller: 'KeylistCtrl',
            });
        $routeProvider.otherwise('/login');
    }]);
})();