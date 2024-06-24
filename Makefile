init:
	go mod init algdatstr
	go mod tidy

test:
	go test -v ./...

bench:
	cd power; go test -bench=. | tee ../bench-power.lst
	cd search; go test -bench=. | tee ../bench-search.lst
	cd search/text; go test -bench=. | tee ../../bench-search-text.lst
	cd sort; go test -bench=. | tee ../../bench-sort.lst

clean:
	find . -name '*~' -exec rm {} \;
	find . -name '*.lst' -exec rm {} \;
	find . -name .DS_Store -exec rm {} \;



