package model

import (
	"fmt"
	"strings"
)

type Directory struct {
	ParentDirectory *Directory
	SubDirectory    map[string]*Directory
	Files           []*File
	Name            string
	TotalSize       int
}

func (d Directory) String() string {
	return d.stringHelper(0)
}

func (d Directory) stringHelper(counter int) string {
	var sb strings.Builder

	tempCtr := counter
	for tempCtr > 0 {
		sb.WriteString("  ")
		tempCtr--
	}
	sb.WriteString("dir " + d.Name + " " + fmt.Sprintf("%d", d.TotalSize) + "\n")
	if len(d.SubDirectory) > 0 {
		for key := range d.SubDirectory {
			sb.WriteString(d.SubDirectory[key].stringHelper(counter+1) + "\n")
		}
	}

	if len(d.Files) > 0 {
		for idx := range d.Files {
			tempCtr = counter + 1
			for tempCtr > 0 {
				sb.WriteString("  ")
				tempCtr--
			}

			sb.WriteString(d.Files[idx].String() + "\n")
		}
	}

	return sb.String()
}
