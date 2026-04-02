# Maintainer: voidVoy7 <voidvoy7@proton.me>
pkgname=Omniwrap
pkgver=1.0
pkgrel=0
pkgdesc="A Package Manager wrapper"
arch=("x86_64")
url="https://github.com/voidVoy7/Omniwrap"
license=('GPL-3.0')
makedepends=('go')
source=("https://github.com/voidVoy7/Styx/archive/refs/tags/${pkgver}.tar.gz")
sha256sums=('f86e13efe340cf4a055d8f3c3d353206ed12e0bc12716714e24dd31a92d0d141')

build() {
	cd "$pkgname-$pkgver"
	go build -buildmode=pie -trimpath
}

package() {
	cd "$pkgname-$pkgver"
	mv Omniwrap ow
	install -Dm0755 -t "$pkgdir/usr/bin/" "ow"
}
