package main

import (
	"context"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}
	privilegeEscalationOptions := &options.AnsiblePrivilegeEscalationOptions{
		Become:       true,
		// BecomeMethod: "sudo",
	}
	cmd := &playbook.AnsiblePlaybookCmd{
		Playbooks:                  []string{"testvm.yaml"},
		ConnectionOptions:          ansiblePlaybookConnectionOptions,
		Options:                    ansiblePlaybookOptions,
		PrivilegeEscalationOptions: privilegeEscalationOptions,
	}

	err := cmd.Run(context.TODO())
	if err != nil {
		panic(err)
	}

}
