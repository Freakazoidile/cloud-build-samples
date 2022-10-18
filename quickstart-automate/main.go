// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
"fmt"
"net/http"
"net/http/httputil"
)

func main () {
  fmt.Println("h4x the plan3t!")

  var serverAddr = "https://google.com"
  
  req, err := http.NewRequest(http.MethodGet, serverAddr, nil)
  req.Header.Add("test-header", "test-header-value")
  reqDump, err := httputil.DumpRequestOut(req, true)

  // http://169.254.169.254/computeMetadata/v1/instance/service-accounts -H "Metadata-Flavor: Google"


  //fmt.Printf("REQUEST:\n%s", string(reqDump))

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
      log.Fatal(err)
  }

  respDump, err := httputil.DumpResponse(resp, true)
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("RESPONSE:\n%s", string(respDump))
  

}
