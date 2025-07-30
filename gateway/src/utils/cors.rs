use tower_http::cors::{CorsLayer};

#[allow(dead_code)]
pub fn init() -> CorsLayer {
    CorsLayer::very_permissive()
}