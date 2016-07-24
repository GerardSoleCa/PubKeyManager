/**
 * Created by gerard on 19/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("KeylistCtrl", ["$scope", "$log", "$location", "backendService", "user",
        function ($scope, $log, $location, backendService, user) {

            backendService.getKeys().then(function (success) {
                $scope.keys = success;
            });

            $scope.deleteKey = function (key) {
                $log.info("KeyListCtrl > DeleteKey > " + key.id);
            };

            $scope.showKey = function (key) {
                $log.info("KeyListCtrl > ShowKey > " + key.id);
            };

        }]);
})();