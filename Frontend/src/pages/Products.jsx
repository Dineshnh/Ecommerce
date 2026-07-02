import { useEffect, useState } from "react";
import "./Products.css";

export default function Products() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetch("/api/products")
      .then((res) => res.json())
      .then((data) => setProducts(data));
  }, []);

  const addToCart = async (id) => {
    await fetch("/api/cart/add", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + localStorage.getItem("token"),
      },
      body: JSON.stringify({ product_id: id, quantity: 1 }),
    });

    alert("Added to cart");
  };

  return (
    <div className="products-page">
      <nav className="navbar">
        <h2>Pacific</h2>

        <div className="nav-links">
          <a href="/products">Products</a>
          <a href="/cart">Cart</a>
          <a href="/orders">Orders</a>
        </div>
      </nav>

      <h1 className="title">Products</h1>

      <div className="products-grid">
        {products.map((p) => (
          <div className="card" key={p.id}>
            <img
              src={
                p.image_url ||
                "https://via.placeholder.com/250x180?text=No+Image"
              }
              alt={p.name}
            />

            <div className="card-body">
              <h3>{p.name}</h3>

              <p className="category">
                {p.category || "Mobile Phone"}
              </p>

              <p className="description">
                {p.description || "No description available."}
              </p>

              <div className="price-stock">
                <span className="price">₹{p.price}</span>

                <span
                  className={
                    p.stock > 0 ? "stock available" : "stock unavailable"
                  }
                >
                  {p.stock > 0
                    ? `${p.stock} Available`
                    : "Out of Stock"}
                </span>
              </div>

              <button
                disabled={p.stock === 0}
                onClick={() => addToCart(p.id)}
              >
                {p.stock > 0 ? "Add to Cart" : "Unavailable"}
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}