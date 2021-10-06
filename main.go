// Copyright © 2021 Genome Research Limited
// Author: Sendu Bala <sb10@sanger.ac.uk>.
//
// This file is part of iseqfile.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	run, lane, tag, ftype := parseArgs(os.Args[1:])

	paths := run_iquest(run)

	path := filter_iquest(paths, lane, tag, ftype)

	fmt.Printf(format_path_like_imeta(path))
}

func parseArgs(args []string) (run, lane, tag, ftype string) {
	for i, val := range args {
		switch val {
		case "id_run":
			run = args[i+2]
		case "lane":
			lane = args[i+2]
		case "tag_index":
			tag = args[i+2]
		case "type":
			ftype = args[i+2]
		}
	}

	if ftype == "" {
		ftype = "cram"
	}

	return
}

func run_iquest(run string) []string {
	out, err := exec.Command("iquest", "-z", "seq", "--no-page", `"%s/%s"`, fmt.Sprintf(`"SELECT COLL_NAME, DATA_NAME WHERE META_DATA_ATTR_NAME = 'id_run' AND META_DATA_ATTR_VALUE = '%s' AND META_DATA_ATTR_NAME = 'target' AND META_DATA_ATTR_VALUE = '1'"`, run)).Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(out), "\n")[1:]
}

func filter_iquest(paths []string, lane, tag, ftype string) string {
	wanted := fmt.Sprintf("_%s#%s.%s\"", lane, tag, ftype)

	for _, path := range paths {
		base := filepath.Base(path)

		if strings.HasSuffix(base, wanted) {
			path = strings.TrimPrefix(path, `"`)
			return strings.TrimSuffix(path, `"`)
		}
	}

	return ""
}

func format_path_like_imeta(path string) string {
	dir := filepath.Dir(path)
	base := filepath.Base(path)

	return fmt.Sprintf("collection: %s\ndataObj: %s\n", dir, base)
}
