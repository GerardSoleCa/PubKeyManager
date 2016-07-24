/**
 * Created by gerard on 24/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("AddKeyModalCtrl", ["$scope", "$uibModalInstance",
        function ($scope, $uibModalInstance) {

            $scope.key = {};

            $scope.ok = function () {
                $uibModalInstance.close($scope.key);
            };

            $scope.cancel = function () {
                $uibModalInstance.dismiss('cancel');
            };
        }]);
})();