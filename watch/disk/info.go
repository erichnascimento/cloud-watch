package disk

import "syscall"

type Info struct {
	Path  string
	Label string
	All   uint64
	Used  uint64
	Free  uint64
}

func (d *Info) UsedPercentage() float64 {
	return (float64(d.Used) / float64(d.All)) * 100
}

func CollectDiskInfo(path string, label string) (*Info, error) {
	buf := syscall.Statfs_t{}
	err := syscall.Statfs(path, &buf)

	if err != nil {
		return nil, err
	}

	all := uint64(buf.Bsize) * buf.Blocks
	free := uint64(buf.Bsize) * buf.Bavail

	info := &Info{
		Path:  path,
		Label: label,
		All:   all,
		Free:  free,
		Used:  all - free,
	}

	return info, nil
}
