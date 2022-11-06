use error_stack::{IntoReport, Report, ResultExt};
use graphql_client::{GraphQLQuery, Response};
use std::error::Error;
use std::fmt;
use std::fmt::{Debug, Display};

/**
 * 
 *  Queries
 * 
 **/


#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetCollectionDetailsQuery.gql",
    response_derives = "Debug"
)]
pub struct CollectionDetails;

type Date = String;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetContractEventLogs.gql",
    response_derives = "Debug"
)]
pub struct ContractEvents;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTDetails.gql",
    response_derives = "Debug"
)]
pub struct NFTDetails;


#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTsByENS.gql",
    response_derives = "Debug"
)]
pub struct WalletNFTsByENS;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTsByWalletAddress.gql",
    response_derives = "Debug"
)]
pub struct WalletNFTs;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetNFTsWalletAndContracts.gql",
    response_derives = "Debug"
)]
pub struct NFTsWalletAndContract;

#[derive(GraphQLQuery)]
#[graphql(
    schema_path = "src/schema.gql",
    query_path = "src/queries/GetContractAddressNFTs.gql",
    response_derives = "Debug"
)]
pub struct ContractNFTs;


/**
 * 
 *  Client configuration
 * 
 **/


#[derive(Debug)]
pub enum QuickNodeError {
    Deserialization,
    Request,
    GraphQL(Vec<graphql_client::Error>),
}

impl Display for QuickNodeError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.write_str("QuickNode error")
    }
}

impl Error for QuickNodeError {}


pub struct QuickNodeSDK {
    base_url: String,
    icy_api_key: String,
    client: reqwest::Client,
}

impl QuickNodeSDK {
    pub fn new(icy_api_key: &str) -> Self {
        Self {
            base_url: String::from("https://graphql.icy.tools/graphql"),
            icy_api_key: icy_api_key.to_string(),
            client: reqwest::Client::new(),
        }
    }

    pub fn set_icy_api_key(&mut self, key: String) {
        self.icy_api_key = key;
    }

