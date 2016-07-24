(function () {
    'use strict';

    var app = angular.module('PubKeyManager');
    app.constant('settings', {
        login: "/auth/login",
        register: "/auth/register"

    });
})();