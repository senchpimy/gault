build:
	go build .
install:
	go build .
	chmod +x gault
	cp gault /usr/bin/gault
	cp passwords /usr/local/gault/passwords
