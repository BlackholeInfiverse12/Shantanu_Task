import React from "react";

// Simple placeholder components for cards and graphs
const StatCard = ({ title, value }) => (
  <div style={{
    background: "#fff",
    borderRadius: "8px",
    padding: "20px",
    boxShadow: "0 2px 8px rgba(0,0,0,0.05)",
    minWidth: "180px",
    flex: 1,
    margin: "10px"
  }}>
    <h4 style={{ margin: 0, color: "#888" }}>{title}</h4>
    <div style={{ fontSize: "2rem", fontWeight: "bold" }}>{value}</div>
  </div>
);

const TrendGraph = () => (
  <div style={{
    background: "#fff",
    borderRadius: "8px",
    padding: "20px",
    minHeight: "180px",
    boxShadow: "0 2px 8px rgba(0,0,0,0.05)",
    margin: "10px",
    flex: 2
  }}>
    <h4 style={{ color: "#888" }}>Trend Graph</h4>
    <div style={{ height: "120px", background: "#f3f3f3", borderRadius: "4px" }} />
  </div>
);

const RecentEvents = () => (
  <div style={{
    background: "#fff",
    borderRadius: "8px",
    padding: "20px",
    minHeight: "180px",
    boxShadow: "0 2px 8px rgba(0,0,0,0.05)",
    margin: "10px",
    flex: 1
  }}>
    <h4 style={{ color: "#888" }}>Recent Events</h4>
    <ul style={{ paddingLeft: "20px" }}>
      <li>Tx #1 - Ethereum</li>
      <li>Tx #2 - Solana</li>
      <li>Tx #3 - Ethereum</li>
    </ul>
  </div>
);

const LiveNetworkStatus = () => (
  <div style={{
    background: "#fff",
    borderRadius: "8px",
    padding: "20px",
    minHeight: "120px",
    boxShadow: "0 2px 8px rgba(0,0,0,0.05)",
    margin: "10px",
    flex: 1
  }}>
    <h4 style={{ color: "#888" }}>Live Network Status</h4>
    <div>
      <span style={{ color: "green", fontWeight: "bold" }}>●</span> All systems operational
    </div>
  </div>
);

function Dashboard() {
  return (
    <div style={{ display: "flex", minHeight: "100vh", background: "#f5f6fa", color: "#333" ,width: "100vw"}}>
      {/* Aside Menu */}
      <aside style={{
        width: "220px",
        background: "#22223b",
        color: "#fff",
        padding: "30px 0",
        display: "flex",
        flexDirection: "column",
        alignItems: "center"
      }}>
        <h2 style={{ marginBottom: "40px", fontWeight: "bold" }}>Menu</h2>
        <nav>
          <ul style={{ listStyle: "none", padding: 0, width: "100%" }}>
            <li style={{ padding: "15px 30px", cursor: "pointer", borderRadius: "6px", marginBottom: "10px", background: "#4a4e69" }} onClick={() => window.location.href = '/'}>
              Blockchain
            </li>
            <li style={{ padding: "15px 30px", cursor: "pointer", borderRadius: "6px", marginBottom: "10px", background: "#4a4e69" }}>
              Ethereum Events
            </li>
            <li style={{ padding: "15px 30px", cursor: "pointer", borderRadius: "6px", background: "#4a4e69" }}>
              Solana Events
            </li>
          </ul>
        </nav>
      </aside>

      {/* Main Content */}
      <main style={{ flex: 1, padding: "40px" }}>
        <h1 style={{ marginBottom: "30px" }}>Dashboard</h1>
        {/* Grid of Stat Cards */}
        <div style={{ display: "flex", gap: "10px", marginBottom: "30px" }}>
          <StatCard title="Transactions" value="12,345" />
          <StatCard title="Cost Volume" value="$1,234,567" />
          <StatCard title="Active Users" value="789" />
        </div>
        {/* Grid for Trend Graph and Recent Events */}
        <div style={{ display: "flex", gap: "10px", marginBottom: "30px" }}>
          <TrendGraph />
          <RecentEvents />
        </div>
        {/* Live Network Status */}
        <div style={{ display: "flex", gap: "10px" }}>
          <LiveNetworkStatus />
        </div>
      </main>
    </div>
  );
}

export default Dashboard;