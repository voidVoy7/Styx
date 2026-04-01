# Maintainer: voidVoy7 <voidvoy7@proton.me>
pkgname=Styx
pkgver=1.0
pkgrel=0
pkgdesc="A Package Manager wrapper"
arch=("x86_64")
url="https://github.com/voidVoy7/Styx"
license=('GPL-3.0')
makedepends=('go')
source=("https://github.com/voidVoy7/Styx/archive/refs/tags/${pkgver}.tar.gz")
sha256sums=('982083504968cb81b82ca5680d4bb588d4f40b64a4dada57dc8a50a5bc39e1ba')

build() {
	cd "$pkgname-$pkgver"
	go build -buildmode=pie -trimpath
}

package() {
	cd "$pkgname-$pkgver"
	mv Styx styx
	install -Dm0755 -t "$pkgdir/usr/bin/" "styx"
}
