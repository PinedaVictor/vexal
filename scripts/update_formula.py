import os

version         = os.environ['VERSION'].lstrip('v')
base_url        = os.environ['BASE_URL']
amd64_sha       = os.environ['DARWIN_AMD64_SHA']
arm64_sha       = os.environ['DARWIN_ARM64_SHA']
linux_amd64_sha = os.environ['LINUX_AMD64_SHA']

formula = f"""class Vexal < Formula
  desc "Developer workflow CLI with dependency graph, AI tooling, and repo automation."
  homepage "https://www.vexal.io/"
  version "{version}"
  license "BSD-3-Clause"

  on_macos do
    if Hardware::CPU.arm?
      url "{base_url}/vexal_darwin_arm64.tar.gz"
      sha256 "{arm64_sha}"
    else
      url "{base_url}/vexal_darwin_amd64.tar.gz"
      sha256 "{amd64_sha}"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "{base_url}/vexal_linux_amd64.tar.gz"
      sha256 "{linux_amd64_sha}"
    end
  end

  def install
    bin.install "vx"
  end

  def caveats
    <<~EOS
      _______________________________
      /                               \\
      |   ___________________________   |
      |  |                           |  |
      |  |   C:\\> vx -v              |  |
      |  |                           |  |
      |  |   VEXAL CLI INSTALLED!    |  |
      |  |                           |  |
      |  |   [ READY TO LAUNCH ]     |  |
      |  |___________________________|  |
      |                                 |
      \\_______________________________/
             \\_______________/

      Thank you for installing Vexal!

      Contact Information:
      - Docs: https://www.vexal.io/
      - GitHub: https://github.com/PinedaVictor/vexal
      - Email: pinedavictor095@gmail.com
    EOS
  end

  test do
    assert_match "vx version", shell_output("#{{bin}}/vx --version")
  end
end
"""

with open('tap/vexal.rb', 'w') as f:
    f.write(formula)
