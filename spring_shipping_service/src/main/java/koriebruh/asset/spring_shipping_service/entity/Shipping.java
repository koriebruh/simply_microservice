package koriebruh.asset.spring_shipping_service.entity;


import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.Set;

@Getter
@Setter
@Entity
@NoArgsConstructor
@AllArgsConstructor
@Table(name = "shipping")
public class Shipping {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    // Relasi OneToMany, dimana Shipping memiliki banyak ProductOrders
    @OneToMany(mappedBy = "shipping", cascade = CascadeType.ALL, fetch = FetchType.LAZY)
    private Set<ProductOrders> items;

    private int amount;

    @Column(name = "order_id")
    private int orderId;

    @Column(name = "payment_method")
    private String paymentMethod;

    @Column(name = "shipping_addr")
    private String shippingAddr;

    @Column(name = "shipping_status")
    private String shippingStatus;

    @Column(name = "created_at")
    private long createdAt;

    @Column(name = "updated_at")
    private long updatedAt;

    @Column(name = "deleted_at")
    private long deletedAt;

}
