/**
 * Created by gerard on 19/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("LoginCtrl", ["$scope", "$log", "$location", "backendService", "user",
        function ($scope, $log, $location, backendService, user) {
            $scope.user = {};

            $scope.login = function () {
                backendService.login($scope.user).then(function (success) {
                    $log.info("Info > LoginCtrl > Login :: Success -> Redirecting to home");
                    user = success;
                    $location.path("/keys");
                }, function (error) {
                    $log.error("Error > LoginCtrl > Login", error);
                    if (error.status == 401) {
                        $scope.error = "User or Password are wrong";
                    }
                });
            };
        }]);
})();