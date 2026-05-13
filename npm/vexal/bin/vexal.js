#!/usr/bin/env node

const path = require('path');
const { execFileSync } = require('child_process');

const PLATFORMS = {
  'darwin arm64': '@dreamlikedigital/vexal-darwin-arm64',
  'darwin x64':   '@dreamlikedigital/vexal-darwin-x64',
  'linux x64':    '@dreamlikedigital/vexal-linux-x64',
  'win32 x64':    '@dreamlikedigital/vexal-win32-x64',
};

const key = `${process.platform} ${process.arch}`;
const pkg = PLATFORMS[key];

if (!pkg) {
  process.stderr.write(`vexal: unsupported platform ${process.platform}/${process.arch}\n`);
  process.exit(1);
}

const binName = process.platform === 'win32' ? 'vx.exe' : 'vx';

let binaryPath;
try {
  binaryPath = require.resolve(path.join(pkg, 'bin', binName));
} catch {
  process.stderr.write(`vexal: platform package ${pkg} not found. Try reinstalling vexal.\n`);
  process.exit(1);
}

try {
  execFileSync(binaryPath, process.argv.slice(2), { stdio: 'inherit' });
} catch (err) {
  process.exit(err.status ?? 1);
}
