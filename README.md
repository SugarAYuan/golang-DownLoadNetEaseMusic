# 下载网易云音乐付费歌曲

> 通过实时监控Mac下网易云音乐的日志，然后分析出相应的下载链接进行下载，启动程序后只需要打开网易云音乐听相应的付费歌曲即可自动下载。

> downLoadNetEaseMusic 为编译好的文件 直接运行即可
> -p参数指定网易云音乐的日志文件夹 -d参数为指定下载好的音乐存入哪个文件
`
./downLoadNetEaseMusic -p ~/Library/Containers/com.netease.163music/Data/Documents/storage/Logs/ -d ./
`
