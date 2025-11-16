# lolo

[![Discord](https://img.shields.io/badge/Join-Discord-blue?logo=discord&logoSize=auto)](https://discord.gg/QMMY8yKh)

开放空间-箱庭游戏服务端部分实现

支持的客户端版本:cn-最新(ios/windows/android/是cn客户端都支持)

## 使用方法
1. ~~编译Lolo `go build main.go` 得到main或main.exe~~ 去[action](https://github.com/BanterSR/ci/actions/workflows/Lolo_Build.yml)下载最新版本
2. ~~下载运行必要文件 [data](https://github.com/BanterSR/Lolo/tree/main/data)和[LoloResource](https://github.com/BanterSR/LoloResource) 分别保存到`data`文件夹和`Resource`文件夹中~~ 去[action](https://github.com/BanterSR/ci/actions/workflows/Lolo_Build.yml)下载最新版本的data 下载Resource(貌似没公开库)
3. 用各种方法代理客户端http:`http://dsp-prod-of.inutan.com:18881/dispatch/region_info` 到服务器http （默认:`http://127.0.0.1:18881`）
4. ~~使用`main/main.exe -g`生成`config.json`修改`config.json`文件~~ 使用`Lolo/Lolo.exe -g`生成`config.json`修改`config.json`文件
5. 启动Lolo和客户端

## 已实现功能
- [x] 登录
  - [ ] sdk登录
  - [x] 服务发现
- [x] 角色
  - [x] 换装
  - [x] 更新装备
  - [x] 角色升级
  - [x] 角色进阶
- [x] 场景
  - [x] 多人游戏
  - [x] 切换房间
  - [x] 切换场景
  - [x] 时间系统
  - [x] 天气系统
  - [ ] 互动
- [x] 背包
- [ ] 聊天
  - [ ] 附近
  - [ ] 世界
  - [ ] 系统
  - [ ] 私聊

## 参与贡献
- 欢迎