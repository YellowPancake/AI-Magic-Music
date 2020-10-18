# 基于Hyperledger Fabric开发的简易区块链Web项目
## 安装与部署
### Step 1
进入`bMRP\Scripts`目录下并执行`npm install`

### Step 2
执行`startFabric.sh`脚本，该脚本会启动Docker、安装链码并实例化链码
+ 如果出现权限权限问题请执行`chmod 777 startFabric.sh`和`chmod 777 ../basic-network/start.sh`
+ 再次执行该脚本后请删除`Scripts`目录下的`hfc-key-store`文件夹（如果有）

### Step 3
执行`node server.js`，并使用接口测试工具（如postman，getman),具体url参见route.js
+ 注册管理员：GET方法，http://localhost:8000/cr_admin, 不需要参数
+ 注册用户：POST方法，http://localhost:8000/cr_user, 参数：userID
+ 新增病历：POST方法，http://localhost:8000/cr_cal, 参数：userID, MRID, MRContent
+ 查询病历：POST方法，http://localhost:8000/se_cal, 参数：userID, MRID
## 详细博客地址
https://blog.csdn.net/qq_43326014/article/details/109139785
