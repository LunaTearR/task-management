use axum::Router;
use std::net::SocketAddr;

use crate::configs::cors::cors_layer;
use crate::routes::project::project_router;
use crate::routes::task::task_router;
use crate::routes::user::user_router;
use crate::utils::fallback::not_found;

pub async fn run() {
    let app = Router::new()
        .merge(user_router())
        .merge(task_router())
        .merge(project_router())
        .layer(cors_layer())
        .fallback(not_found);

    let addr = SocketAddr::from(([0, 0, 0, 0], 8084));
    println!("🚀 API Gateway running on http://{}", addr);

    axum::serve(
        tokio::net::TcpListener::bind(addr).await.unwrap(),
        app.into_make_service(),
    )
    .await
    .unwrap();
}
