import os
import json
import shutil

version = os.environ['VERSION'].lstrip('v')

copies = [
    ('darwin-arm64', os.environ['DARWIN_ARM64_BIN'], 'vx'),
    ('darwin-x64',   os.environ['DARWIN_AMD64_BIN'], 'vx'),
    ('linux-x64',    os.environ['LINUX_AMD64_BIN'],  'vx'),
    ('win32-x64',    os.environ['WINDOWS_AMD64_BIN'], 'vx.exe'),
]

for platform, src, name in copies:
    bin_dir = f'npm/platforms/{platform}/bin'
    os.makedirs(bin_dir, exist_ok=True)
    dst = os.path.join(bin_dir, name)
    shutil.copy2(src, dst)
    if not name.endswith('.exe'):
        os.chmod(dst, 0o755)

for platform, _, _ in copies:
    pkg_path = f'npm/platforms/{platform}/package.json'
    with open(pkg_path) as f:
        pkg = json.load(f)
    pkg['version'] = version
    with open(pkg_path, 'w') as f:
        json.dump(pkg, f, indent=2)
        f.write('\n')

wrapper_path = 'npm/vexal/package.json'
with open(wrapper_path) as f:
    pkg = json.load(f)
pkg['version'] = version
for dep in pkg.get('optionalDependencies', {}):
    pkg['optionalDependencies'][dep] = version
with open(wrapper_path, 'w') as f:
    json.dump(pkg, f, indent=2)
    f.write('\n')

print(f'npm packages prepared for v{version}')
