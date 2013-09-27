package chroot

import (
	"fmt"
	"os/exec"
	"log"
)

func copySingle(dest string, src string, copyCommand string) error {
	cpCommand := fmt.Sprintf("%s %s %s", copyCommand, src, dest)
	localCmd := exec.Command("/bin/sh", "-c", cpCommand)
	log.Printf("Executing copy: %s %#v", localCmd.Path, localCmd.Args)
	if err := localCmd.Run(); err != nil {
		return err
	}
	return nil
}
