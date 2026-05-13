import os
import json

version      = os.environ['VERSION'].lstrip('v')
base_url     = os.environ['BASE_URL']
windows_sha  = os.environ['WINDOWS_AMD64_SHA']

manifest = {
    "version": version,
    "description": "Developer workflow CLI with dependency graph, AI tooling, and repo automation.",
    "homepage": "https://www.vexal.io/",
    "license": "BSD-3-Clause",
    "url": f"{base_url}/vexal_windows_amd64.zip",
    "hash": f"sha256:{windows_sha}",
    "bin": "vx.exe",
    "checkver": {
        "github": "https://github.com/PinedaVictor/vexal"
    },
    "autoupdate": {
        "url": "https://github.com/PinedaVictor/vexal/releases/download/v$version/vexal_windows_amd64.zip",
        "hash": {
            "url": "https://github.com/PinedaVictor/vexal/releases/download/v$version/checksums.txt",
            "regex": "([a-fA-F0-9]+)\\s+vexal_windows_amd64.zip"
        }
    }
}

os.makedirs('scoop/bucket', exist_ok=True)
with open('scoop/bucket/vexal.json', 'w') as f:
    json.dump(manifest, f, indent=4)
    f.write('\n')
