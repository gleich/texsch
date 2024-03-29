FROM golang:1.16

# Meta data:
LABEL maintainer="matthewgleich@gmail.com"
LABEL description="🏫 Automation for techy students that write papers for high school in LaTeX"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing make
RUN apt-get update && apt-get install make=4.2.1-1.2 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

CMD ["make", "local-test"]
