var sdk = require('./sdk.js');
module.exports = function(app){
  app.get('/api/getWallet', function (req, res) {
    var walletid = req.query.walletid;
    let args = [walletid];
    sdk.send(false, 'getWallet', args, res);
  });
  app.get('/api/setWallet', function(req, res){
    var name = req.query.name;
		var id = req.query.id;
    var coin = req.query.coin;
    let args = [name, id, coin];
    sdk.send(true, 'setWallet', args, res);
  });
  app.get('/api/getMusic', function(req, res){
    var musickey = req.query.musickey;
    let args = [musickey];
    sdk.send(false, 'getMusic', args, res);
  });
  app.get('/api/setMusic', function (req, res) {
    var title = req.query.title;
    var singer = req.query.singer;
    var price = req.query.price;
    var walletid = req.query.walletid;
    let args = [title, singer, price, walletid];
    sdk.send(true, 'setMusic', args, res);
  });
  app.get('/api/getAllmusic', function (req, res) {
    let args = [];
    sdk.send(false, 'getAllMusic', args, res);
  });
  app.get('/api/purchaseMusic', function (req, res) {
    var walletid = req.query.walletid;
    var musickey = req.query.musickey;
    let args = [walletid, musickey];
    sdk.send(true, 'purchaseMusic', args, res);
});
  app.get('/api/changeMusicPrice', function(req, res){
    var musickey = req.query.musickey;
    var price = req.query.price;
    let args = [musickey, price];
    sdk.send(true, 'changeMusicPrice', args, res);
  });
  app.get('/api/deleteMusic', function(req, res){
    var musickey = req.query.musickey;
    let args = [musickey];
    sdk.send(true, 'deleteMusic', args, res);
  });
}