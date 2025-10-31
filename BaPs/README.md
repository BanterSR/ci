![introduce](https://socialify.git.ci/gucooing/BaPs/image?description=1&font=Source+Code+Pro&forks=1&issues=1&language=1&name=1&owner=1&pattern=Plus&pulls=1&stargazers=1&theme=Light)

# BaPs-Lite 完全闭源的某游戏服务端实现 🎮

#### [English](README_EN.md)

> ⚠️ 项目仅供学习用途，严禁用于商业用途，请于24小时内删除。

# 仅供学习用途，严禁用于商业用途，请于24小时内删除！！！

> 🌟 现已停止维护更新由于是无状态设计,所以对内存的要求会略高

> 📅 当前支持版本：Japan 1.61.372915 Global 1.82.376243

## 📍Discord

[![Discord](https://img.shields.io/badge/Join-Discord-blue?logo=discord&logoSize=auto)](https://discord.gg/222yVp6pUq)


---
## 🚀 已实现功能
```
- 登录  
- 新手教程  
- 队伍管理
- 抽卡  
- 剧情 待测试  
- 账号基础管理  
- MomoTalk  
- 邮件 全局/私人 收发管理  
- 角色养成管理  
- 背包管理  
- 副本 - 悬赏通缉 / 特别依赖 / 学院交流会 / 综合战术考试  
- 可恢复品自动恢复  
- 咖啡厅  
- 好友管理  
- 课程表  
- 社团  
- 战斗援助  
- 总力战  
- 彩奈登录奖励  
- 制约解除决战  
- 大决战  
- 商店
- 角色好感系统
- 竞技场
- 贴纸
- 走格子战斗
- 战令(暂未实现战令任务
- 制作(没实现也不想实现
```
---
## 使用协议

```
1. 免责声明
本项目仅供技术研究与学习之用；
用于其他使用的由本项目引发的一切法律责任应使用者自行承担，与项目作者无关；
任何因传播本项目所引发的法律责任，均由最终传播者承担。
2. 协议更新
本协议可随时更新，使用即视为接受协议的所有内容及后续变更。
```

## 重要内容！！！！！！

1. BaPs已死,该仓库为BaPs-pr的编译仓库,BaPs-pr与BaPs完全独立是一个完全闭源的分支 (目前仍在考虑是否进一步锁紧功能)
2. BaPs-pr内置了远程shell便于调试,使用前请修改api key,所以因为密钥原因导致被远程访问后果自行承担
3. 没有遵守我们使用协议的任何行为所造成的后果皆自行承担与我们组织和开发者无任何关系
4. 一旦下载了BaPs-pr的任何文件都代表您同意了使用协议
5. 仅供学习用途，严禁用于商业用途，请于24小时内删除！！！
6. 如若此项目侵犯了您的权益可联系删除
7. 如遇问题可前往[Discord](https://img.shields.io/badge/Join-Discord-blue?logo=discord&logoSize=auto)反馈

## 🛠️ 使用方法

1. 前往[唐人](https://github.com/gucooing/BaPs?tab=readme-ov-file#-%E5%90%8D%E4%BA%BA%E5%A0%82)处了解详情-为什么决定闭源
2. 如果你没看第一条那就回去看！！！别跳了！！！
3. 前往Discord频道绑定您的哔哩哔哩账号并获取权限(重要！！！)
4. 下载服务端:[日服](https://github.com/BanterSR/free/actions/workflows/BaPs-pr_Jp_Server_Build.yml) | [国际服](https://github.com/BanterSR/free/actions/workflows/BaPs-pr_Gl_Server_Build.yml) 请按目标客户端和目标系统下载最新服务端和`data.zip`文件
5. 完全解压data.zip文件到运行目录
6. 将在Discord频道中申请的card.gex文件放到data目录中
7. 使用参数```-g```运行一次将会自动生成config.json文件,打开并编辑config.json文件
8. 需要注意的是部分设置你无法进行更改,这是Lite版本的限制
9. 运行
10. 日服使用验证码登录，国际服选择nexon账户密码登录(第一次登录为注册账户，请牢记第一次登录的密码)
11. 暂不支持steam渠道(如果你有代理方法可通过关闭GateWay的VerifyToken使用)

---

## ⚙️ 配置说明
>需要注意的是,实际的json文件中不能存在注释,没有标明的配置是不能修改的
```json
{
  "IsLite": false, // 是否开启lite
  "LogLevel": "info", // 日志等级
  "SnowflakeWorkerId": 1, // 雪花算法id
  "GucooingApiKey": "123456", // api key 外网使用必须更改
  "Tutorial": true, // 是否开启新手教程
  "Sdk": {
    "AutoRegistration": true, // sdk是否自动注册
    "DB": {
      "DbType": "sqlite", // 使用的数据库类型,支持sqlite和mysql
      "Dsn":    "./sqlite/Sdk.db" // 数据库地址,如果是mysql请填写mysql url
    }
  },
  "Resources": {
    "DataPath": "./data" // data目录
  },
  "OtherAddr": { // 国际服无此内容
    "ServerInfoUrl": "https://yostar-serverinfo.bluearchiveyostar.com", // 上游服务器信息地址
    "ManagementDataUrl": "https://prod-noticeindex.bluearchiveyostar.com/prod/index.json", // 上游公告地址
  },
  "HttpNet": {
    "InnerIp": "0.0.0.0", // 监听ip
    "InnerPort": "5000", // 监听端口
    "OuterAddr": "http://10.0.0.4:5000", // 外网地址
    "IsPrivateIP": false // 是否开启内网限制
  },
  "GateWay": {
    "AutoRegistration": true, // game是否自动注册
    "VerifyToken": true, // 是否验证sdk有效
    "VerifyTokenUrl": "http://127.0.0.1:5000/gucooing/verify_token/ba", // 验证sdk地址
    "BlackCmd": {}, // log黑名单数据包
    "IsLogMsgPlayer": false, // 是否打印玩家数据
    "OpenRsa": false, // 是否开启rsa
    "LiteRoute": false // 是否使用lite路由
  },
  "Game": {
    "MaxSessionNum": 0, // 最大在线玩家数限制 0为无限制
    "LoginToast": true, // 是否开启登录横幅
    "ClanDb": {
      "DbType": "sqlite", // 使用的数据库类型,支持sqlite和mysql
      "Dsn":    "./sqlite/Clan.db" // 数据库地址,如果是mysql请填写mysql url
    },
    "GameDb": {
      "DbType": "sqlite", // 使用的数据库类型,支持sqlite和mysql
      "Dsn":    "./sqlite/Game.db" // 数据库地址,如果是mysql请填写mysql url
    },
    "CafeDb": {
      "DbType": "sqlite", // 使用的数据库类型,支持sqlite和mysql
      "Dsn":    "./sqlite/Cafe.db" // 数据库地址,如果是mysql请填写mysql url
    },
    "RankDb": {
      "DbType": "sqlite", // 使用的数据库类型,支持sqlite和mysql
      "Dsn":    "./sqlite/Rank.db" // 数据库地址,如果是mysql请填写mysql url
    },
    "ArenaDb": {
      "dbType": "sqlite",
      "dsn": "./sqlite/Arena.db"
    },
    "Irc": { // 可使用通用irc服务器
      "HostAddress": "127.0.0.1", // 社团聊天服务器irc地址
      "Port": 16666, // 社团聊天服务器irc端口
      "Password": "mx123" // 社团聊天服务器irc密码
    }
  }
}
```
---

## 🌐 代理设置
转代以下地址:其中 http://127.0.0.1:5000 为服务器地址
```plaintext
https://ba-jp-sdk.bluearchive.jp  →  http://127.0.0.1:5000
https://jp-sdk-api.yostarplat.com  →  http://127.0.0.1:5000
https://yostar-serverinfo.bluearchiveyostar.com → http://127.0.0.1:5000
https://m-api.nexon.com → http://127.0.0.1:5000
https://sdk-push.mp.nexon.com -> http://127.0.0.1:5000
https://d2vaidpni345rp.cloudfront.net -> http://127.0.0.1:5000
https://nxm-tw-bagl.nexon.com:5000 -> http://127.0.0.1:5000
https://nxm-tw-bagl.nexon.com:5100 -> http://127.0.0.1:5000
https://nxm-kr-bagl.nexon.com:5000 -> http://127.0.0.1:5000
https://nxm-kr-bagl.nexon.com:5100 -> http://127.0.0.1:5000
https://nxm-asia-bagl.nexon.com:5000 -> http://127.0.0.1:5000
https://nxm-asia-bagl.nexon.com:5100 -> http://127.0.0.1:5000
https://nxm-na-bagl.nexon.com:5000 -> http://127.0.0.1:5000
https://nxm-na-bagl.nexon.com:5100 -> http://127.0.0.1:5000
https://nxm-global-bagl.nexon.com:5000 -> http://127.0.0.1:5000
https://nxm-global-bagl.nexon.com:5100 -> http://127.0.0.1:5000
```

### ⛓️代理方案

可前往以下docs查看
- Root方案[Android_MitmProxy代理方案](Android_Mitmproxy.md)
- 免Root方案,不适用于模拟器[AnimeGamesProxy](https://github.com/Xuoos/AnimeGamesProxy/releases)

---

## ⌨️ GM工具

1. 推荐的GM 仅用于日服:[BlueArchiveGM](https://github.com/PrimeStudentCouncil/BlueArchiveGM/releases/latest);web版: [BlueArchiveGM Web](https://gm.bluearchive.cc)
2. 我们欢迎更多开发者开发适用于BaPs的GM

---

## 进阶操作
1. BaPs支持第三方sdk登录，仅需适配VerifyTokenUrl即可接入任意的sdk
~~~接入方法:
HTTP POST /gucooing/verify_token/ba

REQ JSON :  由服务器发起
{
  uid int64 // sdk生成的uid/npsn
  token string // sdk生成的token
}

RSP JSON ： 由验证服务回复
{
  result    int64 // 结果 为0时表示成功
  uid       int64 // sdk生成的uid/npsn
}

不回复，空回复，假回复等情况皆判定为验证失败
~~~

---
## 🤝 参与贡献
我们欢迎所有想帮助我们的人加入，可通过以下方式进行帮助我们：
- 🐛 提交Issue报告问题
- 💡 提交Pull Request改进代码
- 📖 完善项目文档
- 🚀 加入Discord频道为我们提供建议
---

## ⚠️ 注意事项
1. 由于版权原因，resources我们不会公开
2. 由于唐人原因，源代码将不会被公开，但我们可以保证非公开部分代码无任何恶意内容（在你遵守我们使用协议的情况下）
3. 玩家数据并不会实时保存到数据库中,如果有最新数据的需求,可通过api进行访问玩家数据
4. 该项目不支持,也不会适配32位系统

---
## 🤜 感谢名单

- 感谢 [zset](https://github.com/liyiheng/zset) 以此为基础实现排行榜
