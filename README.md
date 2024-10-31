# wxwebpack-memi

v0.0.1 | 2024-10-31

> 纯Golang，一个普普通通用于自动化反编译微信小程序的工具，小程序挖洞安全利器。
> 在web端可联动burp插件，findsomething快速读取小程序包中的接口配置。
> 可监听文件夹，进行自动解密，解包。

## 声明

**本程序仅供于学习交流，请使用者遵守《中华人民共和国网络安全法》，勿将此工具用于非授权的测试，开发者不负任何连带法律责任。**

## 介绍&使用
>在终端启动程序，程序会在本地起一个web界面，默认端口3031（通过-p 端口号，可进行设置），自动监听本地的WX小程序目录。

![image](https://github.com/user-attachments/assets/f244243e-a502-4edd-87ea-67273752fa8f)

>测试人员可直接点击小程序，终端自动进行解密、解包，并将已经解密完成的文件内容存放在/uploads/目录，在同目录生成一个upload目录

![image](https://github.com/user-attachments/assets/acddba95-65f9-4c38-85c1-b55ac5b01e34)

![image](https://github.com/user-attachments/assets/52c57ae1-3341-4645-ade2-496e9757cc3a)

![image](https://github.com/user-attachments/assets/76b84c10-9015-4d7e-82bf-fa16fdb7699b)

>删除文件在文件夹中删除，也可以通过/del路径删除，是有点多余
