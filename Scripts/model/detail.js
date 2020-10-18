var fs = require('fs');

module.exports = function() {
  // 读取文件中的数据，将其转成一个对象方便使用
  var data = JSON.parse(fs.readFileSync(__dirname + '/../data/detail.json'));
  var foods = {
    getDetail: getDetail
  };

  // 获取对应id的食品详情
  function getDetail(id) {
    for (var i = 0; i < data.detail.length; ++i) {
      if (data.detail[i].id == id)
        return data.detail[i];
    }
  }

  return foods;
};
