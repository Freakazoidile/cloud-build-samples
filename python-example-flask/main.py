# Copyright 2020 Google, LLC.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# [START cloudbuild_python_flask]
import os
#import requests

from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello_world():
    name = os.environ.get("NAME", "World")
    
    #response = requests.get('http://169.254.169.254/computeMetadata/v1/instance/service-accounts/', headers={'Metadata-Flavor':'Google'})
    #print(response.content)
    
    from pip._internal.operations import freeze
    except ImportError: # pip < 10.0
    from pip.operations import freeze

    pkgs = freeze.freeze()
    for pkg in pkgs: print(pkg)


    return f"Hello {name}!"
    



if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=int(os.environ.get("PORT", 8080)))
# [END cloudbuild_python_flask]
