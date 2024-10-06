test:
	go test -v ./...

bench:
	cd power; go test -bench=. | tee ../bench-power.lst
	cd search; go test -bench=. | tee ../bench-search.lst
	cd search/text; go test -bench=. | tee ../../bench-search-text.lst
	cd sort; go test -bench=. | tee ../../bench-sort.lst

init:
	go mod init algdatstr

tidy:
	go mod tidy


clean:
	find . -name '*~' -exec rm {} \;
	find . -name '*.lst' -exec rm {} \;
	find . -name .DS_Store -exec rm {} \;



