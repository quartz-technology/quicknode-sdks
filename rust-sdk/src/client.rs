pub struct QuickNodeClient {
    icy_api_key: String,
}

impl QuickNodeClient {
    pub fn new(icy_api_key: Option<String>) -> Self {
        if icy_api_key.is_none() {
            println!("NO API_KEY");
        }
        
        Self {
            icy_api_key: icy_api_key.unwrap_or(String::from("")),
        }
    }

    pub fn set_icy_api_key(&mut self, key: String) -> &mut self {
        self.icy_api_key = key;
        self
    }
}