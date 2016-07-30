/**
 * Created by gerard on 19/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("RegisterCtrl", ["$scope", "$log", "$location", "backendService",
        function ($scope, $log, $location, backendService) {
            $scope.user = {};

            $scope.register = function () {
                backendService.register($scope.user).then(function (success) {
                    $log.info("Info > RegisterCtrl > Register :: Success -> Redirecting to login");
                    $location.path("/login");
                }, function (error) {
                    $log.error("Error > RegisterCtrl > Register", error);
                    $scope.error = error.error;
                });
            };
        }]);
})();