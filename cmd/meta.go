// Copyright © 2016 Dropbox, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/files"
	"github.com/spf13/cobra"
)

func getFileMetadata(path string) (*files.Metadata, error) {
	arg := files.NewGetMetadataArg(path)

	res, err := dbx.GetMetadata(arg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func meta(cmd *cobra.Command, args []string) (err error) {
	path, err := validatePath(args[0])
	if err != nil {
		return
	}

	metadata, err := getFileMetadata(path)

	fmt.Printf("Revision\tSize\tLast modified\tPath\n")
	printFileMetadata(os.Stdout, metadata.File, true)

	return
}

// metaCmd represents the meta command
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Get metadata for a file",
	RunE:  meta,
}

func init() {
	RootCmd.AddCommand(metaCmd)
}
