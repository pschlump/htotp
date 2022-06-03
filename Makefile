
all:
	go build

deptree:
	gomod graph '**' -a >deps.dot
	sfdp -Tpdf -o deps.pdf deps.dot
