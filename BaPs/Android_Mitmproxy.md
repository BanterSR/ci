## Mitmproxy方案使用教程

## 前置要求
 1. [下载 mitmproxy](https://mitmproxy.org/) 并安装。
 2. 具备 WireGuard 和 Python 脚本的基本知识。
 3. 一台客户端设备（例如 Android 模拟器或手机）以及运行 `mitmproxy` 的主机。
 >本次将以模拟器的方式进行教程

### 第一步：安装mitmproxy
- Linux/Mac
```markdown
  # Ubuntu/Debian

  sudo apt update
  sudo apt install mitmproxy

  # macOS
  brew install mitmproxy
  ```
- **Windows**: 从 [mitmproxy.org](https://mitmproxy.org/) 下载 `.exe` 安装程序，并按说明完成安装。

运行以下命令验证安装是否成功：
```bash
mitmproxy --version
```

---

### 第二步：具体教程

为了解密 HTTPS 流量，客户端需要信任 `mitmproxy` 的 CA 证书。


 1. 启动 `mitmproxy`生成证书 run：`mitmdump`
 2. 在电脑端的C:\Users\用户\ .mitmproxy安装电脑证书（mitmproxy-ca.p12）
 3. mitmproxy的目录下会有mitmproxy-ca-cert.cer
 4. 将 mitmproxy-ca-cert.cer 重命名为 c8750f0d.0
 5. 将证书移动到安卓系统 CA 目录,在adb下依次执行下面指令：（需要注意的是，部分安卓系统不支持该方法，请自行解决证书信任问题）

     `adb root`

     `adb remount`

     `adb shell mv /sdcard/c8750f0d.0 /system/etc/security/cacerts/`

     `adb shell chmod 644 /system/etc/security/cacerts/c8750f0d.0`
 6. 重启设备以生效更改，部分系统无法通过上面的指令信任证书，请采用其他方法信任证书
 7. 客户端设备安装WireGuard客户端，如安卓[WireGuard](https://play.google.com/store/apps/details?id=com.wireguard.android),其他平台请自行查询安装方法和证书信任方法
 8. 修改脚本文件参考:[点击跳转](#修改脚本)
 9. pc在 redirect_server.py 文件目录中运行该指令`mitmweb -m wireguard --no-http2 -s redirect_server.py --set termlog_verbosity=warn --ignore 这里输入你的IP地址(对应的服务器地址)`
 10. pc浏览器打开`http://localhost:8081` 访问 `mitmweb` 界面监控流量。
 11. 打开 WireGuard 客户端，点击左下角＋号，选择扫描Mitmproxy浏览器页面上的二维码（没有的话在设置里面）
 12. 启用添加的WireGuard配置，打开游戏

#### 后续进入只需要执行第 9、12 两步即可

---

### 修改脚本

请将下面内容粘贴到 redirect_server.py 文件中，并修改 http://127.0.0.1:5000 为实际的服务器地址
```python
from mitmproxy import http

redirects = {
    "https://ba-jp-sdk.bluearchive.jp": "http://127.0.0.1:5000",
    "https://jp-sdk-api.yostarplat.com": "http://127.0.0.1:5000",
    "https://yostar-serverinfo.bluearchiveyostar.com": "http://127.0.0.1:5000",
}

def request(flow: http.HTTPFlow) -> None:
    for original_url, redirected_url in redirects.items():
        if flow.request.pretty_url.startswith(original_url):
            flow.request.url = flow.request.pretty_url.replace(original_url, redirected_url)
            print(f"Redirecting {original_url} to {redirected_url}")
            return
```

---

## 故障排查

### Client TLS handshake failed. The client does not trust the proxy's certificate for xxx.com (OpenSSL Error([('SSL routines', '', 'ssl/tls alert certificate unknown')]))
- 确保电脑端以及客户端证书为内容一样的
- 确保双端安装了Mitmproxy证书
- 部分安卓系统不支持上诉方法信任证书，请自行解决，比如使用面具模块此处不做复述

### 安卓端安装后证书消失？
- 可以使用MT管理器授予SU权限
- 然后前往/system/etc/security/cacerts/
- 找到 c8750f0d.0 给予 664 权限。用户组为 root

### （手机端）无权限修改系统目录？
- 安装证书到用户证书
- 然后安装模块自动将用户证书转到系统证书
- 之后重启手机查看有没有相关证书
---