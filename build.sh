rm -r dist
mkdir dist
go build -o dist/GN_Int

pkgbuild --root dist --identifier com.gn.neu --version 1.0.0 --install-location /Applications GN_Int.pkg
