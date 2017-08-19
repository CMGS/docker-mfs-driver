package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/docker/go-plugins-helpers/volume"
)

type mfsDriver struct {
	base string
}

func newMFSDriver(mfsBase string) mfsDriver {
	return mfsDriver{base: mfsBase}
}

func (d mfsDriver) volumeInfo(name string) (string, int, error) {
	// 如果是一个真实存在的路径那就直接返回路径
	// 让docker自己去mount去
	if strings.HasPrefix(name, "/") {
		return name, 0, nil
	}

	// 如果是appname.appuid, 那就按照规则创建mfs的文件夹
	// 并且chown成appuid
	parts := strings.Split(name, ".")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("Volume name should be in format 'appname.appuid'")
	}

	uid, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("App uid must be digit, got %s", parts[1])
	}

	return filepath.Join(d.base, parts[0]), uid, nil
}

// 啥都不搞
func (d mfsDriver) Create(r *volume.CreateRequest) error {
	return nil
}

// 啥都不搞
func (d mfsDriver) Remove(r *volume.RemoveRequest) error {
	return nil
}

// volume 格式是 appname.appuid
// 挂载主机地址是 base+/appname
func (d mfsDriver) Path(r *volume.PathRequest) (*volume.PathResponse, error) {
	mountpoint, _, err := d.volumeInfo(r.Name)
	return &volume.PathResponse{Mountpoint: mountpoint}, err
}

// 挂载上去
// 创建目标目录, 改用户
func (d mfsDriver) Mount(r *volume.MountRequest) (*volume.MountResponse, error) {
	mountpoint, uid, err := d.volumeInfo(r.Name)
	if err != nil {
		return &volume.MountResponse{}, err
	}

	// 看看是不是存在
	// 不存在就给创建一个, 改成uid的owner
	fi, err := os.Lstat(mountpoint)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(mountpoint, 0755); err != nil {
			return &volume.MountResponse{}, err
		}
		if err := os.Chown(mountpoint, uid, uid); err != nil {
			return &volume.MountResponse{}, err
		}
	} else if err != nil {
		return &volume.MountResponse{Err: err.Error()}, err
	}

	if fi != nil && !fi.IsDir() {
		return &volume.MountResponse{}, fmt.Errorf("%v already exist and it's not a directory", mountpoint)
	}

	return &volume.MountResponse{Mountpoint: mountpoint}, nil
}

// 就啥都不做吧
// 毕竟不是 unmoumt 了就删掉文件夹的
func (d mfsDriver) Unmount(r *volume.UnmountRequest) error {
	return nil
}

// 啥都不搞
func (d mfsDriver) Get(r *volume.GetRequest) (*volume.GetResponse, error) {
	return &volume.GetResponse{}, nil
}

// 也啥都不搞
func (d mfsDriver) List() (*volume.ListResponse, error) {
	return &volume.ListResponse{}, nil
}

// 就默认的就行
func (d mfsDriver) Capabilities() *volume.CapabilitiesResponse {
	return &volume.CapabilitiesResponse{
		Err:          "",
		Capabilities: volume.Capability{Scope: "global"},
	}
}
