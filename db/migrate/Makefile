migrate-install:
	brew install whalebrew
	whalebrew install kei2100/docker-ridgepole

migrate:
	ridgepole --apply -c config.yml -f ./schema/Schemafile
	
migrate-test:
	ridgepole --apply -c config.yml -f ./schema/Schemafile -E test
