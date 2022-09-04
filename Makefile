build:
	sam build
	mkdir .aws-sam/build/HelloWorldFunction/configs
	cp configs/* .aws-sam/build/HelloWorldFunction/configs/

d:build
	sam deploy