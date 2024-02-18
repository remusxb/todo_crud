# Build
FROM golang:1.22-bullseye as build

WORKDIR /go/src/app
COPY . .

# configuration to fetch private repos
RUN mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts
RUN git config --global url.git@github.com:.insteadOf https://github.com/ \
    && go env -w GOPRIVATE="github.com/remusxb/*"

COPY go.mod go.sum ./
RUN --mount=type=ssh go mod download
RUN --mount=type=ssh make tools

ARG BIN_NAME=/go/bin/todo_crud
RUN make build

# Deploy
FROM gcr.io/distroless/base-debian11

COPY --from=build /go/bin/todo_crud /
ENTRYPOINT ["/todo_crud"]
