import { useEffect, useState } from "react";
import "./Admin.css";


export default function Admin(){

    const [stats, setStats] = useState({
        revenue:0,
        users:0,
        orders:0
    });


    const [product,setProduct]=useState({
        name:"",
        description:"",
        category:"",
        price:"",
        stock:"",
        image_url:""
    });


    const token = localStorage.getItem("token");


    // ============================
    // Get Admin Dashboard Stats
    // ============================

    useEffect(()=>{

        fetch("/api/admin/stats",{

            headers:{
                Authorization:"Bearer "+token
            }

        })
        .then(res=>res.json())
        .then(data=>{

            setStats(data);

        })
        .catch(error=>{
            console.log(error);
        });


    },[]);



    // ============================
    // Product Input Change
    // ============================

    const handleChange=(e)=>{

        setProduct({

            ...product,

            [e.target.name]:e.target.value

        });

    };



    // ============================
    // Add Product
    // ============================

    const addProduct=()=>{


        fetch("/api/admin/products",{

            method:"POST",

            headers:{

                "Content-Type":"application/json",

                Authorization:"Bearer "+token

            },

            body:JSON.stringify(product)

        })

        .then(res=>res.json())

        .then(data=>{


            alert(data.message || "Product Added");


            setProduct({

                name:"",
                description:"",
                category:"",
                price:"",
                stock:"",
                image_url:""

            });


        })

        .catch(error=>{

            console.log(error);

        });


    };



    return(

        <div className="admin-page">


            <h1>
                Admin Dashboard
            </h1>



            {/* ============================
                Dashboard Statistics
            ============================= */}


            <div className="dashboard-grid">


                <div className="card dashboard-card revenue">

                    <h3>
                        Revenue
                    </h3>

                    <p>
                        ₹ {stats.revenue}
                    </p>

                </div>




                <div className="card dashboard-card users">

                    <h3>
                        Users
                    </h3>

                    <p>
                        {stats.users}
                    </p>

                </div>





                <div className="card dashboard-card orders">

                    <h3>
                        Orders
                    </h3>

                    <p>
                        {stats.orders}
                    </p>

                </div>


            </div>





            {/* ============================
                 Add Product Section
            ============================= */}



            <div className="product-form">


                <h2>
                    Add Product
                </h2>



                <input
                    name="name"
                    placeholder="Product Name"
                    value={product.name}
                    onChange={handleChange}
                />



                <input
                    name="category"
                    placeholder="Category"
                    value={product.category}
                    onChange={handleChange}
                />



                <input
                    name="price"
                    placeholder="Price"
                    type="number"
                    value={product.price}
                    onChange={handleChange}
                />



                <input
                    name="stock"
                    placeholder="Stock"
                    type="number"
                    value={product.stock}
                    onChange={handleChange}
                />



                <input
                    name="image_url"
                    placeholder="Image URL"
                    value={product.image_url}
                    onChange={handleChange}
                />



                <textarea

                    name="description"

                    placeholder="Description"

                    value={product.description}

                    onChange={handleChange}

                />




                <button onClick={addProduct}>

                    Add Product

                </button>


            </div>



        </div>

    );

}