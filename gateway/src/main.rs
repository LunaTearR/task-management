mod app;
mod configs;
mod handlers;
mod routes;
mod utils;

#[tokio::main]
async fn main() {
    app::run().await;
}
