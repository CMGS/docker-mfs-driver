MooseFS Volume Plugin
=====================

使用 [MooseFS](https://moosefs.com/index.html) 作为存储后端的 docker volume 插件。

#### 编译

`make build`

如果需要打包成 rpm 的话执行

`make rpm`

#### 运行

使用这样的命令来运行插件:

`mfs --mfs-base /mnt/mfs/apppermdirs`

那么插件会运行在 `/run/docker/plugins/mfs.sock`, 用 root 跑, 并且 base 指定为上面的参数。

#### 使用

在运行容器时如果 `-v test.10001:/home/test/permdirs`, 那么会在宿主机创建 `/mnt/mfs/apppermdirs/test` 并且 chown 成 10001, 然后挂载到容器内的 `/home/test/permdirs`.
