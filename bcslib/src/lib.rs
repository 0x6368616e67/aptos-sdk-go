
use serde_json;
use aptos_api_types;


pub fn json_to_bsc(json_data: &str) -> serde_json::Result<()> {
    println!("data:{}", json_data);
    let req :aptos_api_types::EncodeSubmissionRequest = serde_json::from_str(json_data)?;
    println!("req:{:?}", req);
    return Ok(());
}

