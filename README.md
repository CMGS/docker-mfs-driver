MooseFS Volume Plugin
=====================

mfs 的 docker volume 插件.

如果这样运行插件:

`mfs --mfs-base /mnt/mfs/apppermdirs`

那么插件会运行在 `/run/docker/plugins/mfs.sock`, 用 root 跑, 并且 base 指定为上面的参数.

-v 的时候如果 `-v test.10001:/home/test/permdirs`, 那么会在宿主机创建 `/mnt/mfs/apppermdirs/test` 并且 chown 成 10001, 然后挂载到容器内的 `/home/test/permdirs`.
