package koriebruh.asset.spring_shipping_service.controller;

import koriebruh.asset.spring_shipping_service.service.ShippingService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/shipping")
public class ShippingController {

    @Autowired
    private ShippingService shippingService;

    // Endpoint untuk mendapatkan status pengiriman berdasarkan orderId
    @GetMapping("/{orderId}")
    public ResponseEntity<Object> getShippingStatus(@PathVariable int orderId) {
        try {
            // Memanggil service untuk mendapatkan status pengiriman
            return ResponseEntity.ok(shippingService.findShippingStatus(orderId));
        } catch (RuntimeException e) {
            // Menangani error jika terjadi saat pencarian
            return ResponseEntity.status(404).body(e.getMessage());  // Mengembalikan status 404 dengan pesan error
        }
    }

    // Endpoint untuk memperbarui status pengiriman berdasarkan orderId
    @PutMapping("status/{orderId}")
    public ResponseEntity<Object> updateShippingStatus(@PathVariable int orderId, @RequestBody String shippingStatus) {
        try {
            // Memanggil service untuk memperbarui status pengiriman
            return ResponseEntity.ok(shippingService.updateShippingStatus(orderId, shippingStatus));
        } catch (RuntimeException e) {
            // Menangani error saat memperbarui status pengiriman
            return ResponseEntity.status(500).body(e.getMessage());  // Mengembalikan status 500 dengan pesan error
        }
    }

}
