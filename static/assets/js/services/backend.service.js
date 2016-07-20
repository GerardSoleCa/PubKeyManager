/**
 * Created by gerard on 19/07/16.
 */
(function () {
    'use strict';
    var app = angular.module('PubKeyManager');
    app.service('backendService', ['$http', 'settings', function ($http, settings) {

        this.register = function (user) {
            return $http.post(settings.remote).then(function (response) {
                return response.data;
            });
        };

        this.login = function (user) {
            return $http.get(settings.remote).then(function (response) {
                return response.data;
            });
        };

    }]);
})();