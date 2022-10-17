FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.18

ENV GO111MODULE=on
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.0
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
RUN wget "https://dlcdn.apache.org/tinkerpop/3.6.1/apache-tinkerpop-gremlin-console-3.6.1-bin.zip" && \
    unzip apache-tinkerpop-gremlin-console-3.6.1-bin.zip && \
    cp -r apache-tinkerpop-gremlin-console-3.6.1 /usr/local/gremlin-console && \
    chmod 700 /usr/local/gremlin-console/bin/gremlin.sh && \
    ln -s /usr/local/gremlin-console/bin/gremlin.sh /usr/local/bin/gremlin