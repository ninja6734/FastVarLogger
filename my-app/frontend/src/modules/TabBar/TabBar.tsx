import { useState } from "react";
import ConnectContent from "../ConnectTab/ConnectTab";
import RawContent from "../RawDataTab/RawDataTab";
import GraphicContent from "../GraphicDataTab/GraphicDataTab";

function TabBar() {
    const [activeTab, setActiveTab] = useState("home")

    return(
        <div>
            <div>
                <button onClick={() => setActiveTab("Connect")}>Connect</button>
                <button onClick={() => setActiveTab("Raw")}>Raw Data</button>
                <button onClick={() => setActiveTab("Graphic")}>Graphical Data</button>
            </div>
            <div>
                {activeTab === 'Connect' && <ConnectContent />}
                {activeTab === 'Raw' && <RawContent />}
                {activeTab === "Graphic" && <GraphicContent/>}
            </div>
        </div>
    )
}

export default TabBar