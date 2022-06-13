
all:
	go build

# See: https://github.com/golang/go/issues/46456
deptree:
	gomod graph '**' -a >deps.dot
	sfdp -Tpdf -o deps.pdf deps.dot
