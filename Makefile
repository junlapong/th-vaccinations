build:
	@go build -o ./bin/app

clean:
	@rm -rf ./bin

up:
	@cd ~/tmp && rm *.csv && \
	wget https://raw.githubusercontent.com/owid/covid-19-data/master/public/data/vaccinations/country_data/Thailand.csv && \
	cd -

