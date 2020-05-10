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
)

func errorExit(location string, msg interface{}) {
	fmt.Printf(`

              __n__n__
       .------'-\00/-'
      /  ##  ## (oo)
      / \## __   ./
        |//YY \|/
        |||   |||

Ermintrude says "Oh my, something went wrong in '%s'"

It looks a bit technical
%s

  `, location, msg)

	os.Exit(1)
}
