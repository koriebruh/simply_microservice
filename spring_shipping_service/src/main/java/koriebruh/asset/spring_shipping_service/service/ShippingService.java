package koriebruh.asset.spring_shipping_service.service;

import koriebruh.asset.spring_shipping_service.entity.Shipping;
import koriebruh.asset.spring_shipping_service.repository.ShippingRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ShippingService {

    @Autowired
    private ShippingRepository shippingRepository;

    public Shipping findShippingStatus(int orderId) {
        try {
            Shipping shipping = shippingRepository.findByOrderId(orderId);
            if (shipping == null) {
                throw new RuntimeException("Shipping with Order ID " + orderId + " not found");
            }
            return shipping;
        } catch (Exception e) {
            throw new RuntimeException("Error while finding shipping status: " + e.getMessage());
        }
    }

    public String updateShippingStatus(int orderId, String shippingStatus) {
        try {
            Shipping shipping = findShippingStatus(orderId);

            shipping.setShippingStatus(shippingStatus);

            return shippingRepository.save(shipping).getShippingStatus();
        } catch (Exception e) {
            throw new RuntimeException("Error while updating shipping status: " + e.getMessage());
        }
    }
}
