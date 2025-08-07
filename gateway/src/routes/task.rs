use crate::handlers::proxy::proxy_service_task;
use axum::{
    Router,
    routing::{get, post},
};

pub fn task_router() -> Router {
    let api_v1 = "/api/v1";
    Router::new()
        .route(
            format!("{}/tasks", api_v1).as_str(),
            get(proxy_service_task),
        )
        .route(
            format!("{}/tasks", api_v1).as_str(),
            post(proxy_service_task),
        )
}
