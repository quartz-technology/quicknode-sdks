use reqwest;

pub enum QuickNodeError {
    DeserializationError(String),
    RequestError(String),
    GraphQLError(String),
}

impl Display for QuickNodeError {
    fn fmt(&self, fmt: &mut fmt::Formatter<'_>) -> std::fmt::Result {
        f.write_str("Quicknode error")
    }
}


#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetCollectionDetailsQuery.gql",
    response_derives = "Debug",
)]
pub struct CollectionDetailsQuery;

pub struct QuickNodeSDK {
    icy_api_key: String,
    client: reqwest::Client,
}

impl QuickNodeSDK {
    pub fn new(icy_api_key: Option<String>) -> Self {
        if icy_api_key.is_none() {
            println!("NO API_KEY");
        }
        
        Self {
            icy_api_key: icy_api_key.unwrap_or(String::from("")),
            client: reqwest::Client::new(),
        }
    }

    pub fn set_icy_api_key(&mut self, key: String) -> &mut self {
        self.icy_api_key = key;
        self
    }

    pub async fn get_collection_details(address: &str) -> () {
        let vars = collection_details_query::Variables {
            address,
        };

        let request_body = CollectionDetailsQuery::build_query(vars);

        let res = client.post("https://graphql.icy.tools/graphql")
            .json(&request_body)
            .send()
            .await?;
        
        let response_body: Response<example_query::ResponseData> = res.json().await?;
    }
}