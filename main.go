/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package main

import (
	"github.com/immanoj16/smart-git/cmd"
)

// var sugarLogger *zap.SugaredLogger

func main() {
	cmd.Execute()
}

// func simpleHTTPGet(url string) {
// 	zap.S().Debugf("Trying to hit GET request for %s", url)
// 	resp, err := http.Get(url)

// 	if err != nil {
// 		zap.S().Errorf("Error fetching url... %s : Error = %s", url, err)
// 	} else {
// 		zap.S().Infof("Success! statusCode = %s for URL %s", resp.Status, url)
// 		resp.Body.Close()
// 	}
// }
