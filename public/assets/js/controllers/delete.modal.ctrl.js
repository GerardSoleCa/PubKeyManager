/**
 * Created by gerard on 24/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("DeleteModalCtrl", ["$scope", "$uibModalInstance",
        function ($scope, $uibModalInstance) {
            $scope.ok = function () {
                $uibModalInstance.close(true);
            };

            $scope.cancel = function () {
                $uibModalInstance.dismiss('cancel');
            };
        }]);
})();