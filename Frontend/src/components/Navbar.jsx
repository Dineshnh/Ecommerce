import { Link } from "react-router-dom";

export default function Navbar(){

    const token = localStorage.getItem("token");

    let role = null;

    if(token){
        try{
            const payload = JSON.parse(
                atob(token.split(".")[1])
            );

            role = payload.role;

        }catch(error){
            console.log(error);
        }
    }


    return(

        <nav className="navbar">

            <h2>Pacific</h2>

            <div className="nav-links">

                <Link to="/products">
                    Products
                </Link>

                <Link to="/cart">
                    Cart
                </Link>

                <Link to="/orders">
                    Orders
                </Link>


                {
                    role === "admin" &&
                    <Link to="/admin">
                        Admin Dashboard
                    </Link>
                }

            </div>

        </nav>

    );
}