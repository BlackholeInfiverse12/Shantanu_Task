import BlockchainViewer from "./components/BlockchainViewer";
import RelayEventsViewer from "./components/RelayEventsViewer";

function App() {
  return (
    <div className="flex flex-col md:flex-row gap-8 p-8">
      <div className="flex-1">
        <BlockchainViewer />
      </div>
      {/* <div className="flex-1">
        <RelayEventsViewer />
      </div> */}
    </div>
  );
}

export default App;