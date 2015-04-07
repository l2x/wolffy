"use strict";

define(['app'], function (app) {

    app.filter('ipFilter', function () {
        return function (ip, machines) {
            if (!machines || !ip) {
                return true
            }
            for (var $i = 0; $i < machines.length; $i++) {
                if (machines[$i].ip == ip) {
                    return false
                }
            }
            return true
        }
    })

    app.filter('clusterFilter', function () {
        return function (id, list) {
            if (!id || !list) {
                return true
            }

            for (var $i = 0; $i < list.length; $i++) {
                if ((list[$i].cid - 0) == (id - 0)) {
                    return false
                }
            }

            return true
        }
    })
})
