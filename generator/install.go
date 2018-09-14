package generator

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// CreateInstallFile create install file
func CreateInstallFile(installfile, binName string) error {
	createDir(installfile)
	insfile := strings.Replace(installtpl, "[[bin]]", binName, -1)
	insfile = strings.Replace(insfile, "[[u]]", binName, -1)
	if err := ioutil.WriteFile(installfile, []byte(insfile), 0644); err != nil {
		fmt.Printf("write file [%s] failed:%v\n", installfile, err)
		return err
	}
	return nil
}

var installtpl = `#!/bin/bash

EXEC_BIN=[[bin]]
SERVICE_FILE=[[bin]].service
SERVICE_PATH=/lib/systemd/system/
CONFIG_PATH=/etc/[[bin]]/
LOCAL_BIN=/usr/local/bin/

CUSER=$(who am i | cut -d' ' -f1)
CGROUP=$(groups $CUSER | cut -d' ' -f1)

##############################################
# Check if run as root
##############################################
if [ xroot != x$(whoami) ]
then
    echo "You must run as root (Hint: Try prefix 'sudo' while execution the script)"
    exit
fi

echo "To install will stop the [[bin]] service. (Y/N)?\c"
read answer
if [ $answer = 'y' ] || [ $answer = 'Y' ]; then
    echo "Contine installation..."
else
    echo "Abort install..."
    exit
fi

if [ ! -x $EXEC_BIN ]; then
    echo "Not executable file found, check it."
    exit
fi

useradd [[u]] -U -M -s /sbin/nologin
cp $SERVICE_FILE $SERVICE_PATH
cp $EXEC_BIN $LOCAL_BIN
service $EXEC_BIN start

echo "[[bin]] has successfully installed."
echo "run service [[bin]] restart to make it work."
`
