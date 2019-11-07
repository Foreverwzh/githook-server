# githook-server

<pre>

git clone https://github.com/Foreverwzh/githook-server.git
</pre>

<pre>
//将id_rsa文件拷贝到本项目.ssh文件夹下

cd githook-server
mkdir .ssh
cp $HOME/.ssh/id_rsa .ssh/id_rsa
</pre>

<pre>
//创建docker镜像githook-server

docker build --rm -f "dockerfile" -t githook-server .
</pre>

将需要自动拉取的项目放到 /data 目录下

<pre>
运行镜像githook-server

docker run --restart=always -v /data:/data -p 3000:8080 -dit githook-server
</pre>
