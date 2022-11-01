# 透過API方式執行ssh命令
## 使用方式POST API 
### 登入認證走 Basic Authentication
輸入SSH 的login and Password
### POST API的範例
```
http
   ### 執行Post CMD
 POST http://localhost:8000/linuxcmd
   Content-Type: application/json
   {
     "Host": "xx.xx.xx.xx",
     "Command": "cd /home;pwd;docker ps -a"
   }
```
