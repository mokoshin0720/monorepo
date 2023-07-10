FROM public.ecr.aws/lambda/go:latest

ENV GO111MODULE=on

WORKDIR ${LAMBDA_TASK_ROOT}

# Copy function code
COPY . ${LAMBDA_TASK_ROOT}

RUN go mod download

CMD [ "cmd.server.main" ]