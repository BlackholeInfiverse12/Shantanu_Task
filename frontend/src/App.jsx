import BlockchainViewer from "./components/BlockchainViewer";
import RelayEventsViewer from "./components/RelayEventsViewer";
import Blockgrid from "./pages/Blockgrid";
import Lightning from "./components/ui/Lightning";
import Particles from "./components/ui/Particles";

function App() {
  return (
    <div className="flex flex-col md:flex-row gap-8 p-8">
      <Particles />
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

export default App;