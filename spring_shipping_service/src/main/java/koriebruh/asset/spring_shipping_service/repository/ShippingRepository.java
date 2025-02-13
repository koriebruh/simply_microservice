package koriebruh.asset.spring_shipping_service.repository;

import koriebruh.asset.spring_shipping_service.entity.Shipping;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ShippingRepository extends JpaRepository<Shipping, Long> {
    Shipping findByOrderId(int orderId);
}
