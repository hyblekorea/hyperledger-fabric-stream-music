'use strict';
var app = angular.module('application', []);
app.controller('AppCtrl', function($scope, appFactory){
        $("#success_setmusic").hide();
        $("#success_getallmusic").hide();
        $("#success_getmusic").hide();
        $("#success_getwallet").hide();
        $("#success_changemusicprice").hide();
        $("#success_deletemusic").hide();
        $scope.getWallet = function(){
                appFactory.getWallet($scope.walletid, function(data){
                        $scope.search_wallet = data;
                        $("#success_getwallet").show();
                });
        }
       $scope.getAllMusic = function(){
                appFactory.getAllMusic(function(data){
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                parseInt(data[i].Key);
                                data[i].Record.Key = data[i].Key;
                                array.push(data[i].Record);
                                $("#success_getallmusic").hide();
                        }
                        array.sort(function(a, b) {
                            return parseFloat(a.Key) - parseFloat(b.Key);
                        });
                        $scope.allMusic = array;
                });
        }
        $scope.getMusic = function(){
                appFactory.getMusic($scope.musickey, function(data){
                        $("#success_getmusic").show();
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                data[i].Key = $scope.musickey;
                                data[i].title = data[i].Title;
                                data[i].singer = data[i].Singer;
                                data[i].price = data[i].Price;
                                data[i].walletid = data[i].WalletID;
                                data[i].count = data[i].Count;
                                array.push(data[i]);
                        }
                        $scope.allMusic = array;
                });
        }
        $scope.setMusic = function(){
            appFactory.setMusic($scope.music, function(data){
                        $scope.create_music = data;
                        $("#success_setmusic").show();
            });
        }
        $scope.purchaseMusic = function(key){
                appFactory.purchaseMusic(key, function(data){
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                parseInt(data[i].Key);
                                data[i].Record.Key = data[i].Key;
                                array.push(data[i].Record);
                                $("#success_getallmusic").hide();
                        }
                        array.sort(function(a, b) {
                            return parseFloat(a.Key) - parseFloat(b.Key);
                        });
                        $scope.allMusic = array;
                });
        }
        $scope.changeMusicPrice = function(){
                appFactory.changeMusicPrice($scope.change, function(data){
                        $scope.change_music_price = data;
                        $("#success_changemusicprice").show();
                });
        }
        $scope.deleteMusic = function(){
                appFactory.deleteMusic($scope.musickeydelete, function(data){
                        $scope.delete_music = data;
                        $("#success_deletemusic").show();
                });
        }
});
 app.factory('appFactory', function($http){
        var factory = {};
        factory.getWallet = function(key, callback){
            $http.get('/api/getWallet?walletid='+key).success(function(output){
                        callback(output)
                });
        }
        factory.getAllMusic = function(callback){
            $http.get('/api/getAllMusic/').success(function(output){
                        callback(output)
                });
        }
        factory.getMusic = function(key, callback){
            $http.get('/api/getMusic?musickey='+key).success(function(output){
                        callback(output)
                });
        }
        factory.setMusic = function(data, callback){
            $http.get('/api/setMusic?title='+data.title+'&singer='+data.singer+'&price='+data.price+'&walletid='+data.walletid).success(function(output){
                        callback(output)
                });
        }
        factory.purchaseMusic = function(key, callback){
            $http.get('/api/purchaseMusic?walletid=5T6Y7U8I&musickey='+key).success(function(output){
                $http.get('/api/getAllMusic/').success(function(output){
                        callback(output)
                });
            });
        }
        factory.changeMusicPrice = function(data, callback){
            $http.get('/api/changeMusicPrice?musickey='+data.musickey+'&price='+data.price).success(function(output){
                        callback(output)
                });
        }
        factory.deleteMusic = function(key, callback){
            $http.get('/api/deleteMusic?musickey='+key).success(function(output){
                        callback(output)
                });
        }
        return factory;
});