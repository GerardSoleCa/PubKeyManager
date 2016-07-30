/**
 * Created by gerard on 30/07/16.
 */
(function () {
    'use strict';

    var pubKeyManager = angular.module("PubKeyManager");

    pubKeyManager.factory("httpResponseInterceptor", ["$q", "$location",
        function ($q, $location) {

            return {
                responseError: function (rejection) {
                    if (rejection.status == 401) {
                        $location.path("/login");
                    }
                    return $q.reject(rejection);
                }
            };

        }]);

    pubKeyManager.config(['$httpProvider', function ($httpProvider) {
        $httpProvider.interceptors.push('httpResponseInterceptor');
    }]);
})();