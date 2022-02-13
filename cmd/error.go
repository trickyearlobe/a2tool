/*
Copyright © 2020 Richard Nixon

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const ermintrudeSays = `

         __n__n__
  .------'-\00/-'
 /  ##  ## (oo)
/ \## __   ./
  |//YY \|/
  |||   |||

Ermintrude says "Oh my gosh, something went wrong near %s line %d"

It looks a bit technical
%s

`

func errorExit(err interface{}) {
	if err == nil {
		return
	}

	_, filename, line, _ := runtime.Caller(1)
	filename = filepath.Base(filename)
	fmt.Printf(ermintrudeSays, filename, line, err)
	os.Exit(1)
}
