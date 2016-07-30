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

            $scope.addKey = function (key) {
                $log.info("KeyListCtrl > AddKey");
                showAddKeyModal(processAddKey);
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

            function showAddKeyModal(key, error, cb) {

                if (!cb) {
                    cb = key;
                    key = undefined;
                    error = undefined;
                }

                var modalInstance = $uibModal.open({
                    animation: true,
                    templateUrl: '/addkey.modal.tpl',
                    controller: 'AddKeyModalCtrl',
                    size: 'lg',
                    resolve: {
                        key: function () {
                            return key;
                        },
                        error: function () {
                            return error;
                        }
                    }
                });
                modalInstance.result.then(function (key) {
                    if (key) {
                        cb(key);
                    }
                });
            }

            function processAddKey(key) {
                backendService.addKey(key).then(function () {
                    backendService.getKeys().then(function (success) {
                        $scope.keys = success;
                    });
                }, function (error) {
                    showAddKeyModal(key, error, processAddKey);
                });
            }

        }]);
})();