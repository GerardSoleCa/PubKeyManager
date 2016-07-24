/**
 * Created by gerard on 19/07/16.
 */
(function () {
    'use strict';
    var app = angular.module('PubKeyManager');
    app.service('backendService', ['$http', '$q', 'settings', function ($http, $q, settings) {

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
            return $http.post(settings.login, user).then(function (response) {
                return response.data;
            });
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
    }]);
})();