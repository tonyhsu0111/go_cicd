# 透過API方式執行ssh命令
## 使用方式POST API 
### 登入認證走 Basic Authentication
輸入SSH 的login and Password
### POST API的範例
```
http
   ### 執行Post CMD
   如果沒有設定Port，預設為22
 POST http://localhost:8400/linuxcmd
   Content-Type: application/json
   {
     "Host": "xx.xx.xx.xx",
     "Port": 30000,
     "Command": "cd /home;pwd;docker ps -a"
   }
```
