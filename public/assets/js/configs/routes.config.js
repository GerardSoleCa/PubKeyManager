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
            .when('/', {
                templateUrl: '/keylist.tpl',
                controller: 'KeylistCtrl',
                requireLogin: true
            });
        $routeProvider.otherwise('/login');
    }]);

    pubKeyManager.run(["$rootScope", "$location", "backendService",
        function ($rootScope, $location, backendService) {
        $rootScope.$on("$routeChangeStart", function (event, next, current) {
            if (next.requireLogin && !backendService.getAuthenticatedUser()) {
                $location.path('/login');
            }
        });
    }]);
})();