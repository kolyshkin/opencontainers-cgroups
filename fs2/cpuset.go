package fs2

import (
	"github.com/opencontainers/cgroups"
	"github.com/opencontainers/cgroups/configs"
)

func isCpusetSet(r *configs.Resources) bool {
	return r.CpusetCpus != "" || r.CpusetMems != ""
}

func setCpuset(dirPath string, r *configs.Resources) error {
	if !isCpusetSet(r) {
		return nil
	}

	if r.CpusetCpus != "" {
		if err := cgroups.WriteFile(dirPath, "cpuset.cpus", r.CpusetCpus); err != nil {
			return err
		}
	}
	if r.CpusetMems != "" {
		if err := cgroups.WriteFile(dirPath, "cpuset.mems", r.CpusetMems); err != nil {
			return err
		}
	}
	return nil
}
