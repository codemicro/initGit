package directoryTree

import (
	"fmt"
	"os"
	"path/filepath"
)

type Difference struct {
	dir     string
	initial []string
}

func NewDifference(dir string) (*Difference, error) {
	nd := new(Difference)
	nd.dir = dir
	return nd, nd.initialScan()
}

func (d *Difference) walkDir() ([]string, error) {
	var o []string
	err := filepath.Walk(d.dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fmt.Println(path, "is not dir")
				o = append(o, path)
			} else {
				fmt.Println(path, "is dir")
			}
			return nil
		})
	return o, err
}

func (d *Difference) initialScan() error {
	initial, err := d.walkDir()
	d.initial = initial
	return err
}

func (d *Difference) Get() ([]string, error) {
	currentState, err := d.walkDir()
	if err != nil {
		return nil, err
	}
	return sliceDifference(currentState, d.initial), nil
}

// slice b should be the slice with more items
func sliceDifference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}