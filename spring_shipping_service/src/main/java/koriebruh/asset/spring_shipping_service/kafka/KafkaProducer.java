package koriebruh.asset.spring_shipping_service.kafka;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.SendResult;
import org.springframework.stereotype.Service;

import java.util.concurrent.CompletableFuture;

@Service
public class KafkaProducer {
    private static final Logger logger = LoggerFactory.getLogger(KafkaProducer.class);
    private final KafkaTemplate<String, String> kafkaTemplate;

    @Autowired
    public KafkaProducer(KafkaTemplate<String, String> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
    }

    public void sendMessage(String topic, String message) {
        logger.info("Producing message to topic {}: {}", topic, message);
        try {
            CompletableFuture<SendResult<String, String>> future = kafkaTemplate.send(topic, message);
            future.whenComplete((result, ex) -> {
                if (ex == null) {
                    logger.info("Message sent successfully to topic {}", topic);
                } else {
                    logger.error("Failed to send message to topic {}: {}", topic, ex.getMessage());
                }
            });
        } catch (Exception e) {
            logger.error("Error while sending message to topic {}: {}", topic, e.getMessage());
            throw new RuntimeException("Error sending message to Kafka", e);
        }
    }
}