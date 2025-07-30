FROM rust:1.88.0-slim-bullseye

ARG APP_NAME=gateway

WORKDIR /app

# ติดตั้ง pkg-config และ OpenSSL dev
RUN apt-get update && \
    apt-get install -y pkg-config libssl-dev && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN cargo build --locked --release
RUN cp ./target/release/$APP_NAME /bin/server

ENV ROCKET_ADDRESS=0.0.0.0
CMD ["/bin/server"]
