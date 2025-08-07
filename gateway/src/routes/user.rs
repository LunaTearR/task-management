use crate::handlers::proxy::proxy_service_user;
use axum::{
    Router,
    routing::{delete, get, post, put},
};

pub fn user_router() -> Router {
    let api_v1 = "/api/v1";
    Router::new()
        .route(
            format!("{}/users", api_v1).as_str(),
            get(proxy_service_user),
        )
        .route(
            format!("{}/users", api_v1).as_str(),
            post(proxy_service_user),
        )
        .route(
            format!("{}/users/by-ids", api_v1).as_str(),
            post(proxy_service_user),
        )
        .route(&format!("{}/users/{{id}}", api_v1), put(proxy_service_user))
        .route(
            &format!("{}/users/{{id}}", api_v1),
            delete(proxy_service_user),
        )
}
