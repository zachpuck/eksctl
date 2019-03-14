package utils

import (
	"os"

	"github.com/kris-nova/logger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	defaultaddons "github.com/weaveworks/eksctl/pkg/addons/default"
	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha4"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/eks"
)

func installCoreDNSCmd(g *cmdutils.Grouping) *cobra.Command {
	p := &api.ProviderConfig{}
	cfg := api.NewClusterConfig()

	cmd := &cobra.Command{
		Use:   "install-coredns",
		Short: "Installs latest version of CoreDNS add-on into clusters, replacing kube-dns",
		Run: func(cmd *cobra.Command, args []string) {
			if err := doInstallCoreDNS(p, cfg, cmdutils.GetNameArg(args), cmd); err != nil {
				logger.Critical("%s\n", err.Error())
				os.Exit(1)
			}
		},
	}

	group := g.New(cmd)

	group.InFlagSet("General", func(fs *pflag.FlagSet) {
		fs.StringVarP(&cfg.Metadata.Name, "name", "n", "", "EKS cluster name (required)")
		cmdutils.AddRegionFlag(fs, p)
		cmdutils.AddConfigFileFlag(&clusterConfigFile, fs)
	})

	cmdutils.AddCommonFlagsForAWS(group, p, false)

	group.AddTo(cmd)

	return cmd
}

func doInstallCoreDNS(p *api.ProviderConfig, cfg *api.ClusterConfig, nameArg string, cmd *cobra.Command) error {
	if err := api.Register(); err != nil {
		return err
	}

	if err := cmdutils.LoadMetadata(p, cfg, clusterConfigFile, nameArg, cmd); err != nil {
		return err
	}

	ctl := eks.New(p, cfg)
	meta := cfg.Metadata

	if !ctl.IsSupportedRegion() {
		return cmdutils.ErrUnsupportedRegion(p)
	}
	logger.Info("using region %s", meta.Region)

	if err := ctl.CheckAuth(); err != nil {
		return err
	}

	if err := ctl.GetCredentials(cfg); err != nil {
		return errors.Wrapf(err, "getting credentials for cluster %q", meta.Name)
	}

	// TODO: find out if we must only operate on upgraded 1.11 clusters

	rawClient, err := ctl.NewRawClient(cfg)
	if err != nil {
		return err
	}

	return defaultaddons.InstallCoreDNS(rawClient, meta.Region, ctl.Provider.WaitTimeout())
}
