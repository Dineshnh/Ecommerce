import { useEffect, useState } from "react";
import "./Admin.css";

export default function Admin() {
  const [stats, setStats] = useState({
    revenue: 0,
    users: 0,
    orders: 0,
  });

  useEffect(() => {
    fetch("/api/admin/stats", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("token"),
      },
    })
      .then((res) => res.json())
      .then((data) => setStats(data || stats));
  }, []);

  return (
    <div className="admin-page">
      <h2>Admin Dashboard</h2>

      <div className="dashboard-grid">
        <div className="card dashboard-card revenue">
          <h3>Revenue</h3>
          <p>₹ {stats.revenue}</p>
        </div>

        <div className="card dashboard-card users">
          <h3>Users</h3>
          <p>{stats.users}</p>
        </div>

        <div className="card dashboard-card orders">
          <h3>Orders</h3>
          <p>{stats.orders}</p>
        </div>
      </div>
    </div>
  );
}