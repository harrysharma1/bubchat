.PHONY: test_cover test_golden_file cover_html cover_list

test_cover: 
	go test ./... -v -coverprofile=coverage.out

test_golden_file: 
	go test ./client/tui -v -update 

cover_html:
	go tool cover -html=coverage.out

cover_list:
	go tool cover -func=coverage.out