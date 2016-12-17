out:
	go get github.com/mitchellh/gox
	rm -rf ./_obj
	gox  -output "_obj/{{.OS}}_{{.Arch}}/{{.Dir}}" ./...
	mv ./_obj/darwin_386 ./_obj/macos_386
	mv ./_obj/darwin_amd64 ./_obj/macos_amd64
	-find ./_obj \! -name "*.zip" -type d -exec zip -r -j {}.zip {} \; -exec rm -R -d {} \;