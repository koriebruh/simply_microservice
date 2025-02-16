use std::time::Duration;
use rdkafka::error::KafkaError::ClientConfig;
use rdkafka::producer::{FutureProducer, FutureRecord};
use rdkafka::util::Timeout;

pub fn create() -> FutureProducer {
    let config = ClientConfig::new()
        .set("bootstrap.servers", "localhost:9092")
        .set("client.id", "my-client")
        .set("acks", "all");

    let producer: FutureProducer = config
        .create().expect("Failure in creating producer");

    producer
}


pub async fn produce(future_producer: FutureProducer, msg: String) {
    let record = FutureRecord::to("test")
        .payload(msg.as_str())
        .key("gtw");

    let status_delivery = future_producer
        .send(record, Timeout::After(Duration::from_secs(5)))
        .await;

    match status_delivery {
        Ok(report) => {println!("Message Send {:?}",report)}
        Err(e) => {println!("Error Producing.. {:?}",e)}
    }
}