package generator

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// CreateServiceFile create service file for contrib
func CreateServiceFile(fileName, binName, desc string) error {
	createDir(fileName)
	svrfile := strings.Replace(svrtpl, "[[bin]]", binName, -1)
	svrfile = strings.Replace(svrfile, "[[desc]]", desc, -1)
	svrfile = strings.Replace(svrfile, "[[u]]", binName, -1)
	if err := ioutil.WriteFile(fileName, []byte(svrfile), 0644); err != nil {
		fmt.Printf("write file [%s] failed:%v\n", fileName, err)
		return err
	}
	return nil
}

var svrtpl = `[Unit]
Description=[[desc]]
After=network-online.target firewalld.service
Wants=network-online.target

[Service]
Type=simple
User=[[u]]
Group=[[u]]
ExecStart=/usr/local/bin/[[bin]]
ExecReload=/bin/kill -s HUP $MAINPID
# LimitNOFILE=1048576
# Having non-zero Limit*s causes performance problems due to accounting overhead
# in the kernel.
LimitNPROC=infinity
LimitCORE=infinity
# Uncomment TasksMax if your systemd version supports it.
# Only systemd 226 and above support this version.
TasksMax=infinity
TimeoutStartSec=0
# set delegate yes so that systemd does not reset the cgroups of docker containers
Delegate=yes
# kill only the process, not all processes in the cgroup
KillMode=process
# restart the process if it exits prematurely
Restart=on-failure
StartLimitBurst=3
StartLimitInterval=60s

[Install]
WantedBy=multi-user.target
`
