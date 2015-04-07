"use strict";

define(['app'], function (app) {
    app.factory('User.GetAll', function ($resource) {
        return $resource("/user/getall", {}, {
            query: {isArray: false}
        });
    })
    app.factory('User.Delete', function ($resource) {
        return $resource("/user/delete", {}, {
            query: {isArray: false}
        });
    })
    app.factory('User.Get', function ($resource) {
        return $resource("/user/get", {}, {
            query: {isArray: false}
        });
    })
    app.factory('User.Save', function ($resource) {
        return $resource("/user/edit", {}, {
            query: {isArray: false}
        });
    })
    app.factory('User.Login', function ($resource) {
        return $resource("/user/login", {}, {
            query: {isArray: false}
        });
    })
    app.factory('User.Changepwd', function ($resource) {
        return $resource("/user/updatepassword", {}, {
            query: {isArray: false}
        });
    })
})
