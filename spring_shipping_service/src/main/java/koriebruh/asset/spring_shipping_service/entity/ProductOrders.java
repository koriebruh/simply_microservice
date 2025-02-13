package koriebruh.asset.spring_shipping_service.entity;


import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "product_orders")
public class ProductOrders {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int id;

    @Column(name = "order_id")
    private int OrderId;

    @Column(name = "product_id")
    private int productId;

    private int quantity;

    // Relasi ManyToOne, dimana banyak ProductOrders memiliki satu Shipping
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "shipping_id", referencedColumnName = "id")
    private Shipping shipping;
}