    pub async fn get_collection_details(
        self,
        address: &str,
    ) -> Result<collection_details::ResponseData, Report<QuickNodeError>> {
        let vars = collection_details::Variables {
            address: address.to_string(),
        };

        let request_body = CollectionDetails::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .header("x-api-key", &self.icy_api_key)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<collection_details::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }
    }

    pub async fn get_contract_address_nfts(
        self,
        address: &str,
        first: Option<i64>,
        after: Option<String>,
    ) -> Result<contract_nf_ts::ResponseData, Report<QuickNodeError>> {
        let vars = contract_nf_ts::Variables {
            address: address.to_string(),
            first,
            after,
        };

        let request_body = ContractNFTs::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .header("x-api-key", &self.icy_api_key)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<contract_nf_ts::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }
    }

    pub async fn get_contract_event_logs(self, address: &str, filter: Option<contract_events::LogsFilterInputType>, first: Option<i64>, after: Option<String>) -> Result<contract_events::ResponseData, Report<QuickNodeError>> {
        let vars = contract_events::Variables {
            address: address.to_string(),
            filter,
            first,
            after,
        };

        let request_body = ContractEvents::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .header("x-api-key", &self.icy_api_key)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<contract_events::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }
    }

    pub async fn get_nft_details(self, contract_address: &str, token_id: &str) -> Result<nft_details::ResponseData, Report<QuickNodeError>> {
        let vars = nft_details::Variables {
            contract_address: contract_address.to_string(),
            token_id: token_id.to_string(),
        };

        let request_body = NFTDetails::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .header("x-api-key", &self.icy_api_key)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<nft_details::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }
    }

    pub async fn get_nfts_by_ens(
        self,
        ens_name: Option<String>,
        first: Option<i64>,
        after: Option<String>,
    ) -> Result<wallet_nf_ts_by_ens::ResponseData, Report<QuickNodeError>> {
        let vars = wallet_nf_ts_by_ens::Variables {
            ens_name,
            first,
            after,
        };

        let request_body = WalletNFTsByENS::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .header("x-api-key", &self.icy_api_key)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<wallet_nf_ts_by_ens::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }

    }

    pub async fn get_nfts_by_wallet_address(
        self,
        address: &str,
        first: Option<i64>,
        after: Option<String>,
    ) -> Result<wallet_nf_ts::ResponseData, Report<QuickNodeError>> {
        let vars = wallet_nf_ts::Variables {
            address: address.to_string(),
            first,
            after,
        };

        let request_body = WalletNFTs::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .header("x-api-key", &self.icy_api_key)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<wallet_nf_ts::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }
    }

    pub async fn get_nfts_wallet_and_contract(
        self,
        filter: Option<nf_ts_wallet_and_contract::TokensFilterInputType>,
        address: Option<String>,
        first: Option<i64>,
        after: Option<String>
    ) -> Result<nf_ts_wallet_and_contract::ResponseData, Report<QuickNodeError>> {
        let vars = nf_ts_wallet_and_contract::Variables {
            filter,
            address,
            first,
            after,
        };

        let request_body = NFTsWalletAndContract::build_query(vars);

        let res = self
            .client
            .post(self.base_url)
            .header("x-api-key", &self.icy_api_key)
            .json(&request_body)
            .send()
            .await
            .into_report()
            .attach_printable("failed to perform POST".to_string())
            .change_context(QuickNodeError::Request)?;

        let response_body: Response<nf_ts_wallet_and_contract::ResponseData> = res
            .json()
            .await
            .into_report()
            .attach_printable("failed to deserialize response".to_string())
            .change_context(QuickNodeError::Deserialization)?;

        if let Some(errors) = response_body.errors {
            return Err(Report::new(QuickNodeError::GraphQL(errors)));
        }

        match response_body.data {
            Some(data) => Ok(data),
            None => Err(Report::new(QuickNodeError::Request)),
        }
    }
    
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn it_get_collection_details() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk
            .get_collection_details("0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D")
            .await;

        match res {
            Ok(collection_details) => println!("{:?}", collection_details),
            Err(err) => println!("{:?}", err),
        }
    }

    #[tokio::test]
    async fn it_gets_contract_address_nfts() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk
            .get_contract_address_nfts(
                "0x2106C00Ac7dA0A3430aE667879139E832307AeAa",
                Option::None,
                Option::None,
            )
            .await;

        match res {
            Ok(nfts) => println!("{:?}", nfts),
            Err(err) => println!("{:?}", err),
        }
    }

    #[tokio::test]
    async fn it_get_nft_details() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk.get_nft_details(
            "ADDR",
            "TOKEN_ID",
        ).await;

        match res {
            Ok(nft_details) => println!("{:?}", nft_details),
            Err(err) => println!("{:?}", err),
        }
    }

    #[tokio::test]
    async fn it_get_nfts_by_ens() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk
            .get_nfts_by_ens(
                Some("FILL_WITH_ENS".to_string()),
                None,
                None,
            )
            .await;

        match res {
            Ok(nfts) => println!("{:?}", nfts),
            Err(err) => println!("{:?}", err),
        }
    }

    #[tokio::test]
    async fn it_get_nfts_by_wallet_address() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk
            .get_nfts_by_wallet_address(
                "FILL_WITH_ADDR",
                None,
                None,
            )
            .await;

        match res {
            Ok(nfts) => println!("{:?}", nfts),
            Err(err) => println!("{:?}", err),
        }
    }

    #[tokio::test]
    async fn it_get_nfts_wallet_and_contract() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk
            .get_nfts_wallet_and_contract(
                None,
                Some("FILL_WITH_ADDR".to_string()),
                None,
                None,
            )
            .await;

        match res {
            Ok(data) => println!("{:?}", data),
            Err(err) => println!("{:?}", err),
        }
    }

    #[tokio::test]
    async fn it_get_contract_event_logs() {
        let sdk = QuickNodeSDK::new("");
        let res = sdk
            .get_contract_event_logs(
                "0x0000000000000000000000000000000000000000",
                None,
                None,
                None,
            )
            .await;

        match res {
            Ok(data) => println!("{:?}", data),
            Err(err) => println!("{:?}", err),
        }
    }

}
