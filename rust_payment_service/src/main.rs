use crate::producer::produce;

mod producer;
mod consumer;

#[tokio::main]
async fn main() {
    let producer1 = producer::create();
    produce(producer1,String::from("HEllow guys")).await;

   consumer::start().await;


}
