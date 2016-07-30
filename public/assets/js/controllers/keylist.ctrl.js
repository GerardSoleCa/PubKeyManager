/**
 * Created by gerard on 19/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.controller("KeylistCtrl", ["$scope", "$log", "$location", "$uibModal", "backendService",
        function ($scope, $log, $location, $uibModal, backendService) {

            backendService.getKeys().then(function (success) {
                $scope.keys = success;
            });

            $scope.addKey = function () {
                $log.info("KeyListCtrl > AddKey");
                showAddKeyModal(function (key) {
                    backendService.addKey(key).then(function () {
                        backendService.getKeys().then(function (success) {
                            $scope.keys = success;
                        });
                    });
                });
            };

            $scope.deleteKey = function (key) {
                $log.info("KeyListCtrl > DeleteKey > " + key.id);
                showDeleteModal(function () {
                    backendService.deleteKey(key.id).then(function (success) {
                        backendService.getKeys().then(function (success) {
                            $scope.keys = success;
                        });
                    });
                });
            };

            $scope.showKey = function (key) {
                $log.info("KeyListCtrl > ShowKey > " + key.id);
                showKeyModal(key);
            };

            function showKeyModal(key) {
                $uibModal.open({
                    animation: true,
                    templateUrl: '/key.modal.tpl',
                    controller: 'KeyModalCtrl',
                    size: 'lg',
                    resolve: {
                        key: function () {
                            return key;
                        }
                    }
                });
            }

            function showDeleteModal(cb) {
                var modalInstance = $uibModal.open({
                    animation: true,
                    templateUrl: '/delete.modal.tpl',
                    controller: 'DeleteModalCtrl',
                    size: 'sm'
                });
                modalInstance.result.then(function (result) {
                    if (result) {
                        cb();
                    }
                });
            }

            function showAddKeyModal(cb) {

                var modalInstance = $uibModal.open({
                    animation: true,
                    templateUrl: '/addkey.modal.tpl',
                    controller: 'AddKeyModalCtrl',
                    size: 'lg'
                });
                modalInstance.result.then(function (key) {
                    if (key) {
                        cb(key);
                    }
                });
            }

        }]);
})();