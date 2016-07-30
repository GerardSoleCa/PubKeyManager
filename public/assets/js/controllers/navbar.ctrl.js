/**
 * Created by gerard on 30/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");
    pubKeyManager.controller("NavbarCtrl", ["$scope", "$rootScope", "$location", "backendService",
        function ($scope, $rootScope, $location, backendService) {
            $scope.logout = function () {
                backendService.removeAuthenticatedUser();
                $location.path("/login");
            };

            $rootScope.$on("$routeChangeStart", function (event, next, current) {
                var user = backendService.getAuthenticatedUser();
                if (next.$$route.originalPath == "/" && user) {
                    $scope.user = user;
                } else {
                    $scope.user = undefined;
                }
            });

            console.log($location.path());
            if ($location.path() == "/") {
                $scope.user = backendService.getAuthenticatedUser();
            }
        }]);
})();