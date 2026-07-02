import { useEffect, useState } from "react";
import "./Cart.css";

export default function Cart() {
  const [cart, setCart] = useState([]);

  useEffect(() => {
    fetch("/api/cart", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("token"),
      },
    })
      .then((res) => res.json())
      .then((data) => setCart(data.items || []));
  }, []);

  return (
    <div className="cart-page">
      <h2>Cart</h2>

      <div className="cart-grid">
        {cart.map((item, i) => (
          <div className="card" key={i}>
            <p>Product ID: {item.product_id}</p>
            <p>Quantity: {item.quantity}</p>
          </div>
        ))}
      </div>
    </div>
  );
}