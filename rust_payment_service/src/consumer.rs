use rdkafka::consumer::{CommitMode, Consumer, StreamConsumer};
use rdkafka::error::KafkaError::ClientConfig;
use rdkafka::Message;

pub async fn start(){
    let consumer  = create();
    consume(consumer).await
}

pub fn create() -> StreamConsumer{
    let mut config = ClientConfig::new()
        .set("bootstrap.servers", "localhost:9092")
        .set("auto.offset.reset", "earliest")
        .set("group.id","payment_status")
        .set("socket.timeout.ms", "4000");

    let consumer : StreamConsumer = config.create()
        .expect("Fail to create consumer");

    consumer
}

async fn consume(consumer : StreamConsumer){
    consumer.subscribe(
        &["test"]
    ).expect("Can't subscribe");

    loop {
        match consumer.recv().await {
            Err(e) => println!("{:?}",e),
            Ok(message) => {
                    match message.payload_view::<str>() {
                        None => println!("None message read"),
                        Some(Ok(msg)) => println!("Message Consumed : {:?}", msg),
                        Some(Err(e)) => println!("Error Parsing : {:?}",e),
                    }
                consumer.commit_message(&message, CommitMode::Async).unwrap()
            }
        }
    }
}