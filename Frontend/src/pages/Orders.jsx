import { useEffect, useState } from "react";
import "./Orders.css";

export default function Orders() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    fetch("/api/orders", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("token"),
      },
    })
      .then((res) => res.json())
      .then((data) => setOrders(data.orders || []));
  }, []);

  return (
    <div className="orders-page">
      <h2>Orders</h2>

      <div className="orders-grid">
        {orders.map((o) => (
          <div className="card order-card" key={o.ID}>
            <p><span>Order ID:</span> {o.ID}</p>
            <p><span>Total:</span> ₹{o.Total}</p>
            <p>
              <span>Status:</span>{" "}
              <span className={`status ${o.Status?.toLowerCase()}`}>
                {o.Status}
              </span>
            </p>
          </div>
        ))}
      </div>
    </div>
  );
}