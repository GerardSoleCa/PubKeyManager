/**
 * Created by gerard on 24/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("KeyModalCtrl", ["$scope", "$uibModalInstance", "key",
        function ($scope, $uibModalInstance, key) {
            $scope.key = key;

            $scope.close = function () {
                $uibModalInstance.dismiss('cancel');
            };
        }]);
})();