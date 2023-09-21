
说明：
    封装一个go的web server，接口和gin 一样，但是不依赖gin，底层是gnet的实现。




git 提交错误

如果没有代理，取消代理设置

git config --global --unset http.proxy
git config --global --unset https.proxy
git config --global --unset  http.https://github.com.proxy
git config --global --unset https.https://github.com.proxy


如果有代理，设置代理
git config --global http.https://github.com.proxy 'socks5://172.16.2.1:10808'
git config --global https.https://github.com.proxy 'socks5://172.16.2.1:10808'

git config --global http.https://github.com.proxy http://127.0.0.1:10809
git config --global https.https://github.com.proxy http://127.0.0.1:10809



git init

git add .
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/jianhl/jhttp.git
git push -u origin main


git remote add origin https://github.com/jianhl/jhttp.git
git branch -M main
git push -u origin main



git add .
git commit -m "first commit"
git push -u origin main
