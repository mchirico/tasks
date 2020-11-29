
docker-build:
	docker build --build-arg build_task_key=${TASK_KEY} --no-cache -t gcr.io/septapig/tasks:test -f Dockerfile .

push:
	docker push gcr.io/septapig/tasks:test

build:
	go build -v .

run:
	docker run --name tasks --rm -it -p 3000:3000  gcr.io/septapig/tasks:test


deploy:
	gcloud beta run deploy tasks  --image gcr.io/septapig/tasks:test --platform managed \
            --allow-unauthenticated --project septapig \
            --vpc-connector=cloudvpc-east \
            --vpc-egress=all \
            --region us-east1 --port 3000 --max-instances 1  --memory 124Mi


