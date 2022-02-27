FROM gitpod/workspace-full

USER root

# Install swagger-codegen
ENV SWAGGER_CODEGEN_VERSION=2.4.8
RUN wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/${SWAGGER_CODEGEN_VERSION}/swagger-codegen-cli-${SWAGGER_CODEGEN_VERSION}.jar \
    -O /usr/local/bin/swagger-codegen-cli.jar && \
    echo -e '#!/bin/bash\njava -jar /usr/local/bin/swagger-codegen-cli.jar "$@"' > /usr/local/bin/swagger-codegen && \
	chmod +x /usr/local/bin/swagger-codegen

# Install Bazelisk as bazel to manage Bazel
RUN go install github.com/bazelbuild/bazelisk@latest && \
    mv $(which bazelisk) /usr/local/bin/bazel

# Install buildifier for bazel formatting
RUN go install github.com/bazelbuild/buildtools/buildifier@latest

USER gitpod
