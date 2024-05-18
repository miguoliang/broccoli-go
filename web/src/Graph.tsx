import ReactFlow, {
  ReactFlowProvider,
  MiniMap,
  Controls,
  Background,
  BackgroundVariant,
  Panel,
} from "reactflow";
import SearchBox from "./components/SearchBox";

const Graph = () => {
  return (
    <ReactFlowProvider>
      <ReactFlow>
        <Panel position="top-left">
          <div className="flex justify-between text-sm shadow-[0_0_2px_1px_rgba(0,0,0,0.08)] bg-white py-1">
            <SearchBox />
            <button className="px-2">
              Add Node
            </button>
          </div>
        </Panel>
        <Controls className="bg-white" />
        <MiniMap />
        <Background variant={BackgroundVariant.Dots} />
      </ReactFlow>
    </ReactFlowProvider>
  );
};

export default Graph;
