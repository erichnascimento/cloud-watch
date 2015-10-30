package watch

import "syscall"

type DiskInfo struct {
	Path  string
	Label string
	All   uint64
	Used  uint64
	Free  uint64
}

func (d *DiskInfo) UsedPercentage() float64 {
	return (float64(d.Used) / float64(d.All)) * 100
}

func CollectDiskInfo(path string, label string) (*DiskInfo, error) {
	buf := syscall.Statfs_t{}
	err := syscall.Statfs(path, &buf)

	if err != nil {
		return nil, err
	}

	all := uint64(buf.Bsize) * buf.Blocks
	free := uint64(buf.Bsize) * buf.Bfree

	info := &DiskInfo{
		Path:  path,
		Label: label,
		All:   all,
		Free:  free,
		Used:  all - free,
	}

	return info, nil
}
