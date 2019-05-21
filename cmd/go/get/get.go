// Copyright 2019 storyicon@foxmail.com
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

package get

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/storyicon/gos/pkg/util"
)

// CmdGet get downloads the packages named by the import paths, along with their dependencies
var CmdGet = &cobra.Command{
	Use:   "get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]",
	Short: "download and install packages and dependencies",
	Long: `
    Get downloads the packages named by the import paths, along with their
    dependencies. It then installs the named packages, like 'go install'.
    
    The -d flag instructs get to stop after downloading the packages; that is,
    it instructs get not to install the packages.
    
    The -f flag, valid only when -u is set, forces get -u not to verify that
    each package has been checked out from the source control repository
    implied by its import path. This can be useful if the source is a local fork
    of the original.
    
    The -fix flag instructs get to run the fix tool on the downloaded packages
    before resolving dependencies or building the code.
    
    The -insecure flag permits fetching from repositories and resolving
    custom domains using insecure schemes such as HTTP. Use with caution.
    
    The -t flag instructs get to also download the packages required to build
    the tests for the specified packages.
    
    The -u flag instructs get to use the network to update the named packages
    and their dependencies. By default, get uses the network to check out
    missing packages but does not use it to look for updates to existing packages.
    
    The -v flag enables verbose progress and debug output.
    
    Get also accepts build flags to control the installation. See 'go help build'.
    
    When checking out a new package, get creates the target directory
    GOPATH/src/<import-path>. If the GOPATH contains multiple entries,
    get uses the first one. For more details see: 'go help gopath'.
    
    When checking out or updating a package, get looks for a branch or tag
    that matches the locally installed version of Go. The most important
    rule is that if the local installation is running version "go1", get
    searches for a branch or tag named "go1". If no such version exists
    it retrieves the default branch of the package.
    
    When go get checks out or updates a Git repository,
    it also updates any git submodules referenced by the repository.
    
    Get never checks out or updates code stored in vendor directories.
    
    For more about specifying packages, see 'go help packages'.
    
    For more about how 'go get' finds source code to
    download, see 'go help importpath'.
    
    This text describes the behavior of get when using GOPATH
    to manage source code and dependencies.
    If instead the go command is running in module-aware mode,
    the details of get's flags and effects change, as does 'go help get'.
    See 'go help modules' and 'go help module-get'.
    
    See also: go build, go install, go clean.
		`,
	DisableFlagParsing: true,
}

func init() {
	CmdGet.Run = func(cmd *cobra.Command, args []string) {
		fd := util.GetGoBinaryCMD("get", args)
		fd.Env = util.GetEnvWithLocalProxy()
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		fd.Run()
	}
}
