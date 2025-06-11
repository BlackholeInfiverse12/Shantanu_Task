import React from "react";
import { BrowserRouter as Router, Routes, Route, useNavigate } from "react-router-dom";
import BlockchainViewer from "./components/BlockchainViewer";
import Particles from "./components/ui/Particles";
import SplashCursor from "./components/SplashCursor";
import Dashboard from "./pages/Dashboard";

function Home() {
  const navigate = useNavigate();
  return (
    <div className="flex flex-col md:flex-row gap-8 p-8">
      <Particles />
      <SplashCursor />
      <button onClick={() => navigate("/dashboard")}> Main Dashboard </button>
      <div className="flex-1">
        <BlockchainViewer style={{ backdropFilter: "blur(10px)" }} />
        {/* <Blockgrid/> */}
      </div>
      {/* <div className="flex-1">
        <RelayEventsViewer />
      </div> */}
    </div>
  );
}

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  );
}

export default App;