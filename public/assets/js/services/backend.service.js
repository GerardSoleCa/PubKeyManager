/**
 * Created by gerard on 19/07/16.
 */
(function () {
    'use strict';
    var app = angular.module('PubKeyManager');
    app.service('backendService', ['$http', '$q', '$cookies', 'settings',
        function ($http, $q, $cookies, settings) {

            this.register = function (user) {
                var deferred = $q.defer();
                $http.post(settings.register, user).success(function (response) {
                    deferred.resolve(response);
                }).error(function (error) {
                    deferred.reject(error);
                });
                return deferred.promise;
            };

            this.login = function (user) {
                var deferred = $q.defer();
                $http.post(settings.login, user).success(function (response) {
                    $cookies.put('user', JSON.stringify(response));
                    deferred.resolve(response);
                }).error(function (error) {
                    deferred.reject(error);
                });
                return deferred.promise;
            };

            this.getKeys = function () {
                var deferred = $q.defer();
                $http.get(settings.keys).success(function (response) {
                    deferred.resolve(response);
                }).error(function (error) {
                    deferred.reject(error);
                });
                return deferred.promise;
            };

            this.deleteKey = function (id) {
                var deferred = $q.defer();
                $http.delete(settings.keys + id).success(function (response) {
                    deferred.resolve(response);
                }).error(function (error) {
                    deferred.reject(error);
                });
                return deferred.promise;
            };

            this.addKey = function (key) {
                var deferred = $q.defer();
                $http.post(settings.keys, key).success(function (response) {
                    deferred.resolve(response);
                }).error(function (error) {
                    deferred.reject(error);
                });
                return deferred.promise;
            };

            this.getAuthenticatedUser = function () {
                if ($cookies.get("authenticated")) {
                    return JSON.parse($cookies.get("user"));
                } else {
                    return undefined;
                }
            };

            this.removeAuthenticatedUser = function () {
                $cookies.remove("authenticated");
                $cookies.remove("user");
            };
        }]);
})();