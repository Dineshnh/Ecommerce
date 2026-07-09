import { useEffect, useState } from "react";
import "./Cart.css";

export default function Cart() {
  const [cart, setCart] =useState([]);
  const [total,setTotal]=useState(0);

  const token=localStorage.getItem("token");

  const loadCart=()=>{
    fetch("/api/cart",{
      headers:{
        Authorization:"Bearer "+token,
      },
    })
      .then(res=>res.json())
      .then(data=>{
        console.log(data);
        console.log(data.items);
        const items=data.items||[];
        setCart(items);

        let grandTotal=0;
        items.forEach(item=>{
          grandTotal+=item.product.price*item.quantity;
        });
        

        setTotal(grandTotal);
      });
  };

  useEffect(()=>{
    loadCart();
  },[]);

  const updateQuantity=(productId,newQuantity)=>{

    if(newQuantity<=0){
      removeItem(productId);
      return;
    }

    fetch("/api/cart/update",{
      method:"PUT",
      headers:{
        "Content-Type":"application/json",
        Authorization:"Bearer "+token,
      },
      body:JSON.stringify({
        product_id:productId,
        quantity:newQuantity,
      }),
    }).then(()=>{
      loadCart();
    });
  };

  const removeItem=(productId)=>{

    fetch(`/api/cart/remove?product_id=${productId}`,{
      method:"DELETE",
      headers:{
        Authorization:"Bearer "+token,
      },
    }).then(()=>{
      loadCart();
    });
  };

  const placeOrder=()=>{

    fetch("/api/orders/place",{
      method:"POST",
      headers:{
        Authorization:"Bearer "+token,
      },
    })
      .then(res=>res.json())
      .then(data=>{
        alert(data.message);
        loadCart();
      });
  };

  if(cart.length===0){
    return(
      <div className="cart-page">
        <h2>Shopping Cart</h2>
        <h3 className="empty-cart">Your cart is empty.</h3>
      </div>
    );
  }

  return(
    <div className="cart-page">

      <h2>Shopping Cart</h2>

      <div className="cart-grid">

        {cart.map(item=>(

          <div className="card" key={item.id}>

            <img
              src={item.product.image_url}
              alt={item.product.name}
              className="product-image"
            />

            <h3>{item.product.name}</h3>

            <p>
              <strong>Category:</strong>{" "}
              {item.product.category}
            </p>

            <p>
              <strong>Price:</strong> ₹
              {item.product.price}
            </p>

            <div className="qty">

              <button
                onClick={()=>
                  updateQuantity(
                    item.product.id,
                    item.quantity-1
                  )
                }
              >
                -
              </button>

              <span>{item.quantity}</span>

              <button
                onClick={()=>
                  updateQuantity(
                    item.product.id,
                    item.quantity+1
                  )
                }
              >
                +
              </button>

            </div>

            <p className="item-total">
              Item Total :
              ₹
              {item.product.price*item.quantity}
            </p>

            <button
              className="remove-btn"
              onClick={()=>
                removeItem(item.product.id)
              }
            >
              Remove
            </button>

          </div>

        ))}

      </div>

      <div className="summary">

        <h2>
          Grand Total : ₹{total}
        </h2>

        <button
          className="order-btn"
          onClick={placeOrder}
        >
          Place Order
        </button>

      </div>

    </div>
  );
}