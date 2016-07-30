/**
 * Created by gerard on 24/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("AddKeyModalCtrl", ["$scope", "$uibModalInstance", "key", "error",
        function ($scope, $uibModalInstance, key, error) {

            $scope.key = key;
            $scope.error = error;

            $scope.ok = function () {
                $uibModalInstance.close($scope.key);
            };

            $scope.cancel = function () {
                $uibModalInstance.dismiss('cancel');
            };
        }]);
})();