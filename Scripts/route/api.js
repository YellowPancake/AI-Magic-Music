var fs = require('fs');
var url = require('url');
var querystring = require('querystring');
var foods = require('../model/foods')();
var detail = require('../model/detail')();

var getHandler = {};
var postHandler = {};

// 处理对主页的请求
getHandler['/'] = function(req, res) {
  var foodMenu = "";
  // 拼装首页数据
  var food = foods.getAllFoods();
  for (var i = 0; i < food.length; ++i) {
      foodMenu += '<div class="food-card" id="' + food[i].id + '"><img src="';
      foodMenu += food[i].image + '"><h1>' + food[i].name + '</h1><h2>' + food[i].price + '</h2></div>';
  }

  res.writeHead(200, {"Content-Type": "text/html"});
  fs.readFile(__dirname + '/../views/index.html', (err, data) => {
    if (err) {
      console.log(err);
      res.end();
    } else {
      // 动态渲染模板
      res.end(data.toString().replace('{{foodMenu}}', foodMenu));
    }
  });
};

// 处理对详情页面的请求
getHandler['/detail'] = function(req, res) {
  var query = querystring.parse(url.parse(req.url).query);
  var foodDetail = detail.getDetail(query.id);
  res.writeHead(200, {"Content-Type": "text/html"});
  fs.readFile(__dirname + '/../views/detail.html', (err, data) => {
    // 动态渲染模板
    res.end(data.toString().replace('{{image}}', foodDetail.image)
      .replace('{{name}}', foodDetail.name)
      .replace('{{description}}', foodDetail.description)
      .replace('{{price}}', foodDetail.price));
  });
};

// 404响应，告知客户端资源未找到
getHandler['/404'] = function(req, res) {
  res.writeHead(404, {"Content-Type": "text/plain"});
  res.end("404 Not Found");
};

// post请求的处理方法示例
postHandler['/'] = function(res, data) {
  // do something
};

// get请求
function get(req, res) {
  var reqUrl = url.parse(req.url);
  if (typeof getHandler[reqUrl.pathname] === "function") {
    getHandler[reqUrl.pathname](req, res);
  } else {
    getHandler["/404"](req, res);
  }
}

// post请求（示例）
function post(req, res) {
  var reqUrl = url.parse(req.url);
  if (typeof postHandler[reqUrl.pathname] === "function") {
    var postData = "";
    req.on('data', (data) => {
      postData += data;
    });
    req.on('end', () => {
      postData = querystring.parse(postData);
      postHandler[reqUrl.pathname](res, postData);
    });
  } else {
    getHandler["/404"](req, res);
  }
}

// 提供给其他模块使用的接口
module.exports = {
  get: get,
  post: post
};
