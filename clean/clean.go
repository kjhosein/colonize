package clean

import (
	"os"

	"github.com/craigmonson/colonize/config"
	"github.com/craigmonson/colonize/log"
	"github.com/craigmonson/colonize/util"
)

func Run(c *config.Config, l log.Logger, args interface{}) error {
	os.Chdir(c.TmplPath)

	l.Log("Cleaning up")
	filesToClean := []string{
		c.CombinedValsFilePath,
		c.CombinedVarsFilePath,
		c.CombinedTfFilePath,
		c.CombinedDerivedValsFilePath,
		c.CombinedDerivedVarsFilePath,
		c.CombinedRemoteFilePath,
		"terraform.tfplan",
		"destroy.tfplan",
		"terraform.tfstate",
		"terraform.tfstate.backup",
		".terraform",
	}

	for _, file := range filesToClean {
		l.Log("rm -f " + file)
		util.RunCmd("rm", "-f", file)
	}

	return nil
}
